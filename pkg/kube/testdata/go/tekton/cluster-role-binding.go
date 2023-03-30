// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package tekton

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PipelinesControllerClusterAccessCRB = &rbacv1.ClusterRoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "controller",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name: "tekton-pipelines-controller-cluster-access",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "ClusterRole",
		Name:     "tekton-pipelines-controller-cluster-access",
	},
	Subjects: []rbacv1.Subject{rbacv1.Subject{
		Kind:      "ServiceAccount",
		Name:      "tekton-pipelines-controller",
		Namespace: "tekton-pipelines",
	}},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
	},
}

var PipelinesControllerTenantAccessCRB = &rbacv1.ClusterRoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "controller",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name: "tekton-pipelines-controller-tenant-access",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "ClusterRole",
		Name:     "tekton-pipelines-controller-tenant-access",
	},
	Subjects: []rbacv1.Subject{rbacv1.Subject{
		Kind:      "ServiceAccount",
		Name:      "tekton-pipelines-controller",
		Namespace: "tekton-pipelines",
	}},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
	},
}

var PipelinesResolversCRB = &rbacv1.ClusterRoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "tekton-pipelines-resolvers",
		Namespace: "tekton-pipelines-resolvers",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "ClusterRole",
		Name:     "tekton-pipelines-resolvers-resolution-request-updates",
	},
	Subjects: []rbacv1.Subject{rbacv1.Subject{
		Kind:      "ServiceAccount",
		Name:      "tekton-pipelines-resolvers",
		Namespace: "tekton-pipelines-resolvers",
	}},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
	},
}

var PipelinesWebhookClusterAccessCRB = &rbacv1.ClusterRoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "webhook",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name: "tekton-pipelines-webhook-cluster-access",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "ClusterRole",
		Name:     "tekton-pipelines-webhook-cluster-access",
	},
	Subjects: []rbacv1.Subject{rbacv1.Subject{
		Kind:      "ServiceAccount",
		Name:      "tekton-pipelines-webhook",
		Namespace: "tekton-pipelines",
	}},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
	},
}