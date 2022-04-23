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
}
