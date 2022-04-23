package policy

var CloudformationPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountLimits&Version=2010-05-15",
		Permission:    "DescribeAccountLimits",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribePublisher&Version=2010-05-15",
		Permission:    "DescribePublisher",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeType&Version=2010-05-15",
		Permission:    "DescribeType",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetTemplateSummary&Version=2010-05-15",
		Permission:    "GetTemplateSummary",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListExports&Version=2010-05-15",
		Permission:    "ListExports",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListStackSets&Version=2010-05-15",
		Permission:    "ListStackSets",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListStacks&Version=2010-05-15",
		Permission:    "ListStacks",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypeRegistrations&Version=2010-05-15",
		Permission:    "ListTypeRegistrations",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypeVersions&Version=2010-05-15",
		Permission:    "ListTypeVersions",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTypes&Version=2010-05-15",
		Permission:    "ListTypes",
	},
}
