package policy

// Microsoft_VideoIndexer_vi policy
var Microsoft_VideoIndexer_vi = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.VideoIndexer/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.VideoIndexer",
	},
	"Accounts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VideoIndexer/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Accounts_CheckNameAvailability",
		Resource:    "Microsoft.VideoIndexer",
	},
	"Accounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VideoIndexer/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Accounts_List",
		Resource:    "Microsoft.VideoIndexer",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VideoIndexer/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.VideoIndexer",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VideoIndexer/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.VideoIndexer",
	},
	"UserClassicAccounts_List": {
		Path:   "/providers/Microsoft.VideoIndexer/locations/{{.location}}/userClassicAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "UserClassicAccounts_List",
		Resource:    "Microsoft.VideoIndexer",
	},
	"ClassicAccounts_GetDetails": {
		Path:   "/providers/Microsoft.VideoIndexer/locations/{{.location}}/classicAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "ClassicAccounts_GetDetails",
		Resource:    "Microsoft.VideoIndexer",
	},
	"Generate_AccessToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VideoIndexer/accounts/{{.accountName}}/generateAccessToken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Generate_AccessToken",
		Resource:    "Microsoft.VideoIndexer",
	},
}
