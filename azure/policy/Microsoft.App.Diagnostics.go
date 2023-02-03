package policy

// Microsoft_App_Diagnostics policy
var Microsoft_App_Diagnostics = map[string]Policy{
	"ContainerAppsDiagnostics_ListDetectors": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsDiagnostics_ListDetectors",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsDiagnostics_GetDetector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsDiagnostics_GetDetector",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsDiagnostics_ListRevisions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/detectorProperties/revisionsApi/revisions/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsDiagnostics_ListRevisions",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsDiagnostics_GetRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/detectorProperties/revisionsApi/revisions/{{.revisionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsDiagnostics_GetRevision",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsDiagnostics_GetRoot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/detectorProperties/rootApi/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsDiagnostics_GetRoot",
		Resource:    "Microsoft.App",
	},
	"ManagedEnvironmentDiagnostics_ListDetectors": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ManagedEnvironmentDiagnostics_ListDetectors",
		Resource:    "Microsoft.App",
	},
	"ManagedEnvironmentDiagnostics_GetDetector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ManagedEnvironmentDiagnostics_GetDetector",
		Resource:    "Microsoft.App",
	},
	"ManagedEnvironmentsDiagnostics_GetRoot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/detectorProperties/rootApi/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ManagedEnvironmentsDiagnostics_GetRoot",
		Resource:    "Microsoft.App",
	},
}
