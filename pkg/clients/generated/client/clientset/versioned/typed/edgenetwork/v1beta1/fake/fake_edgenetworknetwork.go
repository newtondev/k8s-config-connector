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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/edgenetwork/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEdgeNetworkNetworks implements EdgeNetworkNetworkInterface
type FakeEdgeNetworkNetworks struct {
	Fake *FakeEdgenetworkV1beta1
	ns   string
}

var edgenetworknetworksResource = v1beta1.SchemeGroupVersion.WithResource("edgenetworknetworks")

var edgenetworknetworksKind = v1beta1.SchemeGroupVersion.WithKind("EdgeNetworkNetwork")

// Get takes name of the edgeNetworkNetwork, and returns the corresponding edgeNetworkNetwork object, and an error if there is any.
func (c *FakeEdgeNetworkNetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.EdgeNetworkNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(edgenetworknetworksResource, c.ns, name), &v1beta1.EdgeNetworkNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EdgeNetworkNetwork), err
}

// List takes label and field selectors, and returns the list of EdgeNetworkNetworks that match those selectors.
func (c *FakeEdgeNetworkNetworks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EdgeNetworkNetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(edgenetworknetworksResource, edgenetworknetworksKind, c.ns, opts), &v1beta1.EdgeNetworkNetworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EdgeNetworkNetworkList{ListMeta: obj.(*v1beta1.EdgeNetworkNetworkList).ListMeta}
	for _, item := range obj.(*v1beta1.EdgeNetworkNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested edgeNetworkNetworks.
func (c *FakeEdgeNetworkNetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(edgenetworknetworksResource, c.ns, opts))

}

// Create takes the representation of a edgeNetworkNetwork and creates it.  Returns the server's representation of the edgeNetworkNetwork, and an error, if there is any.
func (c *FakeEdgeNetworkNetworks) Create(ctx context.Context, edgeNetworkNetwork *v1beta1.EdgeNetworkNetwork, opts v1.CreateOptions) (result *v1beta1.EdgeNetworkNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(edgenetworknetworksResource, c.ns, edgeNetworkNetwork), &v1beta1.EdgeNetworkNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EdgeNetworkNetwork), err
}

// Update takes the representation of a edgeNetworkNetwork and updates it. Returns the server's representation of the edgeNetworkNetwork, and an error, if there is any.
func (c *FakeEdgeNetworkNetworks) Update(ctx context.Context, edgeNetworkNetwork *v1beta1.EdgeNetworkNetwork, opts v1.UpdateOptions) (result *v1beta1.EdgeNetworkNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(edgenetworknetworksResource, c.ns, edgeNetworkNetwork), &v1beta1.EdgeNetworkNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EdgeNetworkNetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEdgeNetworkNetworks) UpdateStatus(ctx context.Context, edgeNetworkNetwork *v1beta1.EdgeNetworkNetwork, opts v1.UpdateOptions) (*v1beta1.EdgeNetworkNetwork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(edgenetworknetworksResource, "status", c.ns, edgeNetworkNetwork), &v1beta1.EdgeNetworkNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EdgeNetworkNetwork), err
}

// Delete takes name of the edgeNetworkNetwork and deletes it. Returns an error if one occurs.
func (c *FakeEdgeNetworkNetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(edgenetworknetworksResource, c.ns, name, opts), &v1beta1.EdgeNetworkNetwork{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEdgeNetworkNetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(edgenetworknetworksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EdgeNetworkNetworkList{})
	return err
}

// Patch applies the patch and returns the patched edgeNetworkNetwork.
func (c *FakeEdgeNetworkNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.EdgeNetworkNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(edgenetworknetworksResource, c.ns, name, pt, data, subresources...), &v1beta1.EdgeNetworkNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.EdgeNetworkNetwork), err
}
