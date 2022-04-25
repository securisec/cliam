package policy

import "github.com/securisec/cliam/shared"

var CloudformationPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountLimits&Version=2010-05-15",
		Permission:    "DescribeAccountLimits",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribePublisher&Version=2010-05-15",
		Permission:    "DescribePublisher",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeType&Version=2010-05-15",
		Permission:    "DescribeType",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetTemplateSummary&Version=2010-05-15",
		Permission:    "GetTemplateSummary",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListExports&Version=2010-05-15",
		Permission:    "ListExports",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListStackSets&Version=2010-05-15",
		Permission:    "ListStackSets",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListStacks&Version=2010-05-15",
		Permission:    "ListStacks",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypeRegistrations&Version=2010-05-15",
		Permission:    "ListTypeRegistrations",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypeVersions&Version=2010-05-15",
		Permission:    "ListTypeVersions",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypes&Version=2010-05-15",
		Permission:    "ListTypes",
	},

	// extra
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
