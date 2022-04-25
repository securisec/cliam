package policy

import "github.com/securisec/cliam/shared"

var KinesisVideoPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "describeSignalingChannel",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeSignalingChannel",
	},
	{
		Method:        "POST",
		ServiceSuffix: "describeStream",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeStream",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSignalingChannels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSignalingChannels",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listStreams",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListStreams",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listTagsForStream",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListTagsForStream",
	},
}
