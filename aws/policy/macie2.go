package policy

import "github.com/securisec/cliam/shared"

var Macie2Policies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "custom-data-identifiers/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "BatchGetCustomDataIdentifiers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings/sample",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateSampleFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "datasources/s3",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeBuckets",
	},
	{
		Method:        "GET",
		ServiceSuffix: "admin/configuration",
		Permission:    "DescribeOrganizationConfiguration",
	},
	{
		Method:        "POST",
		ServiceSuffix: "administrator/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DisassociateFromAdministratorAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "master/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DisassociateFromMasterAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "macie",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "EnableMacie",
	},
	{
		Method:        "GET",
		ServiceSuffix: "administrator",
		Permission:    "GetAdministratorAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "datasources/s3/statistics",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetBucketStatistics",
	},
	{
		Method:        "GET",
		ServiceSuffix: "classification-export-configuration",
		Permission:    "GetClassificationExportConfiguration",
	},
	{
		Method:        "GET",
		ServiceSuffix: "findings-publication-configuration",
		Permission:    "GetFindingsPublicationConfiguration",
	},
	{
		Method:        "GET",
		ServiceSuffix: "invitations/count",
		Permission:    "GetInvitationsCount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "macie",
		Permission:    "GetMacieSession",
	},
	{
		Method:        "GET",
		ServiceSuffix: "master",
		Permission:    "GetMasterAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "usage/statistics",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetUsageStatistics",
	},
	{
		Method:        "GET",
		ServiceSuffix: "usage",
		Permission:    "GetUsageTotals",
	},
	{
		Method:        "POST",
		ServiceSuffix: "jobs/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListClassificationJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "custom-data-identifiers/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListCustomDataIdentifiers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListFindings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "findingsfilters",
		Permission:    "ListFindingsFilters",
	},
	{
		Method:        "GET",
		ServiceSuffix: "invitations",
		Permission:    "ListInvitations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "managed-data-identifiers/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListManagedDataIdentifiers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "members",
		Permission:    "ListMembers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "admin",
		Permission:    "ListOrganizationAdminAccounts",
	},
	{
		Method:        "POST",
		ServiceSuffix: "datasources/search-resources",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "SearchResources",
	},
}
