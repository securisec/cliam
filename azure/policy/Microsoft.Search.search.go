package policy

// Microsoft_Search_search policy
var Microsoft_Search_search = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Search/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Search",
	},
	"AdminKeys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/listAdminKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "AdminKeys_Get",
		Resource:    "Microsoft.Search",
	},
	"AdminKeys_Regenerate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/regenerateAdminKey/{{.keyKind}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "AdminKeys_Regenerate",
		Resource:    "Microsoft.Search",
	},
	"QueryKeys_Create": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/createQueryKey/{{.name}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "QueryKeys_Create",
		Resource:    "Microsoft.Search",
	},
	"QueryKeys_ListBySearchService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/listQueryKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "QueryKeys_ListBySearchService",
		Resource:    "Microsoft.Search",
	},
	"Services_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.Search",
	},
	"Services_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Services_ListByResourceGroup",
		Resource:    "Microsoft.Search",
	},
	"PrivateLinkResources_ListSupported": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateLinkResources_ListSupported",
		Resource:    "Microsoft.Search",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Search",
	},
	"PrivateEndpointConnections_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateEndpointConnections_ListByService",
		Resource:    "Microsoft.Search",
	},
	"SharedPrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/sharedPrivateLinkResources/{{.sharedPrivateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "SharedPrivateLinkResources_Get",
		Resource:    "Microsoft.Search",
	},
	"SharedPrivateLinkResources_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Search/searchServices/{{.searchServiceName}}/sharedPrivateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "SharedPrivateLinkResources_ListByService",
		Resource:    "Microsoft.Search",
	},
	"Services_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Search/searchServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Services_ListBySubscription",
		Resource:    "Microsoft.Search",
	},
	"Services_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Search/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Services_CheckNameAvailability",
		Resource:    "Microsoft.Search",
	},
}
