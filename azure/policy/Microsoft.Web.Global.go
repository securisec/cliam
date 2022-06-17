package policy

var Microsoft_Web_Global = map[string]Policy{
	"Global_GetDeletedWebApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deletedSites/{{.deletedSiteId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Global_GetDeletedWebApp",
		Resource:    "Microsoft.Web",
	},
	"Global_GetDeletedWebAppSnapshots": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deletedSites/{{.deletedSiteId}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Global_GetDeletedWebAppSnapshots",
		Resource:    "Microsoft.Web",
	},
	"Global_GetSubscriptionOperationWithAsyncResponse": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/locations/{{.location}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Global_GetSubscriptionOperationWithAsyncResponse",
		Resource:    "Microsoft.Web",
	},
}
