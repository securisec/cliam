package policy

import "github.com/securisec/cliam/shared"

// GuardDutyPolicies policy
var GuardDutyPolicies = map[string]Service{
	"GetInvitationsCount": {
		Method:        "GET",
		ServiceSuffix: "invitation/count",
		Permission:    "GetInvitationsCount",
	},
	"ListDetectors": {
		Method:        "GET",
		ServiceSuffix: "detector",
		Permission:    "ListDetectors",
	},
	"ListInvitations": {
		Method:        "GET",
		ServiceSuffix: "invitation",
		Permission:    "ListInvitations",
	},
	"ListOrganizationAdminAccounts": {
		Method:        "GET",
		ServiceSuffix: "admin",
		Permission:    "ListOrganizationAdminAccounts",
	},

	// extra
	"DescribeMalwareScans": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeMalwareScans",
			"Version": "2017-11-28",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeMalwareScans",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DetectorId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "detector_id",
	},
	"DescribeOrganizationConfiguration": {
		ServiceSuffix:          "/detector/{{.detector_id}}/admin",
		Permission:             "DescribeOrganizationConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"GetAdministratorAccount": {
		ServiceSuffix:          "/detector/{{.detector_id}}/administrator",
		Permission:             "GetAdministratorAccount",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"GetDetector": {
		ServiceSuffix:          "/detector/{{.detector_id}}",
		Permission:             "GetDetector",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"GetMalwareScanSettings": {
		ServiceSuffix:          "/detector/{{.detector_id}}/malware-scan-settings",
		Permission:             "GetMalwareScanSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"GetMasterAccount": {
		ServiceSuffix:          "/detector/{{.detector_id}}/master",
		Permission:             "GetMasterAccount",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"GetRemainingFreeTrialDays": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetRemainingFreeTrialDays",
			"Version": "2017-11-28",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetRemainingFreeTrialDays",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DetectorId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "detector_id",
	},
	"ListFilters": {
		ServiceSuffix:          "/detector/{{.detector_id}}/filter",
		Permission:             "ListFilters",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"ListFindings": {
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
	"ListIPSets": {
		ServiceSuffix:          "/detector/{{.detector_id}}/ipset",
		Permission:             "ListIPSets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"ListMembers": {
		ServiceSuffix:          "/detector/{{.detector_id}}/member",
		Permission:             "ListMembers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"ListPublishingDestinations": {
		ServiceSuffix:          "/detector/{{.detector_id}}/publishingDestination",
		Permission:             "ListPublishingDestinations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListThreatIntelSets": {
		ServiceSuffix:          "/detector/{{.detector_id}}/threatintelset",
		Permission:             "ListThreatIntelSets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_id",
	},
}
