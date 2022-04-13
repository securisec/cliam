package policy

import "github.com/securisec/cliam/shared"

var AutoscalingPolicies = []Service{
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountLimits",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountLimits",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAdjustmentTypes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAdjustmentTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAutoScalingGroups",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAutoScalingGroups",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAutoScalingInstances",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAutoScalingInstances",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAutoScalingNotificationTypes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAutoScalingNotificationTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLaunchConfigurations",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLaunchConfigurations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLifecycleHookTypes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLifecycleHookTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeMetricCollectionTypes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeMetricCollectionTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNotificationConfigurations",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNotificationConfigurations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePolicies",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePolicies",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeScalingActivities",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeScalingActivities",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeScalingProcessTypes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeScalingProcessTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeScheduledActions",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeScheduledActions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTags",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTags",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTerminationPolicyTypes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTerminationPolicyTypes",
	},
}
