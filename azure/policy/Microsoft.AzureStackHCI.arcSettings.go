package policy

    var Microsoft_AzureStackHCI_arcSettings = map[string]Policy{
        "ArcSettings_ListByCluster": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/arcSettings",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "ArcSettings_ListByCluster",
    Resource:       "Microsoft.AzureStackHCI",
},
"ArcSettings_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/arcSettings/{{.arcSettingName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "ArcSettings_Get",
    Resource:       "Microsoft.AzureStackHCI",
},
"ArcSettings_GeneratePassword": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/arcSettings/{{.arcSettingName}}/generatePassword",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "ArcSettings_GeneratePassword",
    Resource:       "Microsoft.AzureStackHCI",
},
"ArcSettings_CreateIdentity": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/arcSettings/{{.arcSettingName}}/createArcIdentity",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "ArcSettings_CreateIdentity",
    Resource:       "Microsoft.AzureStackHCI",
},
    }
    