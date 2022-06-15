package policy

var Microsoft_Web_AppServiceEnvironments = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/hostingEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_List",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Get",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/capacities/compute",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListCapacities",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/capacities/virtualip",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetVipInfo",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/changeVirtualNetwork",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ChangeVnet",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/configurations/networking",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetAseV3NetworkingConfiguration",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/diagnostics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListDiagnostics",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/diagnostics/{{.diagnosticsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetDiagnosticsItem",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/inboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetInboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRolePools",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetMultiRolePool",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/instances/{{.instance}}/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRolePoolInstanceMetricDefinitions",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRoleMetricDefinitions",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRolePoolSkus",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRoleUsages",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListOperations",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetPrivateEndpointConnectionList",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetPrivateLinkResources",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/reboot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Reboot",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Resume",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/serverfarms",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListAppServicePlans",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWebApps",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/suspend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Suspend",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListUsages",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWorkerPools",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetWorkerPool",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/instances/{{.instance}}/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWorkerPoolInstanceMetricDefinitions",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWebWorkerMetricDefinitions",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWorkerPoolSkus",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWebWorkerUsages",
		Resource:    "Microsoft.Web",
	},
}
