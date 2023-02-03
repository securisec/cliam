package policy

// APIGatewayV2Policies policy
var APIGatewayV2Policies = map[string]Service{
	"GetApis": {
		Method:        "GET",
		ServiceSuffix: "v2/apis",
		Permission:    "GetApis",
	},
	"GetDomainNames": {
		Method:        "GET",
		ServiceSuffix: "v2/domainnames",
		Permission:    "GetDomainNames",
	},
	"GetVpcLinks": {
		Method:        "GET",
		ServiceSuffix: "v2/vpclinks",
		Permission:    "GetVpcLinks",
	},

	// extra
	"GetApi": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}",
		Permission:             "GetApi",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetApiMappings": {
		ServiceSuffix:          "/v2/domainnames/{{.domain_name}}/apimappings",
		Permission:             "GetApiMappings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetAuthorizers": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/authorizers",
		Permission:             "GetAuthorizers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetDeployments": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/deployments",
		Permission:             "GetDeployments",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetDomainName": {
		ServiceSuffix:          "/v2/domainnames/{{.domain_name}}",
		Permission:             "GetDomainName",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetIntegrations": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/integrations",
		Permission:             "GetIntegrations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetModels": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/models",
		Permission:             "GetModels",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetRoutes": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/routes",
		Permission:             "GetRoutes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetStages": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/stages",
		Permission:             "GetStages",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetTags": {
		ServiceSuffix:          "/v2/tags/{{.resource_arn}}",
		Permission:             "GetTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetVpcLink": {
		ServiceSuffix:          "/v2/vpclinks/{{.vpc_link_id}}",
		Permission:             "GetVpcLink",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "vpc_link_id",
	},
}
