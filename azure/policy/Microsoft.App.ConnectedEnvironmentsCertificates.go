package policy

// Microsoft_App_ConnectedEnvironmentsCertificates policy
var Microsoft_App_ConnectedEnvironmentsCertificates = map[string]Policy{
	"ConnectedEnvironmentsCertificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsCertificates_List",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironmentsCertificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsCertificates_Get",
		Resource:    "Microsoft.App",
	},
}
