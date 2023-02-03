package policy

// Microsoft_Network_networkWatcher policy
var Microsoft_Network_networkWatcher = map[string]Policy{
	"NetworkWatchers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_List",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkWatchers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_ListAll",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetTopology": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/topology",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetTopology",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_VerifyIPFlow": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/ipFlowVerify",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_VerifyIPFlow",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetNextHop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/nextHop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetNextHop",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetVMSecurityRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/securityGroupView",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetVMSecurityRules",
		Resource:    "Microsoft.Network",
	},
	"PacketCaptures_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/packetCaptures/{{.packetCaptureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PacketCaptures_Get",
		Resource:    "Microsoft.Network",
	},
	"PacketCaptures_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/packetCaptures/{{.packetCaptureName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PacketCaptures_Stop",
		Resource:    "Microsoft.Network",
	},
	"PacketCaptures_GetStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/packetCaptures/{{.packetCaptureName}}/queryStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PacketCaptures_GetStatus",
		Resource:    "Microsoft.Network",
	},
	"PacketCaptures_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/packetCaptures",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PacketCaptures_List",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetTroubleshooting": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/troubleshoot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetTroubleshooting",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetTroubleshootingResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/queryTroubleshootResult",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetTroubleshootingResult",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_SetFlowLogConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/configureFlowLog",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_SetFlowLogConfiguration",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetFlowLogStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/queryFlowLogStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetFlowLogStatus",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_CheckConnectivity": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/connectivityCheck",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_CheckConnectivity",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetAzureReachabilityReport": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/azureReachabilityReport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetAzureReachabilityReport",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_ListAvailableProviders": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/availableProvidersList",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_ListAvailableProviders",
		Resource:    "Microsoft.Network",
	},
	"NetworkWatchers_GetNetworkConfigurationDiagnostic": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/networkConfigurationDiagnostic",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkWatchers_GetNetworkConfigurationDiagnostic",
		Resource:    "Microsoft.Network",
	},
	"ConnectionMonitors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/connectionMonitors/{{.connectionMonitorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectionMonitors_Get",
		Resource:    "Microsoft.Network",
	},
	"ConnectionMonitors_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/connectionMonitors/{{.connectionMonitorName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectionMonitors_Stop",
		Resource:    "Microsoft.Network",
	},
	"ConnectionMonitors_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/connectionMonitors/{{.connectionMonitorName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectionMonitors_Start",
		Resource:    "Microsoft.Network",
	},
	"ConnectionMonitors_Query": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/connectionMonitors/{{.connectionMonitorName}}/query",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectionMonitors_Query",
		Resource:    "Microsoft.Network",
	},
	"ConnectionMonitors_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/connectionMonitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectionMonitors_List",
		Resource:    "Microsoft.Network",
	},
	"FlowLogs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/flowLogs/{{.flowLogName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FlowLogs_Get",
		Resource:    "Microsoft.Network",
	},
	"FlowLogs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkWatchers/{{.networkWatcherName}}/flowLogs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FlowLogs_List",
		Resource:    "Microsoft.Network",
	},
}
