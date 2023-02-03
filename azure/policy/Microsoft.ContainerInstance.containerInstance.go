package policy

// Microsoft_ContainerInstance_containerInstance policy
var Microsoft_ContainerInstance_containerInstance = map[string]Policy{
	"ContainerGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/containerGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_List",
		Resource:    "Microsoft.ContainerInstance",
	},
	"ContainerGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_ListByResourceGroup",
		Resource:    "Microsoft.ContainerInstance",
	},
	"ContainerGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_Get",
		Resource:    "Microsoft.ContainerInstance",
	},
	"ContainerGroups_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_Restart",
		Resource:    "Microsoft.ContainerInstance",
	},
	"ContainerGroups_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_Stop",
		Resource:    "Microsoft.ContainerInstance",
	},
	"ContainerGroups_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_Start",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ContainerInstance/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Location_ListUsage": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Location_ListUsage",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Containers_ListLogs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/containers/{{.containerName}}/logs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Containers_ListLogs",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Containers_ExecuteCommand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/containers/{{.containerName}}/exec",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Containers_ExecuteCommand",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Containers_Attach": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/containers/{{.containerName}}/attach",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Containers_Attach",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Location_ListCachedImages": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/locations/{{.location}}/cachedImages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Location_ListCachedImages",
		Resource:    "Microsoft.ContainerInstance",
	},
	"Location_ListCapabilities": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/locations/{{.location}}/capabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Location_ListCapabilities",
		Resource:    "Microsoft.ContainerInstance",
	},
	"ContainerGroups_GetOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ContainerGroups_GetOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.ContainerInstance",
	},
}
