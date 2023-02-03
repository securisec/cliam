package policy

// Microsoft_Network_virtualNetworkGateway policy
var Microsoft_Network_virtualNetworkGateway = map[string]Policy{
	"VirtualNetworkGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_ListConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_ListConnections",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_Reset": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_Reset",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_ResetVpnClientSharedKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/resetvpnclientsharedkey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_ResetVpnClientSharedKey",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_Generatevpnclientpackage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/generatevpnclientpackage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_Generatevpnclientpackage",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GenerateVpnProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/generatevpnprofile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GenerateVpnProfile",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GetVpnProfilePackageUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/getvpnprofilepackageurl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GetVpnProfilePackageUrl",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GetBgpPeerStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/getBgpPeerStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GetBgpPeerStatus",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_SupportedVpnDevices": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/supportedvpndevices",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_SupportedVpnDevices",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GetLearnedRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/getLearnedRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GetLearnedRoutes",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GetAdvertisedRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/getAdvertisedRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GetAdvertisedRoutes",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_SetVpnclientIpsecParameters": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/setvpnclientipsecparameters",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_SetVpnclientIpsecParameters",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GetVpnclientIpsecParameters": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/getvpnclientipsecparameters",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GetVpnclientIpsecParameters",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_VpnDeviceConfigurationScript": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/vpndeviceconfigurationscript",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_VpnDeviceConfigurationScript",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_StartPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/startPacketCapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_StartPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_StopPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/stopPacketCapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_StopPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_GetSharedKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/sharedkey",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_GetSharedKey",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_ResetSharedKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/sharedkey/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_ResetSharedKey",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_StartPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/startPacketCapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_StartPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_StopPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/stopPacketCapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_StopPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_GetIkeSas": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/getikesas",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_GetIkeSas",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayConnections_ResetConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/connections/{{.virtualNetworkGatewayConnectionName}}/resetconnection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayConnections_ResetConnection",
		Resource:    "Microsoft.Network",
	},
	"LocalNetworkGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/localNetworkGateways/{{.localNetworkGatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LocalNetworkGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"LocalNetworkGateways_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/localNetworkGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LocalNetworkGateways_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_GetVpnclientConnectionHealth": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/getVpnClientConnectionHealth",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_GetVpnclientConnectionHealth",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGateways_DisconnectVirtualNetworkGatewayVpnConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/disconnectVirtualNetworkGatewayVpnConnections",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGateways_DisconnectVirtualNetworkGatewayVpnConnections",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayNatRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/natRules/{{.natRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayNatRules_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkGatewayNatRules_ListByVirtualNetworkGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkGateways/{{.virtualNetworkGatewayName}}/natRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkGatewayNatRules_ListByVirtualNetworkGateway",
		Resource:    "Microsoft.Network",
	},
}
