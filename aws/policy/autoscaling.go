package policy

import "github.com/securisec/cliam/shared"

var AutoscalingPolicies = map[string]Service{
	"DescribeAccountLimits": {
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
	"DescribeAdjustmentTypes": {
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
	"DescribeAutoScalingGroups": {
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
	"DescribeAutoScalingInstances": {
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
	"DescribeAutoScalingNotificationTypes": {
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
	"DescribeLaunchConfigurations": {
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
	"DescribeLifecycleHookTypes": {
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
	"DescribeMetricCollectionTypes": {
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
	"DescribeNotificationConfigurations": {
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
	"DescribePolicies": {
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
	"DescribeScalingActivities": {
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
	"DescribeScalingProcessTypes": {
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
	"DescribeScheduledActions": {
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
	"DescribeTags": {
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
	"DescribeTerminationPolicyTypes": {
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

	// extra
	"DescribeInstanceRefreshes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceRefreshes",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeInstanceRefreshes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoScalingGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "auto_scaling_group_name",
	},
	"DescribeLifecycleHooks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLifecycleHooks",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLifecycleHooks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoScalingGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "auto_scaling_group_name",
	},
	"DescribeLoadBalancerTargetGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoadBalancerTargetGroups",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLoadBalancerTargetGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoScalingGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "auto_scaling_group_name",
	},
	"DescribeLoadBalancers": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoadBalancers",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLoadBalancers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoScalingGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "auto_scaling_group_name",
	},
	"DescribeWarmPool": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeWarmPool",
			"Version": "2011-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeWarmPool",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoScalingGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "auto_scaling_group_name",
	},
}
