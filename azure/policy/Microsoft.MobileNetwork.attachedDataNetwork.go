package policy

// Microsoft_MobileNetwork_attachedDataNetwork policy
var Microsoft_MobileNetwork_attachedDataNetwork = map[string]Policy{
	"AttachedDataNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/packetCoreDataPlanes/{{.packetCoreDataPlaneName}}/attachedDataNetworks/{{.attachedDataNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "AttachedDataNetworks_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"AttachedDataNetworks_ListByPacketCoreDataPlane": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/packetCoreDataPlanes/{{.packetCoreDataPlaneName}}/attachedDataNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "AttachedDataNetworks_ListByPacketCoreDataPlane",
		Resource:    "Microsoft.MobileNetwork",
	},
}
