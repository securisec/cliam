package policy

import "github.com/securisec/cliam/shared"

var BackupGatewayPolicies = map[string]Service{
	"ListGateways": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListGateways",
		},
		Permission: "ListGateways",
	},
	"ListHypervisors": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListHypervisors",
		},
		Permission: "ListHypervisors",
	},
	"ListVirtualMachines": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListVirtualMachines",
		},
		Permission: "ListVirtualMachines",
	},

	// extra
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
