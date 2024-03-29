package policy

// Microsoft_ApiManagement_apimconnectivitycheck policy
var Microsoft_ApiManagement_apimconnectivitycheck = map[string]Policy{
	"PerformConnectivityCheckAsync": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/connectivityCheck",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PerformConnectivityCheckAsync",
		Resource:    "Microsoft.ApiManagement",
	},
}
