package policy

// Microsoft_MobileNetwork_sim policy
var Microsoft_MobileNetwork_sim = map[string]Policy{
	"Sims_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups/{{.simGroupName}}/sims/{{.simName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sims_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Sims_ListByGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups/{{.simGroupName}}/sims",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sims_ListByGroup",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Sims_BulkUpload": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups/{{.simGroupName}}/uploadSims",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sims_BulkUpload",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Sims_BulkDelete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups/{{.simGroupName}}/deleteSims",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sims_BulkDelete",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Sims_BulkUploadEncrypted": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/simGroups/{{.simGroupName}}/uploadEncryptedSims",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sims_BulkUploadEncrypted",
		Resource:    "Microsoft.MobileNetwork",
	},
}
