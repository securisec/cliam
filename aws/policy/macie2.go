package policy

import "github.com/securisec/cliam/shared"

var Macie2Policies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "custom-data-identifiers/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchGetCustomDataIdentifiers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings/sample",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateSampleFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "datasources/s3",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromAdministratorAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "master/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromMasterAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "macie",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListClassificationJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "custom-data-identifiers/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListCustomDataIdentifiers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "SearchResources",
	},

	// extra
	{
		ServiceSuffix:          "/jobs/{{.job_id}}",
		Permission:             "DescribeClassificationJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	{
		ServiceSuffix:          "/custom-data-identifiers/{{.id}}",
		Permission:             "GetCustomDataIdentifier",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetFindingStatistics",
			"Version": "2020-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetFindingStatistics",
		IsExtra:                true,
		ExtraComponentBodyKey:  "groupBy",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "group_by",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetFindings",
			"Version": "2020-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetFindings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "findingIds",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "finding_ids",
	},
	{
		ServiceSuffix:          "/findingsfilters/{{.id}}",
		Permission:             "GetFindingsFilter",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/members/{{.id}}",
		Permission:             "GetMember",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
