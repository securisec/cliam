package policy

var Microsoft_Web_AppServiceEnvironments = map[string]Policy{
	"AppServiceEnvironments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/hostingEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_List",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Get",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListCapacities": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/capacities/compute",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListCapacities",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetVipInfo": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/capacities/virtualip",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetVipInfo",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ChangeVnet": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/changeVirtualNetwork",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ChangeVnet",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetAseV3NetworkingConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/configurations/networking",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetAseV3NetworkingConfiguration",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListDiagnostics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/diagnostics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListDiagnostics",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetDiagnosticsItem": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/diagnostics/{{.diagnosticsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetDiagnosticsItem",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetInboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/inboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetInboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListMultiRolePools": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRolePools",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetMultiRolePool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetMultiRolePool",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListMultiRolePoolInstanceMetricDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/instances/{{.instance}}/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRolePoolInstanceMetricDefinitions",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListMultiRoleMetricDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRoleMetricDefinitions",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListMultiRolePoolSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRolePoolSkus",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListMultiRoleUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/multiRolePools/default/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListMultiRoleUsages",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListOperations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListOperations",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetPrivateEndpointConnectionList": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetPrivateEndpointConnectionList",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetPrivateLinkResources",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_Reboot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/reboot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Reboot",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Resume",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListAppServicePlans": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/serverfarms",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListAppServicePlans",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListWebApps": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWebApps",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_Suspend": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/suspend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_Suspend",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListUsages",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListWorkerPools": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWorkerPools",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_GetWorkerPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_GetWorkerPool",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListWorkerPoolInstanceMetricDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/instances/{{.instance}}/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWorkerPoolInstanceMetricDefinitions",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListWebWorkerMetricDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/metricdefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWebWorkerMetricDefinitions",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListWorkerPoolSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWorkerPoolSkus",
		Resource:    "Microsoft.Web",
	},
	"AppServiceEnvironments_ListWebWorkerUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.name}}/workerPools/{{.workerPoolName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "AppServiceEnvironments_ListWebWorkerUsages",
		Resource:    "Microsoft.Web",
	},
}
