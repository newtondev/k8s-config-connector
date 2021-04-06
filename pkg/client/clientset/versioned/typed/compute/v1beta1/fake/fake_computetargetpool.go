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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeTargetPools implements ComputeTargetPoolInterface
type FakeComputeTargetPools struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computetargetpoolsResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Resource: "computetargetpools"}

var computetargetpoolsKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ComputeTargetPool"}

// Get takes name of the computeTargetPool, and returns the corresponding computeTargetPool object, and an error if there is any.
func (c *FakeComputeTargetPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeTargetPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computetargetpoolsResource, c.ns, name), &v1beta1.ComputeTargetPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetPool), err
}

// List takes label and field selectors, and returns the list of ComputeTargetPools that match those selectors.
func (c *FakeComputeTargetPools) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeTargetPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computetargetpoolsResource, computetargetpoolsKind, c.ns, opts), &v1beta1.ComputeTargetPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeTargetPoolList{ListMeta: obj.(*v1beta1.ComputeTargetPoolList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeTargetPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeTargetPools.
func (c *FakeComputeTargetPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computetargetpoolsResource, c.ns, opts))

}

// Create takes the representation of a computeTargetPool and creates it.  Returns the server's representation of the computeTargetPool, and an error, if there is any.
func (c *FakeComputeTargetPools) Create(ctx context.Context, computeTargetPool *v1beta1.ComputeTargetPool, opts v1.CreateOptions) (result *v1beta1.ComputeTargetPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computetargetpoolsResource, c.ns, computeTargetPool), &v1beta1.ComputeTargetPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetPool), err
}

// Update takes the representation of a computeTargetPool and updates it. Returns the server's representation of the computeTargetPool, and an error, if there is any.
func (c *FakeComputeTargetPools) Update(ctx context.Context, computeTargetPool *v1beta1.ComputeTargetPool, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computetargetpoolsResource, c.ns, computeTargetPool), &v1beta1.ComputeTargetPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeTargetPools) UpdateStatus(ctx context.Context, computeTargetPool *v1beta1.ComputeTargetPool, opts v1.UpdateOptions) (*v1beta1.ComputeTargetPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computetargetpoolsResource, "status", c.ns, computeTargetPool), &v1beta1.ComputeTargetPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetPool), err
}

// Delete takes name of the computeTargetPool and deletes it. Returns an error if one occurs.
func (c *FakeComputeTargetPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computetargetpoolsResource, c.ns, name), &v1beta1.ComputeTargetPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeTargetPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computetargetpoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeTargetPoolList{})
	return err
}

// Patch applies the patch and returns the patched computeTargetPool.
func (c *FakeComputeTargetPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computetargetpoolsResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeTargetPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetPool), err
}
