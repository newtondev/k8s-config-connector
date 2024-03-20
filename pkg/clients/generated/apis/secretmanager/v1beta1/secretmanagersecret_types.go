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

type SecretAuto struct {
	/* The customer-managed encryption configuration of the Secret.
	If no configuration is provided, Google-managed default
	encryption is used. */
	// +optional
	CustomerManagedEncryption *SecretCustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`
}

type SecretCustomerManagedEncryption struct {
	/* Customer Managed Encryption for the secret. */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef"`
}

type SecretReplicas struct {
	/* Customer Managed Encryption for the secret. */
	// +optional
	CustomerManagedEncryption *SecretCustomerManagedEncryption `json:"customerManagedEncryption,omitempty"`

	/* Immutable. The canonical IDs of the location to replicate data. For example: "us-east1". */
	Location string `json:"location"`
}

type SecretReplication struct {
	/* The Secret will automatically be replicated without any restrictions. */
	// +optional
	Auto *SecretAuto `json:"auto,omitempty"`

	/* The Secret will automatically be replicated without any restrictions. */
	// +optional
	Automatic *bool `json:"automatic,omitempty"`

	/* Immutable. The Secret will be replicated to the regions specified by the user. */
	// +optional
	UserManaged *SecretUserManaged `json:"userManaged,omitempty"`
}

type SecretRotation struct {
	/* Timestamp in UTC at which the Secret is scheduled to rotate.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	NextRotationTime *string `json:"nextRotationTime,omitempty"`

	/* Immutable. The Duration between rotation notifications. Must be in seconds and at least 3600s (1h) and at most 3153600000s (100 years).
	If rotationPeriod is set, 'next_rotation_time' must be set. 'next_rotation_time' will be advanced by this period when the service automatically sends rotation notifications. */
	// +optional
	RotationPeriod *string `json:"rotationPeriod,omitempty"`
}

type SecretTopics struct {
	/* A list of up to 10 Pub/Sub topics to which messages are
	published when control plane operations are called on the secret
	or its versions. */
	TopicRef v1alpha1.ResourceRef `json:"topicRef"`
}

type SecretUserManaged struct {
	/* Immutable. The list of Replicas for this Secret. Cannot be empty. */
	Replicas []SecretReplicas `json:"replicas"`
}

type SecretManagerSecretSpec struct {
	/* Custom metadata about the secret.

	Annotations are distinct from various forms of labels. Annotations exist to allow
	client tools to store their own state information without requiring a database.

	Annotation keys must be between 1 and 63 characters long, have a UTF-8 encoding of
	maximum 128 bytes, begin and end with an alphanumeric character ([a-z0-9A-Z]), and
	may have dashes (-), underscores (_), dots (.), and alphanumerics in between these
	symbols.

	The total size of annotation keys and values must be less than 16KiB.

	An object containing a list of "key": value pairs. Example:
	{ "name": "wrench", "mass": "1.3kg", "count": "3" }. */
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	/* Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* Immutable. The replication policy of the secret data attached to the Secret. It cannot be changed
	after the Secret has been created. */
	Replication SecretReplication `json:"replication"`

	/* Immutable. Optional. The secretId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The rotation time and period for a Secret. At 'next_rotation_time', Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be set to configure rotation. */
	// +optional
	Rotation *SecretRotation `json:"rotation,omitempty"`

	/* A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions. */
	// +optional
	Topics []SecretTopics `json:"topics,omitempty"`

	/* Immutable. The TTL for the Secret.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	Ttl *string `json:"ttl,omitempty"`

	/* Mapping from version alias to version name.

	A version alias is a string with a maximum length of 63 characters and can contain
	uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_')
	characters. An alias string must start with a letter and cannot be the string
	'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret.

	An object containing a list of "key": value pairs. Example:
	{ "name": "wrench", "mass": "1.3kg", "count": "3" }. */
	// +optional
	VersionAliases map[string]string `json:"versionAliases,omitempty"`
}

type SecretManagerSecretStatus struct {
	/* Conditions represent the latest available observations of the
	   SecretManagerSecret's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time at which the Secret was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The resource name of the Secret. Format:
	'projects/{{project}}/secrets/{{secret_id}}'. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretManagerSecret is the Schema for the secretmanager API
// +k8s:openapi-gen=true
type SecretManagerSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretManagerSecretSpec   `json:"spec,omitempty"`
	Status SecretManagerSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretManagerSecretList contains a list of SecretManagerSecret
type SecretManagerSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretManagerSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretManagerSecret{}, &SecretManagerSecretList{})
}
