package policy

// Microsoft_MobileNetwork_packetCoreDataPlane policy
var Microsoft_MobileNetwork_packetCoreDataPlane = map[string]Policy{
	"PacketCoreDataPlanes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/packetCoreDataPlanes/{{.packetCoreDataPlaneName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreDataPlanes_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreDataPlanes_ListByPacketCoreControlPlane": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/packetCoreDataPlanes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreDataPlanes_ListByPacketCoreControlPlane",
		Resource:    "Microsoft.MobileNetwork",
	},
}
