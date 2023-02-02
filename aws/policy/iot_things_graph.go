package policy

import "github.com/securisec/cliam/shared"

// IOTThingsGraphPolicies policy
var IOTThingsGraphPolicies = map[string]Service{
	"DeploySystemInstance": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.DeploySystemInstance",
		},
		Permission: "DeploySystemInstance",
	},
	"DescribeNamespace": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.DescribeNamespace",
		},
		Permission: "DescribeNamespace",
	},
	"GetNamespaceDeletionStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetNamespaceDeletionStatus",
		},
		Permission: "GetNamespaceDeletionStatus",
	},
	"SearchFlowTemplates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.SearchFlowTemplates",
		},
		Permission: "SearchFlowTemplates",
	},
	"SearchSystemInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.SearchSystemInstances",
		},
		Permission: "SearchSystemInstances",
	},
	"SearchSystemTemplates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.SearchSystemTemplates",
		},
		Permission: "SearchSystemTemplates",
	},
	"UndeploySystemInstance": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.UndeploySystemInstance",
		},
		Permission: "UndeploySystemInstance",
	},
	"UploadEntityDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.UploadEntityDefinitions",
		},
		Permission: "UploadEntityDefinitions",
	},

	// extra
	"GetEntities": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetEntities",
		},
		Permission:             "GetEntities",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ids",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "ids",
	},
	"GetFlowTemplate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetFlowTemplate",
		},
		Permission:             "GetFlowTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"GetFlowTemplateRevisions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetFlowTemplateRevisions",
		},
		Permission:             "GetFlowTemplateRevisions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"GetSystemInstance": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetSystemInstance",
		},
		Permission:             "GetSystemInstance",
		IsExtra:                true,
		ExtraComponentBodyKey:  "id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"GetSystemTemplate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetSystemTemplate",
		},
		Permission:             "GetSystemTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"GetSystemTemplateRevisions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetSystemTemplateRevisions",
		},
		Permission:             "GetSystemTemplateRevisions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"GetUploadStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.GetUploadStatus",
		},
		Permission:             "GetUploadStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "uploadId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "upload_id",
	},
	"ListFlowExecutionMessages": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.ListFlowExecutionMessages",
		},
		Permission:             "ListFlowExecutionMessages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "flowExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "flow_execution_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IotThingsGraphFrontEndService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
