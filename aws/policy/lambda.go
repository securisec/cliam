package policy

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
	},
	"ListLayers": {
		Method:        "GET",
		ServiceSuffix: "2018-10-31/layers",
		Permission:    "ListLayers",
	},

	// extra policies
	"GetFunction": {
		Method:                 "GET",
		ServiceSuffix:          "2015-03-31/functions/{{.function_name}}",
		Permission:             "GetFunction",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
	"GetLayerVersionByArn": {
		Method:                 "GET",
		ServiceSuffix:          "2018-10-31/layers/?find=LayerVersion&Arn={{.arn}}",
		Permission:             "GetLayerVersionByArn",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	"GetFunctionConfiguration": {
		Method:                 "GET",
		ServiceSuffix:          "2015-03-31/functions/{{.function_name}}/configuration",
		Permission:             "GetFunctionConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "function_name",
	},
}
