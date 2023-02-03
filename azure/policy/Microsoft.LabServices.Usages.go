package policy

// Microsoft_LabServices_Usages policy
var Microsoft_LabServices_Usages = map[string]Policy{
	"Usages_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LabServices/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Usages_ListByLocation",
		Resource:    "Microsoft.LabServices",
	},
}
