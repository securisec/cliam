package policy

var SNSPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetSMSAttributes&Version=2010-03-31",
		Permission:    "GetSMSAttributes",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetSMSSandboxAccountStatus&Version=2010-03-31",
		Permission:    "GetSMSSandboxAccountStatus",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListOriginationNumbers&Version=2010-03-31",
		Permission:    "ListOriginationNumbers",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPhoneNumbersOptedOut&Version=2010-03-31",
		Permission:    "ListPhoneNumbersOptedOut",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPlatformApplications&Version=2010-03-31",
		Permission:    "ListPlatformApplications",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListSMSSandboxPhoneNumbers&Version=2010-03-31",
		Permission:    "ListSMSSandboxPhoneNumbers",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListSubscriptions&Version=2010-03-31",
		Permission:    "ListSubscriptions",
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTopics&Version=2010-03-31",
		Permission:    "ListTopics",
	},
}
