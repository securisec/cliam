package policy

var Microsoft_Compute_compute = map[string]Policy{

	"Operations_List": {
		Path:   "/providers/Microsoft.Compute/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Compute",
	},
	"AvailabilitySets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/availabilitySets/{{.availabilitySetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_Get",
		Resource:    "Microsoft.Compute",
	},
	"AvailabilitySets_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/availabilitySets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_ListBySubscription",
		Resource:    "Microsoft.Compute",
	},
	"AvailabilitySets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/availabilitySets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_List",
		Resource:    "Microsoft.Compute",
	},
	"AvailabilitySets_ListAvailableSizes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/availabilitySets/{{.availabilitySetName}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_ListAvailableSizes",
		Resource:    "Microsoft.Compute",
	},
	"ProximityPlacementGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/proximityPlacementGroups/{{.proximityPlacementGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ProximityPlacementGroups_Get",
		Resource:    "Microsoft.Compute",
	},
	"ProximityPlacementGroups_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/proximityPlacementGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ProximityPlacementGroups_ListBySubscription",
		Resource:    "Microsoft.Compute",
	},
	"ProximityPlacementGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/proximityPlacementGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ProximityPlacementGroups_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	},
	"DedicatedHostGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups/{{.hostGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHostGroups_Get",
		Resource:    "Microsoft.Compute",
	},
	"DedicatedHostGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHostGroups_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	},
	"DedicatedHostGroups_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/hostGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHostGroups_ListBySubscription",
		Resource:    "Microsoft.Compute",
	},
	"DedicatedHosts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups/{{.hostGroupName}}/hosts/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHosts_Get",
		Resource:    "Microsoft.Compute",
	},
	"DedicatedHosts_ListByHostGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups/{{.hostGroupName}}/hosts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHosts_ListByHostGroup",
		Resource:    "Microsoft.Compute",
	},
	"SshPublicKeys_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/sshPublicKeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_ListBySubscription",
		Resource:    "Microsoft.Compute",
	},
	"SshPublicKeys_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/sshPublicKeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	},
	"SshPublicKeys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/sshPublicKeys/{{.sshPublicKeyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_Get",
		Resource:    "Microsoft.Compute",
	},
	"SshPublicKeys_GenerateKeyPair": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/sshPublicKeys/{{.sshPublicKeyName}}/generateKeyPair",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_GenerateKeyPair",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineExtensionImages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmextension/types/{{.type}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensionImages_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineExtensionImages_ListTypes": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmextension/types",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensionImages_ListTypes",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineExtensionImages_ListVersions": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmextension/types/{{.type}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensionImages_ListVersions",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineExtensions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/extensions/{{.vmExtensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensions_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineExtensions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensions_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImages_ListOffers": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_ListOffers",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImages_ListPublishers": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_ListPublishers",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImages_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_ListSkus",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImagesEdgeZone_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImagesEdgeZone_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImagesEdgeZone_ListOffers": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_ListOffers",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImagesEdgeZone_ListPublishers": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_ListPublishers",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineImagesEdgeZone_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_ListSkus",
		Resource:    "Microsoft.Compute",
	},
	"Usage_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Usage_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ListByLocation",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/virtualMachineScaleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ListByLocation",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineSizes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineSizes_List",
		Resource:    "Microsoft.Compute",
	},
	"Images_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/images/{{.imageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Images_Get",
		Resource:    "Microsoft.Compute",
	},
	"Images_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Images_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	},
	"Images_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Images_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Capture": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/capture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Capture",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_InstanceView": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/instanceView",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_InstanceView",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_ConvertToManagedDisks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/convertToManagedDisks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ConvertToManagedDisks",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Deallocate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Deallocate",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Generalize": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/generalize",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Generalize",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ListAll",
		Resource:    "Microsoft.Compute",
	},
	"RestorePointCollections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/restorePointCollections/{{.restorePointCollectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePointCollections_Get",
		Resource:    "Microsoft.Compute",
	},
	"RestorePointCollections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/restorePointCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePointCollections_List",
		Resource:    "Microsoft.Compute",
	},
	"RestorePointCollections_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/restorePointCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePointCollections_ListAll",
		Resource:    "Microsoft.Compute",
	},
	"RestorePoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/restorePointCollections/{{.restorePointCollectionName}}/restorePoints/{{.restorePointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePoints_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_ListAvailableSizes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ListAvailableSizes",
		Resource:    "Microsoft.Compute",
	},
	"CapacityReservationGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups/{{.capacityReservationGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservationGroups_Get",
		Resource:    "Microsoft.Compute",
	},
	"CapacityReservationGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservationGroups_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	},
	"CapacityReservationGroups_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/capacityReservationGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservationGroups_ListBySubscription",
		Resource:    "Microsoft.Compute",
	},
	"CapacityReservations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups/{{.capacityReservationGroupName}}/capacityReservations/{{.capacityReservationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservations_Get",
		Resource:    "Microsoft.Compute",
	},
	"CapacityReservations_ListByCapacityReservationGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups/{{.capacityReservationGroupName}}/capacityReservations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservations_ListByCapacityReservationGroup",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_PowerOff": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/powerOff",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_PowerOff",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Reapply": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/reapply",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Reapply",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Restart",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Start",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Redeploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Redeploy",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_Reimage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Reimage",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_RetrieveBootDiagnosticsData": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/retrieveBootDiagnosticsData",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_RetrieveBootDiagnosticsData",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_PerformMaintenance": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/performMaintenance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_PerformMaintenance",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_SimulateEviction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/simulateEviction",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_SimulateEviction",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_AssessPatches": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/assessPatches",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_AssessPatches",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_InstallPatches": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/installPatches",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_InstallPatches",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_Deallocate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Deallocate",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_DeleteInstances": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/delete",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_DeleteInstances",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_GetInstanceView": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/instanceView",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_GetInstanceView",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetExtensions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/extensions/{{.vmssExtensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetExtensions_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetExtensions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetExtensions_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/virtualMachineScaleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ListAll",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ListSkus",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_GetOSUpgradeHistory": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/osUpgradeHistory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_GetOSUpgradeHistory",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_PowerOff": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/poweroff",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_PowerOff",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Restart",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Start",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_Redeploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Redeploy",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_PerformMaintenance": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/performMaintenance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_PerformMaintenance",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_UpdateInstances": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/manualupgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_UpdateInstances",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_Reimage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Reimage",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_ReimageAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/reimageall",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ReimageAll",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetRollingUpgrades_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/rollingUpgrades/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_Cancel",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetRollingUpgrades_StartOSUpgrade": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/osRollingUpgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_StartOSUpgrade",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetRollingUpgrades_StartExtensionUpgrade": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/extensionRollingUpgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_StartExtensionUpgrade",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetRollingUpgrades_GetLatest": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/rollingUpgrades/latest",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_GetLatest",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/forceRecoveryServiceFabricPlatformUpdateDomainWalk",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_ConvertToSinglePlacementGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/convertToSinglePlacementGroup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ConvertToSinglePlacementGroup",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSets_SetOrchestrationServiceState": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/setOrchestrationServiceState",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_SetOrchestrationServiceState",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMExtensions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/extensions/{{.vmExtensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMExtensions_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMExtensions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMExtensions_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_Reimage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Reimage",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_ReimageAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/reimageall",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_ReimageAll",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_Deallocate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Deallocate",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_GetInstanceView": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/instanceView",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_GetInstanceView",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_PowerOff": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/poweroff",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_PowerOff",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Restart",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Start",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_Redeploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Redeploy",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_RetrieveBootDiagnosticsData": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/retrieveBootDiagnosticsData",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_RetrieveBootDiagnosticsData",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_PerformMaintenance": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/performMaintenance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_PerformMaintenance",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_SimulateEviction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/simulateEviction",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_SimulateEviction",
		Resource:    "Microsoft.Compute",
	},
	"LogAnalytics_ExportRequestRateByInterval": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/logAnalytics/apiAccess/getRequestRateByInterval",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "LogAnalytics_ExportRequestRateByInterval",
		Resource:    "Microsoft.Compute",
	},
	"LogAnalytics_ExportThrottledRequests": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/logAnalytics/apiAccess/getThrottledRequests",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "LogAnalytics_ExportThrottledRequests",
		Resource:    "Microsoft.Compute",
	},
}
