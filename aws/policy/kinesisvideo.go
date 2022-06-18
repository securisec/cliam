package policy

import "github.com/securisec/cliam/shared"

var KinesisVideoPolicies = map[string]Service{
	"DescribeSignalingChannel": {
		Method:        "POST",
		ServiceSuffix: "describeSignalingChannel",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeSignalingChannel",
	},
	"DescribeStream": {
		Method:        "POST",
		ServiceSuffix: "describeStream",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeStream",
	},
	"ListSignalingChannels": {
		Method:        "POST",
		ServiceSuffix: "listSignalingChannels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSignalingChannels",
	},
	"ListStreams": {
		Method:        "POST",
		ServiceSuffix: "listStreams",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListStreams",
	},
	"ListTagsForStream": {
		Method:        "POST",
		ServiceSuffix: "listTagsForStream",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListTagsForStream",
	},

	// extra
	"GetDataEndpoint": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetDataEndpoint",
			"Version": "2017-09-30",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetDataEndpoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "APIName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "a_p_i_name",
	},
	"GetSignalingChannelEndpoint": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSignalingChannelEndpoint",
			"Version": "2017-09-30",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetSignalingChannelEndpoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ChannelARN",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "channel_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2017-09-30",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
