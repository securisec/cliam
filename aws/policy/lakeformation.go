package policy

import "github.com/securisec/cliam/shared"

var LakeFormationPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "DeleteDataCellsFilter",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DeleteDataCellsFilter",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ExtendTransaction",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ExtendTransaction",
	},
	{
		Method:        "POST",
		ServiceSuffix: "GetDataLakeSettings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetDataLakeSettings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListDataCellsFilter",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDataCellsFilter",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListLFTags",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListLFTags",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListPermissions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListPermissions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListResources",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListResources",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListTransactions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListTransactions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "StartTransaction",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "StartTransaction",
	},
}
