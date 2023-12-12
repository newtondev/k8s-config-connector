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

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vertexai/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVertexAIEndpoints implements VertexAIEndpointInterface
type FakeVertexAIEndpoints struct {
	Fake *FakeVertexaiV1alpha1
	ns   string
}

var vertexaiendpointsResource = v1alpha1.SchemeGroupVersion.WithResource("vertexaiendpoints")

var vertexaiendpointsKind = v1alpha1.SchemeGroupVersion.WithKind("VertexAIEndpoint")

// Get takes name of the vertexAIEndpoint, and returns the corresponding vertexAIEndpoint object, and an error if there is any.
func (c *FakeVertexAIEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vertexaiendpointsResource, c.ns, name), &v1alpha1.VertexAIEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VertexAIEndpoint), err
}

// List takes label and field selectors, and returns the list of VertexAIEndpoints that match those selectors.
func (c *FakeVertexAIEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VertexAIEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vertexaiendpointsResource, vertexaiendpointsKind, c.ns, opts), &v1alpha1.VertexAIEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VertexAIEndpointList{ListMeta: obj.(*v1alpha1.VertexAIEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.VertexAIEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vertexAIEndpoints.
func (c *FakeVertexAIEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vertexaiendpointsResource, c.ns, opts))

}

// Create takes the representation of a vertexAIEndpoint and creates it.  Returns the server's representation of the vertexAIEndpoint, and an error, if there is any.
func (c *FakeVertexAIEndpoints) Create(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.CreateOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vertexaiendpointsResource, c.ns, vertexAIEndpoint), &v1alpha1.VertexAIEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VertexAIEndpoint), err
}

// Update takes the representation of a vertexAIEndpoint and updates it. Returns the server's representation of the vertexAIEndpoint, and an error, if there is any.
func (c *FakeVertexAIEndpoints) Update(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.UpdateOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vertexaiendpointsResource, c.ns, vertexAIEndpoint), &v1alpha1.VertexAIEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VertexAIEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVertexAIEndpoints) UpdateStatus(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.UpdateOptions) (*v1alpha1.VertexAIEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vertexaiendpointsResource, "status", c.ns, vertexAIEndpoint), &v1alpha1.VertexAIEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VertexAIEndpoint), err
}

// Delete takes name of the vertexAIEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeVertexAIEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vertexaiendpointsResource, c.ns, name, opts), &v1alpha1.VertexAIEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVertexAIEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vertexaiendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VertexAIEndpointList{})
	return err
}

// Patch applies the patch and returns the patched vertexAIEndpoint.
func (c *FakeVertexAIEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VertexAIEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vertexaiendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VertexAIEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VertexAIEndpoint), err
}
