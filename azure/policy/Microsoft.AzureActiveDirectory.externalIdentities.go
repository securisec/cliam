package policy

// Microsoft_AzureActiveDirectory_externalIdentities policy
var Microsoft_AzureActiveDirectory_externalIdentities = map[string]Policy{
	"B2CTenants_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureActiveDirectory/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "B2CTenants_CheckNameAvailability",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"B2CTenants_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureActiveDirectory/b2cDirectories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "B2CTenants_ListByResourceGroup",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"B2CTenants_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureActiveDirectory/b2cDirectories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "B2CTenants_ListBySubscription",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"B2CTenants_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureActiveDirectory/b2cDirectories/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "B2CTenants_Get",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.AzureActiveDirectory/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"GuestUsages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureActiveDirectory/guestUsages/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "GuestUsages_Get",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"GuestUsages_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureActiveDirectory/guestUsages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "GuestUsages_ListBySubscription",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
	"GuestUsages_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureActiveDirectory/guestUsages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "GuestUsages_ListByResourceGroup",
		Resource:    "Microsoft.AzureActiveDirectory",
	},
}
