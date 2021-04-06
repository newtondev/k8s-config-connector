// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodepoolAutoscaling struct {
	/* Maximum number of nodes in the NodePool. Must be >= min_node_count. */
	MaxNodeCount int `json:"maxNodeCount,omitempty"`
	/* Minimum number of nodes in the NodePool. Must be >=0 and <= max_node_count. */
	MinNodeCount int `json:"minNodeCount,omitempty"`
}

type NodepoolEphemeralStorageConfig struct {
	/* Immutable. */
	LocalSsdCount int `json:"localSsdCount,omitempty"`
}

type NodepoolGuestAccelerator struct {
	/* Immutable. The number of the accelerator cards exposed to an instance. */
	Count int `json:"count,omitempty"`
	/* Immutable. The accelerator type resource name. */
	Type string `json:"type,omitempty"`
}

type NodepoolKubeletConfig struct {
	/* Enable CPU CFS quota enforcement for containers that specify CPU limits. */
	CpuCfsQuota bool `json:"cpuCfsQuota,omitempty"`
	/* Set the CPU CFS quota period value 'cpu.cfs_period_us'. */
	CpuCfsQuotaPeriod string `json:"cpuCfsQuotaPeriod,omitempty"`
	/* Control the CPU management policy on the node. */
	CpuManagerPolicy string `json:"cpuManagerPolicy,omitempty"`
}

type NodepoolLinuxNodeConfig struct {
	/* The Linux kernel parameters to be applied to the nodes and all pods running on the nodes. */
	Sysctls map[string]string `json:"sysctls,omitempty"`
}

type NodepoolManagement struct {
	/* Whether the nodes will be automatically repaired. */
	AutoRepair bool `json:"autoRepair,omitempty"`
	/* Whether the nodes will be automatically upgraded. */
	AutoUpgrade bool `json:"autoUpgrade,omitempty"`
}

type NodepoolNodeConfig struct {
	/*  */
	BootDiskKMSCryptoKeyRef v1alpha1.ResourceRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`
	/* Immutable. Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	DiskSizeGb int `json:"diskSizeGb,omitempty"`
	/* Immutable. Type of the disk attached to each node. */
	DiskType string `json:"diskType,omitempty"`
	/* Immutable. */
	EphemeralStorageConfig NodepoolEphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`
	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	GuestAccelerator []NodepoolGuestAccelerator `json:"guestAccelerator,omitempty"`
	/* The image type to use for this node. Note that for a given image type, the latest version of it will be used. */
	ImageType string `json:"imageType,omitempty"`
	/* Node kubelet configs. */
	KubeletConfig NodepoolKubeletConfig `json:"kubeletConfig,omitempty"`
	/* Immutable. The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node. */
	Labels map[string]string `json:"labels,omitempty"`
	/* Parameters that can be configured on Linux nodes. */
	LinuxNodeConfig NodepoolLinuxNodeConfig `json:"linuxNodeConfig,omitempty"`
	/* Immutable. The number of local SSD disks to be attached to the node. */
	LocalSsdCount int `json:"localSsdCount,omitempty"`
	/* Immutable. The name of a Google Compute Engine machine type. */
	MachineType string `json:"machineType,omitempty"`
	/* Immutable. The metadata key/value pairs assigned to instances in the cluster. */
	Metadata map[string]string `json:"metadata,omitempty"`
	/* Immutable. Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. */
	MinCpuPlatform string `json:"minCpuPlatform,omitempty"`
	/* Immutable. The set of Google API scopes to be made available on all of the node VMs. */
	OauthScopes []string `json:"oauthScopes,omitempty"`
	/* Immutable. Whether the nodes are created as preemptible VM instances. */
	Preemptible bool `json:"preemptible,omitempty"`
	/* Immutable. Sandbox configuration for this node. */
	SandboxConfig NodepoolSandboxConfig `json:"sandboxConfig,omitempty"`
	/*  */
	ServiceAccountRef v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
	/* Immutable. Shielded Instance options. */
	ShieldedInstanceConfig NodepoolShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`
	/* Immutable. The list of instance tags applied to all nodes. */
	Tags []string `json:"tags,omitempty"`
	/* Immutable. List of Kubernetes taints to be applied to each node. */
	Taint []NodepoolTaint `json:"taint,omitempty"`
	/* The workload metadata configuration for this node. */
	WorkloadMetadataConfig NodepoolWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

type NodepoolSandboxConfig struct {
	/* Type of the sandbox to use for the node (e.g. 'gvisor') */
	SandboxType string `json:"sandboxType,omitempty"`
}

type NodepoolShieldedInstanceConfig struct {
	/* Immutable. Defines whether the instance has integrity monitoring enabled. */
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty"`
	/* Immutable. Defines whether the instance has Secure Boot enabled. */
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty"`
}

type NodepoolTaint struct {
	/* Immutable. Effect for taint. */
	Effect string `json:"effect,omitempty"`
	/* Immutable. Key for taint. */
	Key string `json:"key,omitempty"`
	/* Immutable. Value for taint. */
	Value string `json:"value,omitempty"`
}

type NodepoolUpgradeSettings struct {
	/* The number of additional nodes that can be added to the node pool during an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater. */
	MaxSurge int `json:"maxSurge,omitempty"`
	/* The number of nodes that can be simultaneously unavailable during an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater. */
	MaxUnavailable int `json:"maxUnavailable,omitempty"`
}

type NodepoolWorkloadMetadataConfig struct {
	/* NodeMetadata is the configuration for how to expose metadata to the workloads running on the node. */
	NodeMetadata string `json:"nodeMetadata,omitempty"`
}

type ContainerNodePoolSpec struct {
	/* Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage. */
	Autoscaling NodepoolAutoscaling `json:"autoscaling,omitempty"`
	/*  */
	ClusterRef v1alpha1.ResourceRef `json:"clusterRef,omitempty"`
	/* Immutable. The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource. */
	InitialNodeCount int `json:"initialNodeCount,omitempty"`
	/* Immutable. The location (region or zone) of the cluster. */
	Location string `json:"location,omitempty"`
	/* Node management configuration, wherein auto-repair and auto-upgrade is configured. */
	Management NodepoolManagement `json:"management,omitempty"`
	/* Immutable. The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled. */
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty"`
	/* Immutable. Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name. */
	NamePrefix string `json:"namePrefix,omitempty"`
	/* Immutable. The configuration of the nodepool */
	NodeConfig NodepoolNodeConfig `json:"nodeConfig,omitempty"`
	/* The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling. */
	NodeCount int `json:"nodeCount,omitempty"`
	/* The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used. */
	NodeLocations []string `json:"nodeLocations,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20. */
	UpgradeSettings NodepoolUpgradeSettings `json:"upgradeSettings,omitempty"`
	/*  */
	Version string `json:"version,omitempty"`
}

type ContainerNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerNodePool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource URLs of the managed instance groups associated with this node pool. */
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`
	/*  */
	Operation string `json:"operation,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerNodePool is the Schema for the container API
// +k8s:openapi-gen=true
type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status ContainerNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerNodePoolList contains a list of ContainerNodePool
type ContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerNodePool{}, &ContainerNodePoolList{})
}
