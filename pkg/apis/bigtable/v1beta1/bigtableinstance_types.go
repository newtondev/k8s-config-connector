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

type InstanceCluster struct {
	/* The ID of the Cloud Bigtable cluster. */
	ClusterId string `json:"clusterId,omitempty"`
	/* The number of nodes in your Cloud Bigtable cluster. For PRODUCTION instances where the numNodes will be managed by Config Connector, this field is required with a minimum of 1. For a DEVELOPMENT instance or for an existing instance where the numNodes is managed outside of Config Connector, this field must be left unset. */
	NumNodes int `json:"numNodes,omitempty"`
	/* The storage type to use. One of "SSD" or "HDD". Defaults to "SSD". */
	StorageType string `json:"storageType,omitempty"`
	/* The zone to create the Cloud Bigtable cluster in. Each cluster must have a different zone in the same region. Zones that support Bigtable instances are noted on the Cloud Bigtable locations page. */
	Zone string `json:"zone,omitempty"`
}

type BigtableInstanceSpec struct {
	/* A block of cluster configuration options. This can be specified at least once. */
	Cluster []InstanceCluster `json:"cluster,omitempty"`
	/*  */
	DeletionProtection bool `json:"deletionProtection,omitempty"`
	/* The human-readable display name of the Bigtable instance. Defaults to the instance name. */
	DisplayName string `json:"displayName,omitempty"`
	/* DEPRECATED — It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away, and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions. The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION". */
	InstanceType string `json:"instanceType,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
}

type BigtableInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   BigtableInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigtableInstance is the Schema for the bigtable API
// +k8s:openapi-gen=true
type BigtableInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableInstanceSpec   `json:"spec,omitempty"`
	Status BigtableInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigtableInstanceList contains a list of BigtableInstance
type BigtableInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableInstance{}, &BigtableInstanceList{})
}
