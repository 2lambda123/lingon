// Code generated by go-kart. EDIT AS MUCH AS YOU LIKE.

package webhook

import (
	"github.com/volvo-cars/lingon/pkg/kubeutil"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecretstoreValidateValidatingWH = &admissionregistrationv1.ValidatingWebhookConfiguration{
	TypeMeta: kubeutil.TypeValidatingWebhookConfigurationV1,
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{"external-secrets.io/component": "webhook"},
		Name:   "secretstore-validate",
	},
	Webhooks: []admissionregistrationv1.ValidatingWebhook{
		{
			AdmissionReviewVersions: []string{"v1", "v1beta1"},
			ClientConfig: admissionregistrationv1.WebhookClientConfig{
				Service: &admissionregistrationv1.ServiceReference{
					Name:      "external-secrets-webhook",
					Namespace: "external-secrets",
					Path:      P("/validate-external-secrets-io-v1beta1-secretstore"),
				},
			},
			Name: "validate.secretstore.external-secrets.io",
			Rules: []admissionregistrationv1.RuleWithOperations{
				{
					Operations: []admissionregistrationv1.OperationType{
						admissionregistrationv1.OperationType("CREATE"),
						admissionregistrationv1.OperationType("UPDATE"),
						admissionregistrationv1.OperationType("DELETE"),
					},
					Rule: admissionregistrationv1.Rule{
						APIGroups:   []string{"external-secrets.io"},
						APIVersions: []string{"v1beta1"},
						Resources:   []string{"secretstores"},
						Scope:       P(admissionregistrationv1.ScopeType("Namespaced")),
					},
				},
			},
			SideEffects:    P(admissionregistrationv1.SideEffectClass("None")),
			TimeoutSeconds: P(int32(5)),
		}, {
			AdmissionReviewVersions: []string{"v1", "v1beta1"},
			ClientConfig: admissionregistrationv1.WebhookClientConfig{
				Service: &admissionregistrationv1.ServiceReference{
					Name:      "external-secrets-webhook",
					Namespace: "external-secrets",
					Path:      P("/validate-external-secrets-io-v1beta1-clustersecretstore"),
				},
			},
			Name: "validate.clustersecretstore.external-secrets.io",
			Rules: []admissionregistrationv1.RuleWithOperations{
				{
					Operations: []admissionregistrationv1.OperationType{
						admissionregistrationv1.OperationType("CREATE"),
						admissionregistrationv1.OperationType("UPDATE"),
						admissionregistrationv1.OperationType("DELETE"),
					},
					Rule: admissionregistrationv1.Rule{
						APIGroups:   []string{"external-secrets.io"},
						APIVersions: []string{"v1beta1"},
						Resources:   []string{"clustersecretstores"},
						Scope:       P(admissionregistrationv1.ScopeType("Cluster")),
					},
				},
			},
			SideEffects:    P(admissionregistrationv1.SideEffectClass("None")),
			TimeoutSeconds: P(int32(5)),
		},
	},
}