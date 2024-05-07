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

type AuthorizationpolicyDestinations struct {
	/* Required. List of host names to match. Matched against HOST header in http requests. Each host can be an exact match, or a prefix match (example, “mydomain.*”) or a suffix match (example, *.myorg.com”) or a presence(any) match “*”. */
	Hosts []string `json:"hosts"`

	/* Optional. Match against key:value pair in http header. Provides a flexible match based on HTTP headers, for potentially advanced use cases. */
	// +optional
	HttpHeaderMatch *AuthorizationpolicyHttpHeaderMatch `json:"httpHeaderMatch,omitempty"`

	/* Optional. A list of HTTP methods to match. Should not be set for gRPC services. */
	// +optional
	Methods []string `json:"methods,omitempty"`

	/* Required. List of destination ports to match. */
	Ports []int64 `json:"ports"`
}

type AuthorizationpolicyHttpHeaderMatch struct {
	/* Required. The name of the HTTP header to match. For matching against the HTTP request's authority, use a headerMatch with the header name ":authority". For matching a request's method, use the headerName ":method". */
	HeaderName string `json:"headerName"`

	/* Required. The value of the header must match the regular expression specified in regexMatch. For regular expression grammar, please see: en.cppreference.com/w/cpp/regex/ecmascript For matching against a port specified in the HTTP request, use a headerMatch with headerName set to Host and a regular expression that satisfies the RFC2616 Host header's port specifier. */
	RegexMatch string `json:"regexMatch"`
}

type AuthorizationpolicyRules struct {
	/* Optional. List of attributes for the traffic destination. If not set, the action specified in the ‘action’ field will be applied without any rule checks for the destination. */
	// +optional
	Destinations []AuthorizationpolicyDestinations `json:"destinations,omitempty"`

	/* Optional. List of attributes for the traffic source. If not set, the action specified in the ‘action’ field will be applied without any rule checks for the source. */
	// +optional
	Sources []AuthorizationpolicySources `json:"sources,omitempty"`
}

type AuthorizationpolicySources struct {
	/* Optional. List of CIDR ranges to match based on source IP address. Single IP (e.g., "1.2.3.4") and CIDR (e.g., "1.2.3.0/24") are supported. */
	// +optional
	IpBlocks []string `json:"ipBlocks,omitempty"`

	/* Optional. List of peer identities to match for authorization. Each peer can be an exact match, or a prefix match (example, “namespace/*”) or a suffix match (example, * /service-account”) or a presence match “*”. */
	// +optional
	Principals []string `json:"principals,omitempty"`
}

type NetworkSecurityAuthorizationPolicySpec struct {
	/* Required. The action to take when a rule match is found. Possible values are "ALLOW" or "DENY". Possible values: ACTION_UNSPECIFIED, ALLOW, DENY */
	Action string `json:"action"`

	/* Optional. Free-text description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional. List of rules to match. If not set, the action specified in the ‘action’ field will be applied without any additional rule checks. */
	// +optional
	Rules []AuthorizationpolicyRules `json:"rules,omitempty"`
}

type NetworkSecurityAuthorizationPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkSecurityAuthorizationPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The timestamp when the resource was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The timestamp when the resource was updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityauthorizationpolicy;gcpnetworksecurityauthorizationpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityAuthorizationPolicy is the Schema for the networksecurity API
// +k8s:openapi-gen=true
type NetworkSecurityAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSecurityAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityAuthorizationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkSecurityAuthorizationPolicyList contains a list of NetworkSecurityAuthorizationPolicy
type NetworkSecurityAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityAuthorizationPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityAuthorizationPolicy{}, &NetworkSecurityAuthorizationPolicyList{})
}
