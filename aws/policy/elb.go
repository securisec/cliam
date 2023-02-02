package policy

import "github.com/securisec/cliam/shared"

// ELBPolicies policy
var ELBPolicies = map[string]Service{
	"DescribeAccountLimits": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountLimits",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountLimits",
	},
	"DescribeListeners": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeListeners",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeListeners",
	},
	"DescribeLoadBalancers": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoadBalancers",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLoadBalancers",
	},
	"DescribeRules": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeRules",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeRules",
	},
	"DescribeSSLPolicies": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSSLPolicies",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSSLPolicies",
	},
	"DescribeTargetGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTargetGroups",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTargetGroups",
	},

	// extra
	"DescribeListenerCertificates": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeListenerCertificates",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeListenerCertificates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ListenerArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "listener_arn",
	},
	"DescribeLoadBalancerAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoadBalancerAttributes",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLoadBalancerAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_arn",
	},
	"DescribeTags": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTags",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArns",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arns",
	},
	"DescribeTargetGroupAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTargetGroupAttributes",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTargetGroupAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TargetGroupArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "target_group_arn",
	},
	"DescribeTargetHealth": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTargetHealth",
			"Version": "2015-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTargetHealth",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TargetGroupArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "target_group_arn",
	},
}
