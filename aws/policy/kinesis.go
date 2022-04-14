package policy

import "github.com/securisec/cliam/shared"

var KinesisPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeLimits",
		},
		Permission: "DescribeLimits",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.DescribeStreamConsumer",
		},
		Permission: "DescribeStreamConsumer",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Kinesis_20131202.ListStreams",
		},
		Permission: "ListStreams",
	},
}
