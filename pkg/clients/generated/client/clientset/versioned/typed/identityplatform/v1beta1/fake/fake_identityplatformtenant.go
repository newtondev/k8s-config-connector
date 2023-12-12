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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityPlatformTenants implements IdentityPlatformTenantInterface
type FakeIdentityPlatformTenants struct {
	Fake *FakeIdentityplatformV1beta1
	ns   string
}

var identityplatformtenantsResource = v1beta1.SchemeGroupVersion.WithResource("identityplatformtenants")

var identityplatformtenantsKind = v1beta1.SchemeGroupVersion.WithKind("IdentityPlatformTenant")

// Get takes name of the identityPlatformTenant, and returns the corresponding identityPlatformTenant object, and an error if there is any.
func (c *FakeIdentityPlatformTenants) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IdentityPlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityplatformtenantsResource, c.ns, name), &v1beta1.IdentityPlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenant), err
}

// List takes label and field selectors, and returns the list of IdentityPlatformTenants that match those selectors.
func (c *FakeIdentityPlatformTenants) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IdentityPlatformTenantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityplatformtenantsResource, identityplatformtenantsKind, c.ns, opts), &v1beta1.IdentityPlatformTenantList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IdentityPlatformTenantList{ListMeta: obj.(*v1beta1.IdentityPlatformTenantList).ListMeta}
	for _, item := range obj.(*v1beta1.IdentityPlatformTenantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityPlatformTenants.
func (c *FakeIdentityPlatformTenants) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityplatformtenantsResource, c.ns, opts))

}

// Create takes the representation of a identityPlatformTenant and creates it.  Returns the server's representation of the identityPlatformTenant, and an error, if there is any.
func (c *FakeIdentityPlatformTenants) Create(ctx context.Context, identityPlatformTenant *v1beta1.IdentityPlatformTenant, opts v1.CreateOptions) (result *v1beta1.IdentityPlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityplatformtenantsResource, c.ns, identityPlatformTenant), &v1beta1.IdentityPlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenant), err
}

// Update takes the representation of a identityPlatformTenant and updates it. Returns the server's representation of the identityPlatformTenant, and an error, if there is any.
func (c *FakeIdentityPlatformTenants) Update(ctx context.Context, identityPlatformTenant *v1beta1.IdentityPlatformTenant, opts v1.UpdateOptions) (result *v1beta1.IdentityPlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityplatformtenantsResource, c.ns, identityPlatformTenant), &v1beta1.IdentityPlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenant), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityPlatformTenants) UpdateStatus(ctx context.Context, identityPlatformTenant *v1beta1.IdentityPlatformTenant, opts v1.UpdateOptions) (*v1beta1.IdentityPlatformTenant, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityplatformtenantsResource, "status", c.ns, identityPlatformTenant), &v1beta1.IdentityPlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenant), err
}

// Delete takes name of the identityPlatformTenant and deletes it. Returns an error if one occurs.
func (c *FakeIdentityPlatformTenants) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(identityplatformtenantsResource, c.ns, name, opts), &v1beta1.IdentityPlatformTenant{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityPlatformTenants) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityplatformtenantsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IdentityPlatformTenantList{})
	return err
}

// Patch applies the patch and returns the patched identityPlatformTenant.
func (c *FakeIdentityPlatformTenants) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IdentityPlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityplatformtenantsResource, c.ns, name, pt, data, subresources...), &v1beta1.IdentityPlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenant), err
}
