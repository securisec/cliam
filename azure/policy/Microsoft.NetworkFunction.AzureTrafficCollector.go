package policy

// Microsoft_NetworkFunction_AzureTrafficCollector policy
var Microsoft_NetworkFunction_AzureTrafficCollector = map[string]Policy{
	"NetworkFunction_ListOperations": {
		Path:   "/providers/Microsoft.NetworkFunction/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "NetworkFunction_ListOperations",
		Resource:    "Microsoft.NetworkFunction",
	},
	"AzureTrafficCollectorsBySubscription_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "AzureTrafficCollectorsBySubscription_List",
		Resource:    "Microsoft.NetworkFunction",
	},
	"AzureTrafficCollectorsByResourceGroup_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "AzureTrafficCollectorsByResourceGroup_List",
		Resource:    "Microsoft.NetworkFunction",
	},
	"AzureTrafficCollectors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetworkFunction/azureTrafficCollectors/{{.azureTrafficCollectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "AzureTrafficCollectors_Get",
		Resource:    "Microsoft.NetworkFunction",
	},
	"CollectorPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetworkFunction/azureTrafficCollectors/{{.azureTrafficCollectorName}}/collectorPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "CollectorPolicies_List",
		Resource:    "Microsoft.NetworkFunction",
	},
	"CollectorPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetworkFunction/azureTrafficCollectors/{{.azureTrafficCollectorName}}/collectorPolicies/{{.collectorPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "CollectorPolicies_Get",
		Resource:    "Microsoft.NetworkFunction",
	},
}
