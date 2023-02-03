package policy

// Microsoft_LabServices_Skus policy
var Microsoft_LabServices_Skus = map[string]Policy{
	"Skus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LabServices/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Skus_List",
		Resource:    "Microsoft.LabServices",
	},
}
