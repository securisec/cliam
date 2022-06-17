package policy

var Microsoft_Web_DeletedWebApps = map[string]Policy{
	"DeletedWebApps_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deletedSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "DeletedWebApps_List",
		Resource:    "Microsoft.Web",
	},
	"DeletedWebApps_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/locations/{{.location}}/deletedSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "DeletedWebApps_ListByLocation",
		Resource:    "Microsoft.Web",
	},
	"DeletedWebApps_GetDeletedWebAppByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/locations/{{.location}}/deletedSites/{{.deletedSiteId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "DeletedWebApps_GetDeletedWebAppByLocation",
		Resource:    "Microsoft.Web",
	},
}
