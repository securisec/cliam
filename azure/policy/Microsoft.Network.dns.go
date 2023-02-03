package policy

// Microsoft_Network_dns policy
var Microsoft_Network_dns = map[string]Policy{
	"RecordSets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/{{.recordType}}/{{.relativeRecordSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_Get",
		Resource:    "Microsoft.Network",
	},
	"RecordSets_ListByType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/{{.recordType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_ListByType",
		Resource:    "Microsoft.Network",
	},
	"RecordSets_ListByDnsZone": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/recordsets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_ListByDnsZone",
		Resource:    "Microsoft.Network",
	},
	"RecordSets_ListAllByDnsZone": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/all",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_ListAllByDnsZone",
		Resource:    "Microsoft.Network",
	},
	"Zones_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "Zones_Get",
		Resource:    "Microsoft.Network",
	},
	"Zones_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "Zones_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"DnsResourceReference_GetByTargetResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/getDnsResourceReference",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "DnsResourceReference_GetByTargetResources",
		Resource:    "Microsoft.Network",
	},
	"Zones_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/dnszones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "Zones_List",
		Resource:    "Microsoft.Network",
	},
}
