package policy

import "github.com/securisec/cliam/shared"

var MediaLivePolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "prod/batch/delete",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "BatchDelete",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/batch/start",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "BatchStart",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/batch/stop",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "BatchStop",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/claimDevice",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ClaimDevice",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/channels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateChannel",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/inputs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateInput",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/inputSecurityGroups",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateInputSecurityGroup",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/channels",
		Permission:    "ListChannels",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/inputDevices",
		Permission:    "ListInputDevices",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/inputSecurityGroups",
		Permission:    "ListInputSecurityGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/inputs",
		Permission:    "ListInputs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/multiplexes",
		Permission:    "ListMultiplexes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/offerings",
		Permission:    "ListOfferings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "prod/reservations",
		Permission:    "ListReservations",
	},
}
