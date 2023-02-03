package policy

// Microsoft_Network_networkInterface policy
var Microsoft_Network_networkInterface = map[string]Policy{
	"NetworkInterfaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListAll",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_List",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_GetEffectiveRouteTable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/effectiveRouteTable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_GetEffectiveRouteTable",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_ListEffectiveNetworkSecurityGroups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/effectiveNetworkSecurityGroups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListEffectiveNetworkSecurityGroups",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaceIPConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/ipConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaceIPConfigurations_List",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaceIPConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/ipConfigurations/{{.ipConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaceIPConfigurations_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaceLoadBalancers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/loadBalancers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaceLoadBalancers_List",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaceTapConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/tapConfigurations/{{.tapConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaceTapConfigurations_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaceTapConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkInterfaces/{{.networkInterfaceName}}/tapConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaceTapConfigurations_List",
		Resource:    "Microsoft.Network",
	},
}
