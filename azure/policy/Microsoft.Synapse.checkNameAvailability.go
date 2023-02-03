package policy

// Microsoft_Synapse_checkNameAvailability policy
var Microsoft_Synapse_checkNameAvailability = map[string]Policy{
	"Operations_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Synapse/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_CheckNameAvailability",
		Resource:    "Microsoft.Synapse",
	},
}
