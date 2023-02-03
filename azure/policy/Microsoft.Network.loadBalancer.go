package policy

// Microsoft_Network_loadBalancer policy
var Microsoft_Network_loadBalancer = map[string]Policy{
	"LoadBalancers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancers_Get",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancers_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/loadBalancers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancers_ListAll",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancers_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerBackendAddressPools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/backendAddressPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerBackendAddressPools_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerBackendAddressPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/backendAddressPools/{{.backendAddressPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerBackendAddressPools_Get",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerFrontendIPConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/frontendIPConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerFrontendIPConfigurations_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerFrontendIPConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/frontendIPConfigurations/{{.frontendIPConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerFrontendIPConfigurations_Get",
		Resource:    "Microsoft.Network",
	},
	"InboundNatRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/inboundNatRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "InboundNatRules_List",
		Resource:    "Microsoft.Network",
	},
	"InboundNatRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/inboundNatRules/{{.inboundNatRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "InboundNatRules_Get",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerLoadBalancingRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/loadBalancingRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerLoadBalancingRules_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerLoadBalancingRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/loadBalancingRules/{{.loadBalancingRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerLoadBalancingRules_Get",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerOutboundRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/outboundRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerOutboundRules_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerOutboundRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/outboundRules/{{.outboundRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerOutboundRules_Get",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerNetworkInterfaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerNetworkInterfaces_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerProbes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/probes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerProbes_List",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancers_SwapPublicIpAddresses": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/setLoadBalancerFrontendPublicIpAddresses",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancers_SwapPublicIpAddresses",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancers_ListInboundNatRulePortMappings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/backendAddressPools/{{.backendPoolName}}/queryInboundNatRulePortMapping",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancers_ListInboundNatRulePortMappings",
		Resource:    "Microsoft.Network",
	},
	"LoadBalancerProbes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/loadBalancers/{{.loadBalancerName}}/probes/{{.probeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "LoadBalancerProbes_Get",
		Resource:    "Microsoft.Network",
	},
}
