package policy

// Microsoft_OffAzure_migrate policy
var Microsoft_OffAzure_migrate = map[string]Policy{
	"HyperVCluster_GetCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVCluster_GetCluster",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVCluster_GetAllClustersInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVCluster_GetAllClustersInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVHost_GetHost": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/hosts/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVHost_GetHost",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVHost_GetAllHostsInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/hosts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVHost_GetAllHostsInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVJobs_GetJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/jobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVJobs_GetJob",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVJobs_GetAllJobsInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVJobs_GetAllJobsInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVMachines_GetMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/machines/{{.machineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVMachines_GetMachine",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVMachines_GetAllMachinesInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/machines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVMachines_GetAllMachinesInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVOperationsStatus_GetOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/operationsStatus/{{.operationStatusName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVOperationsStatus_GetOperationStatus",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVRunAsAccounts_GetRunAsAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/runAsAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVRunAsAccounts_GetRunAsAccount",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVRunAsAccounts_GetAllRunAsAccountsInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/runAsAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVRunAsAccounts_GetAllRunAsAccountsInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVSites_GetSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVSites_GetSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVSites_RefreshSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/refresh",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVSites_RefreshSite",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVSites_GetSiteHealthSummary": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/healthSummary",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVSites_GetSiteHealthSummary",
		Resource:    "Microsoft.OffAzure",
	},
	"Jobs_GetJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/jobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Jobs_GetJob",
		Resource:    "Microsoft.OffAzure",
	},
	"Jobs_GetAllJobsInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Jobs_GetAllJobsInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"Machines_GetMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/machines/{{.machineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Machines_GetMachine",
		Resource:    "Microsoft.OffAzure",
	},
	"Machines_GetAllMachinesInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/machines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Machines_GetAllMachinesInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"Machines_StopMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/machines/{{.machineName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Machines_StopMachine",
		Resource:    "Microsoft.OffAzure",
	},
	"Machines_StartMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/machines/{{.machineName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Machines_StartMachine",
		Resource:    "Microsoft.OffAzure",
	},
	"RunAsAccounts_GetRunAsAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/runAsAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "RunAsAccounts_GetRunAsAccount",
		Resource:    "Microsoft.OffAzure",
	},
	"RunAsAccounts_GetAllRunAsAccountsInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/runAsAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "RunAsAccounts_GetAllRunAsAccountsInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"Sites_GetSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Sites_GetSite",
		Resource:    "Microsoft.OffAzure",
	},
	"Sites_RefreshSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/refresh",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Sites_RefreshSite",
		Resource:    "Microsoft.OffAzure",
	},
	"Sites_GetSiteHealthSummary": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/healthSummary",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Sites_GetSiteHealthSummary",
		Resource:    "Microsoft.OffAzure",
	},
	"VCenter_GetVCenter": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/vCenters/{{.vcenterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "VCenter_GetVCenter",
		Resource:    "Microsoft.OffAzure",
	},
	"VCenter_GetAllVCentersInSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/vCenters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "VCenter_GetAllVCentersInSite",
		Resource:    "Microsoft.OffAzure",
	},
	"VMwareOperationsStatus_GetOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/operationsStatus/{{.operationStatusName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "VMwareOperationsStatus_GetOperationStatus",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVSites_GetSiteUsage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites/{{.siteName}}/summary",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVSites_GetSiteUsage",
		Resource:    "Microsoft.OffAzure",
	},
	"Sites_GetSiteUsage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites/{{.siteName}}/summary",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Sites_GetSiteUsage",
		Resource:    "Microsoft.OffAzure",
	},
	"MasterSites_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.OffAzure/MasterSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "MasterSites_ListBySubscription",
		Resource:    "Microsoft.OffAzure",
	},
	"MasterSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/MasterSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "MasterSites_List",
		Resource:    "Microsoft.OffAzure",
	},
	"MasterSites_GetSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/MasterSites/{{.siteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "MasterSites_GetSite",
		Resource:    "Microsoft.OffAzure",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.OffAzure/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.OffAzure",
	},
	"PrivateEndpointConnection_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/masterSites/{{.siteName}}/privateEndpointConnections/{{.peConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "PrivateEndpointConnection_GetPrivateEndpointConnection",
		Resource:    "Microsoft.OffAzure",
	},
	"PrivateEndpointConnection_GetPrivateEndpointConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/masterSites/{{.siteName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "PrivateEndpointConnection_GetPrivateEndpointConnections",
		Resource:    "Microsoft.OffAzure",
	},
	"PrivateLinkResources_GetPrivateLinkResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/masterSites/{{.siteName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "PrivateLinkResources_GetPrivateLinkResource",
		Resource:    "Microsoft.OffAzure",
	},
	"PrivateLinkResources_GetPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/masterSites/{{.siteName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "PrivateLinkResources_GetPrivateLinkResources",
		Resource:    "Microsoft.OffAzure",
	},
	"VMwareSites_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.OffAzure/VMwareSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "VMwareSites_ListBySubscription",
		Resource:    "Microsoft.OffAzure",
	},
	"VMwareSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/VMwareSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "VMwareSites_List",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVSites_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.OffAzure/HyperVSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVSites_ListBySubscription",
		Resource:    "Microsoft.OffAzure",
	},
	"HyperVSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.OffAzure/HyperVSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-07",
		},
		OperationID: "HyperVSites_List",
		Resource:    "Microsoft.OffAzure",
	},
}
