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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/pubsub/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePubSubTopics implements PubSubTopicInterface
type FakePubSubTopics struct {
	Fake *FakePubsubV1beta1
	ns   string
}

var pubsubtopicsResource = schema.GroupVersionResource{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Resource: "pubsubtopics"}

var pubsubtopicsKind = schema.GroupVersionKind{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubTopic"}

// Get takes name of the pubSubTopic, and returns the corresponding pubSubTopic object, and an error if there is any.
func (c *FakePubSubTopics) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PubSubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsubtopicsResource, c.ns, name), &v1beta1.PubSubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubTopic), err
}

// List takes label and field selectors, and returns the list of PubSubTopics that match those selectors.
func (c *FakePubSubTopics) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PubSubTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsubtopicsResource, pubsubtopicsKind, c.ns, opts), &v1beta1.PubSubTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PubSubTopicList{ListMeta: obj.(*v1beta1.PubSubTopicList).ListMeta}
	for _, item := range obj.(*v1beta1.PubSubTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubSubTopics.
func (c *FakePubSubTopics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsubtopicsResource, c.ns, opts))

}

// Create takes the representation of a pubSubTopic and creates it.  Returns the server's representation of the pubSubTopic, and an error, if there is any.
func (c *FakePubSubTopics) Create(ctx context.Context, pubSubTopic *v1beta1.PubSubTopic, opts v1.CreateOptions) (result *v1beta1.PubSubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsubtopicsResource, c.ns, pubSubTopic), &v1beta1.PubSubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubTopic), err
}

// Update takes the representation of a pubSubTopic and updates it. Returns the server's representation of the pubSubTopic, and an error, if there is any.
func (c *FakePubSubTopics) Update(ctx context.Context, pubSubTopic *v1beta1.PubSubTopic, opts v1.UpdateOptions) (result *v1beta1.PubSubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsubtopicsResource, c.ns, pubSubTopic), &v1beta1.PubSubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubSubTopics) UpdateStatus(ctx context.Context, pubSubTopic *v1beta1.PubSubTopic, opts v1.UpdateOptions) (*v1beta1.PubSubTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsubtopicsResource, "status", c.ns, pubSubTopic), &v1beta1.PubSubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubTopic), err
}

// Delete takes name of the pubSubTopic and deletes it. Returns an error if one occurs.
func (c *FakePubSubTopics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pubsubtopicsResource, c.ns, name), &v1beta1.PubSubTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubSubTopics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsubtopicsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PubSubTopicList{})
	return err
}

// Patch applies the patch and returns the patched pubSubTopic.
func (c *FakePubSubTopics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PubSubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsubtopicsResource, c.ns, name, pt, data, subresources...), &v1beta1.PubSubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubTopic), err
}
