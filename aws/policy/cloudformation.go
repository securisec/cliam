package policy

import "github.com/securisec/cliam/shared"

var CloudformationPolicies = map[string]Service{
	"DescribeAccountLimits": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountLimits&Version=2010-05-15",
		Permission:    "DescribeAccountLimits",
	},
	"DescribePublisher": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribePublisher&Version=2010-05-15",
		Permission:    "DescribePublisher",
	},
	"DescribeType": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeType&Version=2010-05-15",
		Permission:    "DescribeType",
	},
	"GetTemplateSummary": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetTemplateSummary&Version=2010-05-15",
		Permission:    "GetTemplateSummary",
	},
	"ListExports": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListExports&Version=2010-05-15",
		Permission:    "ListExports",
	},
	"ListStackSets": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListStackSets&Version=2010-05-15",
		Permission:    "ListStackSets",
	},
	"ListStacks": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListStacks&Version=2010-05-15",
		Permission:    "ListStacks",
	},
	"ListTypeRegistrations": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypeRegistrations&Version=2010-05-15",
		Permission:    "ListTypeRegistrations",
	},
	"ListTypeVersions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypeVersions&Version=2010-05-15",
		Permission:    "ListTypeVersions",
	},
	"ListTypes": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypes&Version=2010-05-15",
		Permission:    "ListTypes",
	},

	// extra
	"DescribeChangeSet": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeChangeSet",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeChangeSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ChangeSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "change_set_name",
	},
	"DescribeChangeSetHooks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeChangeSetHooks",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeChangeSetHooks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ChangeSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "change_set_name",
	},
	"DescribeStackDriftDetectionStatus": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStackDriftDetectionStatus",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeStackDriftDetectionStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackDriftDetectionId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_drift_detection_id",
	},
	"DescribeStackResourceDrifts": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStackResourceDrifts",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeStackResourceDrifts",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_name",
	},
	"DescribeStackSet": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStackSet",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeStackSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_set_name",
	},
	"DescribeTypeRegistration": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTypeRegistration",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTypeRegistration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RegistrationToken",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "registration_token",
	},
	"GetStackPolicy": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetStackPolicy",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetStackPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_name",
	},
	"ListChangeSets": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListChangeSets",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListChangeSets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_name",
	},
	"ListImports": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListImports",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListImports",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ExportName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "export_name",
	},
	"ListStackInstances": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListStackInstances",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListStackInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_set_name",
	},
	"ListStackResources": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListStackResources",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListStackResources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_name",
	},
	"ListStackSetOperations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListStackSetOperations",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListStackSetOperations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "stack_set_name",
	},
}
