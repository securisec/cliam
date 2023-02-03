package policy

// Microsoft_Authorization_authorization_RoleAssignmentsCalls policy
var Microsoft_Authorization_authorization_RoleAssignmentsCalls = map[string]Policy{
	"RoleAssignments_ListForSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/roleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleAssignments_ListForSubscription",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignments_ListForResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Authorization/roleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleAssignments_ListForResourceGroup",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignments_ListForResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Authorization/roleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleAssignments_ListForResource",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignments_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignments/{{.roleAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleAssignments_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignments_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleAssignments_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignments_GetById": {
		Path:   "/{{.roleAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RoleAssignments_GetById",
		Resource:    "Microsoft.Authorization",
	},
}
