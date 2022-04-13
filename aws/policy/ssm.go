package policy

import "github.com/securisec/cliam/shared"

var SSMPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeActivations",
		},
		Permission: "DescribeActivations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeAvailablePatches",
		},
		Permission: "DescribeAvailablePatches",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeInventoryDeletions",
		},
		Permission: "DescribeInventoryDeletions",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribeMaintenanceWindows",
		},
		Permission: "DescribeMaintenanceWindows",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribePatchBaselines",
		},
		Permission: "DescribePatchBaselines",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.DescribePatchGroups",
		},
		Permission: "DescribePatchGroups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetDefaultPatchBaseline",
		},
		Permission: "GetDefaultPatchBaseline",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.GetInventorySchema",
		},
		Permission: "GetInventorySchema",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListCommandInvocations",
		},
		Permission: "ListCommandInvocations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListCommands",
		},
		Permission: "ListCommands",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListComplianceItems",
		},
		Permission: "ListComplianceItems",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListComplianceSummaries",
		},
		Permission: "ListComplianceSummaries",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListResourceComplianceSummaries",
		},
		Permission: "ListResourceComplianceSummaries",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonSSM.ListResourceDataSync",
		},
		Permission: "ListResourceDataSync",
	},
}
