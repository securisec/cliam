package policy

var Microsoft_ApiManagement_apimoutbounddependency = map[string]Policy{
	"OutboundNetworkDependenciesEndpoints_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "OutboundNetworkDependenciesEndpoints_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
}
