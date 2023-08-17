//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceSpec) DeepCopyInto(out *ContainerResourceSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceSpec.
func (in *ContainerResourceSpec) DeepCopy() *ContainerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResource) DeepCopyInto(out *ControllerResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResource.
func (in *ControllerResource) DeepCopy() *ControllerResource {
	if in == nil {
		return nil
	}
	out := new(ControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceList) DeepCopyInto(out *ControllerResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceList.
func (in *ControllerResourceList) DeepCopy() *ControllerResourceList {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceSpec) DeepCopyInto(out *ControllerResourceSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceSpec.
func (in *ControllerResourceSpec) DeepCopy() *ControllerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceStatus) DeepCopyInto(out *ControllerResourceStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceStatus.
func (in *ControllerResourceStatus) DeepCopy() *ControllerResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookConfigurationCustomization) DeepCopyInto(out *MutatingWebhookConfigurationCustomization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookConfigurationCustomization.
func (in *MutatingWebhookConfigurationCustomization) DeepCopy() *MutatingWebhookConfigurationCustomization {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookConfigurationCustomization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MutatingWebhookConfigurationCustomization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookConfigurationCustomizationList) DeepCopyInto(out *MutatingWebhookConfigurationCustomizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MutatingWebhookConfigurationCustomization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookConfigurationCustomizationList.
func (in *MutatingWebhookConfigurationCustomizationList) DeepCopy() *MutatingWebhookConfigurationCustomizationList {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookConfigurationCustomizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MutatingWebhookConfigurationCustomizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResource) DeepCopyInto(out *NamespacedControllerResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResource.
func (in *NamespacedControllerResource) DeepCopy() *NamespacedControllerResource {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedControllerResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResourceList) DeepCopyInto(out *NamespacedControllerResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedControllerResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResourceList.
func (in *NamespacedControllerResourceList) DeepCopy() *NamespacedControllerResourceList {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedControllerResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResourceSpec) DeepCopyInto(out *NamespacedControllerResourceSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResourceSpec.
func (in *NamespacedControllerResourceSpec) DeepCopy() *NamespacedControllerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResourceStatus) DeepCopyInto(out *NamespacedControllerResourceStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResourceStatus.
func (in *NamespacedControllerResourceStatus) DeepCopy() *NamespacedControllerResourceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookConfigurationCustomization) DeepCopyInto(out *ValidatingWebhookConfigurationCustomization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookConfigurationCustomization.
func (in *ValidatingWebhookConfigurationCustomization) DeepCopy() *ValidatingWebhookConfigurationCustomization {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookConfigurationCustomization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidatingWebhookConfigurationCustomization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookConfigurationCustomizationList) DeepCopyInto(out *ValidatingWebhookConfigurationCustomizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidatingWebhookConfigurationCustomization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookConfigurationCustomizationList.
func (in *ValidatingWebhookConfigurationCustomizationList) DeepCopy() *ValidatingWebhookConfigurationCustomizationList {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookConfigurationCustomizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidatingWebhookConfigurationCustomizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigurationCustomizationSpec) DeepCopyInto(out *WebhookConfigurationCustomizationSpec) {
	*out = *in
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]WebhookCustomizationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigurationCustomizationSpec.
func (in *WebhookConfigurationCustomizationSpec) DeepCopy() *WebhookConfigurationCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigurationCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigurationCustomizationStatus) DeepCopyInto(out *WebhookConfigurationCustomizationStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigurationCustomizationStatus.
func (in *WebhookConfigurationCustomizationStatus) DeepCopy() *WebhookConfigurationCustomizationStatus {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigurationCustomizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookCustomizationSpec) DeepCopyInto(out *WebhookCustomizationSpec) {
	*out = *in
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookCustomizationSpec.
func (in *WebhookCustomizationSpec) DeepCopy() *WebhookCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}
