package policy

// Microsoft_ApiManagement_apimtenant policy
var Microsoft_ApiManagement_apimtenant = map[string]Policy{
	"TenantAccess_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccess_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantAccess_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.accessName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccess_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantAccess_RegeneratePrimaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.accessName}}/regeneratePrimaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccess_RegeneratePrimaryKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantAccess_RegenerateSecondaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.accessName}}/regenerateSecondaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccess_RegenerateSecondaryKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantAccess_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.accessName}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccess_ListSecrets",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantAccessGit_RegeneratePrimaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.accessName}}/git/regeneratePrimaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccessGit_RegeneratePrimaryKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantAccessGit_RegenerateSecondaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.accessName}}/git/regenerateSecondaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantAccessGit_RegenerateSecondaryKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantConfiguration_Deploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.configurationName}}/deploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantConfiguration_Deploy",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantConfiguration_Save": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.configurationName}}/save",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantConfiguration_Save",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantConfiguration_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.configurationName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantConfiguration_Validate",
		Resource:    "Microsoft.ApiManagement",
	},
	"TenantConfiguration_GetSyncState": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tenant/{{.configurationName}}/syncState",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TenantConfiguration_GetSyncState",
		Resource:    "Microsoft.ApiManagement",
	},
}
