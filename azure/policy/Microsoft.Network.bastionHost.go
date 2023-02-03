package policy

// Microsoft_Network_bastionHost policy
var Microsoft_Network_bastionHost = map[string]Policy{
	"BastionHosts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts/{{.bastionHostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "BastionHosts_Get",
		Resource:    "Microsoft.Network",
	},
	"BastionHosts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/bastionHosts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "BastionHosts_List",
		Resource:    "Microsoft.Network",
	},
	"BastionHosts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "BastionHosts_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"PutBastionShareableLink": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts/{{.bastionHostName}}/createShareableLinks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PutBastionShareableLink",
		Resource:    "Microsoft.Network",
	},
	"DeleteBastionShareableLink": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts/{{.bastionHostName}}/deleteShareableLinks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DeleteBastionShareableLink",
		Resource:    "Microsoft.Network",
	},
	"GetBastionShareableLink": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts/{{.bastionHostName}}/getShareableLinks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "GetBastionShareableLink",
		Resource:    "Microsoft.Network",
	},
	"GetActiveSessions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts/{{.bastionHostName}}/getActiveSessions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "GetActiveSessions",
		Resource:    "Microsoft.Network",
	},
	"DisconnectActiveSessions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/bastionHosts/{{.bastionHostName}}/disconnectActiveSessions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DisconnectActiveSessions",
		Resource:    "Microsoft.Network",
	},
}
