package policy

import "github.com/securisec/cliam/shared"

var MediaLivePolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "prod/batch/delete",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchDelete",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/batch/start",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchStart",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/batch/stop",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchStop",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/claimDevice",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ClaimDevice",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/channels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateChannel",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/inputs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateInput",
	},
	{
		Method:        "POST",
		ServiceSuffix: "prod/inputSecurityGroups",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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

	// extra
	{
		ServiceSuffix:          "/prod/channels/{{.channel_id}}",
		Permission:             "DescribeChannel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "channel_id",
	},
	{
		ServiceSuffix:          "/prod/inputs/{{.input_id}}",
		Permission:             "DescribeInput",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_id",
	},
	{
		ServiceSuffix:          "/prod/inputDevices/{{.input_device_id}}",
		Permission:             "DescribeInputDevice",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_device_id",
	},
	{
		ServiceSuffix:          "/prod/inputSecurityGroups/{{.input_security_group_id}}",
		Permission:             "DescribeInputSecurityGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_security_group_id",
	},
	{
		ServiceSuffix:          "/prod/multiplexes/{{.multiplex_id}}",
		Permission:             "DescribeMultiplex",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "multiplex_id",
	},
	{
		ServiceSuffix:          "/prod/offerings/{{.offering_id}}",
		Permission:             "DescribeOffering",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "offering_id",
	},
	{
		ServiceSuffix:          "/prod/reservations/{{.reservation_id}}",
		Permission:             "DescribeReservation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "reservation_id",
	},
	{
		ServiceSuffix:          "/prod/channels/{{.channel_id}}/schedule",
		Permission:             "DescribeSchedule",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "channel_id",
	},
	{
		ServiceSuffix:          "/prod/inputDeviceTransfers",
		Permission:             "ListInputDeviceTransfers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "transfer_type",
	},
	{
		ServiceSuffix:          "/prod/multiplexes/{{.multiplex_id}}/programs",
		Permission:             "ListMultiplexPrograms",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "multiplex_id",
	},
	{
		ServiceSuffix:          "/prod/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
