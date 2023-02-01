package policy

import "github.com/securisec/cliam/shared"

// AppFlowPolicies policy
var AppFlowPolicies = map[string]Service{
	"DescribeConnectorProfiles": {
		Method:        "POST",
		ServiceSuffix: "describe-connector-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeConnectorProfiles",
	},
	"DescribeConnectors": {
		Method:        "POST",
		ServiceSuffix: "describe-connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeConnectors",
	},
	"ListConnectorEntities": {
		Method:        "POST",
		ServiceSuffix: "list-connector-entities",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListConnectorEntities",
	},
	"ListConnectors": {
		Method:        "POST",
		ServiceSuffix: "list-connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListConnectors",
	},
	"ListFlows": {
		Method:        "POST",
		ServiceSuffix: "list-flows",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFlows",
	},
	"RegisterConnector": {
		Method:        "POST",
		ServiceSuffix: "register-connector",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "RegisterConnector",
	},

	// extra
	"DescribeConnector": {
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
	"DescribeConnectorEntity": {
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
	"DescribeFlow": {
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
	"DescribeFlowExecutionRecords": {
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
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
