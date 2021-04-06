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

type ComputeRouterInterfaceSpec struct {
	/*  */
	InterconnectAttachmentRef v1alpha1.ResourceRef `json:"interconnectAttachmentRef,omitempty"`
	/* Immutable. IP address and range of the interface. The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created. */
	IpRange string `json:"ipRange,omitempty"`
	/* Immutable. The region this interface's router sits in. If not specified, the project region will be used. Changing this forces a new interface to be created. */
	Region string `json:"region,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/*  */
	RouterRef v1alpha1.ResourceRef `json:"routerRef,omitempty"`
	/*  */
	VpnTunnelRef v1alpha1.ResourceRef `json:"vpnTunnelRef,omitempty"`
}

type ComputeRouterInterfaceStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRouterInterface's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterInterface is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterInterfaceSpec   `json:"spec,omitempty"`
	Status ComputeRouterInterfaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterInterfaceList contains a list of ComputeRouterInterface
type ComputeRouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouterInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouterInterface{}, &ComputeRouterInterfaceList{})
}
