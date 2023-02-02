package policy

import "github.com/securisec/cliam/shared"

// BackupGatewayPolicies policy
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
	"GetBandwidthRateLimitSchedule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.GetBandwidthRateLimitSchedule",
		},
		Permission:             "GetBandwidthRateLimitSchedule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"GetGateway": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.GetGateway",
		},
		Permission:             "GetGateway",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"GetHypervisor": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.GetHypervisor",
		},
		Permission:             "GetHypervisor",
		IsExtra:                true,
		ExtraComponentBodyKey:  "HypervisorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "hypervisor_arn",
	},
	"GetHypervisorPropertyMappings": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.GetHypervisorPropertyMappings",
		},
		Permission:             "GetHypervisorPropertyMappings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "HypervisorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "hypervisor_arn",
	},
	"GetVirtualMachine": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.GetVirtualMachine",
		},
		Permission:             "GetVirtualMachine",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
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
