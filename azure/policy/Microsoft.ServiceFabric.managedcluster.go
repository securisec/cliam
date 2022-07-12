package policy

    var Microsoft_ServiceFabric_managedcluster = map[string]Policy{
        "ManagedClusters_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusters_ListByResourceGroup",
    Resource:       "Microsoft.ServiceFabric",
},
"ManagedClusters_ListBySubscription": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/managedClusters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusters_ListBySubscription",
    Resource:       "Microsoft.ServiceFabric",
},
"ManagedClusters_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusters_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"ManagedClusterVersion_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/managedClusterVersions/{{.clusterVersion}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusterVersion_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"ManagedClusterVersion_GetByEnvironment": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/environments/{{.environment}}/managedClusterVersions/{{.clusterVersion}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusterVersion_GetByEnvironment",
    Resource:       "Microsoft.ServiceFabric",
},
"ManagedClusterVersion_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/managedClusterVersions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusterVersion_List",
    Resource:       "Microsoft.ServiceFabric",
},
"ManagedClusterVersion_ListByEnvironment": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/environments/{{.environment}}/managedClusterVersions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "ManagedClusterVersion_ListByEnvironment",
    Resource:       "Microsoft.ServiceFabric",
},
"managedUnsupportedVMSizes_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/managedUnsupportedVMSizes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "managedUnsupportedVMSizes_List",
    Resource:       "Microsoft.ServiceFabric",
},
"managedUnsupportedVMSizes_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/managedUnsupportedVMSizes/{{.vmSize}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "managedUnsupportedVMSizes_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"OperationStatus_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/managedClusterOperations/{{.operationId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "OperationStatus_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"OperationResults_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/managedClusterOperationResults/{{.operationId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "OperationResults_Get",
    Resource:       "Microsoft.ServiceFabric",
},
"Operations_List": {
    Path: "/providers/Microsoft.ServiceFabric/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.ServiceFabric",
},
    }
    