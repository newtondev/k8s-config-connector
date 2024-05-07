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

type TopicMessageStoragePolicy struct {
	/* A list of IDs of GCP regions where messages that are published to
	the topic may be persisted in storage. Messages published by
	publishers running in non-allowed GCP regions (or running outside
	of GCP altogether) will be routed for storage in one of the
	allowed regions. An empty list means that no regions are allowed,
	and is not a valid configuration. */
	AllowedPersistenceRegions []string `json:"allowedPersistenceRegions"`
}

type TopicSchemaSettings struct {
	/* The encoding of messages validated against schema. Default value: "ENCODING_UNSPECIFIED" Possible values: ["ENCODING_UNSPECIFIED", "JSON", "BINARY"]. */
	// +optional
	Encoding *string `json:"encoding,omitempty"`

	SchemaRef v1alpha1.ResourceRef `json:"schemaRef"`
}

type PubSubTopicSpec struct {
	/* The KMSCryptoKey to be used to protect access to messages published
	on this topic. Your project's Pub/Sub service account
	('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Indicates the minimum duration to retain a message after it is published
	to the topic. If this field is set, messages published to the topic in
	the last messageRetentionDuration are always available to subscribers.
	For instance, it allows any attached subscription to seek to a timestamp
	that is up to messageRetentionDuration in the past. If this field is not
	set, message retention is controlled by settings on individual subscriptions.
	Cannot be more than 31 days or less than 10 minutes. */
	// +optional
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty"`

	/* Policy constraining the set of Google Cloud Platform regions where
	messages published to the topic may be stored. If not present, then no
	constraints are in effect. */
	// +optional
	MessageStoragePolicy *TopicMessageStoragePolicy `json:"messageStoragePolicy,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Settings for validating messages published against a schema. */
	// +optional
	SchemaSettings *TopicSchemaSettings `json:"schemaSettings,omitempty"`
}

type PubSubTopicStatus struct {
	/* Conditions represent the latest available observations of the
	   PubSubTopic's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsubtopic;gcppubsubtopics
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubTopic is the Schema for the pubsub API
// +k8s:openapi-gen=true
type PubSubTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubTopicSpec   `json:"spec,omitempty"`
	Status PubSubTopicStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PubSubTopicList contains a list of PubSubTopic
type PubSubTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubTopic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubTopic{}, &PubSubTopicList{})
}
