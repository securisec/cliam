package policy

// Microsoft_ApiManagement_apimportalrevisions policy
var Microsoft_ApiManagement_apimportalrevisions = map[string]Policy{
	"PortalRevision_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalRevisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PortalRevision_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"PortalRevision_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/portalRevisions/{{.portalRevisionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PortalRevision_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
