package policy

// Microsoft_LoadTestService_loadtestservice policy
var Microsoft_LoadTestService_loadtestservice = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.LoadTestService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.LoadTestService",
	},
	"Quotas_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LoadTestService/locations/{{.location}}/quotas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Quotas_List",
		Resource:    "Microsoft.LoadTestService",
	},
	"Quotas_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LoadTestService/locations/{{.location}}/quotas/{{.quotaBucketName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Quotas_Get",
		Resource:    "Microsoft.LoadTestService",
	},
	"Quotas_CheckAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LoadTestService/locations/{{.location}}/quotas/{{.quotaBucketName}}/checkAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Quotas_CheckAvailability",
		Resource:    "Microsoft.LoadTestService",
	},
	"LoadTests_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LoadTestService/loadTests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "LoadTests_ListBySubscription",
		Resource:    "Microsoft.LoadTestService",
	},
	"LoadTests_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LoadTestService/loadTests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "LoadTests_ListByResourceGroup",
		Resource:    "Microsoft.LoadTestService",
	},
	"LoadTests_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LoadTestService/loadTests/{{.loadTestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "LoadTests_Get",
		Resource:    "Microsoft.LoadTestService",
	},
	"LoadTests_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LoadTestService/loadTests/{{.loadTestName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "LoadTests_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.LoadTestService",
	},
}
