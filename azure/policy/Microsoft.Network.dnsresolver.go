package policy

// Microsoft_Network_dnsresolver policy
var Microsoft_Network_dnsresolver = map[string]Policy{
	"DnsResolvers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsResolvers/{{.dnsResolverName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsResolvers_Get",
		Resource:    "Microsoft.Network",
	},
	"DnsResolvers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsResolvers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsResolvers_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"DnsResolvers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/dnsResolvers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsResolvers_List",
		Resource:    "Microsoft.Network",
	},
	"InboundEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsResolvers/{{.dnsResolverName}}/inboundEndpoints/{{.inboundEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "InboundEndpoints_Get",
		Resource:    "Microsoft.Network",
	},
	"InboundEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsResolvers/{{.dnsResolverName}}/inboundEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "InboundEndpoints_List",
		Resource:    "Microsoft.Network",
	},
	"OutboundEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsResolvers/{{.dnsResolverName}}/outboundEndpoints/{{.outboundEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "OutboundEndpoints_Get",
		Resource:    "Microsoft.Network",
	},
	"OutboundEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsResolvers/{{.dnsResolverName}}/outboundEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "OutboundEndpoints_List",
		Resource:    "Microsoft.Network",
	},
	"DnsForwardingRulesets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsForwardingRulesets/{{.dnsForwardingRulesetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsForwardingRulesets_Get",
		Resource:    "Microsoft.Network",
	},
	"DnsForwardingRulesets_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsForwardingRulesets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsForwardingRulesets_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"DnsForwardingRulesets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/dnsForwardingRulesets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsForwardingRulesets_List",
		Resource:    "Microsoft.Network",
	},
	"ForwardingRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsForwardingRulesets/{{.dnsForwardingRulesetName}}/forwardingRules/{{.forwardingRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ForwardingRules_Get",
		Resource:    "Microsoft.Network",
	},
	"ForwardingRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsForwardingRulesets/{{.dnsForwardingRulesetName}}/forwardingRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ForwardingRules_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkLinks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsForwardingRulesets/{{.dnsForwardingRulesetName}}/virtualNetworkLinks/{{.virtualNetworkLinkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkLinks_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsForwardingRulesets/{{.dnsForwardingRulesetName}}/virtualNetworkLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkLinks_List",
		Resource:    "Microsoft.Network",
	},
	"DnsResolvers_ListByVirtualNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/listDnsResolvers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsResolvers_ListByVirtualNetwork",
		Resource:    "Microsoft.Network",
	},
	"DnsForwardingRulesets_ListByVirtualNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/listDnsForwardingRulesets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DnsForwardingRulesets_ListByVirtualNetwork",
		Resource:    "Microsoft.Network",
	},
}
