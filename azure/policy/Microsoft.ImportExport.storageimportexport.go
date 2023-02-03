package policy

// Microsoft_ImportExport_storageimportexport policy
var Microsoft_ImportExport_storageimportexport = map[string]Policy{
	"Locations_List": {
		Path:   "/providers/Microsoft.ImportExport/locations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "Locations_List",
		Resource:    "Microsoft.ImportExport",
	},
	"Locations_Get": {
		Path:   "/providers/Microsoft.ImportExport/locations/{{.locationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "Locations_Get",
		Resource:    "Microsoft.ImportExport",
	},
	"Jobs_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ImportExport/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "Jobs_ListBySubscription",
		Resource:    "Microsoft.ImportExport",
	},
	"Jobs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ImportExport/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "Jobs_ListByResourceGroup",
		Resource:    "Microsoft.ImportExport",
	},
	"Jobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ImportExport/jobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "Jobs_Get",
		Resource:    "Microsoft.ImportExport",
	},
	"BitLockerKeys_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ImportExport/jobs/{{.jobName}}/listBitLockerKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "BitLockerKeys_List",
		Resource:    "Microsoft.ImportExport",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ImportExport/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ImportExport",
	},
}
