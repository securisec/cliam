package policy

// Microsoft_ApiManagement_apimpolicydescriptions policy
var Microsoft_ApiManagement_apimpolicydescriptions = map[string]Policy{
	"PolicyDescription_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/policyDescriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PolicyDescription_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
}
