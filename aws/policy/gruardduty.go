package policy

import "github.com/securisec/cliam/shared"

var GuardDutyPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "invitation/count",
		Permission:    "GetInvitationsCount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "detector",
		Permission:    "ListDetectors",
	},
	{
		Method:        "GET",
		ServiceSuffix: "invitation",
		Permission:    "ListInvitations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "admin",
		Permission:    "ListOrganizationAdminAccounts",
	},

	// extra
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/admin",
		Permission:             "DescribeOrganizationConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}",
		Permission:             "GetDetector",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/master",
		Permission:             "GetMasterAccount",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/filter",
		Permission:             "ListFilters",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListFindings",
			"Version": "2017-11-28",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListFindings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DetectorId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/ipset",
		Permission:             "ListIPSets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/member",
		Permission:             "ListMembers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/publishingDestination",
		Permission:             "ListPublishingDestinations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		ServiceSuffix:          "/detector/{{.detector_id}}/threatintelset",
		Permission:             "ListThreatIntelSets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
}
