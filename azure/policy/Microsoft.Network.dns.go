package policy

var Microsoft_Network_dns = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/{{.recordType}}/{{.relativeRecordSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_Get",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/{{.recordType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_ListByType",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/recordsets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_ListByDnsZone",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}/all",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "RecordSets_ListAllByDnsZone",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones/{{.zoneName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "Zones_Get",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dnsZones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "Zones_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/getDnsResourceReference",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "DnsResourceReference_GetByTargetResources",
		Resource:    "Microsoft.Network",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/dnszones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "Zones_List",
		Resource:    "Microsoft.Network",
	},
}
