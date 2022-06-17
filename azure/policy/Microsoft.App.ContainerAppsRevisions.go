package policy

var Microsoft_App_ContainerAppsRevisions = map[string]Policy{
	"ContainerAppsRevisions_ListRevisions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_ListRevisions",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsRevisions_GetRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions/{{.revisionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_GetRevision",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsRevisions_ActivateRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions/{{.revisionName}}/activate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_ActivateRevision",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsRevisions_DeactivateRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions/{{.revisionName}}/deactivate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_DeactivateRevision",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsRevisionReplicas_GetReplica": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions/{{.revisionName}}/replicas/{{.replicaName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisionReplicas_GetReplica",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsRevisionReplicas_ListReplicas": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions/{{.revisionName}}/replicas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisionReplicas_ListReplicas",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsRevisions_RestartRevision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/revisions/{{.revisionName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerAppsRevisions_RestartRevision",
		Resource:    "Microsoft.App",
	},
}
