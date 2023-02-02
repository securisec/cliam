package policy

import "github.com/securisec/cliam/shared"

// KinesisPolicies policy
var KinesisPolicies = map[string]Service{
	"DeregisterStreamConsumer": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DeregisterStreamConsumer",
		},
		Permission: "DeregisterStreamConsumer",
	},
	"DescribeLimits": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeLimits",
		},
		Permission: "DescribeLimits",
	},
	"DescribeStream": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeStream",
		},
		Permission: "DescribeStream",
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
	"DescribeStreamSummary": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeStreamSummary",
		},
		Permission: "DescribeStreamSummary",
	},
	"ListShards": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.ListShards",
		},
		Permission: "ListShards",
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
	"ListTagsForStream": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.ListTagsForStream",
		},
		Permission: "ListTagsForStream",
	},

	// extra
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
}
