package policy

import "github.com/securisec/cliam/shared"

var AppFlowPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "describe-connector-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeConnectorProfiles",
	},
	{
		Method:        "POST",
		ServiceSuffix: "describe-connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeConnectors",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-connector-entities",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListConnectorEntities",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListConnectors",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-flows",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFlows",
	},
	{
		Method:        "POST",
		ServiceSuffix: "register-connector",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "RegisterConnector",
	},

	// extra
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConnector",
			"Version": "2020-08-23",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeConnector",
		IsExtra:                true,
		ExtraComponentBodyKey:  "connectorType",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "connector_type",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConnectorEntity",
			"Version": "2020-08-23",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeConnectorEntity",
		IsExtra:                true,
		ExtraComponentBodyKey:  "connectorEntityName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "connector_entity_name",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFlow",
			"Version": "2020-08-23",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeFlow",
		IsExtra:                true,
		ExtraComponentBodyKey:  "flowName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "flow_name",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFlowExecutionRecords",
			"Version": "2020-08-23",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeFlowExecutionRecords",
		IsExtra:                true,
		ExtraComponentBodyKey:  "flowName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "flow_name",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
