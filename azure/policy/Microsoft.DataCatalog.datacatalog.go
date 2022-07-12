package policy

    var Microsoft_DataCatalog_datacatalog = map[string]Policy{
        "ADCOperations_List": {
    Path: "/providers/Microsoft.DataCatalog/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2016-03-30",
    },
	OperationID:    "ADCOperations_List",
    Resource:       "Microsoft.DataCatalog",
},
"ADCCatalogs_ListtByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataCatalog/catalogs",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2016-03-30",
    },
	OperationID:    "ADCCatalogs_ListtByResourceGroup",
    Resource:       "Microsoft.DataCatalog",
},
"ADCCatalogs_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataCatalog/catalogs/{{.catalogName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2016-03-30",
    },
	OperationID:    "ADCCatalogs_Get",
    Resource:       "Microsoft.DataCatalog",
},
    }
    