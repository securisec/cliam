package policy

import "github.com/securisec/cliam/shared"

var MediaLivePolicies = map[string]Service{
	"BatchDelete": {
		Method:        "POST",
		ServiceSuffix: "prod/batch/delete",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchDelete",
	},
	"BatchStart": {
		Method:        "POST",
		ServiceSuffix: "prod/batch/start",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchStart",
	},
	"BatchStop": {
		Method:        "POST",
		ServiceSuffix: "prod/batch/stop",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchStop",
	},
	"ClaimDevice": {
		Method:        "POST",
		ServiceSuffix: "prod/claimDevice",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ClaimDevice",
	},
	"CreateChannel": {
		Method:        "POST",
		ServiceSuffix: "prod/channels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateChannel",
	},
	"CreateInput": {
		Method:        "POST",
		ServiceSuffix: "prod/inputs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateInput",
	},
	"CreateInputSecurityGroup": {
		Method:        "POST",
		ServiceSuffix: "prod/inputSecurityGroups",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateInputSecurityGroup",
	},
	"ListChannels": {
		Method:        "GET",
		ServiceSuffix: "prod/channels",
		Permission:    "ListChannels",
	},
	"ListInputDevices": {
		Method:        "GET",
		ServiceSuffix: "prod/inputDevices",
		Permission:    "ListInputDevices",
	},
	"ListInputSecurityGroups": {
		Method:        "GET",
		ServiceSuffix: "prod/inputSecurityGroups",
		Permission:    "ListInputSecurityGroups",
	},
	"ListInputs": {
		Method:        "GET",
		ServiceSuffix: "prod/inputs",
		Permission:    "ListInputs",
	},
	"ListMultiplexes": {
		Method:        "GET",
		ServiceSuffix: "prod/multiplexes",
		Permission:    "ListMultiplexes",
	},
	"ListOfferings": {
		Method:        "GET",
		ServiceSuffix: "prod/offerings",
		Permission:    "ListOfferings",
	},
	"ListReservations": {
		Method:        "GET",
		ServiceSuffix: "prod/reservations",
		Permission:    "ListReservations",
	},

	// extra
	"DescribeChannel": {
		ServiceSuffix:          "/prod/channels/{{.channel_id}}",
		Permission:             "DescribeChannel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "channel_id",
	},
	"DescribeInput": {
		ServiceSuffix:          "/prod/inputs/{{.input_id}}",
		Permission:             "DescribeInput",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_id",
	},
	"DescribeInputDevice": {
		ServiceSuffix:          "/prod/inputDevices/{{.input_device_id}}",
		Permission:             "DescribeInputDevice",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_device_id",
	},
	"DescribeInputSecurityGroup": {
		ServiceSuffix:          "/prod/inputSecurityGroups/{{.input_security_group_id}}",
		Permission:             "DescribeInputSecurityGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_security_group_id",
	},
	"DescribeMultiplex": {
		ServiceSuffix:          "/prod/multiplexes/{{.multiplex_id}}",
		Permission:             "DescribeMultiplex",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "multiplex_id",
	},
	"DescribeOffering": {
		ServiceSuffix:          "/prod/offerings/{{.offering_id}}",
		Permission:             "DescribeOffering",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "offering_id",
	},
	"DescribeReservation": {
		ServiceSuffix:          "/prod/reservations/{{.reservation_id}}",
		Permission:             "DescribeReservation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "reservation_id",
	},
	"DescribeSchedule": {
		ServiceSuffix:          "/prod/channels/{{.channel_id}}/schedule",
		Permission:             "DescribeSchedule",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "channel_id",
	},
	"ListInputDeviceTransfers": {
		ServiceSuffix:          "/prod/inputDeviceTransfers",
		Permission:             "ListInputDeviceTransfers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "transfer_type",
	},
	"ListMultiplexPrograms": {
		ServiceSuffix:          "/prod/multiplexes/{{.multiplex_id}}/programs",
		Permission:             "ListMultiplexPrograms",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "multiplex_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/prod/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
