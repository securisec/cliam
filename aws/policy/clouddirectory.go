package policy

import "github.com/securisec/cliam/shared"

var CloudDirectoryPolicies = map[string]Service{
	"ListDevelopmentSchemaArns": {
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/schema/development",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDevelopmentSchemaArns",
	},
	"ListDirectories": {
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/directory/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDirectories",
	},
	"ListManagedSchemaArns": {
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/schema/managed",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListManagedSchemaArns",
	},
	"ListPublishedSchemaArns": {
		Method:        "POST",
		ServiceSuffix: "amazonclouddirectory/2017-01-11/schema/published",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListPublishedSchemaArns",
	},

	// extra
	"GetAppliedSchemaVersion": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAppliedSchemaVersion",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetAppliedSchemaVersion",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SchemaArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "schema_arn",
	},
	"GetDirectory": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetDirectory",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetDirectory",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "directory_arn",
	},
	"GetSchemaAsJson": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSchemaAsJson",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetSchemaAsJson",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SchemaArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "schema_arn",
	},
	"ListAppliedSchemaArns": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListAppliedSchemaArns",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListAppliedSchemaArns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "directory_arn",
	},
	"ListFacetNames": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListFacetNames",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListFacetNames",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SchemaArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "schema_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTypedLinkFacetNames": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTypedLinkFacetNames",
			"Version": "2017-01-11",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTypedLinkFacetNames",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SchemaArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "schema_arn",
	},
}
