//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingChannel) DeepCopyInto(out *FirebaseHostingChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingChannel.
func (in *FirebaseHostingChannel) DeepCopy() *FirebaseHostingChannel {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirebaseHostingChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingChannelList) DeepCopyInto(out *FirebaseHostingChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirebaseHostingChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingChannelList.
func (in *FirebaseHostingChannelList) DeepCopy() *FirebaseHostingChannelList {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirebaseHostingChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingChannelSpec) DeepCopyInto(out *FirebaseHostingChannelSpec) {
	*out = *in
	if in.ExpireTime != nil {
		in, out := &in.ExpireTime, &out.ExpireTime
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RetainedReleaseCount != nil {
		in, out := &in.RetainedReleaseCount, &out.RetainedReleaseCount
		*out = new(int64)
		**out = **in
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingChannelSpec.
func (in *FirebaseHostingChannelSpec) DeepCopy() *FirebaseHostingChannelSpec {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingChannelStatus) DeepCopyInto(out *FirebaseHostingChannelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingChannelStatus.
func (in *FirebaseHostingChannelStatus) DeepCopy() *FirebaseHostingChannelStatus {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingSite) DeepCopyInto(out *FirebaseHostingSite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingSite.
func (in *FirebaseHostingSite) DeepCopy() *FirebaseHostingSite {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingSite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirebaseHostingSite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingSiteList) DeepCopyInto(out *FirebaseHostingSiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirebaseHostingSite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingSiteList.
func (in *FirebaseHostingSiteList) DeepCopy() *FirebaseHostingSiteList {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingSiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirebaseHostingSiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingSiteSpec) DeepCopyInto(out *FirebaseHostingSiteSpec) {
	*out = *in
	if in.AppId != nil {
		in, out := &in.AppId, &out.AppId
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingSiteSpec.
func (in *FirebaseHostingSiteSpec) DeepCopy() *FirebaseHostingSiteSpec {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingSiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirebaseHostingSiteStatus) DeepCopyInto(out *FirebaseHostingSiteStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.DefaultUrl != nil {
		in, out := &in.DefaultUrl, &out.DefaultUrl
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirebaseHostingSiteStatus.
func (in *FirebaseHostingSiteStatus) DeepCopy() *FirebaseHostingSiteStatus {
	if in == nil {
		return nil
	}
	out := new(FirebaseHostingSiteStatus)
	in.DeepCopyInto(out)
	return out
}
