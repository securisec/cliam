package policy

    var Microsoft_VMwareCloudSimple_vmwarecloudsimple = []Policy{
        {
    Path: "/providers/Microsoft.VMwareCloudSimple/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "DedicatedCloudNodes_ListBySubscription",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "DedicatedCloudServices_ListBySubscription",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/availabilities",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "SkusAvailability_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/operationResults/{{.operationId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Operations_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "PrivateClouds_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "PrivateClouds_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/customizationPolicies",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "customizationPolicies_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/customizationPolicies/{{.customizationPolicyName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "customizationPolicies_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/resourcePools",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "ResourcePools_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/resourcePools/{{.resourcePoolName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "ResourcePools_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualMachineTemplates",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachineTemplates_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualMachineTemplates/{{.virtualMachineTemplateName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachineTemplates_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualNetworks",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualNetworks_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/privateClouds/{{.pcName}}/virtualNetworks/{{.virtualNetworkName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualNetworks_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/locations/{{.regionId}}/usages",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "Usages_List",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VMwareCloudSimple/virtualMachines",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachines_ListBySubscription",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "DedicatedCloudNodes_ListByResourceGroup",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{{.dedicatedCloudNodeName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "DedicatedCloudNodes_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "DedicatedCloudServices_ListByResourceGroup",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{{.dedicatedCloudServiceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "DedicatedCloudServices_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachines_ListByResourceGroup",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{{.virtualMachineName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachines_Get",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{{.virtualMachineName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachines_Start",
    Resource:       "Microsoft.VMwareCloudSimple",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{{.virtualMachineName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-04-01",
    },
	OperationID:    "VirtualMachines_Stop",
    Resource:       "Microsoft.VMwareCloudSimple",
},
    }
    