package policy

import "github.com/securisec/enumerate/shared"

var DynamodbStreams = []Service{
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			AWS_X_AMZ_TARGET:           "DynamoDBStreams_20120810.ListStreams",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "ListStreams",
	},
}
