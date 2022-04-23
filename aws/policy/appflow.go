package policy

import "github.com/securisec/cliam/shared"

var AppFlowPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "describe-connector-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeConnectorProfiles",
	},
	{
		Method:        "POST",
		ServiceSuffix: "describe-connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeConnectors",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-connector-entities",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListConnectorEntities",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListConnectors",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-flows",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListFlows",
	},
	{
		Method:        "POST",
		ServiceSuffix: "register-connector",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "RegisterConnector",
	},
}
