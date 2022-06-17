package policy

var Microsoft_ApiManagement_apimnetworkstatus = map[string]Policy{
	"NetworkStatus_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/networkstatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NetworkStatus_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"NetworkStatus_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/locations/{{.locationName}}/networkstatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NetworkStatus_ListByLocation",
		Resource:    "Microsoft.ApiManagement",
	},
}
