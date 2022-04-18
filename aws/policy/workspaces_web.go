package policy

import "github.com/securisec/cliam/shared"

var WorkspacesWebPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "portals",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreatePortal",
	},
	{
		Method:        "GET",
		ServiceSuffix: "browserSettings",
		Permission:    "ListBrowserSettings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "networkSettings",
		Permission:    "ListNetworkSettings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "portals",
		Permission:    "ListPortals",
	},
	{
		Method:        "GET",
		ServiceSuffix: "trustStores",
		Permission:    "ListTrustStores",
	},
	{
		Method:        "GET",
		ServiceSuffix: "userSettings",
		Permission:    "ListUserSettings",
	},
}
