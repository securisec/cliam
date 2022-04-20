package policy

var LambdaPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "2016-08-19/account-settings/",
		Permission:    "GetAccountSettings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2020-04-22/code-signing-configs/",
		Permission:    "ListCodeSigningConfigs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2015-03-31/event-source-mappings/",
		Permission:    "ListEventSourceMappings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2015-03-31/functions/",
		Permission:    "ListFunctions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2018-10-31/layers",
		Permission:    "ListLayers",
	},
}
