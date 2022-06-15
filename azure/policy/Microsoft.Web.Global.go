package policy

var Microsoft_Web_Global = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deletedSites/{{.deletedSiteId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Global_GetDeletedWebApp",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deletedSites/{{.deletedSiteId}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Global_GetDeletedWebAppSnapshots",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/locations/{{.location}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Global_GetSubscriptionOperationWithAsyncResponse",
		Resource:    "Microsoft.Web",
	},
}
