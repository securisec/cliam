package policy

import "github.com/securisec/cliam/shared"

var BackupPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "global-settings",
		Permission:    "DescribeGlobalSettings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "account-settings",
		Permission:    "DescribeRegionSettings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "backup-jobs/",
		Permission:    "ListBackupJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "backup/template/plans",
		Permission:    "ListBackupPlanTemplates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "backup/plans/",
		Permission:    "ListBackupPlans",
	},
	{
		Method:        "GET",
		ServiceSuffix: "backup-vaults/",
		Permission:    "ListBackupVaults",
	},
	{
		Method:        "GET",
		ServiceSuffix: "copy-jobs/",
		Permission:    "ListCopyJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "audit/frameworks",
		Permission:    "ListFrameworks",
	},
	{
		Method:        "GET",
		ServiceSuffix: "resources/",
		Permission:    "ListProtectedResources",
	},
	{
		Method:        "GET",
		ServiceSuffix: "audit/report-jobs",
		Permission:    "ListReportJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "audit/report-plans",
		Permission:    "ListReportPlans",
	},
	{
		Method:        "GET",
		ServiceSuffix: "restore-jobs/",
		Permission:    "ListRestoreJobs",
	},

	// extra
	{
		ServiceSuffix:          "/backup-jobs/{{.backup_job_id}}",
		Permission:             "DescribeBackupJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_job_id",
	},
	{
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}",
		Permission:             "DescribeBackupVault",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	{
		ServiceSuffix:          "/copy-jobs/{{.copy_job_id}}",
		Permission:             "DescribeCopyJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "copy_job_id",
	},
	{
		ServiceSuffix:          "/audit/frameworks/{{.framework_name}}",
		Permission:             "DescribeFramework",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "framework_name",
	},
	{
		ServiceSuffix:          "/resources/{{.resource_arn}}",
		Permission:             "DescribeProtectedResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		ServiceSuffix:          "/audit/report-jobs/{{.report_job_id}}",
		Permission:             "DescribeReportJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "report_job_id",
	},
	{
		ServiceSuffix:          "/audit/report-plans/{{.report_plan_name}}",
		Permission:             "DescribeReportPlan",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "report_plan_name",
	},
	{
		ServiceSuffix:          "/restore-jobs/{{.restore_job_id}}",
		Permission:             "DescribeRestoreJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "restore_job_id",
	},
	{
		ServiceSuffix:          "/backup/plans/{{.backup_plan_id}}/",
		Permission:             "GetBackupPlan",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_id",
	},
	{
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
	{
		ServiceSuffix:          "/backup/template/plans/{{.backup_plan_template_id}}/toPlan",
		Permission:             "GetBackupPlanFromTemplate",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_template_id",
	},
	{
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}/access-policy",
		Permission:             "GetBackupVaultAccessPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	{
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}/notification-configuration",
		Permission:             "GetBackupVaultNotifications",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	{
		ServiceSuffix:          "/backup/plans/{{.backup_plan_id}}/versions/",
		Permission:             "ListBackupPlanVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_id",
	},
	{
		ServiceSuffix:          "/backup/plans/{{.backup_plan_id}}/selections/",
		Permission:             "ListBackupSelections",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_plan_id",
	},
	{
		ServiceSuffix:          "/backup-vaults/{{.backup_vault_name}}/recovery-points/",
		Permission:             "ListRecoveryPointsByBackupVault",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "backup_vault_name",
	},
	{
		ServiceSuffix:          "/resources/{{.resource_arn}}/recovery-points/",
		Permission:             "ListRecoveryPointsByResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}/",
		Permission:             "ListTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
