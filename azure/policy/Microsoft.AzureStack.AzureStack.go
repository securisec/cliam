package policy

var Microsoft_AzureStack_AzureStack = []Policy{
	{
		Path:   "/providers/Microsoft.AzureStack/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/providers/Microsoft.AzureStack/cloudManifestFiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "CloudManifestFile_List",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/providers/Microsoft.AzureStack/cloudManifestFiles/{{.verificationVersion}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "CloudManifestFile_Get",
		Resource:    "Microsoft.AzureStack",
	},
}
