package policy

import "github.com/securisec/cliam/shared"

// RedshiftServerlessPolicies policy
var RedshiftServerlessPolicies = map[string]Service{
	"GetSnapshot": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetSnapshot",
		},
		Permission: "GetSnapshot",
	},
	"ListEndpointAccess": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListEndpointAccess",
		},
		Permission: "ListEndpointAccess",
	},
	"ListNamespaces": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListNamespaces",
		},
		Permission: "ListNamespaces",
	},
	"ListRecoveryPoints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListRecoveryPoints",
		},
		Permission: "ListRecoveryPoints",
	},
	"ListSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListSnapshots",
		},
		Permission: "ListSnapshots",
	},
	"ListTableRestoreStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListTableRestoreStatus",
		},
		Permission: "ListTableRestoreStatus",
	},
	"ListUsageLimits": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListUsageLimits",
		},
		Permission: "ListUsageLimits",
	},
	"ListWorkgroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListWorkgroups",
		},
		Permission: "ListWorkgroups",
	},

	// extra
	"GetCredentials": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetCredentials",
		},
		Permission:             "GetCredentials",
		IsExtra:                true,
		ExtraComponentBodyKey:  "workgroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workgroup_name",
	},
	"GetEndpointAccess": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetEndpointAccess",
		},
		Permission:             "GetEndpointAccess",
		IsExtra:                true,
		ExtraComponentBodyKey:  "endpointName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_name",
	},
	"GetNamespace": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetNamespace",
		},
		Permission:             "GetNamespace",
		IsExtra:                true,
		ExtraComponentBodyKey:  "namespaceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "namespace_name",
	},
	"GetRecoveryPoint": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetRecoveryPoint",
		},
		Permission:             "GetRecoveryPoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "recoveryPointId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "recovery_point_id",
	},
	"GetResourcePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetResourcePolicy",
		},
		Permission:             "GetResourcePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetTableRestoreStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetTableRestoreStatus",
		},
		Permission:             "GetTableRestoreStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "tableRestoreRequestId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_restore_request_id",
	},
	"GetUsageLimit": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetUsageLimit",
		},
		Permission:             "GetUsageLimit",
		IsExtra:                true,
		ExtraComponentBodyKey:  "usageLimitId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "usage_limit_id",
	},
	"GetWorkgroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.GetWorkgroup",
		},
		Permission:             "GetWorkgroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "workgroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workgroup_name",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftServerless.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
