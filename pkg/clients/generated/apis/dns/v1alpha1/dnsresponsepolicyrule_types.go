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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResponsepolicyruleLocalData struct {
	/* All resource record sets for this selector, one per resource record type. The name must match the dns_name. */
	LocalDatas []ResponsepolicyruleLocalDatas `json:"localDatas"`
}

type ResponsepolicyruleLocalDatas struct {
	/* For example, www.example.com. */
	Name string `json:"name"`

	/* As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1). */
	// +optional
	Rrdatas []string `json:"rrdatas,omitempty"`

	/* Number of seconds that this ResourceRecordSet can be cached by
	resolvers. */
	// +optional
	Ttl *int64 `json:"ttl,omitempty"`

	/* One of valid DNS resource types. Possible values: ["A", "AAAA", "CAA", "CNAME", "DNSKEY", "DS", "HTTPS", "IPSECVPNKEY", "MX", "NAPTR", "NS", "PTR", "SOA", "SPF", "SRV", "SSHFP", "SVCB", "TLSA", "TXT"]. */
	Type string `json:"type"`
}

type DNSResponsePolicyRuleSpec struct {
	/* Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'. */
	// +optional
	Behavior *string `json:"behavior,omitempty"`

	/* The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule. */
	DnsName string `json:"dnsName"`

	/* Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
	in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed. */
	// +optional
	LocalData *ResponsepolicyruleLocalData `json:"localData,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The ruleName of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Identifies the response policy addressed by this request. */
	ResponsePolicy string `json:"responsePolicy"`
}

type DNSResponsePolicyRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   DNSResponsePolicyRule's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdnsresponsepolicyrule;gcpdnsresponsepolicyrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DNSResponsePolicyRule is the Schema for the dns API
// +k8s:openapi-gen=true
type DNSResponsePolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSResponsePolicyRuleSpec   `json:"spec,omitempty"`
	Status DNSResponsePolicyRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSResponsePolicyRuleList contains a list of DNSResponsePolicyRule
type DNSResponsePolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSResponsePolicyRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSResponsePolicyRule{}, &DNSResponsePolicyRuleList{})
}
