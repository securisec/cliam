package policy

var Microsoft_ApiManagement_apimreports = map[string]Policy{
	"Reports_ListByApi": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byApi",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByApi",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListByUser": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byUser",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByUser",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListByOperation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byOperation",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByOperation",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListByProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byProduct",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByProduct",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListByGeo": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byGeo",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByGeo",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/bySubscription",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListBySubscription",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListByTime": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byTime",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByTime",
		Resource:    "Microsoft.ApiManagement",
	},
	"Reports_ListByRequest": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/reports/byRequest",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Reports_ListByRequest",
		Resource:    "Microsoft.ApiManagement",
	},
}
