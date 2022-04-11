package policy

var SQSPolicies = []Service{
	{
		ServiceSuffix: "?Action=ListQueues",
		Permission:    "ListQueues",
	},
}
