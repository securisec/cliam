package policy

var Microsoft_VMwareCloudSimple_vmwarecloudsimple = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.VMwareCloudSimple/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"DedicatedCloudNodes_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "DedicatedCloudNodes_ListBySubscription",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"DedicatedCloudServices_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "DedicatedCloudServices_ListBySubscription",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"SkusAvailability_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/availabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "SkusAvailability_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"Operations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Operations_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"PrivateClouds_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "PrivateClouds_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"PrivateClouds_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "PrivateClouds_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"customizationPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/customizationPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "customizationPolicies_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"customizationPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/customizationPolicies/{{.customizationPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "customizationPolicies_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"ResourcePools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/resourcePools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "ResourcePools_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"ResourcePools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/resourcePools/{{.resourcePoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "ResourcePools_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachineTemplates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualMachineTemplates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachineTemplates_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachineTemplates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualMachineTemplates/{{.virtualMachineTemplateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachineTemplates_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualNetworks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualNetworks_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualNetworks/{{.virtualNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualNetworks_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"Usages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Usages_List",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachines_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachines_ListBySubscription",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"DedicatedCloudNodes_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "DedicatedCloudNodes_ListByResourceGroup",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"DedicatedCloudNodes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{{.dedicatedCloudNodeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "DedicatedCloudNodes_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"DedicatedCloudServices_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "DedicatedCloudServices_ListByResourceGroup",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"DedicatedCloudServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{{.dedicatedCloudServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "DedicatedCloudServices_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachines_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachines_ListByResourceGroup",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{{.virtualMachineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachines_Get",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachines_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{{.virtualMachineName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachines_Start",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
	"VirtualMachines_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{{.virtualMachineName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "VirtualMachines_Stop",
		Resource:    "Microsoft.VMwareCloudSimple",
	},
}
