package policy

import "github.com/securisec/cliam/shared"

var DynamodbStreamsPolicies = []Service{
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			aws_X_AMZ_TARGET:           "DynamoDBStreams_20120810.ListStreams",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "ListStreams",
	},

	// extra
	{
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
	{
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
