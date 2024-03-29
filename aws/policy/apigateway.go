package policy

import "github.com/securisec/cliam/shared"

// APIGatewayPolicies policy
var APIGatewayPolicies = map[string]Service{
	"GenerateClientCertificate": {
		Method:        "POST",
		ServiceSuffix: "clientcertificates",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GenerateClientCertificate",
	},
	"GetAccount": {
		Method:        "GET",
		ServiceSuffix: "account",
		Permission:    "GetAccount",
	},
	"GetApiKeys": {
		Method:        "GET",
		ServiceSuffix: "apikeys",
		Permission:    "GetApiKeys",
	},
	"GetClientCertificates": {
		Method:        "GET",
		ServiceSuffix: "clientcertificates",
		Permission:    "GetClientCertificates",
	},
	"GetDomainNames": {
		Method:        "GET",
		ServiceSuffix: "domainnames",
		Permission:    "GetDomainNames",
	},
	"GetRestApis": {
		Method:        "GET",
		ServiceSuffix: "restapis",
		Permission:    "GetRestApis",
	},
	"GetSdkTypes": {
		Method:        "GET",
		ServiceSuffix: "sdktypes",
		Permission:    "GetSdkTypes",
	},
	"GetUsagePlans": {
		Method:        "GET",
		ServiceSuffix: "usageplans",
		Permission:    "GetUsagePlans",
	},
	"GetVpcLinks": {
		Method:        "GET",
		ServiceSuffix: "vpclinks",
		Permission:    "GetVpcLinks",
	},

	// extra
	"GetApiKey": {
		ServiceSuffix:          "/apikeys/{{.api_key}}",
		Permission:             "GetApiKey",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_key",
	},
	"GetAuthorizers": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/authorizers",
		Permission:             "GetAuthorizers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetBasePathMappings": {
		ServiceSuffix:          "/domainnames/{{.domain_name}}/basepathmappings",
		Permission:             "GetBasePathMappings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetClientCertificate": {
		ServiceSuffix:          "/clientcertificates/{{.client_certificate_id}}",
		Permission:             "GetClientCertificate",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "client_certificate_id",
	},
	"GetDeployments": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/deployments",
		Permission:             "GetDeployments",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetDocumentationParts": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/documentation/parts",
		Permission:             "GetDocumentationParts",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetDocumentationVersions": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/documentation/versions",
		Permission:             "GetDocumentationVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetDomainName": {
		ServiceSuffix:          "/domainnames/{{.domain_name}}",
		Permission:             "GetDomainName",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetGatewayResponses": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/gatewayresponses",
		Permission:             "GetGatewayResponses",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetModels": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/models",
		Permission:             "GetModels",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetRequestValidators": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/requestvalidators",
		Permission:             "GetRequestValidators",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetResources": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/resources",
		Permission:             "GetResources",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetRestApi": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}",
		Permission:             "GetRestApi",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetSdkType": {
		ServiceSuffix:          "/sdktypes/{{.id}}",
		Permission:             "GetSdkType",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetStages": {
		ServiceSuffix:          "/restapis/{{.rest_api_id}}/stages",
		Permission:             "GetStages",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "rest_api_id",
	},
	"GetTags": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "GetTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetUsagePlan": {
		ServiceSuffix:          "/usageplans/{{.usage_plan_id}}",
		Permission:             "GetUsagePlan",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "usage_plan_id",
	},
	"GetUsagePlanKeys": {
		ServiceSuffix:          "/usageplans/{{.usage_plan_id}}/keys",
		Permission:             "GetUsagePlanKeys",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "usage_plan_id",
	},
	"GetVpcLink": {
		ServiceSuffix:          "/vpclinks/{{.vpc_link_id}}",
		Permission:             "GetVpcLink",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "vpc_link_id",
	},

	// apigatewayv2
	"GetApis": {
		Method:        "GET",
		ServiceSuffix: "v2/apis",
		Permission:    "GetApis",
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
	"GetIntegrations": {
		ServiceSuffix:          "/v2/apis/{{.api_id}}/integrations",
		Permission:             "GetIntegrations",
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
}
