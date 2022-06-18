package policy

import "github.com/securisec/cliam/shared"

var TransferPolicies = map[string]Service{
	"CreateServer": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.CreateServer",
		},
		Permission: "CreateServer",
	},
	"ListSecurityPolicies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListSecurityPolicies",
		},
		Permission: "ListSecurityPolicies",
	},
	"ListServers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListServers",
		},
		Permission: "ListServers",
	},
	"ListWorkflows": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListWorkflows",
		},
		Permission: "ListWorkflows",
	},

	// extra
	"DescribeSecurityPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.DescribeSecurityPolicy",
		},
		Permission:             "DescribeSecurityPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SecurityPolicyName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "security_policy_name",
	},
	"DescribeServer": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.DescribeServer",
		},
		Permission:             "DescribeServer",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServerId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "server_id",
	},
	"DescribeWorkflow": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.DescribeWorkflow",
		},
		Permission:             "DescribeWorkflow",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkflowId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workflow_id",
	},
	"ListAccesses": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListAccesses",
		},
		Permission:             "ListAccesses",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServerId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "server_id",
	},
	"ListExecutions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListExecutions",
		},
		Permission:             "ListExecutions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkflowId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workflow_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListUsers": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListUsers",
		},
		Permission:             "ListUsers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServerId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "server_id",
	},
}
