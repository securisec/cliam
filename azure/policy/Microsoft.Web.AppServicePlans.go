package policy

// Microsoft_Web_AppServicePlans policy
var Microsoft_Web_AppServicePlans = map[string]Policy{
	"AppServicePlans_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/serverfarms",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_List",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_Get",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListCapabilities": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/capabilities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListCapabilities",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_GetHybridConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_GetHybridConnection",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListHybridConnectionKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListHybridConnectionKeys",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListWebAppsByHybridConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListWebAppsByHybridConnection",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_GetHybridConnectionPlanLimit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionPlanLimits/limit",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_GetHybridConnectionPlanLimit",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListHybridConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/hybridConnectionRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListHybridConnections",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_RestartWebApps": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/restartSites",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_RestartWebApps",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListWebApps": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListWebApps",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_GetServerFarmSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_GetServerFarmSkus",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListUsages",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListVnets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListVnets",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_GetVnetFromServerFarm": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_GetVnetFromServerFarm",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_GetVnetGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}/gateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_GetVnetGateway",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_ListRoutesForVnet": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}/routes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_ListRoutesForVnet",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_GetRouteForVnet": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/virtualNetworkConnections/{{.vnetName}}/routes/{{.routeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_GetRouteForVnet",
		Resource:    "Microsoft.Web",
	},
	"AppServicePlans_RebootWorker": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/serverfarms/{{.name}}/workers/{{.workerName}}/reboot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServicePlans_RebootWorker",
		Resource:    "Microsoft.Web",
	},
}
