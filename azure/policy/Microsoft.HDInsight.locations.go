package policy

// Microsoft_HDInsight_locations policy
var Microsoft_HDInsight_locations = map[string]Policy{
	"Locations_GetCapabilities": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/locations/{{.location}}/capabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Locations_GetCapabilities",
		Resource:    "Microsoft.HDInsight",
	},
	"Locations_ListUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Locations_ListUsages",
		Resource:    "Microsoft.HDInsight",
	},
	"Locations_ListBillingSpecs": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/locations/{{.location}}/billingSpecs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Locations_ListBillingSpecs",
		Resource:    "Microsoft.HDInsight",
	},
	"Locations_GetAzureAsyncOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/locations/{{.location}}/azureasyncoperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Locations_GetAzureAsyncOperationStatus",
		Resource:    "Microsoft.HDInsight",
	},
	"Locations_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Locations_CheckNameAvailability",
		Resource:    "Microsoft.HDInsight",
	},
	"Locations_ValidateClusterCreateRequest": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/locations/{{.location}}/validateCreateRequest",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Locations_ValidateClusterCreateRequest",
		Resource:    "Microsoft.HDInsight",
	},
}
