package policy

var Microsoft_ServiceFabric_application = map[string]Policy{
	"ApplicationTypes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applicationTypes/{{.applicationTypeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ApplicationTypes_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ApplicationTypes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applicationTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ApplicationTypes_List",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ApplicationTypeVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applicationTypes/{{.applicationTypeName}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ApplicationTypeVersions_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ApplicationTypeVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applicationTypes/{{.applicationTypeName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ApplicationTypeVersions_List",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Applications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applications/{{.applicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Applications_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Applications_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Applications_List",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Services_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applications/{{.applicationName}}/services/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Services_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/applications/{{.applicationName}}/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Services_List",
		Resource:    "Microsoft.ServiceFabric",
	},
}
