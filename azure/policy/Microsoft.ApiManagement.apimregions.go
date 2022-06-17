package policy

var Microsoft_ApiManagement_apimregions = map[string]Policy{
	"Region_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/regions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Region_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
}
