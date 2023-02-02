package policy

import "github.com/securisec/cliam/shared"

// ShieldPolicies policy
var ShieldPolicies = map[string]Service{
	"DescribeAttackStatistics": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeAttackStatistics",
		},
		Permission: "DescribeAttackStatistics",
	},
	"DescribeDRTAccess": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeDRTAccess",
		},
		Permission: "DescribeDRTAccess",
	},
	"DescribeEmergencyContactSettings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeEmergencyContactSettings",
		},
		Permission: "DescribeEmergencyContactSettings",
	},
	"DescribeProtection": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeProtection",
		},
		Permission: "DescribeProtection",
	},
	"DescribeSubscription": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.DescribeSubscription",
		},
		Permission: "DescribeSubscription",
	},
	// "DisableProactiveEngagement": {
	// 	Method:   "POST",
	// 	JsonData: map[string]string{},
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
	// 		aws_X_AMZ_TARGET:           "AWSShield_20160616.DisableProactiveEngagement",
	// 	},
	// 	Permission: "DisableProactiveEngagement",
	// },
	// "DisassociateDRTRole": {
	// 	Method:   "POST",
	// 	JsonData: map[string]string{},
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
	// 		aws_X_AMZ_TARGET:           "AWSShield_20160616.DisassociateDRTRole",
	// 	},
	// 	Permission: "DisassociateDRTRole",
	// },
	// "EnableProactiveEngagement": {
	// 	Method:   "POST",
	// 	JsonData: map[string]string{},
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
	// 		aws_X_AMZ_TARGET:           "AWSShield_20160616.EnableProactiveEngagement",
	// 	},
	// 	Permission: "EnableProactiveEngagement",
	// },
	"GetSubscriptionState": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.GetSubscriptionState",
		},
		Permission: "GetSubscriptionState",
	},
	"ListAttacks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListAttacks",
		},
		Permission: "ListAttacks",
	},
	"ListProtectionGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListProtectionGroups",
		},
		Permission: "ListProtectionGroups",
	},
	"ListProtections": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShield_20160616.ListProtections",
		},
		Permission: "ListProtections",
	},
	// "UpdateEmergencyContactSettings": {
	// 	Method:   "POST",
	// 	JsonData: map[string]string{},
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
	// 		aws_X_AMZ_TARGET:           "AWSShield_20160616.UpdateEmergencyContactSettings",
	// 	},
	// 	Permission: "UpdateEmergencyContactSettings",
	// },
	// "UpdateSubscription": {
	// 	Method:   "POST",
	// 	JsonData: map[string]string{},
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
	// 		aws_X_AMZ_TARGET:           "AWSShield_20160616.UpdateSubscription",
	// 	},
	// 	Permission: "UpdateSubscription",
	// },

	// extra
	"DescribeAttack": {
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
	"DescribeProtectionGroup": {
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
	"ListResourcesInProtectionGroup": {
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
	"ListTagsForResource": {
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
