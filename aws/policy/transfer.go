package policy

import "github.com/securisec/cliam/shared"

// TransferPolicies policy
var TransferPolicies = map[string]Service{
	"ListCertificates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListCertificates",
		},
		Permission: "ListCertificates",
	},
	"ListConnectors": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListConnectors",
		},
		Permission: "ListConnectors",
	},
	"ListProfiles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListProfiles",
		},
		Permission: "ListProfiles",
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
	"DescribeCertificate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.DescribeCertificate",
		},
		Permission:             "DescribeCertificate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CertificateId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "certificate_id",
	},
	"DescribeConnector": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.DescribeConnector",
		},
		Permission:             "DescribeConnector",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConnectorId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "connector_id",
	},
	"DescribeProfile": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.DescribeProfile",
		},
		Permission:             "DescribeProfile",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProfileId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "profile_id",
	},
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
	"ListAgreements": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListAgreements",
		},
		Permission:             "ListAgreements",
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
	"ListHostKeys": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TransferService.ListHostKeys",
		},
		Permission:             "ListHostKeys",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServerId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "server_id",
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
