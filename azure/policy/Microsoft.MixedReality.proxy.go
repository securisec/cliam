package policy

// Microsoft_MixedReality_proxy policy
var Microsoft_MixedReality_proxy = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.MixedReality/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.MixedReality",
	},
	"CheckNameAvailabilityLocal": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MixedReality/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "CheckNameAvailabilityLocal",
		Resource:    "Microsoft.MixedReality",
	},
}
