package policy

var APIGatewayPolicies = []Service{
	{
		ServiceSuffix: "account",
		Permission:    "GetAccount",
	},
	{
		ServiceSuffix: "apikeys",
		Permission:    "GetApiKeys",
	},
	{
		ServiceSuffix: "clientcertificates",
		Permission:    "GetClientCertificates",
	},
	{
		ServiceSuffix: "domainnames",
		Permission:    "GetDomainNames",
	},
	{
		ServiceSuffix: "restapis",
		Permission:    "GetRestApis",
	},
	{
		ServiceSuffix: "sdktypes",
		Permission:    "GetSdkTypes",
	},
	{
		ServiceSuffix: "usageplans",
		Permission:    "GetUsagePlans",
	},
	{
		ServiceSuffix: "vpclinks",
		Permission:    "GetVpcLinks",
	},
	{
		ServiceSuffix: "vpclinks",
		Permission:    "GetVpcLinks",
	},
}
