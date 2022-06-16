package policy

var Microsoft_Compute_compute = []Policy{
	{
		Path:   "/providers/Microsoft.Compute/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/availabilitySets/{{.availabilitySetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/availabilitySets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_ListBySubscription",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/availabilitySets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/availabilitySets/{{.availabilitySetName}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "AvailabilitySets_ListAvailableSizes",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/proximityPlacementGroups/{{.proximityPlacementGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ProximityPlacementGroups_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/proximityPlacementGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ProximityPlacementGroups_ListBySubscription",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/proximityPlacementGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ProximityPlacementGroups_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups/{{.hostGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHostGroups_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHostGroups_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/hostGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHostGroups_ListBySubscription",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups/{{.hostGroupName}}/hosts/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHosts_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/hostGroups/{{.hostGroupName}}/hosts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DedicatedHosts_ListByHostGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/sshPublicKeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_ListBySubscription",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/sshPublicKeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/sshPublicKeys/{{.sshPublicKeyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/sshPublicKeys/{{.sshPublicKeyName}}/generateKeyPair",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SshPublicKeys_GenerateKeyPair",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmextension/types/{{.type}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensionImages_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmextension/types",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensionImages_ListTypes",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmextension/types/{{.type}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensionImages_ListVersions",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/extensions/{{.vmExtensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensions_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineExtensions_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_ListOffers",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_ListPublishers",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImages_ListSkus",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus/{{.skus}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_ListOffers",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_ListPublishers",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/edgeZones/{{.edgeZone}}/publishers/{{.publisherName}}/artifacttypes/vmimage/offers/{{.offer}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineImagesEdgeZone_ListSkus",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Usage_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ListByLocation",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/virtualMachineScaleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ListByLocation",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineSizes_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/images/{{.imageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Images_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Images_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Images_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/capture",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Capture",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/instanceView",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_InstanceView",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/convertToManagedDisks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ConvertToManagedDisks",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Deallocate",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/generalize",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Generalize",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ListAll",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/restorePointCollections/{{.restorePointCollectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePointCollections_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/restorePointCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePointCollections_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/restorePointCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePointCollections_ListAll",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/restorePointCollections/{{.restorePointCollectionName}}/restorePoints/{{.restorePointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "RestorePoints_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/vmSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_ListAvailableSizes",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups/{{.capacityReservationGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservationGroups_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservationGroups_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/capacityReservationGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservationGroups_ListBySubscription",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups/{{.capacityReservationGroupName}}/capacityReservations/{{.capacityReservationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservations_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/capacityReservationGroups/{{.capacityReservationGroupName}}/capacityReservations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CapacityReservations_ListByCapacityReservationGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/powerOff",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_PowerOff",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/reapply",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Reapply",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Restart",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Start",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Redeploy",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_Reimage",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/retrieveBootDiagnosticsData",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_RetrieveBootDiagnosticsData",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/performMaintenance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_PerformMaintenance",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/simulateEviction",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_SimulateEviction",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/assessPatches",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_AssessPatches",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/installPatches",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_InstallPatches",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Deallocate",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/delete",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_DeleteInstances",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/instanceView",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_GetInstanceView",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/extensions/{{.vmssExtensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetExtensions_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetExtensions_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/virtualMachineScaleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ListAll",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ListSkus",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/osUpgradeHistory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_GetOSUpgradeHistory",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/poweroff",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_PowerOff",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Restart",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Start",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Redeploy",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/performMaintenance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_PerformMaintenance",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/manualupgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_UpdateInstances",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_Reimage",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/reimageall",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ReimageAll",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/rollingUpgrades/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_Cancel",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/osRollingUpgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_StartOSUpgrade",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/extensionRollingUpgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_StartExtensionUpgrade",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/rollingUpgrades/latest",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetRollingUpgrades_GetLatest",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/forceRecoveryServiceFabricPlatformUpdateDomainWalk",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/convertToSinglePlacementGroup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_ConvertToSinglePlacementGroup",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/setOrchestrationServiceState",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSets_SetOrchestrationServiceState",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/extensions/{{.vmExtensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMExtensions_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMExtensions_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Reimage",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/reimageall",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_ReimageAll",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Deallocate",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/instanceView",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_GetInstanceView",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/poweroff",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_PowerOff",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Restart",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Start",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_Redeploy",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/retrieveBootDiagnosticsData",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_RetrieveBootDiagnosticsData",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/performMaintenance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_PerformMaintenance",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/simulateEviction",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_SimulateEviction",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/logAnalytics/apiAccess/getRequestRateByInterval",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "LogAnalytics_ExportRequestRateByInterval",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/logAnalytics/apiAccess/getThrottledRequests",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "LogAnalytics_ExportThrottledRequests",
		Resource:    "Microsoft.Compute",
	},
}
