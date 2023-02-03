package policy

// Microsoft_Authorization_authorization_DenyAssignmentCalls policy
var Microsoft_Authorization_authorization_DenyAssignmentCalls = map[string]Policy{
	"DenyAssignments_ListForResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourcePath}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Authorization/denyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "DenyAssignments_ListForResource",
		Resource:    "Microsoft.Authorization",
	},
	"DenyAssignments_ListForResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Authorization/denyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "DenyAssignments_ListForResourceGroup",
		Resource:    "Microsoft.Authorization",
	},
	"DenyAssignments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/denyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "DenyAssignments_List",
		Resource:    "Microsoft.Authorization",
	},
	"DenyAssignments_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/denyAssignments/{{.denyAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "DenyAssignments_Get",
		Resource:    "Microsoft.Authorization",
	},
	"DenyAssignments_GetById": {
		Path:   "/{{.denyAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "DenyAssignments_GetById",
		Resource:    "Microsoft.Authorization",
	},
	"DenyAssignments_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/denyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "DenyAssignments_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
}
