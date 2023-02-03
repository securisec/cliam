package policy

// Microsoft_Authorization_authorization_RoleDefinitionsCalls policy
var Microsoft_Authorization_authorization_RoleDefinitionsCalls = map[string]Policy{
	"Permissions_ListForResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Authorization/permissions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Permissions_ListForResourceGroup",
		Resource:    "Microsoft.Authorization",
	},
	"Permissions_ListForResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourcePath}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Authorization/permissions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Permissions_ListForResource",
		Resource:    "Microsoft.Authorization",
	},
	"RoleDefinitions_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleDefinitions/{{.roleDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleDefinitions_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleDefinitions_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleDefinitions_List",
		Resource:    "Microsoft.Authorization",
	},
}
