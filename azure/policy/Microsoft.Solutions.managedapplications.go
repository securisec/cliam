package policy

var Microsoft_Solutions_managedapplications = map[string]Policy{
	"ListOperations": {
		Path:   "/providers/Microsoft.Solutions/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "ListOperations",
		Resource:    "Microsoft.Solutions",
	},
	"Applications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applications/{{.applicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_Get",
		Resource:    "Microsoft.Solutions",
	},
	"ApplicationDefinitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applicationDefinitions/{{.applicationDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "ApplicationDefinitions_Get",
		Resource:    "Microsoft.Solutions",
	},
	"ApplicationDefinitions_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applicationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "ApplicationDefinitions_ListByResourceGroup",
		Resource:    "Microsoft.Solutions",
	},
	"Applications_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_ListByResourceGroup",
		Resource:    "Microsoft.Solutions",
	},
	"Applications_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Solutions/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_ListBySubscription",
		Resource:    "Microsoft.Solutions",
	},
	"Applications_GetById": {
		Path:   "/{{.applicationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_GetById",
		Resource:    "Microsoft.Solutions",
	},
	"JitRequests_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/jitRequests/{{.jitRequestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "JitRequests_Get",
		Resource:    "Microsoft.Solutions",
	},
	"jitRequests_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Solutions/jitRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "jitRequests_ListBySubscription",
		Resource:    "Microsoft.Solutions",
	},
	"jitRequests_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/jitRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "jitRequests_ListByResourceGroup",
		Resource:    "Microsoft.Solutions",
	},
	"Applications_RefreshPermissions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Solutions/applications/{{.applicationName}}/refreshPermissions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-07-01",
		},
		OperationID: "Applications_RefreshPermissions",
		Resource:    "Microsoft.Solutions",
	},
}
