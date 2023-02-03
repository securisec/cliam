package policy

// Microsoft_Dashboard_grafana policy
var Microsoft_Dashboard_grafana = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Dashboard/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Dashboard",
	},
	"Grafana_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Dashboard/grafana",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Grafana_List",
		Resource:    "Microsoft.Dashboard",
	},
	"Grafana_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Dashboard/grafana",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Grafana_ListByResourceGroup",
		Resource:    "Microsoft.Dashboard",
	},
	"Grafana_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Dashboard/grafana/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Grafana_Get",
		Resource:    "Microsoft.Dashboard",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Dashboard/grafana/{{.workspaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Dashboard",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Dashboard/grafana/{{.workspaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Dashboard",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Dashboard/grafana/{{.workspaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.Dashboard",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Dashboard/grafana/{{.workspaceName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Dashboard",
	},
}
