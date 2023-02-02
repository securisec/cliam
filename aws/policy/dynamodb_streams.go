package policy

import "github.com/securisec/cliam/shared"

// DynamodbStreamsPolicies policy
var DynamodbStreamsPolicies = map[string]Service{
	"ListStreams": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DynamoDBStreams_20120810.ListStreams",
		},
		Permission: "ListStreams",
	},

	// extra
	"DescribeStream": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDBStreams_20120810.DescribeStream",
		},
		Permission:             "DescribeStream",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StreamArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stream_arn",
	},
	"GetRecords": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDBStreams_20120810.GetRecords",
		},
		Permission:             "GetRecords",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ShardIterator",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "shard_iterator",
	},
}
