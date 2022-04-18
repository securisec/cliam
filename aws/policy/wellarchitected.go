package policy

import "github.com/securisec/cliam/shared"

var WellarchitectedPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "lenses",
		Permission:    "ListLenses",
	},
	{
		Method:        "POST",
		ServiceSuffix: "notifications",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListNotifications",
	},
	{
		Method:        "GET",
		ServiceSuffix: "shareInvitations",
		Permission:    "ListShareInvitations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "workloadsSummaries",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListWorkloads",
	},
}
