package policy

var Microsoft_MachineLearningServices_machineLearningServices = []Policy{
	{
		Path:   "/providers/Microsoft.MachineLearningServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_Get",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListByResourceGroup",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/diagnose",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_Diagnose",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListKeys",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/resyncKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ResyncKeys",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MachineLearningServices/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Usages_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MachineLearningServices/locations/{{.location}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "VirtualMachineSizes_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MachineLearningServices/locations/{{.location}}/updateQuotas",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Quotas_Update",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MachineLearningServices/locations/{{.location}}/quotas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Quotas_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MachineLearningServices/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListBySubscription",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes/{{.computeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_Get",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes/{{.computeName}}/listNodes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_ListNodes",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/listNotebookAccessToken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListNotebookAccessToken",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes/{{.computeName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_ListKeys",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes/{{.computeName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_Start",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes/{{.computeName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_Stop",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/computes/{{.computeName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Compute_Restart",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/prepareNotebook",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_PrepareNotebook",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/listStorageAccountKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListStorageAccountKeys",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/listNotebookKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListNotebookKeys",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/connections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkspaceConnections_Get",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "WorkspaceConnections_List",
		Resource:    "Microsoft.MachineLearningServices",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Workspaces_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.MachineLearningServices",
	},
}
