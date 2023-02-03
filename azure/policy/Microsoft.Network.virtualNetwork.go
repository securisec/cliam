package policy

// Microsoft_Network_virtualNetwork policy
var Microsoft_Network_virtualNetwork = map[string]Policy{
	"VirtualNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworks_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworks_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworks_ListAll",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworks_List",
		Resource:    "Microsoft.Network",
	},
	"Subnets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/subnets/{{.subnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Subnets_Get",
		Resource:    "Microsoft.Network",
	},
	"Subnets_PrepareNetworkPolicies": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/subnets/{{.subnetName}}/PrepareNetworkPolicies",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Subnets_PrepareNetworkPolicies",
		Resource:    "Microsoft.Network",
	},
	"Subnets_UnprepareNetworkPolicies": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/subnets/{{.subnetName}}/UnprepareNetworkPolicies",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Subnets_UnprepareNetworkPolicies",
		Resource:    "Microsoft.Network",
	},
	"ResourceNavigationLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/subnets/{{.subnetName}}/ResourceNavigationLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ResourceNavigationLinks_List",
		Resource:    "Microsoft.Network",
	},
	"ServiceAssociationLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/subnets/{{.subnetName}}/ServiceAssociationLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ServiceAssociationLinks_List",
		Resource:    "Microsoft.Network",
	},
	"Subnets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/subnets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Subnets_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkPeerings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/virtualNetworkPeerings/{{.virtualNetworkPeeringName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkPeerings_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkPeerings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/virtualNetworkPeerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkPeerings_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworks_CheckIPAddressAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/CheckIPAddressAvailability",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworks_CheckIPAddressAvailability",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworks_ListUsage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworks_ListUsage",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworks_ListDdosProtectionStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/ddosProtectionStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworks_ListDdosProtectionStatus",
		Resource:    "Microsoft.Network",
	},
}
