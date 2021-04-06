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

type IAPBrandSpec struct {
	/* Application name displayed on OAuth consent screen. */
	ApplicationTitle string `json:"applicationTitle,omitempty"`
	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	ResourceID string `json:"resourceID,omitempty"`
	/* Support email displayed on the OAuth consent screen. */
	SupportEmail string `json:"supportEmail,omitempty"`
}

type IAPBrandStatus struct {
	/* Conditions represent the latest available observations of the
	   IAPBrand's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. Whether the brand is only intended for usage inside the G Suite organization only. */
	OrgInternalOnly bool `json:"orgInternalOnly,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAPBrand is the Schema for the iap API
// +k8s:openapi-gen=true
type IAPBrand struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAPBrandSpec   `json:"spec,omitempty"`
	Status IAPBrandStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAPBrandList contains a list of IAPBrand
type IAPBrandList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAPBrand `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAPBrand{}, &IAPBrandList{})
}
