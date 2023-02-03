package policy

// Microsoft_ApiManagement_apimpolicies policy
var Microsoft_ApiManagement_apimpolicies = map[string]Policy{
	"Policy_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/policies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Policy_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Policy_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/policies/{{.policyId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Policy_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
