package policy

// Microsoft_Network_availableDelegations policy
var Microsoft_Network_availableDelegations = map[string]Policy{
	"AvailableDelegations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/availableDelegations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailableDelegations_List",
		Resource:    "Microsoft.Network",
	},
	"AvailableResourceGroupDelegations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/locations/{{.location}}/availableDelegations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailableResourceGroupDelegations_List",
		Resource:    "Microsoft.Network",
	},
}
