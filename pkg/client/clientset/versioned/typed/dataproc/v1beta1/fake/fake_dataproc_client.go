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
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataproc/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDataprocV1beta1 struct {
	*testing.Fake
}

func (c *FakeDataprocV1beta1) DataprocAutoscalingPolicies(namespace string) v1beta1.DataprocAutoscalingPolicyInterface {
	return &FakeDataprocAutoscalingPolicies{c, namespace}
}

func (c *FakeDataprocV1beta1) DataprocClusters(namespace string) v1beta1.DataprocClusterInterface {
	return &FakeDataprocClusters{c, namespace}
}

func (c *FakeDataprocV1beta1) DataprocWorkflowTemplates(namespace string) v1beta1.DataprocWorkflowTemplateInterface {
	return &FakeDataprocWorkflowTemplates{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDataprocV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
