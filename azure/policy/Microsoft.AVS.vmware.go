package policy

// Microsoft_AVS_vmware policy
var Microsoft_AVS_vmware = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.AVS/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AVS",
	},
	"Locations_CheckTrialAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AVS/locations/{{.location}}/checkTrialAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Locations_CheckTrialAvailability",
		Resource:    "Microsoft.AVS",
	},
	"Locations_CheckQuotaAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AVS/locations/{{.location}}/checkQuotaAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Locations_CheckQuotaAvailability",
		Resource:    "Microsoft.AVS",
	},
	"PrivateClouds_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateClouds_List",
		Resource:    "Microsoft.AVS",
	},
	"PrivateClouds_ListInSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AVS/privateClouds",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateClouds_ListInSubscription",
		Resource:    "Microsoft.AVS",
	},
	"PrivateClouds_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateClouds_Get",
		Resource:    "Microsoft.AVS",
	},
	"PrivateClouds_RotateVcenterPassword": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/rotateVcenterPassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateClouds_RotateVcenterPassword",
		Resource:    "Microsoft.AVS",
	},
	"PrivateClouds_RotateNsxtPassword": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/rotateNsxtPassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateClouds_RotateNsxtPassword",
		Resource:    "Microsoft.AVS",
	},
	"Clusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Clusters_List",
		Resource:    "Microsoft.AVS",
	},
	"Clusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Clusters_Get",
		Resource:    "Microsoft.AVS",
	},
	"Clusters_ListZones": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/listZones",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Clusters_ListZones",
		Resource:    "Microsoft.AVS",
	},
	"Datastores_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/datastores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Datastores_List",
		Resource:    "Microsoft.AVS",
	},
	"Datastores_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/datastores/{{.datastoreName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Datastores_Get",
		Resource:    "Microsoft.AVS",
	},
	"PrivateClouds_ListAdminCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/listAdminCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateClouds_ListAdminCredentials",
		Resource:    "Microsoft.AVS",
	},
	"HcxEnterpriseSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/hcxEnterpriseSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "HcxEnterpriseSites_List",
		Resource:    "Microsoft.AVS",
	},
	"HcxEnterpriseSites_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/hcxEnterpriseSites/{{.hcxEnterpriseSiteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "HcxEnterpriseSites_Get",
		Resource:    "Microsoft.AVS",
	},
	"Authorizations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/authorizations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Authorizations_List",
		Resource:    "Microsoft.AVS",
	},
	"Authorizations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/authorizations/{{.authorizationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Authorizations_Get",
		Resource:    "Microsoft.AVS",
	},
	"GlobalReachConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/globalReachConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "GlobalReachConnections_List",
		Resource:    "Microsoft.AVS",
	},
	"GlobalReachConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/globalReachConnections/{{.globalReachConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "GlobalReachConnections_Get",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/{{.workloadNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_Get",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_List",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListSegments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/segments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListSegments",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetSegment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/segments/{{.segmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetSegment",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListDhcp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/dhcpConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListDhcp",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetDhcp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/dhcpConfigurations/{{.dhcpId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetDhcp",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListGateways": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/gateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListGateways",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/gateways/{{.gatewayId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetGateway",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListPortMirroring": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/portMirroringProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListPortMirroring",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetPortMirroring": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/portMirroringProfiles/{{.portMirroringId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetPortMirroring",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListVMGroups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/vmGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListVMGroups",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetVMGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/vmGroups/{{.vmGroupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetVMGroup",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListVirtualMachines": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListVirtualMachines",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetVirtualMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/virtualMachines/{{.virtualMachineId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetVirtualMachine",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListDnsServices": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/dnsServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListDnsServices",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetDnsService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/dnsServices/{{.dnsServiceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetDnsService",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListDnsZones": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/dnsZones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListDnsZones",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetDnsZone": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/dnsZones/{{.dnsZoneId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetDnsZone",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_ListPublicIPs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/publicIPs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_ListPublicIPs",
		Resource:    "Microsoft.AVS",
	},
	"WorkloadNetworks_GetPublicIP": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/workloadNetworks/default/publicIPs/{{.publicIPId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkloadNetworks_GetPublicIP",
		Resource:    "Microsoft.AVS",
	},
	"CloudLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/cloudLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "CloudLinks_List",
		Resource:    "Microsoft.AVS",
	},
	"CloudLinks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/cloudLinks/{{.cloudLinkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "CloudLinks_Get",
		Resource:    "Microsoft.AVS",
	},
	"Addons_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/addons",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Addons_List",
		Resource:    "Microsoft.AVS",
	},
	"Addons_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/addons/{{.addonName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Addons_Get",
		Resource:    "Microsoft.AVS",
	},
	"VirtualMachines_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "VirtualMachines_List",
		Resource:    "Microsoft.AVS",
	},
	"VirtualMachines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/virtualMachines/{{.virtualMachineId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "VirtualMachines_Get",
		Resource:    "Microsoft.AVS",
	},
	"VirtualMachines_RestrictMovement": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/virtualMachines/{{.virtualMachineId}}/restrictMovement",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "VirtualMachines_RestrictMovement",
		Resource:    "Microsoft.AVS",
	},
	"PlacementPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/placementPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PlacementPolicies_List",
		Resource:    "Microsoft.AVS",
	},
	"PlacementPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/clusters/{{.clusterName}}/placementPolicies/{{.placementPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PlacementPolicies_Get",
		Resource:    "Microsoft.AVS",
	},
	"ScriptPackages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptPackages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptPackages_List",
		Resource:    "Microsoft.AVS",
	},
	"ScriptPackages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptPackages/{{.scriptPackageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptPackages_Get",
		Resource:    "Microsoft.AVS",
	},
	"ScriptCmdlets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptPackages/{{.scriptPackageName}}/scriptCmdlets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptCmdlets_List",
		Resource:    "Microsoft.AVS",
	},
	"ScriptCmdlets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptPackages/{{.scriptPackageName}}/scriptCmdlets/{{.scriptCmdletName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptCmdlets_Get",
		Resource:    "Microsoft.AVS",
	},
	"ScriptExecutions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptExecutions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptExecutions_List",
		Resource:    "Microsoft.AVS",
	},
	"ScriptExecutions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptExecutions/{{.scriptExecutionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptExecutions_Get",
		Resource:    "Microsoft.AVS",
	},
	"ScriptExecutions_GetExecutionLogs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AVS/privateClouds/{{.privateCloudName}}/scriptExecutions/{{.scriptExecutionName}}/getExecutionLogs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ScriptExecutions_GetExecutionLogs",
		Resource:    "Microsoft.AVS",
	},
}
