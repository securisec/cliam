package policy

// Microsoft_Synapse_sqlServer policy
var Microsoft_Synapse_sqlServer = map[string]Policy{
	"WorkspaceManagedSqlServerBlobAuditingPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/auditingSettings/{{.blobAuditingPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerBlobAuditingPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerBlobAuditingPolicies_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/auditingSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerBlobAuditingPolicies_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerExtendedBlobAuditingPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/extendedAuditingSettings/{{.blobAuditingPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerExtendedBlobAuditingPolicies_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerExtendedBlobAuditingPolicies_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/extendedAuditingSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerExtendedBlobAuditingPolicies_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerSecurityAlertPolicy_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/securityAlertPolicies/{{.securityAlertPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerSecurityAlertPolicy_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerSecurityAlertPolicy_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/securityAlertPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerSecurityAlertPolicy_List",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerVulnerabilityAssessments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/vulnerabilityAssessments/{{.vulnerabilityAssessmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerVulnerabilityAssessments_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerVulnerabilityAssessments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/vulnerabilityAssessments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerVulnerabilityAssessments_List",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerEncryptionProtector_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/encryptionProtector/{{.encryptionProtectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerEncryptionProtector_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerEncryptionProtector_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/encryptionProtector",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerEncryptionProtector_List",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerEncryptionProtector_Revalidate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/encryptionProtector/{{.encryptionProtectorName}}/revalidate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerEncryptionProtector_Revalidate",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerUsages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlUsages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerUsages_List",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerRecoverableSqlPools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/recoverableSqlPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerRecoverableSqlPools_List",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerRecoverableSqlPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/recoverableSqlPools/{{.sqlPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerRecoverableSqlPools_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/dedicatedSQLminimalTlsSettings/{{.dedicatedSQLminimalTlsSettingsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettings_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/dedicatedSQLminimalTlsSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettings_List",
		Resource:    "Microsoft.Synapse",
	},
}
