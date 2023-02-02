package policy

import "github.com/securisec/cliam/shared"

// CloudFormationPolicies policy
var CloudFormationPolicies = map[string]Service{
	"ActivateType": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ActivateType",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ActivateType",
	},
	"DeactivateType": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeactivateType",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeactivateType",
	},
	"DeregisterType": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeregisterType",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeregisterType",
	},
	"DescribeAccountLimits": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountLimits",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountLimits",
	},
	"DescribePublisher": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePublisher",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePublisher",
	},
	"DescribeStackEvents": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStackEvents",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeStackEvents",
	},
	"DescribeStackResources": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStackResources",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeStackResources",
	},
	"DescribeStacks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStacks",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeStacks",
	},
	"DescribeType": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeType",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeType",
	},
	"EstimateTemplateCost": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "EstimateTemplateCost",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "EstimateTemplateCost",
	},
	"GetTemplate": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTemplate",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetTemplate",
	},
	"GetTemplateSummary": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTemplateSummary",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetTemplateSummary",
	},
	"ListExports": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListExports",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListExports",
	},
	"ListStackSets": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListStackSets",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListStackSets",
	},
	"ListStacks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListStacks",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListStacks",
	},
	"ListTypeRegistrations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTypeRegistrations",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListTypeRegistrations",
	},
	"ListTypeVersions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTypeVersions",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListTypeVersions",
	},
	"ListTypes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTypes",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListTypes",
	},
	"PublishType": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "PublishType",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "PublishType",
	},
	"RegisterPublisher": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RegisterPublisher",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RegisterPublisher",
	},
	"SetTypeDefaultVersion": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "SetTypeDefaultVersion",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "SetTypeDefaultVersion",
	},
	"TestType": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "TestType",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "TestType",
	},
	"ValidateTemplate": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ValidateTemplate",
			"Version": "2010-05-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ValidateTemplate",
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
