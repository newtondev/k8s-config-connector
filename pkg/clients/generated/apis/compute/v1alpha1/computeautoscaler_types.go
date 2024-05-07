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

type AutoscalerAutoscalingPolicy struct {
	/* The number of seconds that the autoscaler should wait before it
	starts collecting information from a new instance. This prevents
	the autoscaler from collecting information when the instance is
	initializing, during which the collected usage would not be
	reliable. The default time autoscaler waits is 60 seconds.

	Virtual machine initialization times might vary because of
	numerous factors. We recommend that you test how long an
	instance may take to initialize. To do this, create an instance
	and time the startup process. */
	// +optional
	CooldownPeriod *int64 `json:"cooldownPeriod,omitempty"`

	/* Defines the CPU utilization policy that allows the autoscaler to
	scale based on the average CPU utilization of a managed instance
	group. */
	// +optional
	CpuUtilization *AutoscalerCpuUtilization `json:"cpuUtilization,omitempty"`

	/* Configuration parameters of autoscaling based on a load balancer. */
	// +optional
	LoadBalancingUtilization *AutoscalerLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty"`

	/* The maximum number of instances that the autoscaler can scale up
	to. This is required when creating or updating an autoscaler. The
	maximum number of replicas should not be lower than minimal number
	of replicas. */
	MaxReplicas int64 `json:"maxReplicas"`

	/* Configuration parameters of autoscaling based on a custom metric. */
	// +optional
	Metric []AutoscalerMetric `json:"metric,omitempty"`

	/* The minimum number of replicas that the autoscaler can scale down
	to. This cannot be less than 0. If not provided, autoscaler will
	choose a default value depending on maximum number of instances
	allowed. */
	MinReplicas int64 `json:"minReplicas"`

	/* Defines operating mode for this policy. */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/* Defines scale down controls to reduce the risk of response latency
	and outages due to abrupt scale-in events. */
	// +optional
	ScaleDownControl *AutoscalerScaleDownControl `json:"scaleDownControl,omitempty"`

	/* Defines scale in controls to reduce the risk of response latency
	and outages due to abrupt scale-in events. */
	// +optional
	ScaleInControl *AutoscalerScaleInControl `json:"scaleInControl,omitempty"`

	/* Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler and they can overlap. */
	// +optional
	ScalingSchedules []AutoscalerScalingSchedules `json:"scalingSchedules,omitempty"`
}

type AutoscalerCpuUtilization struct {
	/* Indicates whether predictive autoscaling based on CPU metric is enabled. Valid values are:

	- NONE (default). No predictive method is used. The autoscaler scales the group to meet current demand based on real-time metrics.

	- OPTIMIZE_AVAILABILITY. Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand. */
	// +optional
	PredictiveMethod *string `json:"predictiveMethod,omitempty"`

	/* The target CPU utilization that the autoscaler should maintain.
	Must be a float value in the range (0, 1]. If not specified, the
	default is 0.6.

	If the CPU level is below the target utilization, the autoscaler
	scales down the number of instances until it reaches the minimum
	number of instances you specified or until the average CPU of
	your instances reaches the target utilization.

	If the average CPU is above the target utilization, the autoscaler
	scales up until it reaches the maximum number of instances you
	specified or until the average utilization reaches the target
	utilization. */
	Target float64 `json:"target"`
}

type AutoscalerLoadBalancingUtilization struct {
	/* Fraction of backend capacity utilization (set in HTTP(s) load
	balancing configuration) that autoscaler should maintain. Must
	be a positive float value. If not defined, the default is 0.8. */
	Target float64 `json:"target"`
}

type AutoscalerMaxScaledDownReplicas struct {
	/* Specifies a fixed number of VM instances. This must be a positive
	integer. */
	// +optional
	Fixed *int64 `json:"fixed,omitempty"`

	/* Specifies a percentage of instances between 0 to 100%, inclusive.
	For example, specify 80 for 80%. */
	// +optional
	Percent *int64 `json:"percent,omitempty"`
}

type AutoscalerMaxScaledInReplicas struct {
	/* Specifies a fixed number of VM instances. This must be a positive
	integer. */
	// +optional
	Fixed *int64 `json:"fixed,omitempty"`

	/* Specifies a percentage of instances between 0 to 100%, inclusive.
	For example, specify 80 for 80%. */
	// +optional
	Percent *int64 `json:"percent,omitempty"`
}

type AutoscalerMetric struct {
	/* A filter string to be used as the filter string for
	a Stackdriver Monitoring TimeSeries.list API call.
	This filter is used to select a specific TimeSeries for
	the purpose of autoscaling and to determine whether the metric
	is exporting per-instance or per-group data.

	You can only use the AND operator for joining selectors.
	You can only use direct equality comparison operator (=) without
	any functions for each selector.
	You can specify the metric in both the filter string and in the
	metric field. However, if specified in both places, the metric must
	be identical.

	The monitored resource type determines what kind of values are
	expected for the metric. If it is a gce_instance, the autoscaler
	expects the metric to include a separate TimeSeries for each
	instance in a group. In such a case, you cannot filter on resource
	labels.

	If the resource type is any other value, the autoscaler expects
	this metric to contain values that apply to the entire autoscaled
	instance group and resource label filtering can be performed to
	point autoscaler at the correct TimeSeries to scale upon.
	This is called a per-group metric for the purpose of autoscaling.

	If not specified, the type defaults to gce_instance.

	You should provide a filter that is selective enough to pick just
	one TimeSeries for the autoscaled group or for each of the instances
	(if you are using gce_instance resource type). If multiple
	TimeSeries are returned upon the query execution, the autoscaler
	will sum their respective values to obtain its scaling value. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* The identifier (type) of the Stackdriver Monitoring metric.
	The metric cannot have negative values.

	The metric must have a value type of INT64 or DOUBLE. */
	Name string `json:"name"`

	/* If scaling is based on a per-group metric value that represents the
	total amount of work to be done or resource usage, set this value to
	an amount assigned for a single instance of the scaled group.
	The autoscaler will keep the number of instances proportional to the
	value of this metric, the metric itself should not change value due
	to group resizing.

	For example, a good metric to use with the target is
	'pubsub.googleapis.com/subscription/num_undelivered_messages'
	or a custom metric exporting the total number of requests coming to
	your instances.

	A bad example would be a metric exporting an average or median
	latency, since this value can't include a chunk assignable to a
	single instance, it could be better used with utilization_target
	instead. */
	// +optional
	SingleInstanceAssignment *float64 `json:"singleInstanceAssignment,omitempty"`

	/* The target value of the metric that autoscaler should
	maintain. This must be a positive value. A utilization
	metric scales number of virtual machines handling requests
	to increase or decrease proportionally to the metric.

	For example, a good metric to use as a utilizationTarget is
	www.googleapis.com/compute/instance/network/received_bytes_count.
	The autoscaler will work to keep this value constant for each
	of the instances. */
	// +optional
	Target *float64 `json:"target,omitempty"`

	/* Defines how target utilization value is expressed for a
	Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"]. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type AutoscalerScaleDownControl struct {
	/* A nested object resource. */
	// +optional
	MaxScaledDownReplicas *AutoscalerMaxScaledDownReplicas `json:"maxScaledDownReplicas,omitempty"`

	/* How long back autoscaling should look when computing recommendations
	to include directives regarding slower scale down, as described above. */
	// +optional
	TimeWindowSec *int64 `json:"timeWindowSec,omitempty"`
}

type AutoscalerScaleInControl struct {
	/* A nested object resource. */
	// +optional
	MaxScaledInReplicas *AutoscalerMaxScaledInReplicas `json:"maxScaledInReplicas,omitempty"`

	/* How long back autoscaling should look when computing recommendations
	to include directives regarding slower scale down, as described above. */
	// +optional
	TimeWindowSec *int64 `json:"timeWindowSec,omitempty"`
}

type AutoscalerScalingSchedules struct {
	/* A description of a scaling schedule. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* A boolean value that specifies if a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* The duration of time intervals (in seconds) for which this scaling schedule will be running. The minimum allowed value is 300. */
	DurationSec int64 `json:"durationSec"`

	/* Minimum number of VM instances that autoscaler will recommend in time intervals starting according to schedule. */
	MinRequiredReplicas int64 `json:"minRequiredReplicas"`

	Name string `json:"name"`

	/* The start timestamps of time intervals when this scaling schedule should provide a scaling signal. This field uses the extended cron format (with an optional year field). */
	Schedule string `json:"schedule"`

	/* The time zone to be used when interpreting the schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database. */
	// +optional
	TimeZone *string `json:"timeZone,omitempty"`
}

type ComputeAutoscalerSpec struct {
	/* The configuration parameters for the autoscaling algorithm. You can
	define one or more of the policies for an autoscaler: cpuUtilization,
	customMetricUtilizations, and loadBalancingUtilization.

	If none of these are specified, the default will be to autoscale based
	on cpuUtilization to 0.6 or 60%. */
	AutoscalingPolicy AutoscalerAutoscalingPolicy `json:"autoscalingPolicy"`

	/* An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	TargetRef v1alpha1.ResourceRef `json:"targetRef"`

	/* Immutable. URL of the zone where the instance group resides. */
	Zone string `json:"zone"`
}

type ComputeAutoscalerStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeAutoscaler's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeautoscaler;gcpcomputeautoscalers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeAutoscaler is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeAutoscalerSpec   `json:"spec,omitempty"`
	Status ComputeAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeAutoscalerList contains a list of ComputeAutoscaler
type ComputeAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeAutoscaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeAutoscaler{}, &ComputeAutoscalerList{})
}
