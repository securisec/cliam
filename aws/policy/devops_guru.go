package policy

import "github.com/securisec/cliam/shared"

var DevopsGuruPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "accounts/health",
		Permission:    "DescribeAccountHealth",
	},
	{
		Method:        "POST",
		ServiceSuffix: "event-sources",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeEventSourcesConfig",
	},
	{
		Method:        "POST",
		ServiceSuffix: "feedback",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeFeedback",
	},
	{
		Method:        "POST",
		ServiceSuffix: "organization/health",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeOrganizationHealth",
	},
	{
		Method:        "GET",
		ServiceSuffix: "service-integrations",
		Permission:    "DescribeServiceIntegration",
	},
	{
		Method:        "GET",
		ServiceSuffix: "cost-estimation",
		Permission:    "GetCostEstimation",
	},
	{
		Method:        "POST",
		ServiceSuffix: "channels",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListNotificationChannels",
	},
}
