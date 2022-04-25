package policy

import "github.com/securisec/cliam/shared"

var ShieldPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeAttackStatistics",
		},
		Permission: "DescribeAttackStatistics",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeDrtAccess",
		},
		Permission: "DescribeDrtAccess",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeEmergencyContactSettings",
		},
		Permission: "DescribeEmergencyContactSettings",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeProtection",
		},
		Permission: "DescribeProtection",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeSubscription",
		},
		Permission: "DescribeSubscription",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.GetSubscriptionState",
		},
		Permission: "GetSubscriptionState",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListAttacks",
		},
		Permission: "ListAttacks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListProtectionGroups",
		},
		Permission: "ListProtectionGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListProtections",
		},
		Permission: "ListProtections",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeAttack",
		},
		Permission:             "DescribeAttack",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AttackId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "attack_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeProtectionGroup",
		},
		Permission:             "DescribeProtectionGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProtectionGroupId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "protection_group_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListResourcesInProtectionGroup",
		},
		Permission:             "ListResourcesInProtectionGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProtectionGroupId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "protection_group_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
