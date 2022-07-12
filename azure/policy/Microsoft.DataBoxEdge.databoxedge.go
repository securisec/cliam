package policy

    var Microsoft_DataBoxEdge_databoxedge = map[string]Policy{
        "Operations_List": {
    Path: "/providers/Microsoft.DataBoxEdge/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.DataBoxEdge",
},
"AvailableSkus_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataBoxEdge/availableSkus",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "AvailableSkus_List",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_ListBySubscription": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_ListBySubscription",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_ListByResourceGroup",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Alerts_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/alerts",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Alerts_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"Alerts_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/alerts/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Alerts_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"BandwidthSchedules_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/bandwidthSchedules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "BandwidthSchedules_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"BandwidthSchedules_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/bandwidthSchedules/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "BandwidthSchedules_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"DeviceCapacityCheck_CheckResourceCreationFeasibility": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/deviceCapacityCheck",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DeviceCapacityCheck_CheckResourceCreationFeasibility",
    Resource:       "Microsoft.DataBoxEdge",
},
"DeviceCapacityInfo_GetDeviceCapacityInfo": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/deviceCapacityInfo/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DeviceCapacityInfo_GetDeviceCapacityInfo",
    Resource:       "Microsoft.DataBoxEdge",
},
"DiagnosticSettings_GetDiagnosticProactiveLogCollectionSettings": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/diagnosticProactiveLogCollectionSettings/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DiagnosticSettings_GetDiagnosticProactiveLogCollectionSettings",
    Resource:       "Microsoft.DataBoxEdge",
},
"DiagnosticSettings_GetDiagnosticRemoteSupportSettings": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/diagnosticRemoteSupportSettings/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DiagnosticSettings_GetDiagnosticRemoteSupportSettings",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_DownloadUpdates": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/downloadUpdates",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_DownloadUpdates",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_GenerateCertificate": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/generateCertificate",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_GenerateCertificate",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_GetExtendedInformation": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/getExtendedInformation",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_GetExtendedInformation",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_InstallUpdates": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/installUpdates",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_InstallUpdates",
    Resource:       "Microsoft.DataBoxEdge",
},
"Jobs_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/jobs/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Jobs_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_GetNetworkSettings": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/networkSettings/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_GetNetworkSettings",
    Resource:       "Microsoft.DataBoxEdge",
},
"Nodes_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/nodes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Nodes_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"OperationsStatus_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/operationsStatus/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "OperationsStatus_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Orders_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/orders",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Orders_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"Orders_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/orders/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Orders_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Orders_ListDCAccessCode": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/orders/default/listDCAccessCode",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Orders_ListDCAccessCode",
    Resource:       "Microsoft.DataBoxEdge",
},
"Roles_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/roles",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Roles_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"Roles_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/roles/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Roles_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Addons_ListByRole": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/roles/{{.roleName}}/addons",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Addons_ListByRole",
    Resource:       "Microsoft.DataBoxEdge",
},
"Addons_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/roles/{{.roleName}}/addons/{{.addonName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Addons_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"MonitoringConfig_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/roles/{{.roleName}}/monitoringConfig",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "MonitoringConfig_List",
    Resource:       "Microsoft.DataBoxEdge",
},
"MonitoringConfig_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/roles/{{.roleName}}/monitoringConfig/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "MonitoringConfig_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_ScanForUpdates": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/scanForUpdates",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_ScanForUpdates",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_CreateOrUpdateSecuritySettings": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/securitySettings/default/update",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_CreateOrUpdateSecuritySettings",
    Resource:       "Microsoft.DataBoxEdge",
},
"Shares_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/shares",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Shares_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"Shares_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/shares/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Shares_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Shares_Refresh": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/shares/{{.name}}/refresh",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Shares_Refresh",
    Resource:       "Microsoft.DataBoxEdge",
},
"StorageAccountCredentials_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccountCredentials",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "StorageAccountCredentials_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"StorageAccountCredentials_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccountCredentials/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "StorageAccountCredentials_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"StorageAccounts_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccounts",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "StorageAccounts_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"StorageAccounts_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccounts/{{.storageAccountName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "StorageAccounts_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Containers_ListByStorageAccount": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccounts/{{.storageAccountName}}/containers",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Containers_ListByStorageAccount",
    Resource:       "Microsoft.DataBoxEdge",
},
"Containers_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccounts/{{.storageAccountName}}/containers/{{.containerName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Containers_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"Containers_Refresh": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/storageAccounts/{{.storageAccountName}}/containers/{{.containerName}}/refresh",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Containers_Refresh",
    Resource:       "Microsoft.DataBoxEdge",
},
"Triggers_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/triggers",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Triggers_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"Triggers_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/triggers/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Triggers_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
"SupportPackages_TriggerSupportPackage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/triggerSupportPackage",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "SupportPackages_TriggerSupportPackage",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_UpdateExtendedInformation": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/updateExtendedInformation",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_UpdateExtendedInformation",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_GetUpdateSummary": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/updateSummary/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_GetUpdateSummary",
    Resource:       "Microsoft.DataBoxEdge",
},
"Devices_UploadCertificate": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/uploadCertificate",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Devices_UploadCertificate",
    Resource:       "Microsoft.DataBoxEdge",
},
"Users_ListByDataBoxEdgeDevice": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/users",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Users_ListByDataBoxEdgeDevice",
    Resource:       "Microsoft.DataBoxEdge",
},
"Users_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{.deviceName}}/users/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Users_Get",
    Resource:       "Microsoft.DataBoxEdge",
},
    }
    