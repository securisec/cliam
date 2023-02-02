package policy

// AppSyncPolicies policy
var AppSyncPolicies = map[string]Service{
	"ListDomainNames": {
		Method:        "GET",
		ServiceSuffix: "v1/domainnames",
		Permission:    "ListDomainNames",
	},
	"ListGraphqlApis": {
		Method:        "GET",
		ServiceSuffix: "v1/apis",
		Permission:    "ListGraphqlApis",
	},

	// extra
	"GetApiAssociation": {
		ServiceSuffix:          "/v1/domainnames/{{.domain_name}}/apiassociation",
		Permission:             "GetApiAssociation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetApiCache": {
		ServiceSuffix:          "/v1/apis/{{.api_id}}/ApiCaches",
		Permission:             "GetApiCache",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetDomainName": {
		ServiceSuffix:          "/v1/domainnames/{{.domain_name}}",
		Permission:             "GetDomainName",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetGraphqlApi": {
		ServiceSuffix:          "/v1/apis/{{.api_id}}",
		Permission:             "GetGraphqlApi",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"GetSchemaCreationStatus": {
		ServiceSuffix:          "/v1/apis/{{.api_id}}/schemacreation",
		Permission:             "GetSchemaCreationStatus",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"ListApiKeys": {
		ServiceSuffix:          "/v1/apis/{{.api_id}}/apikeys",
		Permission:             "ListApiKeys",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"ListDataSources": {
		ServiceSuffix:          "/v1/apis/{{.api_id}}/datasources",
		Permission:             "ListDataSources",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"ListFunctions": {
		ServiceSuffix:          "/v1/apis/{{.api_id}}/functions",
		Permission:             "ListFunctions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "api_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/v1/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
