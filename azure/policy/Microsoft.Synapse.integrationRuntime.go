package policy

// Microsoft_Synapse_integrationRuntime policy
var Microsoft_Synapse_integrationRuntime = map[string]Policy{
	"IntegrationRuntimes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_Get",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeNodeIpAddress_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/nodes/{{.nodeName}}/ipAddress",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeNodeIpAddress_Get",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_Upgrade": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/upgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_Upgrade",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeObjectMetadata_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/getObjectMetadata",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeObjectMetadata_List",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeNodes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/nodes/{{.nodeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeNodes_Get",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeCredentials_Sync": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/syncCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeCredentials_Sync",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeConnectionInfos_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/getConnectionInfo",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeConnectionInfos_Get",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_Start",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeAuthKeys_Regenerate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/regenerateAuthKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeAuthKeys_Regenerate",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeAuthKeys_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/listAuthKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeAuthKeys_List",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeMonitoringData_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/monitoringData",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeMonitoringData_List",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_Stop",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeObjectMetadata_Refresh": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/refreshObjectMetadata",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeObjectMetadata_Refresh",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimeStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/getStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimeStatus_Get",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_EnableInteractiveQuery": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/enableInteractiveQuery",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_EnableInteractiveQuery",
		Resource:    "Microsoft.Synapse",
	},
	"IntegrationRuntimes_DisableInteractiveQuery": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/integrationRuntimes/{{.integrationRuntimeName}}/disableInteractiveQuery",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IntegrationRuntimes_DisableInteractiveQuery",
		Resource:    "Microsoft.Synapse",
	},
}
