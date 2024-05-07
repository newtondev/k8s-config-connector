// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeTargetSSLProxySpec struct {
	/* A reference to the ComputeBackendService resource. */
	BackendServiceRef v1alpha1.ResourceRef `json:"backendServiceRef"`

	/* A reference to the CertificateMap resource uri that identifies a
	certificate map associated with the given target proxy. This
	field can only be set for global target proxies. Accepted format is
	'//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'. */
	// +optional
	CertificateMapRef *v1alpha1.ResourceRef `json:"certificateMapRef,omitempty"`

	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Specifies the type of proxy header to append before sending data to
	the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]. */
	// +optional
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	SslCertificates []v1alpha1.ResourceRef `json:"sslCertificates,omitempty"`

	/* A reference to the ComputeSSLPolicy resource that will be
	associated with the TargetSslProxy resource. If not set, the
	ComputeTargetSSLProxy resource will not have any SSL policy
	configured. */
	// +optional
	SslPolicyRef *v1alpha1.ResourceRef `json:"sslPolicyRef,omitempty"`
}

type ComputeTargetSSLProxyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeTargetSSLProxy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The unique identifier for the resource. */
	// +optional
	ProxyId *int64 `json:"proxyId,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetsslproxy;gcpcomputetargetsslproxies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetSSLProxy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeTargetSSLProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetSSLProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetSSLProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeTargetSSLProxyList contains a list of ComputeTargetSSLProxy
type ComputeTargetSSLProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetSSLProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetSSLProxy{}, &ComputeTargetSSLProxyList{})
}
