package policy

// Microsoft_Web_ContainerAppsRevisions policy
var Microsoft_Web_ContainerAppsRevisions = map[string]Policy{
	"ContainerAppsRevisions_ListRevisions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.containerAppName}}/revisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_ListRevisions",
		Resource:    "Microsoft.Web",
	},
	"ContainerAppsRevisions_GetRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.containerAppName}}/revisions/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_GetRevision",
		Resource:    "Microsoft.Web",
	},
	"ContainerAppsRevisions_ActivateRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.containerAppName}}/revisions/{{.name}}/activate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_ActivateRevision",
		Resource:    "Microsoft.Web",
	},
	"ContainerAppsRevisions_DeactivateRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.containerAppName}}/revisions/{{.name}}/deactivate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_DeactivateRevision",
		Resource:    "Microsoft.Web",
	},
	"ContainerAppsRevisions_RestartRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.containerAppName}}/revisions/{{.name}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_RestartRevision",
		Resource:    "Microsoft.Web",
	},
}
