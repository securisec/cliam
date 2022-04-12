package policy

var CloudformationPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeAccountLimits&Version=2010-05-15",
		Permission:    "DescribeAccountLimits",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribePublisher&Version=2010-05-15",
		Permission:    "DescribePublisher",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeType&Version=2010-05-15",
		Permission:    "DescribeType",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=GetTemplateSummary&Version=2010-05-15",
		Permission:    "GetTemplateSummary",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListExports&Version=2010-05-15",
		Permission:    "ListExports",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListStackSets&Version=2010-05-15",
		Permission:    "ListStackSets",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListStacks&Version=2010-05-15",
		Permission:    "ListStacks",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListTypeRegistrations&Version=2010-05-15",
		Permission:    "ListTypeRegistrations",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListTypeVersions&Version=2010-05-15",
		Permission:    "ListTypeVersions",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListTypes&Version=2010-05-15",
		Permission:    "ListTypes",
	},
}
