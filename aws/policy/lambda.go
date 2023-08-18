package policy

import (
	"github.com/buger/jsonparser"
	"github.com/securisec/cliam/shared"
)

// LambdaPolicies policy
var LambdaPolicies = map[string]Service{
	"GetAccountSettings": {
		Method:        "GET",
		ServiceSuffix: "2016-08-19/account-settings/",
		Permission:    "GetAccountSettings",
	},
	"ListCodeSigningConfigs": {
		Method:        "GET",
		ServiceSuffix: "2020-04-22/code-signing-configs/",
		Permission:    "ListCodeSigningConfigs",
	},
	"ListEventSourceMappings": {
		Method:        "GET",
		ServiceSuffix: "2015-03-31/event-source-mappings/",
		Permission:    "ListEventSourceMappings",
	},
	"ListFunctions": {
		Method:        "GET",
		ServiceSuffix: "2015-03-31/functions/",
		Permission:    "ListFunctions",
		ResponseParser: &ResponseParser{
			ExtractedExtraCommandLineFlag: "function_name",
			ResponseFormat:                "json",
			ResponseParser: func(respBytes []byte) ([]string, error) {
				converted, err := shared.ResponseToJSON("json", respBytes)
				if err != nil {
					return nil, err
				}
				hold := make([]string, 0)
				_, err = jsonparser.ArrayEach(converted, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					bucketName, _ := jsonparser.GetString(value, "FunctionName")
					hold = append(hold, bucketName)
				}, "Functions")
				return hold, err
			},
		},
	},
	"ListLayers": {
		Method:        "GET",
		ServiceSuffix: "2018-10-31/layers",
		Permission:    "ListLayers",
	},

	// extra
	"GetCodeSigningConfig": {
		ServiceSuffix:          "/2020-04-22/code-signing-configs/{{.code_signing_config_arn}}",
		Permission:             "GetCodeSigningConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "code_signing_config_arn",
	},
	"GetEventSourceMapping": {
		ServiceSuffix:          "/2015-03-31/event-source-mappings/{{.u_u_i_d}}",
		Permission:             "GetEventSourceMapping",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "u_u_i_d",
	},
	"GetFunction": {
		ServiceSuffix:          "/2015-03-31/functions/{{.function_name}}",
		Permission:             "GetFunction",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetFunctionCodeSigningConfig": {
		ServiceSuffix:          "/2020-06-30/functions/{{.function_name}}/code-signing-config",
		Permission:             "GetFunctionCodeSigningConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetFunctionConcurrency": {
		ServiceSuffix:          "/2019-09-30/functions/{{.function_name}}/concurrency",
		Permission:             "GetFunctionConcurrency",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetFunctionConfiguration": {
		ServiceSuffix:          "/2015-03-31/functions/{{.function_name}}/configuration",
		Permission:             "GetFunctionConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetFunctionEventInvokeConfig": {
		ServiceSuffix:          "/2019-09-25/functions/{{.function_name}}/event-invoke-config",
		Permission:             "GetFunctionEventInvokeConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetFunctionUrlConfig": {
		ServiceSuffix:          "/2021-10-31/functions/{{.function_name}}/url",
		Permission:             "GetFunctionUrlConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetLayerVersionByArn": {
		ServiceSuffix:          "/2018-10-31/layers?find=LayerVersion",
		Permission:             "GetLayerVersionByArn",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	"GetPolicy": {
		ServiceSuffix:          "/2015-03-31/functions/{{.function_name}}/policy",
		Permission:             "GetPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetRuntimeManagementConfig": {
		ServiceSuffix:          "/2021-07-20/functions/{{.function_name}}/runtime-management-config",
		Permission:             "GetRuntimeManagementConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"ListAliases": {
		ServiceSuffix:          "/2015-03-31/functions/{{.function_name}}/aliases",
		Permission:             "ListAliases",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"ListFunctionEventInvokeConfigs": {
		ServiceSuffix:          "/2019-09-25/functions/{{.function_name}}/event-invoke-config/list",
		Permission:             "ListFunctionEventInvokeConfigs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"ListFunctionUrlConfigs": {
		ServiceSuffix:          "/2021-10-31/functions/{{.function_name}}/urls",
		Permission:             "ListFunctionUrlConfigs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"ListFunctionsByCodeSigningConfig": {
		ServiceSuffix:          "/2020-04-22/code-signing-configs/{{.code_signing_config_arn}}/functions",
		Permission:             "ListFunctionsByCodeSigningConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "code_signing_config_arn",
	},
	"ListLayerVersions": {
		ServiceSuffix:          "/2018-10-31/layers/{{.layer_name}}/versions",
		Permission:             "ListLayerVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "layer_name",
	},
	"ListProvisionedConcurrencyConfigs": {
		ServiceSuffix:          "/2019-09-30/functions/{{.function_name}}/provisioned-concurrency?List=ALL",
		Permission:             "ListProvisionedConcurrencyConfigs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"ListTags": {
		ServiceSuffix:          "/2017-03-31/tags/{{.resource}}",
		Permission:             "ListTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource",
	},
	"ListVersionsByFunction": {
		ServiceSuffix:          "/2015-03-31/functions/{{.function_name}}/versions",
		Permission:             "ListVersionsByFunction",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},

	"ListEventSources": {
		Method:        "GET",
		ServiceSuffix: "2014-11-13/event-source-mappings/",
		Permission:    "ListEventSources",
	},

	// extra
	"GetEventSource": {
		ServiceSuffix:          "/2014-11-13/event-source-mappings/{{.u_u_i_d}}",
		Permission:             "GetEventSource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "u_u_i_d",
	},
}
