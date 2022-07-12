package policy

    var Microsoft_ServiceFabric_managedapplication = map[string]Policy{
        "ApplicationTypes_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applicationTypes/{{.applicationTypeName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ApplicationTypes_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"ApplicationTypes_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applicationTypes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ApplicationTypes_List",
    Resource:       "Microsoft.ServiceFabric",
},
"ApplicationTypeVersions_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applicationTypes/{{.applicationTypeName}}/versions/{{.version}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ApplicationTypeVersions_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"ApplicationTypeVersions_ListByApplicationTypes": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applicationTypes/{{.applicationTypeName}}/versions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ApplicationTypeVersions_ListByApplicationTypes",
    Resource:       "Microsoft.ServiceFabric",
},
"Applications_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applications/{{.applicationName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "Applications_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"Applications_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applications",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "Applications_List",
    Resource:       "Microsoft.ServiceFabric",
},
"Services_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applications/{{.applicationName}}/services/{{.serviceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "Services_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"Services_ListByApplications": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedclusters/{{.clusterName}}/applications/{{.applicationName}}/services",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "Services_ListByApplications",
    Resource:       "Microsoft.ServiceFabric",
},
    }
    