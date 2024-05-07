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

type ConnectorSubnet struct {
	/* Immutable. Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is
	https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}" */
	// +optional
	NameRef *v1alpha1.ResourceRef `json:"nameRef,omitempty"`

	/* Immutable. Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`
}

type VPCAccessConnectorSpec struct {
	/* Immutable. The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'. */
	// +optional
	IpCidrRange *string `json:"ipCidrRange,omitempty"`

	/* Location represents the geographical location of the VPCAccessConnector. Specify a region name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. Machine type of VM Instance underlying connector. Default is e2-micro. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Maximum value of instances in autoscaling group underlying the connector. */
	// +optional
	MaxInstances *int64 `json:"maxInstances,omitempty"`

	/* Immutable. Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 300. */
	// +optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty"`

	/* Immutable. Minimum value of instances in autoscaling group underlying the connector. */
	// +optional
	MinInstances *int64 `json:"minInstances,omitempty"`

	/* Immutable. Minimum throughput of the connector in Mbps. Default and min is 200. */
	// +optional
	MinThroughput *int64 `json:"minThroughput,omitempty"`

	/* Immutable. Name or self_link of the VPC network. Required if 'ip_cidr_range' is set. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The subnet in which to house the connector. */
	// +optional
	Subnet *ConnectorSubnet `json:"subnet,omitempty"`
}

type VPCAccessConnectorStatus struct {
	/* Conditions represent the latest available observations of the
	   VPCAccessConnector's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* List of projects using the connector. */
	// +optional
	ConnectedProjects []string `json:"connectedProjects,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The fully qualified name of this VPC connector. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* State of the VPC access connector. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvpcaccessconnector;gcpvpcaccessconnectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VPCAccessConnector is the Schema for the vpcaccess API
// +k8s:openapi-gen=true
type VPCAccessConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPCAccessConnectorSpec   `json:"spec,omitempty"`
	Status VPCAccessConnectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VPCAccessConnectorList contains a list of VPCAccessConnector
type VPCAccessConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCAccessConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCAccessConnector{}, &VPCAccessConnectorList{})
}
