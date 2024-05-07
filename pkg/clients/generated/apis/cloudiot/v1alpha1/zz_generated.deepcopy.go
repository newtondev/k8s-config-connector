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
func (in *CloudIOTDevice) DeepCopyInto(out *CloudIOTDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDevice.
func (in *CloudIOTDevice) DeepCopy() *CloudIOTDevice {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudIOTDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceList) DeepCopyInto(out *CloudIOTDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudIOTDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceList.
func (in *CloudIOTDeviceList) DeepCopy() *CloudIOTDeviceList {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudIOTDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceRegistry) DeepCopyInto(out *CloudIOTDeviceRegistry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceRegistry.
func (in *CloudIOTDeviceRegistry) DeepCopy() *CloudIOTDeviceRegistry {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudIOTDeviceRegistry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceRegistryList) DeepCopyInto(out *CloudIOTDeviceRegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudIOTDeviceRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceRegistryList.
func (in *CloudIOTDeviceRegistryList) DeepCopy() *CloudIOTDeviceRegistryList {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceRegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudIOTDeviceRegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceRegistrySpec) DeepCopyInto(out *CloudIOTDeviceRegistrySpec) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]DeviceregistryCredentials, len(*in))
		copy(*out, *in)
	}
	if in.EventNotificationConfigs != nil {
		in, out := &in.EventNotificationConfigs, &out.EventNotificationConfigs
		*out = make([]DeviceregistryEventNotificationConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpConfig != nil {
		in, out := &in.HttpConfig, &out.HttpConfig
		*out = new(DeviceregistryHttpConfig)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.MqttConfig != nil {
		in, out := &in.MqttConfig, &out.MqttConfig
		*out = new(DeviceregistryMqttConfig)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.StateNotificationConfig != nil {
		in, out := &in.StateNotificationConfig, &out.StateNotificationConfig
		*out = new(DeviceregistryStateNotificationConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceRegistrySpec.
func (in *CloudIOTDeviceRegistrySpec) DeepCopy() *CloudIOTDeviceRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceRegistryStatus) DeepCopyInto(out *CloudIOTDeviceRegistryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceRegistryStatus.
func (in *CloudIOTDeviceRegistryStatus) DeepCopy() *CloudIOTDeviceRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceRegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceSpec) DeepCopyInto(out *CloudIOTDeviceSpec) {
	*out = *in
	if in.Blocked != nil {
		in, out := &in.Blocked, &out.Blocked
		*out = new(bool)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]DeviceCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayConfig != nil {
		in, out := &in.GatewayConfig, &out.GatewayConfig
		*out = new(DeviceGatewayConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceSpec.
func (in *CloudIOTDeviceSpec) DeepCopy() *CloudIOTDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudIOTDeviceStatus) DeepCopyInto(out *CloudIOTDeviceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]DeviceConfigStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastConfigAckTime != nil {
		in, out := &in.LastConfigAckTime, &out.LastConfigAckTime
		*out = new(string)
		**out = **in
	}
	if in.LastConfigSendTime != nil {
		in, out := &in.LastConfigSendTime, &out.LastConfigSendTime
		*out = new(string)
		**out = **in
	}
	if in.LastErrorStatus != nil {
		in, out := &in.LastErrorStatus, &out.LastErrorStatus
		*out = make([]DeviceLastErrorStatusStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastErrorTime != nil {
		in, out := &in.LastErrorTime, &out.LastErrorTime
		*out = new(string)
		**out = **in
	}
	if in.LastEventTime != nil {
		in, out := &in.LastEventTime, &out.LastEventTime
		*out = new(string)
		**out = **in
	}
	if in.LastHeartbeatTime != nil {
		in, out := &in.LastHeartbeatTime, &out.LastHeartbeatTime
		*out = new(string)
		**out = **in
	}
	if in.LastStateTime != nil {
		in, out := &in.LastStateTime, &out.LastStateTime
		*out = new(string)
		**out = **in
	}
	if in.NumId != nil {
		in, out := &in.NumId, &out.NumId
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = make([]DeviceStateStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudIOTDeviceStatus.
func (in *CloudIOTDeviceStatus) DeepCopy() *CloudIOTDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudIOTDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceConfigStatus) DeepCopyInto(out *DeviceConfigStatus) {
	*out = *in
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = new(string)
		**out = **in
	}
	if in.CloudUpdateTime != nil {
		in, out := &in.CloudUpdateTime, &out.CloudUpdateTime
		*out = new(string)
		**out = **in
	}
	if in.DeviceAckTime != nil {
		in, out := &in.DeviceAckTime, &out.DeviceAckTime
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceConfigStatus.
func (in *DeviceConfigStatus) DeepCopy() *DeviceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceCredentials) DeepCopyInto(out *DeviceCredentials) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	out.PublicKey = in.PublicKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceCredentials.
func (in *DeviceCredentials) DeepCopy() *DeviceCredentials {
	if in == nil {
		return nil
	}
	out := new(DeviceCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceDetailsStatus) DeepCopyInto(out *DeviceDetailsStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceDetailsStatus.
func (in *DeviceDetailsStatus) DeepCopy() *DeviceDetailsStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceDetailsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceGatewayConfig) DeepCopyInto(out *DeviceGatewayConfig) {
	*out = *in
	if in.GatewayAuthMethod != nil {
		in, out := &in.GatewayAuthMethod, &out.GatewayAuthMethod
		*out = new(string)
		**out = **in
	}
	if in.GatewayType != nil {
		in, out := &in.GatewayType, &out.GatewayType
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedGatewayId != nil {
		in, out := &in.LastAccessedGatewayId, &out.LastAccessedGatewayId
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedGatewayTime != nil {
		in, out := &in.LastAccessedGatewayTime, &out.LastAccessedGatewayTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceGatewayConfig.
func (in *DeviceGatewayConfig) DeepCopy() *DeviceGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(DeviceGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceLastErrorStatusStatus) DeepCopyInto(out *DeviceLastErrorStatusStatus) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]DeviceDetailsStatus, len(*in))
		copy(*out, *in)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceLastErrorStatusStatus.
func (in *DeviceLastErrorStatusStatus) DeepCopy() *DeviceLastErrorStatusStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceLastErrorStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePublicKey) DeepCopyInto(out *DevicePublicKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePublicKey.
func (in *DevicePublicKey) DeepCopy() *DevicePublicKey {
	if in == nil {
		return nil
	}
	out := new(DevicePublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStateStatus) DeepCopyInto(out *DeviceStateStatus) {
	*out = *in
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStateStatus.
func (in *DeviceStateStatus) DeepCopy() *DeviceStateStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceregistryCredentials) DeepCopyInto(out *DeviceregistryCredentials) {
	*out = *in
	out.PublicKeyCertificate = in.PublicKeyCertificate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceregistryCredentials.
func (in *DeviceregistryCredentials) DeepCopy() *DeviceregistryCredentials {
	if in == nil {
		return nil
	}
	out := new(DeviceregistryCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceregistryEventNotificationConfigs) DeepCopyInto(out *DeviceregistryEventNotificationConfigs) {
	*out = *in
	if in.SubfolderMatches != nil {
		in, out := &in.SubfolderMatches, &out.SubfolderMatches
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceregistryEventNotificationConfigs.
func (in *DeviceregistryEventNotificationConfigs) DeepCopy() *DeviceregistryEventNotificationConfigs {
	if in == nil {
		return nil
	}
	out := new(DeviceregistryEventNotificationConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceregistryHttpConfig) DeepCopyInto(out *DeviceregistryHttpConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceregistryHttpConfig.
func (in *DeviceregistryHttpConfig) DeepCopy() *DeviceregistryHttpConfig {
	if in == nil {
		return nil
	}
	out := new(DeviceregistryHttpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceregistryMqttConfig) DeepCopyInto(out *DeviceregistryMqttConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceregistryMqttConfig.
func (in *DeviceregistryMqttConfig) DeepCopy() *DeviceregistryMqttConfig {
	if in == nil {
		return nil
	}
	out := new(DeviceregistryMqttConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceregistryPublicKeyCertificate) DeepCopyInto(out *DeviceregistryPublicKeyCertificate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceregistryPublicKeyCertificate.
func (in *DeviceregistryPublicKeyCertificate) DeepCopy() *DeviceregistryPublicKeyCertificate {
	if in == nil {
		return nil
	}
	out := new(DeviceregistryPublicKeyCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceregistryStateNotificationConfig) DeepCopyInto(out *DeviceregistryStateNotificationConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceregistryStateNotificationConfig.
func (in *DeviceregistryStateNotificationConfig) DeepCopy() *DeviceregistryStateNotificationConfig {
	if in == nil {
		return nil
	}
	out := new(DeviceregistryStateNotificationConfig)
	in.DeepCopyInto(out)
	return out
}
