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
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeEventSourcesConfig",
	},
	{
		Method:        "POST",
		ServiceSuffix: "feedback",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeFeedback",
	},
	{
		Method:        "POST",
		ServiceSuffix: "organization/health",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListNotificationChannels",
	},
}
