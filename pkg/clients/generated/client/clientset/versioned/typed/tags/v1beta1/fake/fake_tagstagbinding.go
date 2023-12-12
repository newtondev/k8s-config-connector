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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tags/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTagsTagBindings implements TagsTagBindingInterface
type FakeTagsTagBindings struct {
	Fake *FakeTagsV1beta1
	ns   string
}

var tagstagbindingsResource = v1beta1.SchemeGroupVersion.WithResource("tagstagbindings")

var tagstagbindingsKind = v1beta1.SchemeGroupVersion.WithKind("TagsTagBinding")

// Get takes name of the tagsTagBinding, and returns the corresponding tagsTagBinding object, and an error if there is any.
func (c *FakeTagsTagBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.TagsTagBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tagstagbindingsResource, c.ns, name), &v1beta1.TagsTagBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TagsTagBinding), err
}

// List takes label and field selectors, and returns the list of TagsTagBindings that match those selectors.
func (c *FakeTagsTagBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.TagsTagBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tagstagbindingsResource, tagstagbindingsKind, c.ns, opts), &v1beta1.TagsTagBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.TagsTagBindingList{ListMeta: obj.(*v1beta1.TagsTagBindingList).ListMeta}
	for _, item := range obj.(*v1beta1.TagsTagBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tagsTagBindings.
func (c *FakeTagsTagBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tagstagbindingsResource, c.ns, opts))

}

// Create takes the representation of a tagsTagBinding and creates it.  Returns the server's representation of the tagsTagBinding, and an error, if there is any.
func (c *FakeTagsTagBindings) Create(ctx context.Context, tagsTagBinding *v1beta1.TagsTagBinding, opts v1.CreateOptions) (result *v1beta1.TagsTagBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tagstagbindingsResource, c.ns, tagsTagBinding), &v1beta1.TagsTagBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TagsTagBinding), err
}

// Update takes the representation of a tagsTagBinding and updates it. Returns the server's representation of the tagsTagBinding, and an error, if there is any.
func (c *FakeTagsTagBindings) Update(ctx context.Context, tagsTagBinding *v1beta1.TagsTagBinding, opts v1.UpdateOptions) (result *v1beta1.TagsTagBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tagstagbindingsResource, c.ns, tagsTagBinding), &v1beta1.TagsTagBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TagsTagBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTagsTagBindings) UpdateStatus(ctx context.Context, tagsTagBinding *v1beta1.TagsTagBinding, opts v1.UpdateOptions) (*v1beta1.TagsTagBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tagstagbindingsResource, "status", c.ns, tagsTagBinding), &v1beta1.TagsTagBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TagsTagBinding), err
}

// Delete takes name of the tagsTagBinding and deletes it. Returns an error if one occurs.
func (c *FakeTagsTagBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(tagstagbindingsResource, c.ns, name, opts), &v1beta1.TagsTagBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTagsTagBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tagstagbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.TagsTagBindingList{})
	return err
}

// Patch applies the patch and returns the patched tagsTagBinding.
func (c *FakeTagsTagBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TagsTagBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tagstagbindingsResource, c.ns, name, pt, data, subresources...), &v1beta1.TagsTagBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TagsTagBinding), err
}
