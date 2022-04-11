package policy

var ECRPolicies = []Service{
	{
		ServiceSuffix: "",
		Permission:    "DescribeRepositories",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			"Content-Type": aws_JSON_CONTENT_TYPE,
			"x-amz-target": "AmazonEC2ContainerRegistry_V20150921.DescribeRepositories",
		},
	},
	{
		ServiceSuffix: "",
		Permission:    "GetAuthorizationToken",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			"Content-Type": aws_JSON_CONTENT_TYPE,
			"x-amz-target": "AmazonEC2ContainerRegistry_V20150921.GetAuthorizationToken",
		},
	},
}
