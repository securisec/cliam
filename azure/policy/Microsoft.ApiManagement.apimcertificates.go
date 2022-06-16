package policy

var Microsoft_ApiManagement_apimcertificates = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/certificates",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "Certificate_ListByService",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/certificates/{{.certificateId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "Certificate_Get",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/certificates/{{.certificateId}}/refreshSecret",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "Certificate_RefreshSecret",
    Resource:       "Microsoft.ApiManagement",
},
}
