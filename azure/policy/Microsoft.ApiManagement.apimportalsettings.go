package policy

// Microsoft_ApiManagement_apimportalsettings policy
var Microsoft_ApiManagement_apimportalsettings = map[string]Policy{
	"PortalSettings_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalsettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PortalSettings_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"SignInSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalsettings/signin",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "SignInSettings_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"SignUpSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalsettings/signup",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "SignUpSettings_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"DelegationSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalsettings/delegation",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DelegationSettings_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"DelegationSettings_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalsettings/delegation/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DelegationSettings_ListSecrets",
		Resource:    "Microsoft.ApiManagement",
	},
}
