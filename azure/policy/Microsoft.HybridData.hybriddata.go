package policy

// Microsoft_HybridData_hybriddata policy
var Microsoft_HybridData_hybriddata = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.HybridData/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.HybridData",
	},
	"DataManagers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridData/dataManagers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataManagers_List",
		Resource:    "Microsoft.HybridData",
	},
	"DataManagers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataManagers_ListByResourceGroup",
		Resource:    "Microsoft.HybridData",
	},
	"DataManagers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataManagers_Get",
		Resource:    "Microsoft.HybridData",
	},
	"DataServices_ListByDataManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataServices_ListByDataManager",
		Resource:    "Microsoft.HybridData",
	},
	"DataServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataServices_Get",
		Resource:    "Microsoft.HybridData",
	},
	"JobDefinitions_ListByDataService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "JobDefinitions_ListByDataService",
		Resource:    "Microsoft.HybridData",
	},
	"JobDefinitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions/{{.jobDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "JobDefinitions_Get",
		Resource:    "Microsoft.HybridData",
	},
	"Jobs_ListByJobDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions/{{.jobDefinitionName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Jobs_ListByJobDefinition",
		Resource:    "Microsoft.HybridData",
	},
	"Jobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions/{{.jobDefinitionName}}/jobs/{{.jobId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Jobs_Get",
		Resource:    "Microsoft.HybridData",
	},
	"Jobs_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions/{{.jobDefinitionName}}/jobs/{{.jobId}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Jobs_Cancel",
		Resource:    "Microsoft.HybridData",
	},
	"Jobs_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions/{{.jobDefinitionName}}/jobs/{{.jobId}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Jobs_Resume",
		Resource:    "Microsoft.HybridData",
	},
	"JobDefinitions_Run": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobDefinitions/{{.jobDefinitionName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "JobDefinitions_Run",
		Resource:    "Microsoft.HybridData",
	},
	"Jobs_ListByDataService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataServices/{{.dataServiceName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Jobs_ListByDataService",
		Resource:    "Microsoft.HybridData",
	},
	"DataStores_ListByDataManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataStores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataStores_ListByDataManager",
		Resource:    "Microsoft.HybridData",
	},
	"DataStores_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataStores/{{.dataStoreName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataStores_Get",
		Resource:    "Microsoft.HybridData",
	},
	"DataStoreTypes_ListByDataManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataStoreTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataStoreTypes_ListByDataManager",
		Resource:    "Microsoft.HybridData",
	},
	"DataStoreTypes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/dataStoreTypes/{{.dataStoreTypeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "DataStoreTypes_Get",
		Resource:    "Microsoft.HybridData",
	},
	"JobDefinitions_ListByDataManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/jobDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "JobDefinitions_ListByDataManager",
		Resource:    "Microsoft.HybridData",
	},
	"Jobs_ListByDataManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "Jobs_ListByDataManager",
		Resource:    "Microsoft.HybridData",
	},
	"PublicKeys_ListByDataManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/publicKeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "PublicKeys_ListByDataManager",
		Resource:    "Microsoft.HybridData",
	},
	"PublicKeys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridData/dataManagers/{{.dataManagerName}}/publicKeys/{{.publicKeyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-06-01",
		},
		OperationID: "PublicKeys_Get",
		Resource:    "Microsoft.HybridData",
	},
}
