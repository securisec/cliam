package policy

var ELBPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeAccountLimits&Version=2015-12-01",
		Permission:    "DescribeAccountLimits",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeLoadBalancers&Version=2015-12-01",
		Permission:    "DescribeLoadBalancers",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeSSLPolicies&Version=2015-12-01",
		Permission:    "DescribeSSLPolicies",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeTargetGroups&Version=2015-12-01",
		Permission:    "DescribeTargetGroups",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeLoadBalancerPolicies&Version=2012-06-01",
		Permission:    "DescribeLoadBalancerPolicies",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeLoadBalancerPolicyTypes&Version=2012-06-01",
		Permission:    "DescribeLoadBalancerPolicyTypes",
	},
}
