package policy

var Microsoft_Web_AppServicePlans = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/serverfarms",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_List",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_Get",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/capabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListCapabilities",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_GetHybridConnection",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListHybridConnectionKeys",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListWebAppsByHybridConnection",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionPlanLimits/limit",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_GetHybridConnectionPlanLimit",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListHybridConnections",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/restartSites",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_RestartWebApps",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListWebApps",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_GetServerFarmSkus",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListUsages",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListVnets",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_GetVnetFromServerFarm",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}/gateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_GetVnetGateway",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}/routes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_ListRoutesForVnet",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}/routes/{{.routeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_GetRouteForVnet",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/workers/{{.workerName}}/reboot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServicePlans_RebootWorker",
		Resource:    "Microsoft.Web",
	},
}
