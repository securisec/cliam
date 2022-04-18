package policy

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
}
