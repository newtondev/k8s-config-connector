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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeInterconnectAttachmentSpec struct {
	/* Whether the VLAN attachment is enabled or disabled.  When using
	PARTNER type this will Pre-Activate the interconnect attachment */
	AdminEnabled bool `json:"adminEnabled,omitempty"`
	/* Provisioned bandwidth capacity for the interconnect attachment.
	For attachments of type DEDICATED, the user can set the bandwidth.
	For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	Defaults to BPS_10G Possible values: ["BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G"] */
	Bandwidth string `json:"bandwidth,omitempty"`
	/* Immutable. Up to 16 candidate prefixes that can be used to restrict the allocation
	of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	All prefixes must be within link-local address space (169.254.0.0/16)
	and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	an unused /29 from the supplied candidate prefix(es). The request will
	fail if all possible /29s are in use on Google's edge. If not supplied,
	Google will randomly select an unused /29 from all of link-local space. */
	CandidateSubnets []string `json:"candidateSubnets,omitempty"`
	/* An optional description of this resource. */
	Description string `json:"description,omitempty"`
	/* Immutable. Desired availability domain for the attachment. Only available for type
	PARTNER, at creation time. For improved reliability, customers should
	configure a pair of attachments with one per availability domain. The
	selected availability domain will be provided to the Partner via the
	pairing key so that the provisioned circuit will lie in the specified
	domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY. */
	EdgeAvailabilityDomain string `json:"edgeAvailabilityDomain,omitempty"`
	/* Immutable. URL of the underlying Interconnect object that this attachment's
	traffic will traverse through. Required if type is DEDICATED, must not
	be set if type is PARTNER. */
	Interconnect string `json:"interconnect,omitempty"`
	/* Maximum Transmission Unit (MTU), in bytes, of packets passing through
	this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440. */
	Mtu string `json:"mtu,omitempty"`
	/* Region where the regional interconnect attachment resides. */
	Region string `json:"region,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* The Cloud Router to be used for dynamic routing. This router must
	be in the same region as this ComputeInterconnectAttachment. The
	ComputeInterconnectAttachment will automatically connect the
	interconnect to the network & region within which the Cloud Router
	is configured. */
	RouterRef v1alpha1.ResourceRef `json:"routerRef,omitempty"`
	/* Immutable. The type of InterconnectAttachment you wish to create. Defaults to
	DEDICATED. Possible values: ["DEDICATED", "PARTNER", "PARTNER_PROVIDER"] */
	Type string `json:"type,omitempty"`
	/* Immutable. The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
	using PARTNER type this will be managed upstream. */
	VlanTag8021q int `json:"vlanTag8021q,omitempty"`
}

type InterconnectattachmentPrivateInterconnectInfoStatus struct {
	/* 802.1q encapsulation tag to be used for traffic between
	Google and the customer, going to and from this network and region. */
	Tag8021q int `json:"tag8021q,omitempty"`
}

type ComputeInterconnectAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeInterconnectAttachment's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* IPv4 address + prefix length to be configured on Cloud Router
	Interface for this interconnect attachment. */
	CloudRouterIpAddress string `json:"cloudRouterIpAddress,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* IPv4 address + prefix length to be configured on the customer
	router subinterface for this interconnect attachment. */
	CustomerRouterIpAddress string `json:"customerRouterIpAddress,omitempty"`
	/* Google reference ID, to be used when raising support tickets with
	Google or otherwise to debug backend connectivity issues. */
	GoogleReferenceId string `json:"googleReferenceId,omitempty"`
	/* [Output only for type PARTNER. Not present for DEDICATED]. The opaque
	identifier of an PARTNER attachment used to initiate provisioning with
	a selected partner. Of the form "XXXXX/region/domain" */
	PairingKey string `json:"pairingKey,omitempty"`
	/* [Output only for type PARTNER. Not present for DEDICATED]. Optional
	BGP ASN for the router that should be supplied by a layer 3 Partner if
	they configured BGP on behalf of the customer. */
	PartnerAsn string `json:"partnerAsn,omitempty"`
	/* Information specific to an InterconnectAttachment. This property
	is populated if the interconnect that this is attached to is of type DEDICATED. */
	PrivateInterconnectInfo InterconnectattachmentPrivateInterconnectInfoStatus `json:"privateInterconnectInfo,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
	/* [Output Only] The current state of this attachment's functionality. */
	State string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInterconnectAttachment is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeInterconnectAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInterconnectAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeInterconnectAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInterconnectAttachmentList contains a list of ComputeInterconnectAttachment
type ComputeInterconnectAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInterconnectAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInterconnectAttachment{}, &ComputeInterconnectAttachmentList{})
}
