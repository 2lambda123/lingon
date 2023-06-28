// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package vmk8s

import (
	"github.com/VictoriaMetrics/operator/api/victoriametrics/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CadvisorNodeScrape = &v1beta1.VMNodeScrape{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "vmk8s",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "victoria-metrics-k8s-stack",
			"app.kubernetes.io/version":    "v1.91.2",
			"helm.sh/chart":                "victoria-metrics-k8s-stack-0.16.3",
		},
		Name:      "vmk8s-victoria-metrics-k8s-stack-cadvisor",
		Namespace: "monitoring",
	},
	Spec: v1beta1.VMNodeScrapeSpec{
		BearerTokenFile: "/var/run/secrets/kubernetes.io/serviceaccount/token",
		HonorLabels:     true,
		Interval:        "30s",
		MetricRelabelConfigs: []*v1beta1.RelabelConfig{
			{
				Action: "labeldrop",
				Regex:  "(uid)",
			}, {
				Action: "labeldrop",
				Regex:  "(id|name)",
			}, {
				Action:                 "drop",
				Regex:                  "(rest_client_request_duration_seconds_bucket|rest_client_request_duration_seconds_sum|rest_client_request_duration_seconds_count)",
				SourceLabels:           []string{"__name__"},
				UnderScoreSourceLabels: []string{"__name__"},
			},
		},
		Path: "/metrics/cadvisor",
		RelabelConfigs: []*v1beta1.RelabelConfig{
			{
				Action: "labelmap",
				Regex:  "__meta_kubernetes_node_label_(.+)",
			}, {
				SourceLabels:           []string{"__metrics_path__"},
				TargetLabel:            "metrics_path",
				UnderScoreSourceLabels: []string{"__metrics_path__"},
				UnderScoreTargetLabel:  "metrics_path",
			}, {
				Replacement:           "kubelet",
				TargetLabel:           "job",
				UnderScoreTargetLabel: "job",
			},
		},
		Scheme:        "https",
		ScrapeTimeout: "5s",
		TLSConfig: &v1beta1.TLSConfig{
			CAFile:             "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
			InsecureSkipVerify: true,
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "operator.victoriametrics.com/v1beta1",
		Kind:       "VMNodeScrape",
	},
}

var KubeletNodeScrape = &v1beta1.VMNodeScrape{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "vmk8s",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "victoria-metrics-k8s-stack",
			"app.kubernetes.io/version":    "v1.91.2",
			"helm.sh/chart":                "victoria-metrics-k8s-stack-0.16.3",
		},
		Name:      "vmk8s-victoria-metrics-k8s-stack-kubelet",
		Namespace: "monitoring",
	},
	Spec: v1beta1.VMNodeScrapeSpec{
		BearerTokenFile: "/var/run/secrets/kubernetes.io/serviceaccount/token",
		HonorLabels:     true,
		Interval:        "30s",
		MetricRelabelConfigs: []*v1beta1.RelabelConfig{
			{
				Action: "labeldrop",
				Regex:  "(uid)",
			}, {
				Action: "labeldrop",
				Regex:  "(id|name)",
			}, {
				Action:                 "drop",
				Regex:                  "(rest_client_request_duration_seconds_bucket|rest_client_request_duration_seconds_sum|rest_client_request_duration_seconds_count)",
				SourceLabels:           []string{"__name__"},
				UnderScoreSourceLabels: []string{"__name__"},
			},
		},
		RelabelConfigs: []*v1beta1.RelabelConfig{
			{
				Action: "labelmap",
				Regex:  "__meta_kubernetes_node_label_(.+)",
			}, {
				SourceLabels:           []string{"__metrics_path__"},
				TargetLabel:            "metrics_path",
				UnderScoreSourceLabels: []string{"__metrics_path__"},
				UnderScoreTargetLabel:  "metrics_path",
			}, {
				Replacement:           "kubelet",
				TargetLabel:           "job",
				UnderScoreTargetLabel: "job",
			},
		},
		Scheme:        "https",
		ScrapeTimeout: "5s",
		TLSConfig: &v1beta1.TLSConfig{
			CAFile:             "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
			InsecureSkipVerify: true,
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "operator.victoriametrics.com/v1beta1",
		Kind:       "VMNodeScrape",
	},
}

var ProbesNodeScrape = &v1beta1.VMNodeScrape{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "vmk8s",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "victoria-metrics-k8s-stack",
			"app.kubernetes.io/version":    "v1.91.2",
			"helm.sh/chart":                "victoria-metrics-k8s-stack-0.16.3",
		},
		Name:      "vmk8s-victoria-metrics-k8s-stack-probes",
		Namespace: "monitoring",
	},
	Spec: v1beta1.VMNodeScrapeSpec{
		BearerTokenFile: "/var/run/secrets/kubernetes.io/serviceaccount/token",
		HonorLabels:     true,
		Interval:        "30s",
		MetricRelabelConfigs: []*v1beta1.RelabelConfig{
			{
				Action: "labeldrop",
				Regex:  "(uid)",
			}, {
				Action: "labeldrop",
				Regex:  "(id|name)",
			}, {
				Action:                 "drop",
				Regex:                  "(rest_client_request_duration_seconds_bucket|rest_client_request_duration_seconds_sum|rest_client_request_duration_seconds_count)",
				SourceLabels:           []string{"__name__"},
				UnderScoreSourceLabels: []string{"__name__"},
			},
		},
		Path: "/metrics/probes",
		RelabelConfigs: []*v1beta1.RelabelConfig{
			{
				Action: "labelmap",
				Regex:  "__meta_kubernetes_node_label_(.+)",
			}, {
				SourceLabels:           []string{"__metrics_path__"},
				TargetLabel:            "metrics_path",
				UnderScoreSourceLabels: []string{"__metrics_path__"},
				UnderScoreTargetLabel:  "metrics_path",
			}, {
				Replacement:           "kubelet",
				TargetLabel:           "job",
				UnderScoreTargetLabel: "job",
			},
		},
		Scheme:        "https",
		ScrapeTimeout: "5s",
		TLSConfig: &v1beta1.TLSConfig{
			CAFile:             "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
			InsecureSkipVerify: true,
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "operator.victoriametrics.com/v1beta1",
		Kind:       "VMNodeScrape",
	},
}