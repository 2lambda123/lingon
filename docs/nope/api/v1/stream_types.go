/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StreamSpec defines the desired state of Stream
type StreamSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Account is the name of the account to use for this stream
	Account string `json:"account,omitempty"`
	// Name is the name of the stream
	Name string `json:"name,omitempty"`
	// Subjects of the stream
	Subjects []string `json:"subjects,omitempty"`
}

// StreamStatus defines the observed state of Stream
type StreamStatus struct {
	Name string `json:"name,omitempty"`
	// Subjects []string `json:"subjects,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Stream is the Schema for the streams API
type Stream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StreamSpec   `json:"spec,omitempty"`
	Status StreamStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StreamList contains a list of Stream
type StreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stream{}, &StreamList{})
}