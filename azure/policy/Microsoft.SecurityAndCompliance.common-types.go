package policy

var Microsoft_SecurityAndCompliance_common_types = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.SecurityAndCompliance/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"OperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/locations/{{.locationName}}/operationresults/{{.operationResultId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "OperationResults_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
}
