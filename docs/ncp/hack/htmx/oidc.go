package htmx

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"golang.org/x/oauth2"
)

func SetupOIDCHandler(ctx context.Context, nc *nats.Conn, r *chi.Mux) error {
	name := "oidc"
	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("jetstream: %w", err)
	}
	bucket, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: name,
	})
	if err != nil {
		return fmt.Errorf("create key value: %w", err)
	}
	// Setup OIDC server
	provider, err := oidc.NewProvider(ctx, "http://localhost:9998/")
	if err != nil {
		return fmt.Errorf("new oidc provider: %w", err)
	}
	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     "web",
		ClientSecret: "secret",
		RedirectURL:  "http://localhost:9999/auth/callback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile"},
	}
	oidcConfig := &oidc.Config{
		ClientID: oauth2Config.ClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	or := &oidcHandler{
		ctx:      ctx,
		config:   &oauth2Config,
		verifier: verifier,
		bucket:   bucket,
	}

	r.Use(or.authMiddleware)
	r.Get("/login", or.login)
	r.Get("/logout", or.logout)
	r.Get("/auth/callback", or.authCallback)

	// l, err := net.Listen("tcp", ":9999")
	// if err := Setup(r); err != nil {
	// 	return fmt.Errorf("start htmx: %w", err)
	// }

	// if err != nil {
	// 	return fmt.Errorf("listen: %w", err)
	// }
	// actor.listener = l

	// go func() {
	// 	if err := http.Serve(l, r); err != nil {
	// 		slog.Error("serve", "error", err)
	// 	}
	// }()

	return nil
}

type oidcHandler struct {
	ctx    context.Context
	config *oauth2.Config
	// provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
	bucket   nats.KeyValue
}

func (or *oidcHandler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/auth/callback" {
			next.ServeHTTP(w, r)
		}
		sessionID, err := getSessionCookie(r)
		if err != nil {
			http.Error(w, "getting session cookie: "+err.Error(), http.StatusUnauthorized)
			return
		}
		kve, err := or.bucket.Get(sessionID.String())
		if err != nil {
			http.Error(w, "invalid session: "+err.Error(), http.StatusUnauthorized)
			return
		}
		var userInfo UserInfo
		if err := json.Unmarshal(kve.Value(), &userInfo); err != nil {
			http.Error(w, "unmarshalling claims: "+err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), AuthContext, userInfo)
		slog.Info("CONTEXT WITH VALUE", "key", AuthContext, "value", userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (or *oidcHandler) login(w http.ResponseWriter, req *http.Request) {
	state := uuid.New().String()
	nonce := uuid.New().String()
	c := &http.Cookie{
		Name:     "nonce",
		Value:    nonce,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   req.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	http.Redirect(w, req, or.config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
}

func (or *oidcHandler) logout(w http.ResponseWriter, req *http.Request) {
	// sessionID, err := getSessionCookie(r)
	// if err != nil {
	// 	http.Error(w, "Invalid session: "+err.Error(), http.StatusUnauthorized)
	// 	return
	// }
	// if err := p.store.EndSession(sessionID); err != nil {
	// 	http.Error(w, "Cannot logout: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }
	http.Redirect(w, req, "http://localhost:9999", http.StatusSeeOther)
}

func (or *oidcHandler) authCallback(w http.ResponseWriter, req *http.Request) {
	// Verify the state
	state := req.FormValue("state")
	if state == "" {
		http.Error(w, "State not found", http.StatusUnauthorized)
		return
	}
	expectedState := req.URL.Query().Get("state")
	if expectedState != state {
		http.Error(w, "Invalid state", http.StatusUnauthorized)
		return
	}

	oauth2Token, err := or.config.Exchange(or.ctx, req.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "exchange: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	// Parse and verify ID Token payload.
	idToken, err := or.verifier.Verify(or.ctx, rawIDToken)
	if err != nil {
		http.Error(w, "token verification failed: "+err.Error(), http.StatusUnauthorized)
		return
	}
	// Verify the nonce
	nonce, err := req.Cookie("nonce")
	if err != nil {
		http.Error(w, "Nonce not found", http.StatusUnauthorized)
		return
	}
	if idToken.Nonce != nonce.Value {
		http.Error(w, "Nonce did not match token nonce", http.StatusUnauthorized)
		return
	}

	var claims UserInfo
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "unmarshalling claims: "+err.Error(), http.StatusUnauthorized)
		return
	}
	// HACK: remove this as it only applies while testing against zitadel oidc provider
	if claims.Email == "" {
		claims.Email = "admin@localhost"
	}
	if claims.Name == "" {
		claims.Name = "admin"
	}
	// END HACK
	bClaims, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, "marhsalling claims: "+err.Error(), http.StatusUnauthorized)
		return
	}
	slog.Info("claims", "claims", claims)
	sessionID := uuid.New()
	if _, err := or.bucket.Put(sessionID.String(), bClaims); err != nil {
		http.Error(w, "cannot create session: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeSessionCookieHeader(w, sessionID)
	http.Redirect(w, req, "http://localhost:9999", http.StatusSeeOther)
}

const cookieName = "ncp_session"

func getSessionCookie(r *http.Request) (uuid.UUID, error) {
	sessionCookie, err := r.Cookie(cookieName)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.Parse(sessionCookie.Value)
}

func writeSessionCookieHeader(w http.ResponseWriter, sessionID uuid.UUID) {
	// Create cookie
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    sessionID.String(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(time.Hour.Seconds() * 8),
		Path:     "/",
	})
}