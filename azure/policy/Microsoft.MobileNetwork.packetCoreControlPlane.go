package policy

// Microsoft_MobileNetwork_packetCoreControlPlane policy
var Microsoft_MobileNetwork_packetCoreControlPlane = map[string]Policy{
	"PacketCoreControlPlanes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlanes_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlanes_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlanes_ListBySubscription",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlanes_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlanes_ListByResourceGroup",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlanes_Rollback": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/rollback",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlanes_Rollback",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlanes_Reinstall": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/reinstall",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlanes_Reinstall",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlanes_CollectDiagnosticsPackage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{{.packetCoreControlPlaneName}}/collectDiagnosticsPackage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlanes_CollectDiagnosticsPackage",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlaneVersions_Get": {
		Path:   "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/{{.versionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlaneVersions_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"PacketCoreControlPlaneVersions_List": {
		Path:   "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "PacketCoreControlPlaneVersions_List",
		Resource:    "Microsoft.MobileNetwork",
	},
}
