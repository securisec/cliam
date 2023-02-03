package policy

// Microsoft_Network_privatedns policy
var Microsoft_Network_privatedns = map[string]Policy{
	"PrivateZones_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones/{{.privateZoneName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateZones_Get",
		Resource:    "Microsoft.Network",
	},
	"PrivateZones_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/privateDnsZones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateZones_List",
		Resource:    "Microsoft.Network",
	},
	"PrivateZones_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateZones_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkLinks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones/{{.privateZoneName}}/virtualNetworkLinks/{{.virtualNetworkLinkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "VirtualNetworkLinks_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones/{{.privateZoneName}}/virtualNetworkLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "VirtualNetworkLinks_List",
		Resource:    "Microsoft.Network",
	},
	"RecordSets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones/{{.privateZoneName}}/{{.recordType}}/{{.relativeRecordSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "RecordSets_Get",
		Resource:    "Microsoft.Network",
	},
	"RecordSets_ListByType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones/{{.privateZoneName}}/{{.recordType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "RecordSets_ListByType",
		Resource:    "Microsoft.Network",
	},
	"RecordSets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateDnsZones/{{.privateZoneName}}/ALL",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "RecordSets_List",
		Resource:    "Microsoft.Network",
	},
}
