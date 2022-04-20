package policy

import "github.com/securisec/cliam/shared"

var KinesisVideoPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "describeSignalingChannel",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeSignalingChannel",
	},
	{
		Method:        "POST",
		ServiceSuffix: "describeStream",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeStream",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSignalingChannels",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListSignalingChannels",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listStreams",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListStreams",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listTagsForStream",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListTagsForStream",
	},
}