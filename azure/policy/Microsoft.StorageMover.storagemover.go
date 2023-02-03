package policy

// Microsoft_StorageMover_storagemover policy
var Microsoft_StorageMover_storagemover = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.StorageMover/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.StorageMover",
	},
	"StorageMovers_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageMover/storageMovers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "StorageMovers_ListBySubscription",
		Resource:    "Microsoft.StorageMover",
	},
	"StorageMovers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "StorageMovers_List",
		Resource:    "Microsoft.StorageMover",
	},
	"StorageMovers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "StorageMovers_Get",
		Resource:    "Microsoft.StorageMover",
	},
	"Agents_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/agents",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Agents_List",
		Resource:    "Microsoft.StorageMover",
	},
	"Agents_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/agents/{{.agentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Agents_Get",
		Resource:    "Microsoft.StorageMover",
	},
	"Endpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/endpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Endpoints_List",
		Resource:    "Microsoft.StorageMover",
	},
	"Endpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/endpoints/{{.endpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Endpoints_Get",
		Resource:    "Microsoft.StorageMover",
	},
	"Projects_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Projects_List",
		Resource:    "Microsoft.StorageMover",
	},
	"Projects_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "Projects_Get",
		Resource:    "Microsoft.StorageMover",
	},
	"JobDefinitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}/jobDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "JobDefinitions_List",
		Resource:    "Microsoft.StorageMover",
	},
	"JobDefinitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}/jobDefinitions/{{.jobDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "JobDefinitions_Get",
		Resource:    "Microsoft.StorageMover",
	},
	"JobDefinitions_StartJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}/jobDefinitions/{{.jobDefinitionName}}/startJob",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "JobDefinitions_StartJob",
		Resource:    "Microsoft.StorageMover",
	},
	"JobDefinitions_StopJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}/jobDefinitions/{{.jobDefinitionName}}/stopJob",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "JobDefinitions_StopJob",
		Resource:    "Microsoft.StorageMover",
	},
	"JobRuns_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}/jobDefinitions/{{.jobDefinitionName}}/jobRuns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "JobRuns_List",
		Resource:    "Microsoft.StorageMover",
	},
	"JobRuns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageMover/storageMovers/{{.storageMoverName}}/projects/{{.projectName}}/jobDefinitions/{{.jobDefinitionName}}/jobRuns/{{.jobRunName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-03-01",
		},
		OperationID: "JobRuns_Get",
		Resource:    "Microsoft.StorageMover",
	},
}
