package policy

import "github.com/securisec/cliam/shared"

var LakeFormationPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "DeleteDataCellsFilter",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DeleteDataCellsFilter",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ExtendTransaction",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ExtendTransaction",
	},
	{
		Method:        "POST",
		ServiceSuffix: "GetDataLakeSettings",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetDataLakeSettings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListDataCellsFilter",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDataCellsFilter",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListLFTags",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListLFTags",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListPermissions",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListPermissions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListResources",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListResources",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListTransactions",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListTransactions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "StartTransaction",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "StartTransaction",
	},
}
