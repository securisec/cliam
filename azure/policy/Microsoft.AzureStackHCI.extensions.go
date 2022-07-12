package policy

    var Microsoft_AzureStackHCI_extensions = map[string]Policy{
        "Extensions_ListByArcSetting": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/arcSettings/{{.arcSettingName}}/extensions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Extensions_ListByArcSetting",
    Resource:       "Microsoft.AzureStackHCI",
},
"Extensions_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/arcSettings/{{.arcSettingName}}/extensions/{{.extensionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Extensions_Get",
    Resource:       "Microsoft.AzureStackHCI",
},
    }
    