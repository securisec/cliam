package policy

var Microsoft_Solutions_managedapplications = []Policy{
	{
		Path:   "/providers/Microsoft.Solutions/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "ListOperations",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applications/{{.applicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_Get",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applicationDefinitions/{{.applicationDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "ApplicationDefinitions_Get",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applicationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "ApplicationDefinitions_ListByResourceGroup",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_ListByResourceGroup",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Solutions/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_ListBySubscription",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/{{.applicationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_GetById",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/jitRequests/{{.jitRequestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "JitRequests_Get",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Solutions/jitRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "jitRequests_ListBySubscription",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/jitRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "jitRequests_ListByResourceGroup",
		Resource:    "Microsoft.Solutions",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applications/{{.applicationName}}/refreshPermissions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_RefreshPermissions",
		Resource:    "Microsoft.Solutions",
	},
}
