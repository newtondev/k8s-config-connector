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

type BackendbucketCdnPolicy struct {
	/* Specifies the cache setting for all responses from this backend.
	The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"] */
	CacheMode string `json:"cacheMode,omitempty"`
	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	ClientTtl int `json:"clientTtl,omitempty"`
	/* Specifies the default TTL for cached content served by this origin for responses
	that do not have an existing valid TTL (max-age or s-max-age). */
	DefaultTtl int `json:"defaultTtl,omitempty"`
	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	MaxTtl int `json:"maxTtl,omitempty"`
	/* Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects. */
	NegativeCaching bool `json:"negativeCaching,omitempty"`
	/* Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs. */
	NegativeCachingPolicy []BackendbucketNegativeCachingPolicy `json:"negativeCachingPolicy,omitempty"`
	/* Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache. */
	ServeWhileStale int `json:"serveWhileStale,omitempty"`
	/* Maximum number of seconds the response to a signed URL request will
	be considered fresh. After this time period,
	the response will be revalidated before being served.
	When serving responses to signed URL requests,
	Cloud CDN will internally behave as though
	all responses from this backend had a "Cache-Control: public,
	max-age=[TTL]" header, regardless of any existing Cache-Control
	header. The actual headers served in responses will not be altered. */
	SignedUrlCacheMaxAgeSec int `json:"signedUrlCacheMaxAgeSec,omitempty"`
}

type BackendbucketNegativeCachingPolicy struct {
	/* The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	can be specified as values, and you cannot specify a status code more than once. */
	Code int `json:"code,omitempty"`
	/* The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
	(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL. */
	Ttl int `json:"ttl,omitempty"`
}

type ComputeBackendBucketSpec struct {
	/* Reference to the bucket. */
	BucketRef v1alpha1.ResourceRef `json:"bucketRef,omitempty"`
	/* Cloud CDN configuration for this Backend Bucket. */
	CdnPolicy BackendbucketCdnPolicy `json:"cdnPolicy,omitempty"`
	/* Headers that the HTTP/S load balancer should add to proxied responses. */
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty"`
	/* An optional textual description of the resource; provided by the
	client when the resource is created. */
	Description string `json:"description,omitempty"`
	/* If true, enable Cloud CDN for this BackendBucket. */
	EnableCdn bool `json:"enableCdn,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
}

type ComputeBackendBucketStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeBackendBucket's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeBackendBucket is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeBackendBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendBucketSpec   `json:"spec,omitempty"`
	Status ComputeBackendBucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeBackendBucketList contains a list of ComputeBackendBucket
type ComputeBackendBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendBucket{}, &ComputeBackendBucketList{})
}
