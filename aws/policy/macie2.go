package policy

import "github.com/securisec/cliam/shared"

var Macie2Policies = map[string]Service{
	"BatchGetCustomDataIdentifiers": {
		Method:        "POST",
		ServiceSuffix: "custom-data-identifiers/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchGetCustomDataIdentifiers",
	},
	"CreateSampleFindings": {
		Method:        "POST",
		ServiceSuffix: "findings/sample",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateSampleFindings",
	},
	"DescribeBuckets": {
		Method:        "POST",
		ServiceSuffix: "datasources/s3",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeBuckets",
	},
	"DescribeOrganizationConfiguration": {
		Method:        "GET",
		ServiceSuffix: "admin/configuration",
		Permission:    "DescribeOrganizationConfiguration",
	},
	"DisassociateFromAdministratorAccount": {
		Method:        "POST",
		ServiceSuffix: "administrator/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromAdministratorAccount",
	},
	"DisassociateFromMasterAccount": {
		Method:        "POST",
		ServiceSuffix: "master/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromMasterAccount",
	},
	"EnableMacie": {
		Method:        "POST",
		ServiceSuffix: "macie",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "EnableMacie",
	},
	"GetAdministratorAccount": {
		Method:        "GET",
		ServiceSuffix: "administrator",
		Permission:    "GetAdministratorAccount",
	},
	"GetBucketStatistics": {
		Method:        "POST",
		ServiceSuffix: "datasources/s3/statistics",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetBucketStatistics",
	},
	"GetClassificationExportConfiguration": {
		Method:        "GET",
		ServiceSuffix: "classification-export-configuration",
		Permission:    "GetClassificationExportConfiguration",
	},
	"GetFindingsPublicationConfiguration": {
		Method:        "GET",
		ServiceSuffix: "findings-publication-configuration",
		Permission:    "GetFindingsPublicationConfiguration",
	},
	"GetInvitationsCount": {
		Method:        "GET",
		ServiceSuffix: "invitations/count",
		Permission:    "GetInvitationsCount",
	},
	"GetMacieSession": {
		Method:        "GET",
		ServiceSuffix: "macie",
		Permission:    "GetMacieSession",
	},
	"GetMasterAccount": {
		Method:        "GET",
		ServiceSuffix: "master",
		Permission:    "GetMasterAccount",
	},
	"GetUsageStatistics": {
		Method:        "POST",
		ServiceSuffix: "usage/statistics",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetUsageStatistics",
	},
	"GetUsageTotals": {
		Method:        "GET",
		ServiceSuffix: "usage",
		Permission:    "GetUsageTotals",
	},
	"ListClassificationJobs": {
		Method:        "POST",
		ServiceSuffix: "jobs/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListClassificationJobs",
	},
	"ListCustomDataIdentifiers": {
		Method:        "POST",
		ServiceSuffix: "custom-data-identifiers/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListCustomDataIdentifiers",
	},
	"ListFindings": {
		Method:        "POST",
		ServiceSuffix: "findings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFindings",
	},
	"ListFindingsFilters": {
		Method:        "GET",
		ServiceSuffix: "findingsfilters",
		Permission:    "ListFindingsFilters",
	},
	"ListInvitations": {
		Method:        "GET",
		ServiceSuffix: "invitations",
		Permission:    "ListInvitations",
	},
	"ListManagedDataIdentifiers": {
		Method:        "POST",
		ServiceSuffix: "managed-data-identifiers/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListManagedDataIdentifiers",
	},
	"ListMembers": {
		Method:        "GET",
		ServiceSuffix: "members",
		Permission:    "ListMembers",
	},
	"ListOrganizationAdminAccounts": {
		Method:        "GET",
		ServiceSuffix: "admin",
		Permission:    "ListOrganizationAdminAccounts",
	},
	"SearchResources": {
		Method:        "POST",
		ServiceSuffix: "datasources/search-resources",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "SearchResources",
	},

	// extra
	"DescribeClassificationJob": {
		ServiceSuffix:          "/jobs/{{.job_id}}",
		Permission:             "DescribeClassificationJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	"GetCustomDataIdentifier": {
		ServiceSuffix:          "/custom-data-identifiers/{{.id}}",
		Permission:             "GetCustomDataIdentifier",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetFindingStatistics": {
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
	"GetFindings": {
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
	"GetFindingsFilter": {
		ServiceSuffix:          "/findingsfilters/{{.id}}",
		Permission:             "GetFindingsFilter",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetMember": {
		ServiceSuffix:          "/members/{{.id}}",
		Permission:             "GetMember",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
