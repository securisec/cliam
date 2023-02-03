package policy

// Microsoft_HealthcareApis_healthcare_apis policy
var Microsoft_HealthcareApis_healthcare_apis = map[string]Policy{
	"Services_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/services/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Services_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HealthcareApis/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Services_List",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Services_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Services_ListByResourceGroup",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Services_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HealthcareApis/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Services_CheckNameAvailability",
		Resource:    "Microsoft.HealthcareApis",
	},
	"PrivateEndpointConnections_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/services/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateEndpointConnections_ListByService",
		Resource:    "Microsoft.HealthcareApis",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/services/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"PrivateLinkResources_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/services/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateLinkResources_ListByService",
		Resource:    "Microsoft.HealthcareApis",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/services/{{.resourceName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Workspaces_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HealthcareApis/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Workspaces_ListBySubscription",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Workspaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Workspaces_ListByResourceGroup",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Workspaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Workspaces_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"DicomServices_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/dicomservices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "DicomServices_ListByWorkspace",
		Resource:    "Microsoft.HealthcareApis",
	},
	"DicomServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/dicomservices/{{.dicomServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "DicomServices_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"IotConnectors_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/iotconnectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "IotConnectors_ListByWorkspace",
		Resource:    "Microsoft.HealthcareApis",
	},
	"IotConnectors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/iotconnectors/{{.iotConnectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "IotConnectors_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"FhirDestinations_ListByIotConnector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/iotconnectors/{{.iotConnectorName}}/fhirdestinations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "FhirDestinations_ListByIotConnector",
		Resource:    "Microsoft.HealthcareApis",
	},
	"IotConnectorFhirDestination_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/iotconnectors/{{.iotConnectorName}}/fhirdestinations/{{.fhirDestinationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "IotConnectorFhirDestination_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"FhirServices_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/fhirservices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "FhirServices_ListByWorkspace",
		Resource:    "Microsoft.HealthcareApis",
	},
	"FhirServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/fhirservices/{{.fhirServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "FhirServices_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"WorkspacePrivateEndpointConnections_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "WorkspacePrivateEndpointConnections_ListByWorkspace",
		Resource:    "Microsoft.HealthcareApis",
	},
	"WorkspacePrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "WorkspacePrivateEndpointConnections_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"WorkspacePrivateLinkResources_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "WorkspacePrivateLinkResources_ListByWorkspace",
		Resource:    "Microsoft.HealthcareApis",
	},
	"WorkspacePrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthcareApis/workspaces/{{.workspaceName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "WorkspacePrivateLinkResources_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.HealthcareApis/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.HealthcareApis",
	},
	"OperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HealthcareApis/locations/{{.locationName}}/operationresults/{{.operationResultId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "OperationResults_Get",
		Resource:    "Microsoft.HealthcareApis",
	},
}
