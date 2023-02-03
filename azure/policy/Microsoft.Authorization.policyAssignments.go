package policy

// Microsoft_Authorization_policyAssignments policy
var Microsoft_Authorization_policyAssignments = map[string]Policy{
	"PolicyAssignments_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/policyAssignments/{{.policyAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PolicyAssignments_Get",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyAssignments_ListForResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PolicyAssignments_ListForResourceGroup",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyAssignments_ListForResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourcePath}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PolicyAssignments_ListForResource",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyAssignments_ListForManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PolicyAssignments_ListForManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyAssignments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PolicyAssignments_List",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyAssignments_GetById": {
		Path:   "/{{.policyAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PolicyAssignments_GetById",
		Resource:    "Microsoft.Authorization",
	},
}
