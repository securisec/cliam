package policy

var Microsoft_ApiManagement_apimsettings = map[string]Policy{
	"TenantSettings_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/settings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantSettings_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/settings/{{.settingsType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantSettings_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
