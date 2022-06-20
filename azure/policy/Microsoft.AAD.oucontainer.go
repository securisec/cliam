package policy

    var Microsoft_AAD_oucontainer = map[string]Policy{
        "OuContainerOperations_List": {
    Path: "/providers/Microsoft.Aad/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "OuContainerOperations_List",
    Resource:       "Microsoft.AAD",
},
"OuContainer_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Aad/domainServices/{{.domainServiceName}}/ouContainer",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "OuContainer_List",
    Resource:       "Microsoft.AAD",
},
"OuContainer_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Aad/domainServices/{{.domainServiceName}}/ouContainer/{{.ouContainerName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "OuContainer_Get",
    Resource:       "Microsoft.AAD",
},
    }
    