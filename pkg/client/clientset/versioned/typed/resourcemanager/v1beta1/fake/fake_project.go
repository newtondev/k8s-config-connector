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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/resourcemanager/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjects implements ProjectInterface
type FakeProjects struct {
	Fake *FakeResourcemanagerV1beta1
	ns   string
}

var projectsResource = schema.GroupVersionResource{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Resource: "projects"}

var projectsKind = schema.GroupVersionKind{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Project"}

// Get takes name of the project, and returns the corresponding project object, and an error if there is any.
func (c *FakeProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectsResource, c.ns, name), &v1beta1.Project{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// List takes label and field selectors, and returns the list of Projects that match those selectors.
func (c *FakeProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectsResource, projectsKind, c.ns, opts), &v1beta1.ProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ProjectList{ListMeta: obj.(*v1beta1.ProjectList).ListMeta}
	for _, item := range obj.(*v1beta1.ProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projects.
func (c *FakeProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectsResource, c.ns, opts))

}

// Create takes the representation of a project and creates it.  Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Create(ctx context.Context, project *v1beta1.Project, opts v1.CreateOptions) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectsResource, c.ns, project), &v1beta1.Project{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// Update takes the representation of a project and updates it. Returns the server's representation of the project, and an error, if there is any.
func (c *FakeProjects) Update(ctx context.Context, project *v1beta1.Project, opts v1.UpdateOptions) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectsResource, c.ns, project), &v1beta1.Project{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjects) UpdateStatus(ctx context.Context, project *v1beta1.Project, opts v1.UpdateOptions) (*v1beta1.Project, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectsResource, "status", c.ns, project), &v1beta1.Project{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}

// Delete takes name of the project and deletes it. Returns an error if one occurs.
func (c *FakeProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(projectsResource, c.ns, name), &v1beta1.Project{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ProjectList{})
	return err
}

// Patch applies the patch and returns the patched project.
func (c *FakeProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Project, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectsResource, c.ns, name, pt, data, subresources...), &v1beta1.Project{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Project), err
}
