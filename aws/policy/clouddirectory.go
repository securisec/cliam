package policy

import "github.com/securisec/cliam/shared"

var CloudDirectoryPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/schema/development",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDevelopmentSchemaArns",
	},
	{
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/directory/list",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDirectories",
	},
	{
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/schema/managed",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListManagedSchemaArns",
	},
	{
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/schema/published",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListPublishedSchemaArns",
	},
}
