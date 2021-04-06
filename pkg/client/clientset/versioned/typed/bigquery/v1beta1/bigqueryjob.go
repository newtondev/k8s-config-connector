// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/bigquery/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BigQueryJobsGetter has a method to return a BigQueryJobInterface.
// A group's client should implement this interface.
type BigQueryJobsGetter interface {
	BigQueryJobs(namespace string) BigQueryJobInterface
}

// BigQueryJobInterface has methods to work with BigQueryJob resources.
type BigQueryJobInterface interface {
	Create(ctx context.Context, bigQueryJob *v1beta1.BigQueryJob, opts v1.CreateOptions) (*v1beta1.BigQueryJob, error)
	Update(ctx context.Context, bigQueryJob *v1beta1.BigQueryJob, opts v1.UpdateOptions) (*v1beta1.BigQueryJob, error)
	UpdateStatus(ctx context.Context, bigQueryJob *v1beta1.BigQueryJob, opts v1.UpdateOptions) (*v1beta1.BigQueryJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BigQueryJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BigQueryJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BigQueryJob, err error)
	BigQueryJobExpansion
}

// bigQueryJobs implements BigQueryJobInterface
type bigQueryJobs struct {
	client rest.Interface
	ns     string
}

// newBigQueryJobs returns a BigQueryJobs
func newBigQueryJobs(c *BigqueryV1beta1Client, namespace string) *bigQueryJobs {
	return &bigQueryJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bigQueryJob, and returns the corresponding bigQueryJob object, and an error if there is any.
func (c *bigQueryJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BigQueryJob, err error) {
	result = &v1beta1.BigQueryJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BigQueryJobs that match those selectors.
func (c *bigQueryJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BigQueryJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BigQueryJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bigQueryJobs.
func (c *bigQueryJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bigQueryJob and creates it.  Returns the server's representation of the bigQueryJob, and an error, if there is any.
func (c *bigQueryJobs) Create(ctx context.Context, bigQueryJob *v1beta1.BigQueryJob, opts v1.CreateOptions) (result *v1beta1.BigQueryJob, err error) {
	result = &v1beta1.BigQueryJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigQueryJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bigQueryJob and updates it. Returns the server's representation of the bigQueryJob, and an error, if there is any.
func (c *bigQueryJobs) Update(ctx context.Context, bigQueryJob *v1beta1.BigQueryJob, opts v1.UpdateOptions) (result *v1beta1.BigQueryJob, err error) {
	result = &v1beta1.BigQueryJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		Name(bigQueryJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigQueryJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bigQueryJobs) UpdateStatus(ctx context.Context, bigQueryJob *v1beta1.BigQueryJob, opts v1.UpdateOptions) (result *v1beta1.BigQueryJob, err error) {
	result = &v1beta1.BigQueryJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		Name(bigQueryJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigQueryJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bigQueryJob and deletes it. Returns an error if one occurs.
func (c *bigQueryJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bigQueryJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigqueryjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bigQueryJob.
func (c *bigQueryJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BigQueryJob, err error) {
	result = &v1beta1.BigQueryJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bigqueryjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
