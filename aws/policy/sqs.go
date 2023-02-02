package policy

import "github.com/securisec/cliam/shared"

// SQSPolicies policy
var SQSPolicies = map[string]Service{
	"ListQueues": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListQueues",
			"Version": "2012-11-05",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListQueues",
	},

	// extra
	"GetQueueAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetQueueAttributes",
			"Version": "2012-11-05",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetQueueAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueueUrl",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "queue_url",
	},
	"GetQueueUrl": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetQueueUrl",
			"Version": "2012-11-05",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetQueueUrl",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueueName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "queue_name",
	},
	"ListDeadLetterSourceQueues": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListDeadLetterSourceQueues",
			"Version": "2012-11-05",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListDeadLetterSourceQueues",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueueUrl",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "queue_url",
	},
	"ListQueueTags": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListQueueTags",
			"Version": "2012-11-05",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListQueueTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueueUrl",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "queue_url",
	},
}
