package policy

// Microsoft_DelegatedNetwork_delegatedSubnets policy
var Microsoft_DelegatedNetwork_delegatedSubnets = map[string]Policy{
	"DelegatedSubnetService_GetDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DelegatedNetwork/delegatedSubnets/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "DelegatedSubnetService_GetDetails",
		Resource:    "Microsoft.DelegatedNetwork",
	},
	"DelegatedSubnetService_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DelegatedNetwork/delegatedSubnets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "DelegatedSubnetService_ListBySubscription",
		Resource:    "Microsoft.DelegatedNetwork",
	},
	"DelegatedSubnetService_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DelegatedNetwork/delegatedSubnets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "DelegatedSubnetService_ListByResourceGroup",
		Resource:    "Microsoft.DelegatedNetwork",
	},
}
