package policy

var Microsoft_ContainerService_containerService = map[string]Policy{
	"ContainerServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/containerServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-01-31",
		},
		OperationID: "ContainerServices_List",
		Resource:    "Microsoft.ContainerService",
	},
	"ContainerServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/containerServices/{{.containerServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-01-31",
		},
		OperationID: "ContainerServices_Get",
		Resource:    "Microsoft.ContainerService",
	},
	"ContainerServices_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/containerServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-01-31",
		},
		OperationID: "ContainerServices_ListByResourceGroup",
		Resource:    "Microsoft.ContainerService",
	},
}
