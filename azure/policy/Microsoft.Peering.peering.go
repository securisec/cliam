package policy

// Microsoft_Peering_peering policy
var Microsoft_Peering_peering = map[string]Policy{
	"CdnPeeringPrefixes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/cdnPeeringPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CdnPeeringPrefixes_List",
		Resource:    "Microsoft.Peering",
	},
	"CheckServiceProviderAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/checkServiceProviderAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CheckServiceProviderAvailability",
		Resource:    "Microsoft.Peering",
	},
	"LegacyPeerings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/legacyPeerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "LegacyPeerings_List",
		Resource:    "Microsoft.Peering",
	},
	"LookingGlass_Invoke": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/lookingGlass",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "LookingGlass_Invoke",
		Resource:    "Microsoft.Peering",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Peering/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Peering",
	},
	"PeerAsns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peerAsns/{{.peerAsnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeerAsns_Get",
		Resource:    "Microsoft.Peering",
	},
	"PeerAsns_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peerAsns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeerAsns_ListBySubscription",
		Resource:    "Microsoft.Peering",
	},
	"PeeringLocations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peeringLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringLocations_List",
		Resource:    "Microsoft.Peering",
	},
	"RegisteredAsns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/registeredAsns/{{.registeredAsnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegisteredAsns_Get",
		Resource:    "Microsoft.Peering",
	},
	"RegisteredAsns_ListByPeering": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/registeredAsns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegisteredAsns_ListByPeering",
		Resource:    "Microsoft.Peering",
	},
	"RegisteredPrefixes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/registeredPrefixes/{{.registeredPrefixName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegisteredPrefixes_Get",
		Resource:    "Microsoft.Peering",
	},
	"RegisteredPrefixes_ListByPeering": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/registeredPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegisteredPrefixes_ListByPeering",
		Resource:    "Microsoft.Peering",
	},
	"RegisteredPrefixes_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/registeredPrefixes/{{.registeredPrefixName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegisteredPrefixes_Validate",
		Resource:    "Microsoft.Peering",
	},
	"Peerings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Peerings_Get",
		Resource:    "Microsoft.Peering",
	},
	"Peerings_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Peerings_ListByResourceGroup",
		Resource:    "Microsoft.Peering",
	},
	"Peerings_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Peerings_ListBySubscription",
		Resource:    "Microsoft.Peering",
	},
	"ReceivedRoutes_ListByPeering": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/receivedRoutes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReceivedRoutes_ListByPeering",
		Resource:    "Microsoft.Peering",
	},
	"ConnectionMonitorTests_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peeringServices/{{.peeringServiceName}}/connectionMonitorTests/{{.connectionMonitorTestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectionMonitorTests_Get",
		Resource:    "Microsoft.Peering",
	},
	"ConnectionMonitorTests_ListByPeeringService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peeringServices/{{.peeringServiceName}}/connectionMonitorTests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectionMonitorTests_ListByPeeringService",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServiceCountries_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peeringServiceCountries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServiceCountries_List",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServiceLocations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peeringServiceLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServiceLocations_List",
		Resource:    "Microsoft.Peering",
	},
	"Prefixes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peeringServices/{{.peeringServiceName}}/prefixes/{{.prefixName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Prefixes_Get",
		Resource:    "Microsoft.Peering",
	},
	"Prefixes_ListByPeeringService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peeringServices/{{.peeringServiceName}}/prefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Prefixes_ListByPeeringService",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServiceProviders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peeringServiceProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServiceProviders_List",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peeringServices/{{.peeringServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServices_Get",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServices_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peeringServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServices_ListByResourceGroup",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServices_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/peeringServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServices_ListBySubscription",
		Resource:    "Microsoft.Peering",
	},
	"PeeringServices_InitializeConnectionMonitor": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Peering/initializeConnectionMonitor",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PeeringServices_InitializeConnectionMonitor",
		Resource:    "Microsoft.Peering",
	},
	"RpUnbilledPrefixes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Peering/peerings/{{.peeringName}}/rpUnbilledPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RpUnbilledPrefixes_List",
		Resource:    "Microsoft.Peering",
	},
}
