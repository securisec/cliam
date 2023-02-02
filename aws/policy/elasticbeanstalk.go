package policy

import "github.com/securisec/cliam/shared"

// ElasticBeanStalkPolicies policy
var ElasticBeanStalkPolicies = map[string]Service{
	"AbortEnvironmentUpdate": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AbortEnvironmentUpdate",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AbortEnvironmentUpdate",
	},
	"ComposeEnvironments": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ComposeEnvironments",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ComposeEnvironments",
	},
	"DescribeAccountAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountAttributes",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountAttributes",
	},
	"DescribeApplicationVersions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeApplicationVersions",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeApplicationVersions",
	},
	"DescribeApplications": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeApplications",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeApplications",
	},
	"DescribeConfigurationOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConfigurationOptions",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeConfigurationOptions",
	},
	"DescribeEnvironmentHealth": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEnvironmentHealth",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEnvironmentHealth",
	},
	"DescribeEnvironmentManagedActionHistory": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEnvironmentManagedActionHistory",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEnvironmentManagedActionHistory",
	},
	"DescribeEnvironmentManagedActions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEnvironmentManagedActions",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEnvironmentManagedActions",
	},
	"DescribeEnvironmentResources": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEnvironmentResources",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEnvironmentResources",
	},
	"DescribeEnvironments": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEnvironments",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEnvironments",
	},
	"DescribeEvents": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEvents",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEvents",
	},
	"DescribeInstancesHealth": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstancesHealth",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstancesHealth",
	},
	"DescribePlatformVersion": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePlatformVersion",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePlatformVersion",
	},
	"ListAvailableSolutionStacks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListAvailableSolutionStacks",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListAvailableSolutionStacks",
	},
	"ListPlatformBranches": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListPlatformBranches",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListPlatformBranches",
	},
	"ListPlatformVersions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListPlatformVersions",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListPlatformVersions",
	},
	"RebuildEnvironment": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RebuildEnvironment",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RebuildEnvironment",
	},
	"RestartAppServer": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RestartAppServer",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RestartAppServer",
	},
	"SwapEnvironmentCNAMEs": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "SwapEnvironmentCNAMEs",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "SwapEnvironmentCNAMEs",
	},
	"TerminateEnvironment": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "TerminateEnvironment",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "TerminateEnvironment",
	},
	"UpdateEnvironment": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "UpdateEnvironment",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "UpdateEnvironment",
	},

	// extra
	"DescribeConfigurationSettings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConfigurationSettings",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeConfigurationSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ApplicationName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "application_name",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
