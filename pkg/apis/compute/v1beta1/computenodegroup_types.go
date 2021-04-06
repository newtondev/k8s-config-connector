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

type NodegroupAutoscalingPolicy struct {
	/* Immutable. Maximum size of the node group. Set to a value less than or equal
	to 100 and greater than or equal to min-nodes. */
	MaxNodes int `json:"maxNodes,omitempty"`
	/* Immutable. Minimum size of the node group. Must be less
	than or equal to max-nodes. The default value is 0. */
	MinNodes int `json:"minNodes,omitempty"`
	/* Immutable. The autoscaling mode. Set to one of the following:
	- OFF: Disables the autoscaler.
	- ON: Enables scaling in and scaling out.
	- ONLY_SCALE_OUT: Enables only scaling out.
	You must use this mode if your node groups are configured to
	restart their hosted VMs on minimal servers. Possible values: ["OFF", "ON", "ONLY_SCALE_OUT"] */
	Mode string `json:"mode,omitempty"`
}

type ComputeNodeGroupSpec struct {
	/* Immutable. If you use sole-tenant nodes for your workloads, you can use the node
	group autoscaler to automatically manage the sizes of your node groups. */
	AutoscalingPolicy NodegroupAutoscalingPolicy `json:"autoscalingPolicy,omitempty"`
	/* Immutable. An optional textual description of the resource. */
	Description string `json:"description,omitempty"`
	/* Immutable. Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT. */
	MaintenancePolicy string `json:"maintenancePolicy,omitempty"`
	/* The node template to which this node group belongs. */
	NodeTemplateRef v1alpha1.ResourceRef `json:"nodeTemplateRef,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* Immutable. The total number of nodes in the node group. */
	Size int `json:"size,omitempty"`
	/* Immutable. Zone where this node group is located */
	Zone string `json:"zone,omitempty"`
}

type ComputeNodeGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNodeGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNodeGroup is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNodeGroupSpec   `json:"spec,omitempty"`
	Status ComputeNodeGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeNodeGroupList contains a list of ComputeNodeGroup
type ComputeNodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNodeGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNodeGroup{}, &ComputeNodeGroupList{})
}
