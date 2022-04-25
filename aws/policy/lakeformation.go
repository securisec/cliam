package policy

import "github.com/securisec/cliam/shared"

var LakeFormationPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "DeleteDataCellsFilter",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DeleteDataCellsFilter",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ExtendTransaction",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ExtendTransaction",
	},
	{
		Method:        "POST",
		ServiceSuffix: "GetDataLakeSettings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetDataLakeSettings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListDataCellsFilter",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDataCellsFilter",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListLFTags",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListLFTags",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListPermissions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListPermissions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListResources",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListResources",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListTransactions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListTransactions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "StartTransaction",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "StartTransaction",
	},
}
