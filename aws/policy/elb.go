package policy

import "github.com/securisec/cliam/shared"

var ELBPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountLimits&Version=2015-12-01",
		Permission:    "DescribeAccountLimits",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeLoadBalancers&Version=2015-12-01",
		Permission:    "DescribeLoadBalancers",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeSSLPolicies&Version=2015-12-01",
		Permission:    "DescribeSSLPolicies",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeTargetGroups&Version=2015-12-01",
		Permission:    "DescribeTargetGroups",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeLoadBalancerPolicies&Version=2012-06-01",
		Permission:    "DescribeLoadBalancerPolicies",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeLoadBalancerPolicyTypes&Version=2012-06-01",
		Permission:    "DescribeLoadBalancerPolicyTypes",
	},

	// extra
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceHealth",
			"Version": "2012-06-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeInstanceHealth",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_name",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoadBalancerAttributes",
			"Version": "2012-06-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLoadBalancerAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_name",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTags",
			"Version": "2012-06-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerNames",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_names",
	},
}
