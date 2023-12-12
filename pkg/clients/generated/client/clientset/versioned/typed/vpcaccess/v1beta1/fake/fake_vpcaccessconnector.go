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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vpcaccess/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVPCAccessConnectors implements VPCAccessConnectorInterface
type FakeVPCAccessConnectors struct {
	Fake *FakeVpcaccessV1beta1
	ns   string
}

var vpcaccessconnectorsResource = v1beta1.SchemeGroupVersion.WithResource("vpcaccessconnectors")

var vpcaccessconnectorsKind = v1beta1.SchemeGroupVersion.WithKind("VPCAccessConnector")

// Get takes name of the vPCAccessConnector, and returns the corresponding vPCAccessConnector object, and an error if there is any.
func (c *FakeVPCAccessConnectors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VPCAccessConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcaccessconnectorsResource, c.ns, name), &v1beta1.VPCAccessConnector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VPCAccessConnector), err
}

// List takes label and field selectors, and returns the list of VPCAccessConnectors that match those selectors.
func (c *FakeVPCAccessConnectors) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VPCAccessConnectorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcaccessconnectorsResource, vpcaccessconnectorsKind, c.ns, opts), &v1beta1.VPCAccessConnectorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VPCAccessConnectorList{ListMeta: obj.(*v1beta1.VPCAccessConnectorList).ListMeta}
	for _, item := range obj.(*v1beta1.VPCAccessConnectorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vPCAccessConnectors.
func (c *FakeVPCAccessConnectors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcaccessconnectorsResource, c.ns, opts))

}

// Create takes the representation of a vPCAccessConnector and creates it.  Returns the server's representation of the vPCAccessConnector, and an error, if there is any.
func (c *FakeVPCAccessConnectors) Create(ctx context.Context, vPCAccessConnector *v1beta1.VPCAccessConnector, opts v1.CreateOptions) (result *v1beta1.VPCAccessConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcaccessconnectorsResource, c.ns, vPCAccessConnector), &v1beta1.VPCAccessConnector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VPCAccessConnector), err
}

// Update takes the representation of a vPCAccessConnector and updates it. Returns the server's representation of the vPCAccessConnector, and an error, if there is any.
func (c *FakeVPCAccessConnectors) Update(ctx context.Context, vPCAccessConnector *v1beta1.VPCAccessConnector, opts v1.UpdateOptions) (result *v1beta1.VPCAccessConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcaccessconnectorsResource, c.ns, vPCAccessConnector), &v1beta1.VPCAccessConnector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VPCAccessConnector), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVPCAccessConnectors) UpdateStatus(ctx context.Context, vPCAccessConnector *v1beta1.VPCAccessConnector, opts v1.UpdateOptions) (*v1beta1.VPCAccessConnector, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcaccessconnectorsResource, "status", c.ns, vPCAccessConnector), &v1beta1.VPCAccessConnector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VPCAccessConnector), err
}

// Delete takes name of the vPCAccessConnector and deletes it. Returns an error if one occurs.
func (c *FakeVPCAccessConnectors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vpcaccessconnectorsResource, c.ns, name, opts), &v1beta1.VPCAccessConnector{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVPCAccessConnectors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcaccessconnectorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VPCAccessConnectorList{})
	return err
}

// Patch applies the patch and returns the patched vPCAccessConnector.
func (c *FakeVPCAccessConnectors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VPCAccessConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcaccessconnectorsResource, c.ns, name, pt, data, subresources...), &v1beta1.VPCAccessConnector{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VPCAccessConnector), err
}
