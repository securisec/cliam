package policy

var IAM = []Service{
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListRoles&Version=2010-05-08",
		Permission:    "ListRoles",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountAuthorizationDetails&Version=2010-05-08",
		Permission:    "GetAccountAuthorizationDetails",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountPasswordPolicy&Version=2010-05-08",
		Permission:    "GetAccountPasswordPolicy",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountSummary&Version=2010-05-08",
		Permission:    "GetAccountSummary",
	},
}
