package policy

// Microsoft_Network_serviceTags policy
var Microsoft_Network_serviceTags = map[string]Policy{
	"ServiceTags_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/serviceTags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceTags_List",
		Resource:    "Microsoft.Network",
	},
	"ServiceTagInformation_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/serviceTagDetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceTagInformation_List",
		Resource:    "Microsoft.Network",
	},
}
