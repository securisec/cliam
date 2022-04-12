package policy

var SNSPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=GetSMSAttributes&Version=2010-03-31",
		Permission:    "GetSMSAttributes",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=GetSMSSandboxAccountStatus&Version=2010-03-31",
		Permission:    "GetSMSSandboxAccountStatus",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListOriginationNumbers&Version=2010-03-31",
		Permission:    "ListOriginationNumbers",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListPhoneNumbersOptedOut&Version=2010-03-31",
		Permission:    "ListPhoneNumbersOptedOut",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListPlatformApplications&Version=2010-03-31",
		Permission:    "ListPlatformApplications",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListSMSSandboxPhoneNumbers&Version=2010-03-31",
		Permission:    "ListSMSSandboxPhoneNumbers",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListSubscriptions&Version=2010-03-31",
		Permission:    "ListSubscriptions",
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListTopics&Version=2010-03-31",
		Permission:    "ListTopics",
	},
}
