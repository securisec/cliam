package policy

// Microsoft_MachineLearningServices_mfe policy
var Microsoft_MachineLearningServices_mfe = map[string]Policy{
	"BatchEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/batchEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BatchEndpoints_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"BatchEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/batchEndpoints/{{.endpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BatchEndpoints_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"BatchDeployments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/batchEndpoints/{{.endpointName}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BatchDeployments_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"BatchDeployments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/batchEndpoints/{{.endpointName}}/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BatchDeployments_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"BatchEndpoints_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/batchEndpoints/{{.endpointName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BatchEndpoints_ListKeys",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"CodeContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/codes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CodeContainers_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"CodeContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/codes/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CodeContainers_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"CodeVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/codes/{{.name}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CodeVersions_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"CodeVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/codes/{{.name}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CodeVersions_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ComponentContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/components",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ComponentContainers_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ComponentContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/components/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ComponentContainers_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ComponentVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/components/{{.name}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ComponentVersions_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ComponentVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/components/{{.name}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ComponentVersions_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"DataContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/data",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "DataContainers_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"DataContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/data/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "DataContainers_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"DataVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/data/{{.name}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "DataVersions_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"DataVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/data/{{.name}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "DataVersions_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Datastores_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/datastores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Datastores_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Datastores_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/datastores/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Datastores_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Datastores_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/datastores/{{.name}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Datastores_ListSecrets",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"EnvironmentContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/environments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "EnvironmentContainers_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"EnvironmentContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/environments/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "EnvironmentContainers_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"EnvironmentVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/environments/{{.name}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "EnvironmentVersions_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"EnvironmentVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/environments/{{.name}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "EnvironmentVersions_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Jobs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Jobs_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Jobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/jobs/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Jobs_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Jobs_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/jobs/{{.id}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Jobs_Cancel",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ModelContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/models",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ModelContainers_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ModelContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/models/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ModelContainers_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ModelVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/models/{{.name}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ModelVersions_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"ModelVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/models/{{.name}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ModelVersions_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineEndpoints_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineEndpoints_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineDeployments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineDeployments_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineDeployments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineDeployments_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineDeployments_GetLogs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/deployments/{{.deploymentName}}/getLogs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineDeployments_GetLogs",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineDeployments_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/deployments/{{.deploymentName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineDeployments_ListSkus",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineEndpoints_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineEndpoints_ListKeys",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineEndpoints_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineEndpoints_RegenerateKeys",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"OnlineEndpoints_GetToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/onlineEndpoints/{{.endpointName}}/token",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OnlineEndpoints_GetToken",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Schedules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Schedules_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
	"Schedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/schedules/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Schedules_Get",
		Resource:    "Microsoft.MachineLearningServices",
	},
}
