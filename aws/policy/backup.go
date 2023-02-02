package policy

import "github.com/securisec/cliam/shared"

// BackupPolicies policy
var BackupPolicies = map[string]Service{
	"DescribeGlobalSettings": {
		Method:        "GET",
		ServiceSuffix: "global-settings",
		Permission:    "DescribeGlobalSettings",
	},
	"DescribeRegionSettings": {
		Method:        "GET",
		ServiceSuffix: "account-settings",
		Permission:    "DescribeRegionSettings",
	},
	"GetSupportedResourceTypes": {
		Method:        "GET",
		ServiceSuffix: "supported-resource-types",
		Permission:    "GetSupportedResourceTypes",
	},
	"ListBackupJobs": {
		Method:        "GET",
		ServiceSuffix: "backup-jobs/",
		Permission:    "ListBackupJobs",
	},
	"ListBackupPlanTemplates": {
		Method:        "GET",
		ServiceSuffix: "backup/template/plans",
		Permission:    "ListBackupPlanTemplates",
	},
	"ListBackupPlans": {
		Method:        "GET",
		ServiceSuffix: "backup/plans/",
		Permission:    "ListBackupPlans",
	},
	"ListBackupVaults": {
		Method:        "GET",
		ServiceSuffix: "backup-vaults/",
		Permission:    "ListBackupVaults",
	},
	"ListCopyJobs": {
		Method:        "GET",
		ServiceSuffix: "copy-jobs/",
		Permission:    "ListCopyJobs",
	},
	"ListFrameworks": {
		Method:        "GET",
		ServiceSuffix: "audit/frameworks",
		Permission:    "ListFrameworks",
	},
	"ListLegalHolds": {
		Method:        "GET",
		ServiceSuffix: "legal-holds/",
		Permission:    "ListLegalHolds",
	},
	"ListProtectedResources": {
		Method:        "GET",
		ServiceSuffix: "resources/",
		Permission:    "ListProtectedResources",
	},
	"ListReportJobs": {
		Method:        "GET",
		ServiceSuffix: "audit/report-jobs",
		Permission:    "ListReportJobs",
	},
	"ListReportPlans": {
		Method:        "GET",
		ServiceSuffix: "audit/report-plans",
		Permission:    "ListReportPlans",
	},
	"ListRestoreJobs": {
		Method:        "GET",
		ServiceSuffix: "restore-jobs/",
		Permission:    "ListRestoreJobs",
	},

	// extra
	"DescribeBackupJob": {
		ServiceSuffix:          "/backup-jobs/{{.backup_job_id}}",
		Permission:             "DescribeBackupJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_job_id",
	},
	"DescribeBackupVault": {
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}",
		Permission:             "DescribeBackupVault",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	"DescribeCopyJob": {
		ServiceSuffix:          "/copy-jobs/{{.copy_job_id}}",
		Permission:             "DescribeCopyJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "copy_job_id",
	},
	"DescribeFramework": {
		ServiceSuffix:          "/audit/frameworks/{{.framework_name}}",
		Permission:             "DescribeFramework",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "framework_name",
	},
	"DescribeProtectedResource": {
		ServiceSuffix:          "/resources/{{.resource_arn}}",
		Permission:             "DescribeProtectedResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"DescribeReportJob": {
		ServiceSuffix:          "/audit/report-jobs/{{.report_job_id}}",
		Permission:             "DescribeReportJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "report_job_id",
	},
	"DescribeReportPlan": {
		ServiceSuffix:          "/audit/report-plans/{{.report_plan_name}}",
		Permission:             "DescribeReportPlan",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "report_plan_name",
	},
	"DescribeRestoreJob": {
		ServiceSuffix:          "/restore-jobs/{{.restore_job_id}}",
		Permission:             "DescribeRestoreJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "restore_job_id",
	},
	"GetBackupPlan": {
		ServiceSuffix:          "/backup/plans/{{.backup_plan_id}}/",
		Permission:             "GetBackupPlan",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_id",
	},
	"GetBackupPlanFromJSON": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetBackupPlanFromJSON",
			"Version": "2018-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetBackupPlanFromJSON",
		IsExtra:                true,
		ExtraComponentBodyKey:  "BackupPlanTemplateJson",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "backup_plan_template_json",
	},
	"GetBackupPlanFromTemplate": {
		ServiceSuffix:          "/backup/template/plans/{{.backup_plan_template_id}}/toPlan",
		Permission:             "GetBackupPlanFromTemplate",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_template_id",
	},
	"GetBackupVaultAccessPolicy": {
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}/access-policy",
		Permission:             "GetBackupVaultAccessPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	"GetBackupVaultNotifications": {
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}/notification-configuration",
		Permission:             "GetBackupVaultNotifications",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	"GetLegalHold": {
		ServiceSuffix:          "/legal-holds/{{.legal_hold_id}}/",
		Permission:             "GetLegalHold",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "legal_hold_id",
	},
	"ListBackupPlanVersions": {
		ServiceSuffix:          "/backup/plans/{{.backup_plan_id}}/versions/",
		Permission:             "ListBackupPlanVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_id",
	},
	"ListBackupSelections": {
		ServiceSuffix:          "/backup/plans/{{.backup_plan_id}}/selections/",
		Permission:             "ListBackupSelections",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_id",
	},
	"ListRecoveryPointsByBackupVault": {
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}/recovery-points/",
		Permission:             "ListRecoveryPointsByBackupVault",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	"ListRecoveryPointsByLegalHold": {
		ServiceSuffix:          "/legal-holds/{{.legal_hold_id}}/recovery-points",
		Permission:             "ListRecoveryPointsByLegalHold",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "legal_hold_id",
	},
	"ListRecoveryPointsByResource": {
		ServiceSuffix:          "/resources/{{.resource_arn}}/recovery-points/",
		Permission:             "ListRecoveryPointsByResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTags": {
		ServiceSuffix:          "/tags/{{.resource_arn}}/",
		Permission:             "ListTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
