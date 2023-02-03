package policy

// Microsoft_Logic_logic policy
var Microsoft_Logic_logic = map[string]Policy{
	"Workflows_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Logic/workflows",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_ListBySubscription",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_ListByResourceGroup",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_Get",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_Disable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_Disable",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_Enable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/enable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_Enable",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_GenerateUpgradedDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/generateUpgradedDefinition",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_GenerateUpgradedDefinition",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_ListCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/listCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_ListCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_ListSwagger": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/listSwagger",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_ListSwagger",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_Move": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/move",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_Move",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_RegenerateAccessKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/regenerateAccessKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_RegenerateAccessKey",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_ValidateByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_ValidateByResourceGroup",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowVersions_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/versions/{{.versionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowVersions_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_Reset": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_Reset",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_Run": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_Run",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_GetSchemaJson": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/schemas/json",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_GetSchemaJson",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_SetState": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/setState",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_SetState",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggers_ListCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/listCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggers_ListCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowVersionTriggers_ListCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/versions/{{.versionId}}/triggers/{{.triggerName}}/listCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowVersionTriggers_ListCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggerHistories_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/histories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggerHistories_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggerHistories_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/histories/{{.historyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggerHistories_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowTriggerHistories_Resubmit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/triggers/{{.triggerName}}/histories/{{.historyName}}/resubmit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowTriggerHistories_Resubmit",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRuns_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRuns_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRuns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRuns_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRuns_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRuns_Cancel",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActions_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActions_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActions_ListExpressionTraces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/listExpressionTraces",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActions_ListExpressionTraces",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRepetitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRepetitions_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRepetitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRepetitions_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRepetitions_ListExpressionTraces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}/listExpressionTraces",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRepetitions_ListExpressionTraces",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRepetitionsRequestHistories_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}/requestHistories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRepetitionsRequestHistories_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRepetitionsRequestHistories_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}/requestHistories/{{.requestHistoryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRepetitionsRequestHistories_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRequestHistories_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/requestHistories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRequestHistories_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionRequestHistories_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/requestHistories/{{.requestHistoryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionRequestHistories_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionScopeRepetitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/scopeRepetitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionScopeRepetitions_List",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunActionScopeRepetitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/scopeRepetitions/{{.repetitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunActionScopeRepetitions_Get",
		Resource:    "Microsoft.Logic",
	},
	"WorkflowRunOperations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/workflows/{{.workflowName}}/runs/{{.runName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "WorkflowRunOperations_Get",
		Resource:    "Microsoft.Logic",
	},
	"Workflows_ValidateByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/locations/{{.location}}/workflows/{{.workflowName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Workflows_ValidateByLocation",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Logic/integrationAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_ListBySubscription",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_ListByResourceGroup",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountAssemblies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/assemblies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountAssemblies_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountAssemblies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/assemblies/{{.assemblyArtifactName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountAssemblies_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountAssemblies_ListContentCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/assemblies/{{.assemblyArtifactName}}/listContentCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountAssemblies_ListContentCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountBatchConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/batchConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountBatchConfigurations_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountBatchConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/batchConfigurations/{{.batchConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountBatchConfigurations_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_ListCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/listCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_ListCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_ListKeyVaultKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/listKeyVaultKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_ListKeyVaultKeys",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_LogTrackingEvents": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/logTrackingEvents",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_LogTrackingEvents",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccounts_RegenerateAccessKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/regenerateAccessKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccounts_RegenerateAccessKey",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountSchemas_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/schemas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountSchemas_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountSchemas_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/schemas/{{.schemaName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountSchemas_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountSchemas_ListContentCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/schemas/{{.schemaName}}/listContentCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountSchemas_ListContentCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountMaps_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/maps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountMaps_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountMaps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/maps/{{.mapName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountMaps_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountMaps_ListContentCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/maps/{{.mapName}}/listContentCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountMaps_ListContentCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountPartners_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/partners",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountPartners_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountPartners_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/partners/{{.partnerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountPartners_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountPartners_ListContentCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/partners/{{.partnerName}}/listContentCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountPartners_ListContentCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountAgreements_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/agreements",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountAgreements_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountAgreements_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/agreements/{{.agreementName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountAgreements_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountAgreements_ListContentCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/agreements/{{.agreementName}}/listContentCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountAgreements_ListContentCallbackUrl",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountCertificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountCertificates_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountCertificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountCertificates_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountSessions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/sessions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountSessions_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationAccountSessions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logic/integrationAccounts/{{.integrationAccountName}}/sessions/{{.sessionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationAccountSessions_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironments_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Logic/integrationServiceEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironments_ListBySubscription",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironments_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironmentSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironmentSkus_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironments_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironments_Restart",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironmentNetworkHealth_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}/health/network",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironmentNetworkHealth_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironmentManagedApis_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}/managedApis",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironmentManagedApis_List",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironmentManagedApis_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}/managedApis/{{.apiName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironmentManagedApis_Get",
		Resource:    "Microsoft.Logic",
	},
	"IntegrationServiceEnvironmentManagedApiOperations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Logic/integrationServiceEnvironments/{{.integrationServiceEnvironmentName}}/managedApis/{{.apiName}}/apiOperations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "IntegrationServiceEnvironmentManagedApiOperations_List",
		Resource:    "Microsoft.Logic",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Logic/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Logic",
	},
}
