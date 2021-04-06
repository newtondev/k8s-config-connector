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

// ComputeImagesGetter has a method to return a ComputeImageInterface.
// A group's client should implement this interface.
type ComputeImagesGetter interface {
	ComputeImages(namespace string) ComputeImageInterface
}

// ComputeImageInterface has methods to work with ComputeImage resources.
type ComputeImageInterface interface {
	Create(ctx context.Context, computeImage *v1beta1.ComputeImage, opts v1.CreateOptions) (*v1beta1.ComputeImage, error)
	Update(ctx context.Context, computeImage *v1beta1.ComputeImage, opts v1.UpdateOptions) (*v1beta1.ComputeImage, error)
	UpdateStatus(ctx context.Context, computeImage *v1beta1.ComputeImage, opts v1.UpdateOptions) (*v1beta1.ComputeImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ComputeImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ComputeImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeImage, err error)
	ComputeImageExpansion
}

// computeImages implements ComputeImageInterface
type computeImages struct {
	client rest.Interface
	ns     string
}

// newComputeImages returns a ComputeImages
func newComputeImages(c *ComputeV1beta1Client, namespace string) *computeImages {
	return &computeImages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeImage, and returns the corresponding computeImage object, and an error if there is any.
func (c *computeImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeImage, err error) {
	result = &v1beta1.ComputeImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeimages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeImages that match those selectors.
func (c *computeImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeImageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ComputeImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeImages.
func (c *computeImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeImage and creates it.  Returns the server's representation of the computeImage, and an error, if there is any.
func (c *computeImages) Create(ctx context.Context, computeImage *v1beta1.ComputeImage, opts v1.CreateOptions) (result *v1beta1.ComputeImage, err error) {
	result = &v1beta1.ComputeImage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeImage).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeImage and updates it. Returns the server's representation of the computeImage, and an error, if there is any.
func (c *computeImages) Update(ctx context.Context, computeImage *v1beta1.ComputeImage, opts v1.UpdateOptions) (result *v1beta1.ComputeImage, err error) {
	result = &v1beta1.ComputeImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeimages").
		Name(computeImage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeImage).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeImages) UpdateStatus(ctx context.Context, computeImage *v1beta1.ComputeImage, opts v1.UpdateOptions) (result *v1beta1.ComputeImage, err error) {
	result = &v1beta1.ComputeImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeimages").
		Name(computeImage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeImage).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeImage and deletes it. Returns an error if one occurs.
func (c *computeImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeimages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeimages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeImage.
func (c *computeImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeImage, err error) {
	result = &v1beta1.ComputeImage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeimages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
