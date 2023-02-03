package policy

// Microsoft_Synapse_sqlPool policy
var Microsoft_Synapse_sqlPool = map[string]Policy{
	"SqlPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPools_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPools_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPools_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPools_Pause": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/pause",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPools_Pause",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPools_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPools_Resume",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPools_Rename": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/move",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPools_Rename",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolMetadataSyncConfigs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/metadataSync/config",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolMetadataSyncConfigs_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolOperationResults_GetLocationHeaderResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolOperationResults_GetLocationHeaderResult",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolGeoBackupPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/geoBackupPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolGeoBackupPolicies_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolGeoBackupPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/geoBackupPolicies/{{.geoBackupPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolGeoBackupPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolDataWarehouseUserActivities_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/dataWarehouseUserActivities/{{.dataWarehouseUserActivityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolDataWarehouseUserActivities_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolRestorePoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/restorePoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolRestorePoints_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolRestorePoints_Create": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/restorePoints",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolRestorePoints_Create",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolReplicationLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/replicationLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolReplicationLinks_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolReplicationLinks_GetByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/replicationLinks/{{.linkId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolReplicationLinks_GetByName",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolMaintenanceWindows_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/maintenancewindows/current",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolMaintenanceWindows_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolMaintenanceWindowOptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/maintenanceWindowOptions/current",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolMaintenanceWindowOptions_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolTransparentDataEncryptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/transparentDataEncryption/{{.transparentDataEncryptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolTransparentDataEncryptions_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolTransparentDataEncryptions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/transparentDataEncryption",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolTransparentDataEncryptions_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolBlobAuditingPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/auditingSettings/{{.blobAuditingPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolBlobAuditingPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolOperations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolOperations_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolUsages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolUsages_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSensitivityLabels_ListCurrent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/currentSensitivityLabels",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSensitivityLabels_ListCurrent",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSensitivityLabels_ListRecommended": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/recommendedSensitivityLabels",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSensitivityLabels_ListRecommended",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSensitivityLabels_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables/{{.tableName}}/columns/{{.columnName}}/sensitivityLabels/{{.sensitivityLabelSource}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSensitivityLabels_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSensitivityLabels_EnableRecommendation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables/{{.tableName}}/columns/{{.columnName}}/sensitivityLabels/{{.sensitivityLabelSource}}/enable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSensitivityLabels_EnableRecommendation",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSensitivityLabels_DisableRecommendation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables/{{.tableName}}/columns/{{.columnName}}/sensitivityLabels/{{.sensitivityLabelSource}}/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSensitivityLabels_DisableRecommendation",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSchemas_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSchemas_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolTables_ListBySchema": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolTables_ListBySchema",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolTableColumns_ListByTableName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables/{{.tableName}}/columns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolTableColumns_ListByTableName",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolConnectionPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/connectionPolicies/{{.connectionPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolConnectionPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessments_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessmentScans_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}/scans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessmentScans_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessmentScans_InitiateScan": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}/scans/{{.scanId}}/initiateScan",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessmentScans_InitiateScan",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessmentScans_Export": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}/scans/{{.scanId}}/export",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessmentScans_Export",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSecurityAlertPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/securityAlertPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSecurityAlertPolicies_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSecurityAlertPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/securityAlertPolicies/{{.securityAlertPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSecurityAlertPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessmentRuleBaselines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}/rules/{{.ruleId}}/baselines/{{.baselineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessmentRuleBaselines_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessments_Get",
		Resource:    "Microsoft.Synapse",
	},
	"ExtendedSqlPoolBlobAuditingPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/extendedAuditingSettings/{{.blobAuditingPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ExtendedSqlPoolBlobAuditingPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"ExtendedSqlPoolBlobAuditingPolicies_ListBySqlPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/extendedAuditingSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ExtendedSqlPoolBlobAuditingPolicies_ListBySqlPool",
		Resource:    "Microsoft.Synapse",
	},
	"DataMaskingPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/dataMaskingPolicies/{{.dataMaskingPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "DataMaskingPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"DataMaskingRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/dataMaskingPolicies/{{.dataMaskingPolicyName}}/rules/{{.dataMaskingRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "DataMaskingRules_Get",
		Resource:    "Microsoft.Synapse",
	},
	"DataMaskingRules_ListBySqlPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/dataMaskingPolicies/{{.dataMaskingPolicyName}}/rules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "DataMaskingRules_ListBySqlPool",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolBlobAuditingPolicies_ListBySqlPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/auditingSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolBlobAuditingPolicies_ListBySqlPool",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolSchemas_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolSchemas_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolTables_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables/{{.tableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolTables_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolColumns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/schemas/{{.schemaName}}/tables/{{.tableName}}/columns/{{.columnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolColumns_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolVulnerabilityAssessmentScans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}/scans/{{.scanId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolVulnerabilityAssessmentScans_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolRestorePoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/restorePoints/{{.restorePointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolRestorePoints_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolWorkloadGroup_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/workloadGroups/{{.workloadGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolWorkloadGroup_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolWorkloadGroup_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/workloadGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolWorkloadGroup_List",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolWorkloadClassifier_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/workloadGroups/{{.workloadGroupName}}/workloadClassifiers/{{.workloadClassifierName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolWorkloadClassifier_Get",
		Resource:    "Microsoft.Synapse",
	},
	"SqlPoolWorkloadClassifier_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlPools/{{.sqlPoolName}}/workloadGroups/{{.workloadGroupName}}/workloadClassifiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "SqlPoolWorkloadClassifier_List",
		Resource:    "Microsoft.Synapse",
	},
}
