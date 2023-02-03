package policy

// Microsoft_Network_virtualWan policy
var Microsoft_Network_virtualWan = map[string]Policy{
	"VirtualWans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualWans/{{.VirtualWANName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualWans_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualWans_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualWans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualWans_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"VirtualWans_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/virtualWans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualWans_List",
		Resource:    "Microsoft.Network",
	},
	"VpnSites_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnSites/{{.vpnSiteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSites_Get",
		Resource:    "Microsoft.Network",
	},
	"VpnSites_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSites_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"VpnSiteLinks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnSites/{{.vpnSiteName}}/vpnSiteLinks/{{.vpnSiteLinkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSiteLinks_Get",
		Resource:    "Microsoft.Network",
	},
	"VpnSiteLinks_ListByVpnSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnSites/{{.vpnSiteName}}/vpnSiteLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSiteLinks_ListByVpnSite",
		Resource:    "Microsoft.Network",
	},
	"VpnSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/vpnSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSites_List",
		Resource:    "Microsoft.Network",
	},
	"VpnSitesConfiguration_Download": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualWans/{{.virtualWANName}}/vpnConfiguration",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSitesConfiguration_Download",
		Resource:    "Microsoft.Network",
	},
	"SupportedSecurityProviders": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualWans/{{.virtualWANName}}/supportedSecurityProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SupportedSecurityProviders",
		Resource:    "Microsoft.Network",
	},
	"VpnServerConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnServerConfigurations/{{.vpnServerConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnServerConfigurations_Get",
		Resource:    "Microsoft.Network",
	},
	"VpnServerConfigurations_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnServerConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnServerConfigurations_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"ConfigurationPolicyGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnServerConfigurations/{{.vpnServerConfigurationName}}/configurationPolicyGroups/{{.configurationPolicyGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConfigurationPolicyGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"configurationPolicyGroups_ListByVpnServerConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnServerConfigurations/{{.vpnServerConfigurationName}}/configurationPolicyGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "configurationPolicyGroups_ListByVpnServerConfiguration",
		Resource:    "Microsoft.Network",
	},
	"VpnServerConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/vpnServerConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnServerConfigurations_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubs_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubs_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/virtualHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubs_List",
		Resource:    "Microsoft.Network",
	},
	"RouteMaps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/routeMaps/{{.routeMapName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteMaps_Get",
		Resource:    "Microsoft.Network",
	},
	"RouteMaps_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/routeMaps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteMaps_List",
		Resource:    "Microsoft.Network",
	},
	"HubVirtualNetworkConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/hubVirtualNetworkConnections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "HubVirtualNetworkConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"HubVirtualNetworkConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/hubVirtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "HubVirtualNetworkConnections_List",
		Resource:    "Microsoft.Network",
	},
	"VpnGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"VpnGateways_Reset": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnGateways_Reset",
		Resource:    "Microsoft.Network",
	},
	"VpnGateways_StartPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/startpacketcapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnGateways_StartPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VpnGateways_StopPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/stoppacketcapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnGateways_StopPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VpnLinkConnections_ResetConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.connectionName}}/vpnLinkConnections/{{.linkConnectionName}}/resetconnection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnLinkConnections_ResetConnection",
		Resource:    "Microsoft.Network",
	},
	"VpnGateways_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnGateways_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"VpnGateways_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/vpnGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnGateways_List",
		Resource:    "Microsoft.Network",
	},
	"VpnConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"VpnSiteLinkConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.connectionName}}/vpnLinkConnections/{{.linkConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnSiteLinkConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"VpnLinkConnections_GetIkeSas": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.connectionName}}/vpnLinkConnections/{{.linkConnectionName}}/getikesas",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnLinkConnections_GetIkeSas",
		Resource:    "Microsoft.Network",
	},
	"VpnConnections_StartPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.vpnConnectionName}}/startpacketcapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnConnections_StartPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VpnConnections_StopPacketCapture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.vpnConnectionName}}/stoppacketcapture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnConnections_StopPacketCapture",
		Resource:    "Microsoft.Network",
	},
	"VpnConnections_ListByVpnGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnConnections_ListByVpnGateway",
		Resource:    "Microsoft.Network",
	},
	"VpnLinkConnections_ListByVpnConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/vpnConnections/{{.connectionName}}/vpnLinkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnLinkConnections_ListByVpnConnection",
		Resource:    "Microsoft.Network",
	},
	"NatRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/natRules/{{.natRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NatRules_Get",
		Resource:    "Microsoft.Network",
	},
	"NatRules_ListByVpnGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/vpnGateways/{{.gatewayName}}/natRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NatRules_ListByVpnGateway",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/p2svpnGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_List",
		Resource:    "Microsoft.Network",
	},
	"P2SVpnGateways_Reset": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways/{{.gatewayName}}/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2SVpnGateways_Reset",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_GenerateVpnProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways/{{.gatewayName}}/generatevpnprofile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_GenerateVpnProfile",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_GetP2sVpnConnectionHealth": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways/{{.gatewayName}}/getP2sVpnConnectionHealth",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_GetP2sVpnConnectionHealth",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_GetP2sVpnConnectionHealthDetailed": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways/{{.gatewayName}}/getP2sVpnConnectionHealthDetailed",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_GetP2sVpnConnectionHealthDetailed",
		Resource:    "Microsoft.Network",
	},
	"VpnServerConfigurationsAssociatedWithVirtualWan_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualWans/{{.virtualWANName}}/vpnServerConfigurations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VpnServerConfigurationsAssociatedWithVirtualWan_List",
		Resource:    "Microsoft.Network",
	},
	"generatevirtualwanvpnserverconfigurationvpnprofile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualWans/{{.virtualWANName}}/GenerateVpnProfile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "generatevirtualwanvpnserverconfigurationvpnprofile",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubRouteTableV2s_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/routeTables/{{.routeTableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubRouteTableV2s_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubRouteTableV2s_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/routeTables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubRouteTableV2s_List",
		Resource:    "Microsoft.Network",
	},
	"P2sVpnGateways_DisconnectP2sVpnConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/p2svpnGateways/{{.p2sVpnGatewayName}}/disconnectP2sVpnConnections",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "P2sVpnGateways_DisconnectP2sVpnConnections",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteGateways_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/expressRouteGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteGateways_ListBySubscription",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteGateways_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteGateways_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteGateways/{{.expressRouteGatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteGateways/{{.expressRouteGatewayName}}/expressRouteConnections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteGateways/{{.expressRouteGatewayName}}/expressRouteConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteConnections_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubBgpConnection_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/bgpConnections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubBgpConnection_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubBgpConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/bgpConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubBgpConnections_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubBgpConnections_ListLearnedRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.hubName}}/bgpConnections/{{.connectionName}}/learnedRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubBgpConnections_ListLearnedRoutes",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubBgpConnections_ListAdvertisedRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.hubName}}/bgpConnections/{{.connectionName}}/advertisedRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubBgpConnections_ListAdvertisedRoutes",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubIpConfiguration_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/ipConfigurations/{{.ipConfigName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubIpConfiguration_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubIpConfiguration_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/ipConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubIpConfiguration_List",
		Resource:    "Microsoft.Network",
	},
	"HubRouteTables_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/hubRouteTables/{{.routeTableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "HubRouteTables_Get",
		Resource:    "Microsoft.Network",
	},
	"HubRouteTables_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/hubRouteTables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "HubRouteTables_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubs_GetEffectiveVirtualHubRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/effectiveRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubs_GetEffectiveVirtualHubRoutes",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubs_GetInboundRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/inboundRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubs_GetInboundRoutes",
		Resource:    "Microsoft.Network",
	},
	"VirtualHubs_GetOutboundRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/outboundRoutes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualHubs_GetOutboundRoutes",
		Resource:    "Microsoft.Network",
	},
	"RoutingIntent_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/routingIntent/{{.routingIntentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RoutingIntent_Get",
		Resource:    "Microsoft.Network",
	},
	"RoutingIntent_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualHubs/{{.virtualHubName}}/routingIntent",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RoutingIntent_List",
		Resource:    "Microsoft.Network",
	},
}
