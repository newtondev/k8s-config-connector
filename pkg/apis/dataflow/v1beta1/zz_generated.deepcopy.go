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

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowFlexTemplateJob) DeepCopyInto(out *DataflowFlexTemplateJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowFlexTemplateJob.
func (in *DataflowFlexTemplateJob) DeepCopy() *DataflowFlexTemplateJob {
	if in == nil {
		return nil
	}
	out := new(DataflowFlexTemplateJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataflowFlexTemplateJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowFlexTemplateJobList) DeepCopyInto(out *DataflowFlexTemplateJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataflowFlexTemplateJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowFlexTemplateJobList.
func (in *DataflowFlexTemplateJobList) DeepCopy() *DataflowFlexTemplateJobList {
	if in == nil {
		return nil
	}
	out := new(DataflowFlexTemplateJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataflowFlexTemplateJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowFlexTemplateJobSpec) DeepCopyInto(out *DataflowFlexTemplateJobSpec) {
	*out = *in
	out.Parameters = in.Parameters
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowFlexTemplateJobSpec.
func (in *DataflowFlexTemplateJobSpec) DeepCopy() *DataflowFlexTemplateJobSpec {
	if in == nil {
		return nil
	}
	out := new(DataflowFlexTemplateJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowFlexTemplateJobStatus) DeepCopyInto(out *DataflowFlexTemplateJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowFlexTemplateJobStatus.
func (in *DataflowFlexTemplateJobStatus) DeepCopy() *DataflowFlexTemplateJobStatus {
	if in == nil {
		return nil
	}
	out := new(DataflowFlexTemplateJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowJob) DeepCopyInto(out *DataflowJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowJob.
func (in *DataflowJob) DeepCopy() *DataflowJob {
	if in == nil {
		return nil
	}
	out := new(DataflowJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataflowJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowJobList) DeepCopyInto(out *DataflowJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataflowJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowJobList.
func (in *DataflowJobList) DeepCopy() *DataflowJobList {
	if in == nil {
		return nil
	}
	out := new(DataflowJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataflowJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowJobSpec) DeepCopyInto(out *DataflowJobSpec) {
	*out = *in
	if in.AdditionalExperiments != nil {
		in, out := &in.AdditionalExperiments, &out.AdditionalExperiments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.KmsKeyRef = in.KmsKeyRef
	out.NetworkRef = in.NetworkRef
	out.Parameters = in.Parameters
	out.ServiceAccountRef = in.ServiceAccountRef
	out.SubnetworkRef = in.SubnetworkRef
	out.TransformNameMapping = in.TransformNameMapping
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowJobSpec.
func (in *DataflowJobSpec) DeepCopy() *DataflowJobSpec {
	if in == nil {
		return nil
	}
	out := new(DataflowJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataflowJobStatus) DeepCopyInto(out *DataflowJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataflowJobStatus.
func (in *DataflowJobStatus) DeepCopy() *DataflowJobStatus {
	if in == nil {
		return nil
	}
	out := new(DataflowJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlextemplatejobParameters) DeepCopyInto(out *FlextemplatejobParameters) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlextemplatejobParameters.
func (in *FlextemplatejobParameters) DeepCopy() *FlextemplatejobParameters {
	if in == nil {
		return nil
	}
	out := new(FlextemplatejobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobParameters) DeepCopyInto(out *JobParameters) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobParameters.
func (in *JobParameters) DeepCopy() *JobParameters {
	if in == nil {
		return nil
	}
	out := new(JobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTransformNameMapping) DeepCopyInto(out *JobTransformNameMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTransformNameMapping.
func (in *JobTransformNameMapping) DeepCopy() *JobTransformNameMapping {
	if in == nil {
		return nil
	}
	out := new(JobTransformNameMapping)
	in.DeepCopyInto(out)
	return out
}
