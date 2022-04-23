package policy

import "github.com/securisec/cliam/shared"

var BackupGatewayPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListGateways",
		},
		Permission: "ListGateways",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListHypervisors",
		},
		Permission: "ListHypervisors",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "BackupOnPremises_v20210101.ListVirtualMachines",
		},
		Permission: "ListVirtualMachines",
	},
}
