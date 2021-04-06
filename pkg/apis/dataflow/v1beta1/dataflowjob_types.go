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

type JobParameters struct {
}

type JobTransformNameMapping struct {
}

type DataflowJobSpec struct {
	/* List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"]. */
	AdditionalExperiments []string `json:"additionalExperiments,omitempty"`
	/* Indicates if the job should use the streaming engine feature. */
	EnableStreamingEngine bool `json:"enableStreamingEngine,omitempty"`
	/* The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE". */
	IpConfiguration string `json:"ipConfiguration,omitempty"`
	/* The name for the Cloud KMS key for the job. */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`
	/* The machine type to use for the job. */
	MachineType string `json:"machineType,omitempty"`
	/* Immutable. The number of workers permitted to work on the job. More workers may improve processing speed at additional cost. */
	MaxWorkers int `json:"maxWorkers,omitempty"`
	/*  */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef,omitempty"`
	/* Key/Value pairs to be passed to the Dataflow job (as used in the template). */
	Parameters JobParameters `json:"parameters,omitempty"`
	/* Immutable. The region in which the created job should run. */
	Region string `json:"region,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/*  */
	ServiceAccountRef v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
	/*  */
	SubnetworkRef v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
	/* A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data. */
	TempGcsLocation string `json:"tempGcsLocation,omitempty"`
	/* The Google Cloud Storage path to the Dataflow job template. */
	TemplateGcsPath string `json:"templateGcsPath,omitempty"`
	/* Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. */
	TransformNameMapping JobTransformNameMapping `json:"transformNameMapping,omitempty"`
	/* Immutable. The zone in which the created job should run. If it is not provided, the provider zone is used. */
	Zone string `json:"zone,omitempty"`
}

type DataflowJobStatus struct {
	/* Conditions represent the latest available observations of the
	   DataflowJob's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The unique ID of this job. */
	JobId string `json:"jobId,omitempty"`
	/* The current state of the resource, selected from the JobState enum. */
	State string `json:"state,omitempty"`
	/* The type of this job, selected from the JobType enum. */
	Type string `json:"type,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataflowJob is the Schema for the dataflow API
// +k8s:openapi-gen=true
type DataflowJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataflowJobSpec   `json:"spec,omitempty"`
	Status DataflowJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataflowJobList contains a list of DataflowJob
type DataflowJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataflowJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataflowJob{}, &DataflowJobList{})
}
