package policy

// Microsoft_Network_trafficmanager policy
var Microsoft_Network_trafficmanager = map[string]Policy{
	"Endpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/trafficmanagerprofiles/{{.profileName}}/{{.endpointType}}/{{.endpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "Endpoints_Get",
		Resource:    "Microsoft.Network",
	},
	"Profiles_CheckTrafficManagerRelativeDnsNameAvailability": {
		Path:   "/providers/Microsoft.Network/checkTrafficManagerNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "Profiles_CheckTrafficManagerRelativeDnsNameAvailability",
		Resource:    "Microsoft.Network",
	},
	"Profiles_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/trafficmanagerprofiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "Profiles_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"Profiles_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/trafficmanagerprofiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "Profiles_ListBySubscription",
		Resource:    "Microsoft.Network",
	},
	"Profiles_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/trafficmanagerprofiles/{{.profileName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "Profiles_Get",
		Resource:    "Microsoft.Network",
	},
	"GeographicHierarchies_GetDefault": {
		Path:   "/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "GeographicHierarchies_GetDefault",
		Resource:    "Microsoft.Network",
	},
	"HeatMap_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/trafficmanagerprofiles/{{.profileName}}/heatMaps/{{.heatMapType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "HeatMap_Get",
		Resource:    "Microsoft.Network",
	},
	"TrafficManagerUserMetricsKeys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-08-01",
		},
		OperationID: "TrafficManagerUserMetricsKeys_Get",
		Resource:    "Microsoft.Network",
	},
}
