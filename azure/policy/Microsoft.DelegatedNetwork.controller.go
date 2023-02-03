package policy

// Microsoft_DelegatedNetwork_controller policy
var Microsoft_DelegatedNetwork_controller = map[string]Policy{
	"Controller_GetDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DelegatedNetwork/controller/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "Controller_GetDetails",
		Resource:    "Microsoft.DelegatedNetwork",
	},
	"DelegatedNetwork_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DelegatedNetwork/controllers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "DelegatedNetwork_ListBySubscription",
		Resource:    "Microsoft.DelegatedNetwork",
	},
	"DelegatedNetwork_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DelegatedNetwork/controllers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "DelegatedNetwork_ListByResourceGroup",
		Resource:    "Microsoft.DelegatedNetwork",
	},
}
