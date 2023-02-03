package policy

// Microsoft_ApiManagement_apimdeletedservices policy
var Microsoft_ApiManagement_apimdeletedservices = map[string]Policy{
	"DeletedServices_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/deletedservices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DeletedServices_ListBySubscription",
		Resource:    "Microsoft.ApiManagement",
	},
	"DeletedServices_GetByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/locations/{{.location}}/deletedservices/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DeletedServices_GetByName",
		Resource:    "Microsoft.ApiManagement",
	},
}
