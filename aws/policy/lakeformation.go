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

	// extra
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeResource",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransaction",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTransaction",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransactionId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transaction_id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetEffectivePermissionsForPath",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetEffectivePermissionsForPath",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetLFTag",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetLFTag",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TagKey",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "tag_key",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetQueryState",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetQueryState",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "query_id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetQueryStatistics",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetQueryStatistics",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "query_id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetResourceLFTags",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetResourceLFTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Resource",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetWorkUnits",
			"Version": "2017-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetWorkUnits",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "query_id",
	},
}
