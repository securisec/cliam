package policy

// Microsoft_Maps_maps_management policy
var Microsoft_Maps_maps_management = map[string]Policy{
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maps/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.Maps",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maps/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.Maps",
	},
	"Accounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Maps/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Accounts_ListBySubscription",
		Resource:    "Microsoft.Maps",
	},
	"Accounts_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maps/accounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Accounts_ListKeys",
		Resource:    "Microsoft.Maps",
	},
	"Accounts_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maps/accounts/{{.accountName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Accounts_RegenerateKeys",
		Resource:    "Microsoft.Maps",
	},
	"Maps_ListOperations": {
		Path:   "/providers/Microsoft.Maps/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Maps_ListOperations",
		Resource:    "Microsoft.Maps",
	},
	"Maps_ListSubscriptionOperations": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Maps/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Maps_ListSubscriptionOperations",
		Resource:    "Microsoft.Maps",
	},
	"Creators_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maps/accounts/{{.accountName}}/creators",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Creators_ListByAccount",
		Resource:    "Microsoft.Maps",
	},
	"Creators_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maps/accounts/{{.accountName}}/creators/{{.creatorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-02-01",
		},
		OperationID: "Creators_Get",
		Resource:    "Microsoft.Maps",
	},
}
