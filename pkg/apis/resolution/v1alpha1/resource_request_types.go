/*
Copyright 2022 The Tekton Authors

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceRequest is an object for requesting the content of
// a Tekton resource like a pipeline.yaml.
//
// +genclient
// +genreconciler
type ResourceRequest struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the information for the request part of the resource request.
	// +optional
	Spec ResourceRequestSpec `json:"spec,omitempty"`

	// Status communicates the state of the request and, ultimately,
	// the content of the resolved resource.
	// +optional
	Status ResourceRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceRequestList is a list of ResourceRequests.
type ResourceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata"`
	Items           []ResourceRequest `json:"items"`
}

type ResourceRequestSpec struct {
	// Parameters are the runtime attributes passed to
	// the resolver to help it figure out how to resolve the
	// resource being requested. For example: repo URL, commit SHA,
	// path to file, the kind of authentication to leverage, etc.
	// +optional
	Parameters map[string]string `json:"params,omitempty"`
}

type ResourceRequestStatus struct {
	duckv1.Status               `json:",inline"`
	ResourceRequestStatusFields `json:",inline"`
}

type ResourceRequestStatusFields struct {
	// Data is a string representation of the resolved content
	// of the requested resource in-lined into the ResourceRequest
	// object.
	Data string `json:"data"`
}

// GetStatus implements KRShaped.
func (rr *ResourceRequest) GetStatus() *duckv1.Status {
	return &rr.Status.Status
}