package policy

// Microsoft_Network_usage policy
var Microsoft_Network_usage = map[string]Policy{
	"Usages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Usages_List",
		Resource:    "Microsoft.Network",
	},
}
