package policy

// Microsoft_Network_availableServiceAliases policy
var Microsoft_Network_availableServiceAliases = map[string]Policy{
	"AvailableServiceAliases_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/availableServiceAliases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailableServiceAliases_List",
		Resource:    "Microsoft.Network",
	},
	"AvailableServiceAliases_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/locations/{{.location}}/availableServiceAliases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailableServiceAliases_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
}
