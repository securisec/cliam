package policy

var Microsoft_Authorization_policyAssignments = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/policyAssignments/{{.policyAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyAssignments_Get",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyAssignments_ListForResourceGroup",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourcePath}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyAssignments_ListForResource",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyAssignments_ListForManagementGroup",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyAssignments_List",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.policyAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyAssignments_GetById",
		Resource:    "Microsoft.Authorization",
	},
}
