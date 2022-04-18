package policy

import "github.com/securisec/cliam/shared"

var MediaStorePolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.DescribeContainer",
		},
		Permission: "DescribeContainers",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.ListContainers",
		},
		Permission: "ListContainers",
	},
}
