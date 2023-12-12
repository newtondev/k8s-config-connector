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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeRouterPeers implements ComputeRouterPeerInterface
type FakeComputeRouterPeers struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computerouterpeersResource = v1beta1.SchemeGroupVersion.WithResource("computerouterpeers")

var computerouterpeersKind = v1beta1.SchemeGroupVersion.WithKind("ComputeRouterPeer")

// Get takes name of the computeRouterPeer, and returns the corresponding computeRouterPeer object, and an error if there is any.
func (c *FakeComputeRouterPeers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computerouterpeersResource, c.ns, name), &v1beta1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRouterPeer), err
}

// List takes label and field selectors, and returns the list of ComputeRouterPeers that match those selectors.
func (c *FakeComputeRouterPeers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeRouterPeerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computerouterpeersResource, computerouterpeersKind, c.ns, opts), &v1beta1.ComputeRouterPeerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeRouterPeerList{ListMeta: obj.(*v1beta1.ComputeRouterPeerList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeRouterPeerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeRouterPeers.
func (c *FakeComputeRouterPeers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computerouterpeersResource, c.ns, opts))

}

// Create takes the representation of a computeRouterPeer and creates it.  Returns the server's representation of the computeRouterPeer, and an error, if there is any.
func (c *FakeComputeRouterPeers) Create(ctx context.Context, computeRouterPeer *v1beta1.ComputeRouterPeer, opts v1.CreateOptions) (result *v1beta1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computerouterpeersResource, c.ns, computeRouterPeer), &v1beta1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRouterPeer), err
}

// Update takes the representation of a computeRouterPeer and updates it. Returns the server's representation of the computeRouterPeer, and an error, if there is any.
func (c *FakeComputeRouterPeers) Update(ctx context.Context, computeRouterPeer *v1beta1.ComputeRouterPeer, opts v1.UpdateOptions) (result *v1beta1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computerouterpeersResource, c.ns, computeRouterPeer), &v1beta1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRouterPeer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeRouterPeers) UpdateStatus(ctx context.Context, computeRouterPeer *v1beta1.ComputeRouterPeer, opts v1.UpdateOptions) (*v1beta1.ComputeRouterPeer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computerouterpeersResource, "status", c.ns, computeRouterPeer), &v1beta1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRouterPeer), err
}

// Delete takes name of the computeRouterPeer and deletes it. Returns an error if one occurs.
func (c *FakeComputeRouterPeers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computerouterpeersResource, c.ns, name, opts), &v1beta1.ComputeRouterPeer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeRouterPeers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computerouterpeersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeRouterPeerList{})
	return err
}

// Patch applies the patch and returns the patched computeRouterPeer.
func (c *FakeComputeRouterPeers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computerouterpeersResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRouterPeer), err
}
