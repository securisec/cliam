package policy

import "github.com/securisec/cliam/shared"

// SecurityLakePolicies policy
var SecurityLakePolicies = map[string]Service{
	"GetDatalake": {
		Method:        "GET",
		ServiceSuffix: "v1/datalake",
		Permission:    "GetDatalake",
	},
	"GetDatalakeAutoEnable": {
		Method:        "GET",
		ServiceSuffix: "v1/datalake/autoenable",
		Permission:    "GetDatalakeAutoEnable",
	},
	"GetDatalakeExceptionsExpiry": {
		Method:        "GET",
		ServiceSuffix: "v1/datalake/exceptions/expiry",
		Permission:    "GetDatalakeExceptionsExpiry",
	},
	"GetDatalakeExceptionsSubscription": {
		Method:        "GET",
		ServiceSuffix: "v1/datalake/exceptions/subscription",
		Permission:    "GetDatalakeExceptionsSubscription",
	},
	"GetDatalakeStatus": {
		Method:        "POST",
		ServiceSuffix: "v1/datalake/status",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetDatalakeStatus",
	},
	"ListDatalakeExceptions": {
		Method:        "POST",
		ServiceSuffix: "v1/datalake/exceptions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDatalakeExceptions",
	},
	"ListLogSources": {
		Method:        "POST",
		ServiceSuffix: "v1/logsources/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListLogSources",
	},
	"ListSubscribers": {
		Method:        "GET",
		ServiceSuffix: "v1/subscribers",
		Permission:    "ListSubscribers",
	},

	// extra
	"GetSubscriber": {
		ServiceSuffix:          "/v1/subscribers/{{.id}}",
		Permission:             "GetSubscriber",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
}
