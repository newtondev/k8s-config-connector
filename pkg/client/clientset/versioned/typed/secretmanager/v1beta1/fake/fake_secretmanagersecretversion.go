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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/secretmanager/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecretManagerSecretVersions implements SecretManagerSecretVersionInterface
type FakeSecretManagerSecretVersions struct {
	Fake *FakeSecretmanagerV1beta1
	ns   string
}

var secretmanagersecretversionsResource = schema.GroupVersionResource{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Resource: "secretmanagersecretversions"}

var secretmanagersecretversionsKind = schema.GroupVersionKind{Group: "secretmanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SecretManagerSecretVersion"}

// Get takes name of the secretManagerSecretVersion, and returns the corresponding secretManagerSecretVersion object, and an error if there is any.
func (c *FakeSecretManagerSecretVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SecretManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretmanagersecretversionsResource, c.ns, name), &v1beta1.SecretManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SecretManagerSecretVersion), err
}

// List takes label and field selectors, and returns the list of SecretManagerSecretVersions that match those selectors.
func (c *FakeSecretManagerSecretVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SecretManagerSecretVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretmanagersecretversionsResource, secretmanagersecretversionsKind, c.ns, opts), &v1beta1.SecretManagerSecretVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SecretManagerSecretVersionList{ListMeta: obj.(*v1beta1.SecretManagerSecretVersionList).ListMeta}
	for _, item := range obj.(*v1beta1.SecretManagerSecretVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretManagerSecretVersions.
func (c *FakeSecretManagerSecretVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretmanagersecretversionsResource, c.ns, opts))

}

// Create takes the representation of a secretManagerSecretVersion and creates it.  Returns the server's representation of the secretManagerSecretVersion, and an error, if there is any.
func (c *FakeSecretManagerSecretVersions) Create(ctx context.Context, secretManagerSecretVersion *v1beta1.SecretManagerSecretVersion, opts v1.CreateOptions) (result *v1beta1.SecretManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretmanagersecretversionsResource, c.ns, secretManagerSecretVersion), &v1beta1.SecretManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SecretManagerSecretVersion), err
}

// Update takes the representation of a secretManagerSecretVersion and updates it. Returns the server's representation of the secretManagerSecretVersion, and an error, if there is any.
func (c *FakeSecretManagerSecretVersions) Update(ctx context.Context, secretManagerSecretVersion *v1beta1.SecretManagerSecretVersion, opts v1.UpdateOptions) (result *v1beta1.SecretManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretmanagersecretversionsResource, c.ns, secretManagerSecretVersion), &v1beta1.SecretManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SecretManagerSecretVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretManagerSecretVersions) UpdateStatus(ctx context.Context, secretManagerSecretVersion *v1beta1.SecretManagerSecretVersion, opts v1.UpdateOptions) (*v1beta1.SecretManagerSecretVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretmanagersecretversionsResource, "status", c.ns, secretManagerSecretVersion), &v1beta1.SecretManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SecretManagerSecretVersion), err
}

// Delete takes name of the secretManagerSecretVersion and deletes it. Returns an error if one occurs.
func (c *FakeSecretManagerSecretVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretmanagersecretversionsResource, c.ns, name), &v1beta1.SecretManagerSecretVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretManagerSecretVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretmanagersecretversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SecretManagerSecretVersionList{})
	return err
}

// Patch applies the patch and returns the patched secretManagerSecretVersion.
func (c *FakeSecretManagerSecretVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SecretManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretmanagersecretversionsResource, c.ns, name, pt, data, subresources...), &v1beta1.SecretManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SecretManagerSecretVersion), err
}
