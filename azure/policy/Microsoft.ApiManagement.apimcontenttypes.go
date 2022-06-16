package policy

var Microsoft_ApiManagement_apimcontenttypes = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/contentTypes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "ContentType_ListByService",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/contentTypes/{{.contentTypeId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "ContentType_Get",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/contentTypes/{{.contentTypeId}}/contentItems",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "ContentItem_ListByService",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/contentTypes/{{.contentTypeId}}/contentItems/{{.contentItemId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "ContentItem_Get",
    Resource:       "Microsoft.ApiManagement",
},
}
