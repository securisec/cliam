package policy

var Microsoft_ApiManagement_apimnamedvalues = map[string]Policy{
	"NamedValue_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/namedValues",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NamedValue_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"NamedValue_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/namedValues/{{.namedValueId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NamedValue_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"NamedValue_ListValue": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/namedValues/{{.namedValueId}}/listValue",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NamedValue_ListValue",
		Resource:    "Microsoft.ApiManagement",
	},
	"NamedValue_RefreshSecret": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/namedValues/{{.namedValueId}}/refreshSecret",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NamedValue_RefreshSecret",
		Resource:    "Microsoft.ApiManagement",
	},
}
