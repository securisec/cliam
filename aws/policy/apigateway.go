package policy

import "github.com/securisec/cliam/shared"

var APIGatewayPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "apikeys",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateApiKey",
	},
	{
		Method:        "POST",
		ServiceSuffix: "clientcertificates",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GenerateClientCertificate",
	},
	{
		Method:        "GET",
		ServiceSuffix: "account",
		Permission:    "GetAccount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "apikeys",
		Permission:    "GetApiKeys",
	},
	{
		Method:        "GET",
		ServiceSuffix: "clientcertificates",
		Permission:    "GetClientCertificates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "domainnames",
		Permission:    "GetDomainNames",
	},
	{
		Method:        "GET",
		ServiceSuffix: "restapis",
		Permission:    "GetRestApis",
	},
	{
		Method:        "GET",
		ServiceSuffix: "sdktypes",
		Permission:    "GetSdkTypes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "usageplans",
		Permission:    "GetUsagePlans",
	},
	{
		Method:        "GET",
		ServiceSuffix: "vpclinks",
		Permission:    "GetVpcLinks",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v2/apis",
		Permission:    "GetApis",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v2/domainnames",
		Permission:    "GetDomainNames",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v2/vpclinks",
		Permission:    "GetVpcLinks",
	},
}
