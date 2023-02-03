package policy

// Microsoft_AzureStack_AzureStack policy
var Microsoft_AzureStack_AzureStack = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.AzureStack/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AzureStack",
	},
	"CloudManifestFile_List": {
		Path:   "/providers/Microsoft.AzureStack/cloudManifestFiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudManifestFile_List",
		Resource:    "Microsoft.AzureStack",
	},
	"CloudManifestFile_Get": {
		Path:   "/providers/Microsoft.AzureStack/cloudManifestFiles/{{.verificationVersion}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudManifestFile_Get",
		Resource:    "Microsoft.AzureStack",
	},
	"DeploymentLicense_Create": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureStack/generateDeploymentLicense",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "DeploymentLicense_Create",
		Resource:    "Microsoft.AzureStack",
	},
}
