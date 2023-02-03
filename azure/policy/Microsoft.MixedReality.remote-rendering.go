package policy

// Microsoft_MixedReality_remote_rendering policy
var Microsoft_MixedReality_remote_rendering = map[string]Policy{
	"RemoteRenderingAccounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MixedReality/remoteRenderingAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "RemoteRenderingAccounts_ListBySubscription",
		Resource:    "Microsoft.MixedReality",
	},
	"RemoteRenderingAccounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/remoteRenderingAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "RemoteRenderingAccounts_ListByResourceGroup",
		Resource:    "Microsoft.MixedReality",
	},
	"RemoteRenderingAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "RemoteRenderingAccounts_Get",
		Resource:    "Microsoft.MixedReality",
	},
	"RemoteRenderingAccounts_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "RemoteRenderingAccounts_ListKeys",
		Resource:    "Microsoft.MixedReality",
	},
	"RemoteRenderingAccounts_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{{.accountName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "RemoteRenderingAccounts_RegenerateKeys",
		Resource:    "Microsoft.MixedReality",
	},
}
