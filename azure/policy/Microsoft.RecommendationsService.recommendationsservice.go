package policy

// Microsoft_RecommendationsService_recommendationsservice policy
var Microsoft_RecommendationsService_recommendationsservice = map[string]Policy{
	"Accounts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RecommendationsService/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Accounts_CheckNameAvailability",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Accounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RecommendationsService/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Accounts_ListBySubscription",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Accounts_GetStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts/{{.accountName}}/status",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Accounts_GetStatus",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Modeling_ListByAccountResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts/{{.accountName}}/modeling",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Modeling_ListByAccountResource",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Modeling_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts/{{.accountName}}/modeling/{{.modelingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Modeling_Get",
		Resource:    "Microsoft.RecommendationsService",
	},
	"ServiceEndpoints_ListByAccountResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts/{{.accountName}}/serviceEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "ServiceEndpoints_ListByAccountResource",
		Resource:    "Microsoft.RecommendationsService",
	},
	"ServiceEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecommendationsService/accounts/{{.accountName}}/serviceEndpoints/{{.serviceEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "ServiceEndpoints_Get",
		Resource:    "Microsoft.RecommendationsService",
	},
	"OperationStatuses_Get": {
		Path:   "/providers/Microsoft.RecommendationsService/locations/{{.location}}/operationStatuses/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "OperationStatuses_Get",
		Resource:    "Microsoft.RecommendationsService",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.RecommendationsService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.RecommendationsService",
	},
}
