package policy

// Microsoft_MobileNetwork_simGroup policy
var Microsoft_MobileNetwork_simGroup = map[string]Policy{
	"SimGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups/{{.simGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SimGroups_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"SimGroups_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MobileNetwork/simGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SimGroups_ListBySubscription",
		Resource:    "Microsoft.MobileNetwork",
	},
	"SimGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SimGroups_ListByResourceGroup",
		Resource:    "Microsoft.MobileNetwork",
	},
}
