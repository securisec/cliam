package policy

import "github.com/securisec/cliam/shared"

var SQSPolicies = []Service{
	{
		ServiceSuffix: "?Action=ListQueues",
		Permission:    "ListQueues",
	},

	// extras
	{
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
	{
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
	{
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
	{
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
