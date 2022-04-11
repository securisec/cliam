package policy

var Lambda = []Service{
	{
		ServiceSuffix: "2015-03-31/functions",
		Permission:    "ListFunctions",
	},
	{
		ServiceSuffix: "2016-08-19/account-settings/",
		Permission:    "GetAccountSettings",
	},
	{
		ServiceSuffix: "/2015-03-31/event-source-mappings/",
		Permission:    "ListEventSourceMappings",
	},
	{
		ServiceSuffix: "/2018-10-31/layers/",
		Permission:    "ListLayers",
	},
}
