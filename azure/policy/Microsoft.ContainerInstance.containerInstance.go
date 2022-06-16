package policy

var Microsoft_ContainerInstance_containerInstance = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/containerGroups",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_List",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_ListByResourceGroup",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_Get",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/restart",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_Restart",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_Stop",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_Start",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/providers/Microsoft.ContainerInstance/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/locations/{{.location}}/usages",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Location_ListUsage",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/containers/{{.containerName}}/logs",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Containers_ListLogs",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/containers/{{.containerName}}/exec",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Containers_ExecuteCommand",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/containers/{{.containerName}}/attach",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Containers_Attach",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/locations/{{.location}}/cachedImages",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Location_ListCachedImages",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/locations/{{.location}}/capabilities",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Location_ListCapabilities",
    Resource:       "Microsoft.ContainerInstance",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/outboundNetworkDependenciesEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "ContainerGroups_GetOutboundNetworkDependenciesEndpoints",
    Resource:       "Microsoft.ContainerInstance",
},
}
