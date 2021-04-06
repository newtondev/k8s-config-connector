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

type ClusterAddonsConfig struct {
	/* The status of the CloudRun addon. It is disabled by default. Set disabled = false to enable. */
	CloudrunConfig ClusterCloudrunConfig `json:"cloudrunConfig,omitempty"`
	/* The of the Config Connector addon. */
	ConfigConnectorConfig ClusterConfigConnectorConfig `json:"configConnectorConfig,omitempty"`
	/* The status of the NodeLocal DNSCache addon. It is disabled by default. Set enabled = true to enable. */
	DnsCacheConfig ClusterDnsCacheConfig `json:"dnsCacheConfig,omitempty"`
	/* Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to disabled; set enabled = true to enable. */
	GcePersistentDiskCsiDriverConfig ClusterGcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfig,omitempty"`
	/* The status of the Horizontal Pod Autoscaling addon, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. It ensures that a Heapster pod is running in the cluster, which is also used by the Cloud Monitoring service. It is enabled by default; set disabled = true to disable. */
	HorizontalPodAutoscaling ClusterHorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`
	/* The status of the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster. It is enabled by default; set disabled = true to disable. */
	HttpLoadBalancing ClusterHttpLoadBalancing `json:"httpLoadBalancing,omitempty"`
	/* The status of the Istio addon. */
	IstioConfig ClusterIstioConfig `json:"istioConfig,omitempty"`
	/* Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set enabled = true to enable. */
	KalmConfig ClusterKalmConfig `json:"kalmConfig,omitempty"`
	/* Whether we should enable the network policy addon for the master. This must be enabled in order to enable network policy for the nodes. To enable this, you must also define a network_policy block, otherwise nothing will happen. It can only be disabled if the nodes already do not have network policies enabled. Defaults to disabled; set disabled = false to enable. */
	NetworkPolicyConfig ClusterNetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`
}

type ClusterAuthenticatorGroupsConfig struct {
	/* Immutable. The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format gke-security-groups@yourdomain.com. */
	SecurityGroup string `json:"securityGroup,omitempty"`
}

type ClusterAutoProvisioningDefaults struct {
	/* Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell. */
	MinCpuPlatform string `json:"minCpuPlatform,omitempty"`
	/* Scopes that are used by NAP when creating node pools. */
	OauthScopes []string `json:"oauthScopes,omitempty"`
	/*  */
	ServiceAccountRef v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
}

type ClusterBigqueryDestination struct {
	/* The ID of a BigQuery Dataset. */
	DatasetId string `json:"datasetId,omitempty"`
}

type ClusterCidrBlocks struct {
	/* External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation. */
	CidrBlock string `json:"cidrBlock,omitempty"`
	/* Field for users to identify CIDR blocks. */
	DisplayName string `json:"displayName,omitempty"`
}

type ClusterClientCertificateConfig struct {
	/* Immutable. Whether client certificate authorization is enabled for this cluster. */
	IssueClientCertificate bool `json:"issueClientCertificate,omitempty"`
}

type ClusterCloudrunConfig struct {
	/*  */
	Disabled bool `json:"disabled,omitempty"`
	/*  */
	LoadBalancerType string `json:"loadBalancerType,omitempty"`
}

type ClusterClusterAutoscaling struct {
	/* Contains defaults for a node pool created by NAP. */
	AutoProvisioningDefaults ClusterAutoProvisioningDefaults `json:"autoProvisioningDefaults,omitempty"`
	/* Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster. Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED. */
	AutoscalingProfile string `json:"autoscalingProfile,omitempty"`
	/* Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning. */
	Enabled bool `json:"enabled,omitempty"`
	/* Global constraints for machine resources in the cluster. Configuring the cpu and memory types is required if node auto-provisioning is enabled. These limits will apply to node pool autoscaling in addition to node auto-provisioning. */
	ResourceLimits []ClusterResourceLimits `json:"resourceLimits,omitempty"`
}

type ClusterClusterTelemetry struct {
	/* Type of the integration. */
	Type string `json:"type,omitempty"`
}

type ClusterConfidentialNodes struct {
	/* Immutable. Whether Confidential Nodes feature is enabled for all nodes in this cluster. */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterConfigConnectorConfig struct {
	/*  */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterDailyMaintenanceWindow struct {
	/*  */
	Duration string `json:"duration,omitempty"`
	/*  */
	StartTime string `json:"startTime,omitempty"`
}

type ClusterDatabaseEncryption struct {
	/* The key to use to encrypt/decrypt secrets. */
	KeyName string `json:"keyName,omitempty"`
	/* ENCRYPTED or DECRYPTED. */
	State string `json:"state,omitempty"`
}

type ClusterDefaultSnatStatus struct {
	/* When disabled is set to false, default IP masquerade rules will be applied to the nodes to prevent sNAT on cluster internal traffic. */
	Disabled bool `json:"disabled,omitempty"`
}

type ClusterDnsCacheConfig struct {
	/*  */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterEphemeralStorageConfig struct {
	/* Immutable. */
	LocalSsdCount int `json:"localSsdCount,omitempty"`
}

type ClusterGcePersistentDiskCsiDriverConfig struct {
	/*  */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterGuestAccelerator struct {
	/* Immutable. The number of the accelerator cards exposed to an instance. */
	Count int `json:"count,omitempty"`
	/* Immutable. The accelerator type resource name. */
	Type string `json:"type,omitempty"`
}

type ClusterHorizontalPodAutoscaling struct {
	/*  */
	Disabled bool `json:"disabled,omitempty"`
}

type ClusterHttpLoadBalancing struct {
	/*  */
	Disabled bool `json:"disabled,omitempty"`
}

type ClusterIpAllocationPolicy struct {
	/* Immutable. The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. */
	ClusterIpv4CidrBlock string `json:"clusterIpv4CidrBlock,omitempty"`
	/* Immutable. The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_ipv4_cidr_block can be used to automatically create a GKE-managed one. */
	ClusterSecondaryRangeName string `json:"clusterSecondaryRangeName,omitempty"`
	/* Immutable. The IP address range of the services IPs in this cluster. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. */
	ServicesIpv4CidrBlock string `json:"servicesIpv4CidrBlock,omitempty"`
	/* Immutable. The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs. Alternatively, services_ipv4_cidr_block can be used to automatically create a GKE-managed one. */
	ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName,omitempty"`
}

type ClusterIstioConfig struct {
	/* The authentication type between services in Istio. Available options include AUTH_MUTUAL_TLS. */
	Auth string `json:"auth,omitempty"`
	/* The status of the Istio addon, which makes it easy to set up Istio for services in a cluster. It is disabled by default. Set disabled = false to enable. */
	Disabled bool `json:"disabled,omitempty"`
}

type ClusterKalmConfig struct {
	/*  */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterKubeletConfig struct {
	/* Enable CPU CFS quota enforcement for containers that specify CPU limits. */
	CpuCfsQuota bool `json:"cpuCfsQuota,omitempty"`
	/* Set the CPU CFS quota period value 'cpu.cfs_period_us'. */
	CpuCfsQuotaPeriod string `json:"cpuCfsQuotaPeriod,omitempty"`
	/* Control the CPU management policy on the node. */
	CpuManagerPolicy string `json:"cpuManagerPolicy,omitempty"`
}

type ClusterLinuxNodeConfig struct {
	/* The Linux kernel parameters to be applied to the nodes and all pods running on the nodes. */
	Sysctls map[string]string `json:"sysctls,omitempty"`
}

type ClusterMaintenanceExclusion struct {
	/*  */
	EndTime string `json:"endTime,omitempty"`
	/*  */
	ExclusionName string `json:"exclusionName,omitempty"`
	/*  */
	StartTime string `json:"startTime,omitempty"`
}

type ClusterMaintenancePolicy struct {
	/* Time window specified for daily maintenance operations. Specify start_time in RFC3339 format "HH:MM”, where HH : [00-23] and MM : [00-59] GMT. */
	DailyMaintenanceWindow ClusterDailyMaintenanceWindow `json:"dailyMaintenanceWindow,omitempty"`
	/* Exceptions to maintenance window. Non-emergency maintenance should not occur in these windows. */
	MaintenanceExclusion []ClusterMaintenanceExclusion `json:"maintenanceExclusion,omitempty"`
	/* Time window for recurring maintenance operations. */
	RecurringWindow ClusterRecurringWindow `json:"recurringWindow,omitempty"`
}

type ClusterMasterAuth struct {
	/* Base64 encoded public certificate used by clients to authenticate to the cluster endpoint. */
	ClientCertificate string `json:"clientCertificate,omitempty"`
	/* Immutable. Whether client certificate authorization is enabled for this cluster. */
	ClientCertificateConfig ClusterClientCertificateConfig `json:"clientCertificateConfig,omitempty"`
	/* Base64 encoded private key used by clients to authenticate to the cluster endpoint. */
	ClientKey string `json:"clientKey,omitempty"`
	/* Base64 encoded public certificate that is the root of trust for the cluster. */
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty"`
	/* The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint. */
	Password ClusterPassword `json:"password,omitempty"`
	/* The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint. If not present basic auth will be disabled. */
	Username string `json:"username,omitempty"`
}

type ClusterMasterAuthorizedNetworksConfig struct {
	/* External networks that can access the Kubernetes cluster master through HTTPS. */
	CidrBlocks []ClusterCidrBlocks `json:"cidrBlocks,omitempty"`
}

type ClusterMasterGlobalAccessConfig struct {
	/* Whether the cluster master is accessible globally or not. */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterNetworkPolicy struct {
	/* Whether network policy is enabled on the cluster. */
	Enabled bool `json:"enabled,omitempty"`
	/* The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED. */
	Provider string `json:"provider,omitempty"`
}

type ClusterNetworkPolicyConfig struct {
	/*  */
	Disabled bool `json:"disabled,omitempty"`
}

type ClusterNodeConfig struct {
	/*  */
	BootDiskKMSCryptoKeyRef v1alpha1.ResourceRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`
	/* Immutable. Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. */
	DiskSizeGb int `json:"diskSizeGb,omitempty"`
	/* Immutable. Type of the disk attached to each node. */
	DiskType string `json:"diskType,omitempty"`
	/* Immutable. */
	EphemeralStorageConfig ClusterEphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`
	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	GuestAccelerator []ClusterGuestAccelerator `json:"guestAccelerator,omitempty"`
	/* The image type to use for this node. Note that for a given image type, the latest version of it will be used. */
	ImageType string `json:"imageType,omitempty"`
	/* Node kubelet configs. */
	KubeletConfig ClusterKubeletConfig `json:"kubeletConfig,omitempty"`
	/* Immutable. The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node. */
	Labels map[string]string `json:"labels,omitempty"`
	/* Parameters that can be configured on Linux nodes. */
	LinuxNodeConfig ClusterLinuxNodeConfig `json:"linuxNodeConfig,omitempty"`
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
	SandboxConfig ClusterSandboxConfig `json:"sandboxConfig,omitempty"`
	/*  */
	ServiceAccountRef v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
	/* Immutable. Shielded Instance options. */
	ShieldedInstanceConfig ClusterShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`
	/* Immutable. The list of instance tags applied to all nodes. */
	Tags []string `json:"tags,omitempty"`
	/* Immutable. List of Kubernetes taints to be applied to each node. */
	Taint []ClusterTaint `json:"taint,omitempty"`
	/* Immutable. The workload metadata configuration for this node. */
	WorkloadMetadataConfig ClusterWorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
}

type ClusterNotificationConfig struct {
	/* Notification config for Cloud Pub/Sub */
	Pubsub ClusterPubsub `json:"pubsub,omitempty"`
}

type ClusterPassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	Value string `json:"value,omitempty"`
	/* Source for the field's value. Cannot be used if 'value' is specified. */
	ValueFrom ClusterValueFrom `json:"valueFrom,omitempty"`
}

type ClusterPodSecurityPolicyConfig struct {
	/* Enable the PodSecurityPolicy controller for this cluster. If enabled, pods must be valid under a PodSecurityPolicy to be created. */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterPrivateClusterConfig struct {
	/* Immutable. Enables the private cluster feature, creating a private endpoint on the cluster. In a private cluster, nodes only have RFC 1918 private addresses and communicate with the master's private endpoint via private networking. */
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint,omitempty"`
	/* Immutable. When true, the cluster's private endpoint is used as the cluster endpoint and access through the public endpoint is disabled. When false, either endpoint can be used. This field only applies to private clusters, when enable_private_nodes is true. */
	EnablePrivateNodes bool `json:"enablePrivateNodes,omitempty"`
	/* Controls cluster master global access settings. */
	MasterGlobalAccessConfig ClusterMasterGlobalAccessConfig `json:"masterGlobalAccessConfig,omitempty"`
	/* Immutable. The IP range in CIDR notation to use for the hosted master network. This range will be used for assigning private IP addresses to the cluster master(s) and the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network, and it must be a /28 subnet. See Private Cluster Limitations for more details. This field only applies to private clusters, when enable_private_nodes is true. */
	MasterIpv4CidrBlock string `json:"masterIpv4CidrBlock,omitempty"`
	/* The name of the peering between this cluster and the Google owned VPC. */
	PeeringName string `json:"peeringName,omitempty"`
	/* The internal IP address of this cluster's master endpoint. */
	PrivateEndpoint string `json:"privateEndpoint,omitempty"`
	/* The external IP address of this cluster's master endpoint. */
	PublicEndpoint string `json:"publicEndpoint,omitempty"`
}

type ClusterPubsub struct {
	/* Whether or not the notification config is enabled */
	Enabled bool `json:"enabled,omitempty"`
	/* The PubSubTopic to send the notification to. */
	TopicRef v1alpha1.ResourceRef `json:"topicRef,omitempty"`
}

type ClusterRecurringWindow struct {
	/*  */
	EndTime string `json:"endTime,omitempty"`
	/*  */
	Recurrence string `json:"recurrence,omitempty"`
	/*  */
	StartTime string `json:"startTime,omitempty"`
}

type ClusterReleaseChannel struct {
	/* The selected release channel. Accepted values are:
	* UNSPECIFIED: Not set.
	* RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
	* REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
	* STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky. */
	Channel string `json:"channel,omitempty"`
}

type ClusterResourceLimits struct {
	/* Maximum amount of the resource in the cluster. */
	Maximum int `json:"maximum,omitempty"`
	/* Minimum amount of the resource in the cluster. */
	Minimum int `json:"minimum,omitempty"`
	/* The type of the resource. For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types. */
	ResourceType string `json:"resourceType,omitempty"`
}

type ClusterResourceUsageExportConfig struct {
	/* Parameters for using BigQuery as the destination of resource usage export. */
	BigqueryDestination ClusterBigqueryDestination `json:"bigqueryDestination,omitempty"`
	/* Whether to enable network egress metering for this cluster. If enabled, a daemonset will be created in the cluster to meter network egress traffic. */
	EnableNetworkEgressMetering bool `json:"enableNetworkEgressMetering,omitempty"`
	/* Whether to enable resource consumption metering on this cluster. When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true. */
	EnableResourceConsumptionMetering bool `json:"enableResourceConsumptionMetering,omitempty"`
}

type ClusterSandboxConfig struct {
	/* Type of the sandbox to use for the node (e.g. 'gvisor') */
	SandboxType string `json:"sandboxType,omitempty"`
}

type ClusterShieldedInstanceConfig struct {
	/* Immutable. Defines whether the instance has integrity monitoring enabled. */
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty"`
	/* Immutable. Defines whether the instance has Secure Boot enabled. */
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty"`
}

type ClusterTaint struct {
	/* Immutable. Effect for taint. */
	Effect string `json:"effect,omitempty"`
	/* Immutable. Key for taint. */
	Key string `json:"key,omitempty"`
	/* Immutable. Value for taint. */
	Value string `json:"value,omitempty"`
}

type ClusterValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	SecretKeyRef v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type ClusterVerticalPodAutoscaling struct {
	/* Enables vertical pod autoscaling. */
	Enabled bool `json:"enabled,omitempty"`
}

type ClusterWorkloadIdentityConfig struct {
	/* Enables workload identity. */
	IdentityNamespace string `json:"identityNamespace,omitempty"`
}

type ClusterWorkloadMetadataConfig struct {
	/* NodeMetadata is the configuration for how to expose metadata to the workloads running on the node. */
	NodeMetadata string `json:"nodeMetadata,omitempty"`
}

type ContainerClusterSpec struct {
	/* The configuration for addons supported by GKE. */
	AddonsConfig ClusterAddonsConfig `json:"addonsConfig,omitempty"`
	/* Immutable. Configuration for the Google Groups for GKE feature. */
	AuthenticatorGroupsConfig ClusterAuthenticatorGroupsConfig `json:"authenticatorGroupsConfig,omitempty"`
	/* Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the guide to using Node Auto-Provisioning for more details. */
	ClusterAutoscaling ClusterClusterAutoscaling `json:"clusterAutoscaling,omitempty"`
	/* Immutable. The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined. */
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty"`
	/* Telemetry integration for the cluster. */
	ClusterTelemetry ClusterClusterTelemetry `json:"clusterTelemetry,omitempty"`
	/* Immutable. Configuration for the confidential nodes feature, which makes nodes run on confidential VMs. Warning: This configuration can't be changed (or added/removed) after cluster creation without deleting and recreating the entire cluster. */
	ConfidentialNodes ClusterConfidentialNodes `json:"confidentialNodes,omitempty"`
	/* Application-layer Secrets Encryption settings. The object format is {state = string, key_name = string}. Valid values of state are: "ENCRYPTED"; "DECRYPTED". key_name is the name of a CloudKMS key. */
	DatabaseEncryption ClusterDatabaseEncryption `json:"databaseEncryption,omitempty"`
	/* The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation. */
	DatapathProvider string `json:"datapathProvider,omitempty"`
	/* Immutable. The default maximum number of pods per node in this cluster. This doesn't work on "routes-based" clusters, clusters that don't have IP Aliasing enabled. */
	DefaultMaxPodsPerNode int `json:"defaultMaxPodsPerNode,omitempty"`
	/* Whether the cluster disables default in-node sNAT rules. In-node sNAT rules will be disabled when defaultSnatStatus is disabled. */
	DefaultSnatStatus ClusterDefaultSnatStatus `json:"defaultSnatStatus,omitempty"`
	/* Immutable.  Description of the cluster. */
	Description string `json:"description,omitempty"`
	/* Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization. */
	EnableBinaryAuthorization bool `json:"enableBinaryAuthorization,omitempty"`
	/* Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network. */
	EnableIntranodeVisibility bool `json:"enableIntranodeVisibility,omitempty"`
	/* Immutable. Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days. */
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty"`
	/* Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false. */
	EnableLegacyAbac bool `json:"enableLegacyAbac,omitempty"`
	/* Enable Shielded Nodes features on all nodes in this cluster. Defaults to false. */
	EnableShieldedNodes bool `json:"enableShieldedNodes,omitempty"`
	/* Immutable. Whether to enable Cloud TPU resources in this cluster. */
	EnableTpu bool `json:"enableTpu,omitempty"`
	/* Immutable. The number of nodes to create in this cluster's default node pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true. */
	InitialNodeCount int `json:"initialNodeCount,omitempty"`
	/* Immutable. Configuration of cluster IP allocation for VPC-native clusters. Adding this block enables IP aliasing, making the cluster VPC-native instead of routes-based. */
	IpAllocationPolicy ClusterIpAllocationPolicy `json:"ipAllocationPolicy,omitempty"`
	/* Immutable. The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well. */
	Location string `json:"location,omitempty"`
	/* The logging service that the cluster should write logs to. Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes. */
	LoggingService string `json:"loggingService,omitempty"`
	/* The maintenance policy to use for the cluster. */
	MaintenancePolicy ClusterMaintenancePolicy `json:"maintenancePolicy,omitempty"`
	/* The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff removing a username/password or unsetting your client cert, ensure you have the container.clusters.getCredentials permission. */
	MasterAuth ClusterMasterAuth `json:"masterAuth,omitempty"`
	/* The desired configuration options for master authorized networks. Omit the nested cidr_blocks attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists). */
	MasterAuthorizedNetworksConfig ClusterMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`
	/* The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version). */
	MinMasterVersion string `json:"minMasterVersion,omitempty"`
	/* The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes. */
	MonitoringService string `json:"monitoringService,omitempty"`
	/* Configuration options for the NetworkPolicy feature. */
	NetworkPolicy ClusterNetworkPolicy `json:"networkPolicy,omitempty"`
	/*  */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef,omitempty"`
	/* Immutable. Determines whether alias IPs or routes will be used for pod IPs in the cluster. */
	NetworkingMode string `json:"networkingMode,omitempty"`
	/* Immutable. The configuration of the nodepool */
	NodeConfig ClusterNodeConfig `json:"nodeConfig,omitempty"`
	/* The list of zones in which the cluster's nodes are located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone. */
	NodeLocations []string `json:"nodeLocations,omitempty"`
	/*  */
	NodeVersion string `json:"nodeVersion,omitempty"`
	/* The notification config for sending cluster upgrade notifications */
	NotificationConfig ClusterNotificationConfig `json:"notificationConfig,omitempty"`
	/* Configuration for the PodSecurityPolicy feature. */
	PodSecurityPolicyConfig ClusterPodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty"`
	/* Configuration for private clusters, clusters with private nodes. */
	PrivateClusterConfig ClusterPrivateClusterConfig `json:"privateClusterConfig,omitempty"`
	/* Configuration options for the Release channel feature, which provide more control over automatic upgrades of your GKE clusters. Note that removing this field from your config will not unenroll it. Instead, use the "UNSPECIFIED" channel. */
	ReleaseChannel ClusterReleaseChannel `json:"releaseChannel,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* Configuration for the ResourceUsageExportConfig feature. */
	ResourceUsageExportConfig ClusterResourceUsageExportConfig `json:"resourceUsageExportConfig,omitempty"`
	/*  */
	SubnetworkRef v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
	/* Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it. */
	VerticalPodAutoscaling ClusterVerticalPodAutoscaling `json:"verticalPodAutoscaling,omitempty"`
	/* Configuration for the use of Kubernetes Service Accounts in GCP IAM policies. */
	WorkloadIdentityConfig ClusterWorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`
}

type ContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The IP address of this cluster's Kubernetes master. */
	Endpoint string `json:"endpoint,omitempty"`
	/* List of instance group URLs which have been assigned to the cluster. */
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`
	/* The fingerprint of the set of labels for this cluster. */
	LabelFingerprint string `json:"labelFingerprint,omitempty"`
	/* The current version of the master in the cluster. This may be different than the min_master_version set in the config if the master has been updated by GKE. */
	MasterVersion string `json:"masterVersion,omitempty"`
	/*  */
	Operation string `json:"operation,omitempty"`
	/* Server-defined URL for the resource. */
	SelfLink string `json:"selfLink,omitempty"`
	/* The IP address range of the Kubernetes services in this cluster, in CIDR notation (e.g. 1.2.3.4/29). Service addresses are typically put in the last /16 from the container CIDR. */
	ServicesIpv4Cidr string `json:"servicesIpv4Cidr,omitempty"`
	/* The IP address range of the Cloud TPUs in this cluster, in CIDR notation (e.g. 1.2.3.4/29). */
	TpuIpv4CidrBlock string `json:"tpuIpv4CidrBlock,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerCluster is the Schema for the container API
// +k8s:openapi-gen=true
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerClusterSpec   `json:"spec,omitempty"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}
