package policy

var SQS = []Service{
	{
		ServiceSuffix: "?Action=ListQueues",
		Permission:    "ListQueues",
	},
}
