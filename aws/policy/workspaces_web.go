package policy

import "github.com/securisec/cliam/shared"

var WorkspacesWebPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "portals",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
