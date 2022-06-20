package policy

    var Microsoft_HybridNetwork_networkFunction = map[string]Policy{
        "NetworkFunctions_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridNetwork/networkFunctions/{{.networkFunctionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "NetworkFunctions_Get",
    Resource:       "Microsoft.HybridNetwork",
},
"NetworkFunctions_ListBySubscription": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/networkFunctions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "NetworkFunctions_ListBySubscription",
    Resource:       "Microsoft.HybridNetwork",
},
"NetworkFunctions_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridNetwork/networkFunctions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "NetworkFunctions_ListByResourceGroup",
    Resource:       "Microsoft.HybridNetwork",
},
    }
    