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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iap/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIAPBrands implements IAPBrandInterface
type FakeIAPBrands struct {
	Fake *FakeIapV1beta1
	ns   string
}

var iapbrandsResource = schema.GroupVersionResource{Group: "iap.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iapbrands"}

var iapbrandsKind = schema.GroupVersionKind{Group: "iap.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAPBrand"}

// Get takes name of the iAPBrand, and returns the corresponding iAPBrand object, and an error if there is any.
func (c *FakeIAPBrands) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IAPBrand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iapbrandsResource, c.ns, name), &v1beta1.IAPBrand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAPBrand), err
}

// List takes label and field selectors, and returns the list of IAPBrands that match those selectors.
func (c *FakeIAPBrands) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IAPBrandList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iapbrandsResource, iapbrandsKind, c.ns, opts), &v1beta1.IAPBrandList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAPBrandList{ListMeta: obj.(*v1beta1.IAPBrandList).ListMeta}
	for _, item := range obj.(*v1beta1.IAPBrandList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAPBrands.
func (c *FakeIAPBrands) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iapbrandsResource, c.ns, opts))

}

// Create takes the representation of a iAPBrand and creates it.  Returns the server's representation of the iAPBrand, and an error, if there is any.
func (c *FakeIAPBrands) Create(ctx context.Context, iAPBrand *v1beta1.IAPBrand, opts v1.CreateOptions) (result *v1beta1.IAPBrand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iapbrandsResource, c.ns, iAPBrand), &v1beta1.IAPBrand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAPBrand), err
}

// Update takes the representation of a iAPBrand and updates it. Returns the server's representation of the iAPBrand, and an error, if there is any.
func (c *FakeIAPBrands) Update(ctx context.Context, iAPBrand *v1beta1.IAPBrand, opts v1.UpdateOptions) (result *v1beta1.IAPBrand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iapbrandsResource, c.ns, iAPBrand), &v1beta1.IAPBrand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAPBrand), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIAPBrands) UpdateStatus(ctx context.Context, iAPBrand *v1beta1.IAPBrand, opts v1.UpdateOptions) (*v1beta1.IAPBrand, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iapbrandsResource, "status", c.ns, iAPBrand), &v1beta1.IAPBrand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAPBrand), err
}

// Delete takes name of the iAPBrand and deletes it. Returns an error if one occurs.
func (c *FakeIAPBrands) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iapbrandsResource, c.ns, name), &v1beta1.IAPBrand{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAPBrands) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iapbrandsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IAPBrandList{})
	return err
}

// Patch applies the patch and returns the patched iAPBrand.
func (c *FakeIAPBrands) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAPBrand, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iapbrandsResource, c.ns, name, pt, data, subresources...), &v1beta1.IAPBrand{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAPBrand), err
}
