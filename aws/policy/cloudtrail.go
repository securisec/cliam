package policy

import "github.com/securisec/cliam/shared"

var CloudtrailPolicies = []Service{
	{
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "DescribeTrails",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.DescribeTrails",
		},
	},
	{
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "ListPublicKeys",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListPublicKeys",
		},
	},
}
