package policy

import "github.com/securisec/cliam/shared"

var KinesisPolicies = map[string]Service{
	"DescribeLimits": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeLimits",
		},
		Permission: "DescribeLimits",
	},
	"DescribeStreamConsumer": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeStreamConsumer",
		},
		Permission: "DescribeStreamConsumer",
	},
	"ListStreams": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.ListStreams",
		},
		Permission: "ListStreams",
	},

	// extra
	"DescribeStream": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeStream",
		},
		Permission:             "DescribeStream",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StreamName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stream_name",
	},
	"DescribeStreamSummary": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeStreamSummary",
		},
		Permission:             "DescribeStreamSummary",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StreamName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stream_name",
	},
	"GetRecords": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.GetRecords",
		},
		Permission:             "GetRecords",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ShardIterator",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "shard_iterator",
	},
	"ListStreamConsumers": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.ListStreamConsumers",
		},
		Permission:             "ListStreamConsumers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StreamARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stream_arn",
	},
	"ListTagsForStream": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.ListTagsForStream",
		},
		Permission:             "ListTagsForStream",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StreamName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stream_name",
	},
}
