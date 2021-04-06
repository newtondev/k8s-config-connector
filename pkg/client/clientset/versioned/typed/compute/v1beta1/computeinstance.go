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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeInstancesGetter has a method to return a ComputeInstanceInterface.
// A group's client should implement this interface.
type ComputeInstancesGetter interface {
	ComputeInstances(namespace string) ComputeInstanceInterface
}

// ComputeInstanceInterface has methods to work with ComputeInstance resources.
type ComputeInstanceInterface interface {
	Create(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.CreateOptions) (*v1beta1.ComputeInstance, error)
	Update(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.UpdateOptions) (*v1beta1.ComputeInstance, error)
	UpdateStatus(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.UpdateOptions) (*v1beta1.ComputeInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ComputeInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ComputeInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeInstance, err error)
	ComputeInstanceExpansion
}

// computeInstances implements ComputeInstanceInterface
type computeInstances struct {
	client rest.Interface
	ns     string
}

// newComputeInstances returns a ComputeInstances
func newComputeInstances(c *ComputeV1beta1Client, namespace string) *computeInstances {
	return &computeInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeInstance, and returns the corresponding computeInstance object, and an error if there is any.
func (c *computeInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeInstance, err error) {
	result = &v1beta1.ComputeInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeInstances that match those selectors.
func (c *computeInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ComputeInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeInstances.
func (c *computeInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeInstance and creates it.  Returns the server's representation of the computeInstance, and an error, if there is any.
func (c *computeInstances) Create(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.CreateOptions) (result *v1beta1.ComputeInstance, err error) {
	result = &v1beta1.ComputeInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeInstance and updates it. Returns the server's representation of the computeInstance, and an error, if there is any.
func (c *computeInstances) Update(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.UpdateOptions) (result *v1beta1.ComputeInstance, err error) {
	result = &v1beta1.ComputeInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstances").
		Name(computeInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeInstances) UpdateStatus(ctx context.Context, computeInstance *v1beta1.ComputeInstance, opts v1.UpdateOptions) (result *v1beta1.ComputeInstance, err error) {
	result = &v1beta1.ComputeInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstances").
		Name(computeInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeInstance and deletes it. Returns an error if one occurs.
func (c *computeInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeInstance.
func (c *computeInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeInstance, err error) {
	result = &v1beta1.ComputeInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
