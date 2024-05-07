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

type JobAppEngineHttpTarget struct {
	/* App Engine Routing setting for the job. */
	// +optional
	AppEngineRouting *JobAppEngineRouting `json:"appEngineRouting,omitempty"`

	/* Body. HTTP request body. A request body is allowed only if the HTTP method is POST or PUT. It will result in invalid argument error to set a body on a job with an incompatible HttpMethod. */
	// +optional
	Body *string `json:"body,omitempty"`

	/* HTTP request headers. This map contains the header field names and values. Headers can be set when the job is created. Cloud Scheduler sets some headers to default values: * `User-Agent`: By default, this header is `"App Engine-Google; (+http://code.google.com/appengine)"`. This header can be modified, but Cloud Scheduler will append `"App Engine-Google; (+http://code.google.com/appengine)"` to the modified `User-Agent`. * `X-CloudScheduler`: This header will be set to true. The headers below are output only. They cannot be set or overridden: * `X-Google-*`: For Google internal use only. * `X-App Engine-*`: For Google internal use only. In addition, some App Engine headers, which contain job-specific information, are also be sent to the job handler. */
	// +optional
	Headers map[string]string `json:"headers,omitempty"`

	/* The HTTP method to use for the request. PATCH and OPTIONS are not permitted. Possible values: HTTP_METHOD_UNSPECIFIED, POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS */
	// +optional
	HttpMethod *string `json:"httpMethod,omitempty"`

	/* The relative URI. The relative URL must begin with "/" and must be a valid HTTP relative URL. It can contain a path, query string arguments, and `#` fragments. If the relative URL is empty, then the root path "/" will be used. No spaces are allowed, and the maximum length allowed is 2083 characters. */
	// +optional
	RelativeUri *string `json:"relativeUri,omitempty"`
}

type JobAppEngineRouting struct {
	/* App instance. By default, the job is sent to an instance which is available when the job is attempted. Requests can only be sent to a specific instance if [manual scaling is used in App Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes). App Engine Flex does not support instances. For more information, see [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed) and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed). */
	// +optional
	Instance *string `json:"instance,omitempty"`

	/* App service. By default, the job is sent to the service which is the default service when the job is attempted. */
	// +optional
	Service *string `json:"service,omitempty"`

	/* App version. By default, the job is sent to the version which is the default version when the job is attempted. */
	// +optional
	Version *string `json:"version,omitempty"`
}

type JobHttpTarget struct {
	/* HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a job with an incompatible HttpMethod. */
	// +optional
	Body *string `json:"body,omitempty"`

	/* The user can specify HTTP request headers to send with the job's HTTP request. This map contains the header field names and values. Repeated headers are not supported, but a header value can contain commas. These headers represent a subset of the headers that will accompany the job's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is below: - Host: This will be computed by Cloud Scheduler and derived from uri. * `Content-Length`: This will be computed by Cloud Scheduler. * `User-Agent`: This will be set to `"Google-Cloud-Scheduler"`. * `X-Google-*`: Google internal use only. * `X-appengine-*`: Google internal use only. The total size of headers must be less than 80KB. */
	// +optional
	Headers map[string]string `json:"headers,omitempty"`

	/* Which HTTP method to use for the request. Possible values: HTTP_METHOD_UNSPECIFIED, POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS */
	// +optional
	HttpMethod *string `json:"httpMethod,omitempty"`

	/* If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com. */
	// +optional
	OauthToken *JobOauthToken `json:"oauthToken,omitempty"`

	/* If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself. */
	// +optional
	OidcToken *JobOidcToken `json:"oidcToken,omitempty"`

	/* Required. The full URI path that the request will be sent to. This string must begin with either "http://" or "https://". Some examples of valid values for uri are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will encode some characters for safety and compatibility. The maximum allowed URL length is 2083 characters after encoding. */
	Uri string `json:"uri"`
}

type JobOauthToken struct {
	/* OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used. */
	// +optional
	Scope *string `json:"scope,omitempty"`

	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
}

type JobOidcToken struct {
	/* Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used. */
	// +optional
	Audience *string `json:"audience,omitempty"`

	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
}

type JobPubsubTarget struct {
	/* Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute. */
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`

	/* The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute. */
	// +optional
	Data *string `json:"data,omitempty"`

	TopicRef v1alpha1.ResourceRef `json:"topicRef"`
}

type JobRetryConfig struct {
	/* The maximum amount of time to wait before retrying a job after it fails. The default value of this field is 1 hour. */
	// +optional
	MaxBackoffDuration *string `json:"maxBackoffDuration,omitempty"`

	/* The time between retries will double `max_doublings` times. A job's retry interval starts at min_backoff_duration, then doubles `max_doublings` times, then increases linearly, and finally retries at intervals of max_backoff_duration up to retry_count times. For example, if min_backoff_duration is 10s, max_backoff_duration is 300s, and `max_doublings` is 3, then the a job will first be retried in 10s. The retry interval will double three times, and then increase linearly by 2^3 * 10s. Finally, the job will retry at intervals of max_backoff_duration until the job has been attempted retry_count times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, .... The default value of this field is 5. */
	// +optional
	MaxDoublings *int64 `json:"maxDoublings,omitempty"`

	/* The time limit for retrying a failed job, measured from time when an execution was first attempted. If specified with retry_count, the job will be retried until both limits are reached. The default value for max_retry_duration is zero, which means retry duration is unlimited. */
	// +optional
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty"`

	/* The minimum amount of time to wait before retrying a job after it fails. The default value of this field is 5 seconds. */
	// +optional
	MinBackoffDuration *string `json:"minBackoffDuration,omitempty"`

	/* The number of attempts that the system will make to run a job using the exponential backoff procedure described by max_doublings. The default value of retry_count is zero. If retry_count is zero, a job attempt will *not* be retried if it fails. Instead the Cloud Scheduler system will wait for the next scheduled execution time. If retry_count is set to a non-zero number then Cloud Scheduler will retry failed attempts, using exponential backoff, retry_count times, or until the next scheduled execution time, whichever comes first. Values greater than 5 and negative values are not allowed. */
	// +optional
	RetryCount *int64 `json:"retryCount,omitempty"`
}

type CloudSchedulerJobSpec struct {
	/* App Engine HTTP target. */
	// +optional
	AppEngineHttpTarget *JobAppEngineHttpTarget `json:"appEngineHttpTarget,omitempty"`

	/* The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The allowed duration for this deadline is: * For HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets, between 15 seconds and 24 hours. */
	// +optional
	AttemptDeadline *string `json:"attemptDeadline,omitempty"`

	/* Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* HTTP target. */
	// +optional
	HttpTarget *JobHttpTarget `json:"httpTarget,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Pub/Sub target. */
	// +optional
	PubsubTarget *JobPubsubTarget `json:"pubsubTarget,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Settings that determine the retry behavior. */
	// +optional
	RetryConfig *JobRetryConfig `json:"retryConfig,omitempty"`

	/* Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time. */
	// +optional
	Schedule *string `json:"schedule,omitempty"`

	/* Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT). */
	// +optional
	TimeZone *string `json:"timeZone,omitempty"`
}

type JobAppEngineHttpTargetStatus struct {
	// +optional
	AppEngineRouting *JobAppEngineRoutingStatus `json:"appEngineRouting,omitempty"`
}

type JobAppEngineRoutingStatus struct {
	/* Output only. The host that the job is sent to. For more information about how App Engine requests are routed, see [here](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed). The host is constructed as: * `host = [application_domain_name]` `| [service] + '.' + [application_domain_name]` `| [version] + '.' + [application_domain_name]` `| [version_dot_service]+ '.' + [application_domain_name]` `| [instance] + '.' + [application_domain_name]` `| [instance_dot_service] + '.' + [application_domain_name]` `| [instance_dot_version] + '.' + [application_domain_name]` `| [instance_dot_version_dot_service] + '.' + [application_domain_name]` * `application_domain_name` = The domain name of the app, for example .appspot.com, which is associated with the job's project ID. * `service =` service * `version =` version * `version_dot_service =` version `+ '.' +` service * `instance =` instance * `instance_dot_service =` instance `+ '.' +` service * `instance_dot_version =` instance `+ '.' +` version * `instance_dot_version_dot_service =` instance `+ '.' +` version `+ '.' +` service If service is empty, then the job will be sent to the service which is the default service when the job is attempted. If version is empty, then the job will be sent to the version which is the default version when the job is attempted. If instance is empty, then the job will be sent to an instance which is available when the job is attempted. If service, version, or instance is invalid, then the job will be sent to the default version of the default service when the job is attempted. */
	// +optional
	Host *string `json:"host,omitempty"`
}

type JobDetailsStatus struct {
	/* A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one "/" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading "." is not accepted). In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows: * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a google.protobuf.Type value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.) Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com. Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics. */
	// +optional
	TypeUrl *string `json:"typeUrl,omitempty"`

	/* Must be a valid serialized protocol buffer of the above specified type. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type JobStatusStatus struct {
	/* The status code, which should be an enum value of google.rpc.Code. */
	// +optional
	Code *int64 `json:"code,omitempty"`

	/* A list of messages that carry the error details. There is a common set of message types for APIs to use. */
	// +optional
	Details []JobDetailsStatus `json:"details,omitempty"`

	/* A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client. */
	// +optional
	Message *string `json:"message,omitempty"`
}

type CloudSchedulerJobStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudSchedulerJob's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// +optional
	AppEngineHttpTarget *JobAppEngineHttpTargetStatus `json:"appEngineHttpTarget,omitempty"`

	/* Output only. The time the last job attempt started. */
	// +optional
	LastAttemptTime *string `json:"lastAttemptTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule. */
	// +optional
	ScheduleTime *string `json:"scheduleTime,omitempty"`

	/* Output only. State of the job. Possible values: STATE_UNSPECIFIED, ENABLED, PAUSED, DISABLED, UPDATE_FAILED */
	// +optional
	State *string `json:"state,omitempty"`

	/* Output only. The response from the target for the last attempted execution. */
	// +optional
	Status *JobStatusStatus `json:"status,omitempty"`

	/* Output only. The creation time of the job. */
	// +optional
	UserUpdateTime *string `json:"userUpdateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudschedulerjob;gcpcloudschedulerjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudSchedulerJob is the Schema for the cloudscheduler API
// +k8s:openapi-gen=true
type CloudSchedulerJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudSchedulerJobSpec   `json:"spec,omitempty"`
	Status CloudSchedulerJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudSchedulerJobList contains a list of CloudSchedulerJob
type CloudSchedulerJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudSchedulerJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudSchedulerJob{}, &CloudSchedulerJobList{})
}
