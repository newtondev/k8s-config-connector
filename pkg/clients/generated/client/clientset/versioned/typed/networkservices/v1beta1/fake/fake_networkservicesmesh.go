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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesMeshes implements NetworkServicesMeshInterface
type FakeNetworkServicesMeshes struct {
	Fake *FakeNetworkservicesV1beta1
	ns   string
}

var networkservicesmeshesResource = v1beta1.SchemeGroupVersion.WithResource("networkservicesmeshes")

var networkservicesmeshesKind = v1beta1.SchemeGroupVersion.WithKind("NetworkServicesMesh")

// Get takes name of the networkServicesMesh, and returns the corresponding networkServicesMesh object, and an error if there is any.
func (c *FakeNetworkServicesMeshes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkServicesMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicesmeshesResource, c.ns, name), &v1beta1.NetworkServicesMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesMesh), err
}

// List takes label and field selectors, and returns the list of NetworkServicesMeshes that match those selectors.
func (c *FakeNetworkServicesMeshes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkServicesMeshList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicesmeshesResource, networkservicesmeshesKind, c.ns, opts), &v1beta1.NetworkServicesMeshList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkServicesMeshList{ListMeta: obj.(*v1beta1.NetworkServicesMeshList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkServicesMeshList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesMeshes.
func (c *FakeNetworkServicesMeshes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicesmeshesResource, c.ns, opts))

}

// Create takes the representation of a networkServicesMesh and creates it.  Returns the server's representation of the networkServicesMesh, and an error, if there is any.
func (c *FakeNetworkServicesMeshes) Create(ctx context.Context, networkServicesMesh *v1beta1.NetworkServicesMesh, opts v1.CreateOptions) (result *v1beta1.NetworkServicesMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicesmeshesResource, c.ns, networkServicesMesh), &v1beta1.NetworkServicesMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesMesh), err
}

// Update takes the representation of a networkServicesMesh and updates it. Returns the server's representation of the networkServicesMesh, and an error, if there is any.
func (c *FakeNetworkServicesMeshes) Update(ctx context.Context, networkServicesMesh *v1beta1.NetworkServicesMesh, opts v1.UpdateOptions) (result *v1beta1.NetworkServicesMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicesmeshesResource, c.ns, networkServicesMesh), &v1beta1.NetworkServicesMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesMesh), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesMeshes) UpdateStatus(ctx context.Context, networkServicesMesh *v1beta1.NetworkServicesMesh, opts v1.UpdateOptions) (*v1beta1.NetworkServicesMesh, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicesmeshesResource, "status", c.ns, networkServicesMesh), &v1beta1.NetworkServicesMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesMesh), err
}

// Delete takes name of the networkServicesMesh and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesMeshes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicesmeshesResource, c.ns, name, opts), &v1beta1.NetworkServicesMesh{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesMeshes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicesmeshesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkServicesMeshList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesMesh.
func (c *FakeNetworkServicesMeshes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkServicesMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicesmeshesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkServicesMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesMesh), err
}
