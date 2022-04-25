package policy

import "github.com/securisec/cliam/shared"

var ConfigServicePolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeAggregationAuthorizations",
		},
		Permission: "DescribeAggregationAuthorizations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeComplianceByConfigRule",
		},
		Permission: "DescribeComplianceByConfigRule",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeComplianceByResource",
		},
		Permission: "DescribeComplianceByResource",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigRuleEvaluationStatus",
		},
		Permission: "DescribeConfigRuleEvaluationStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigRules",
		},
		Permission: "DescribeConfigRules",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationAggregators",
		},
		Permission: "DescribeConfigurationAggregators",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationRecorderStatus",
		},
		Permission: "DescribeConfigurationRecorderStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationRecorders",
		},
		Permission: "DescribeConfigurationRecorders",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConformancePackStatus",
		},
		Permission: "DescribeConformancePackStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConformancePacks",
		},
		Permission: "DescribeConformancePacks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeDeliveryChannelStatus",
		},
		Permission: "DescribeDeliveryChannelStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeDeliveryChannels",
		},
		Permission: "DescribeDeliveryChannels",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConfigRuleStatuses",
		},
		Permission: "DescribeOrganizationConfigRuleStatuses",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConfigRules",
		},
		Permission: "DescribeOrganizationConfigRules",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConformancePackStatuses",
		},
		Permission: "DescribeOrganizationConformancePackStatuses",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConformancePacks",
		},
		Permission: "DescribeOrganizationConformancePacks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribePendingAggregationRequests",
		},
		Permission: "DescribePendingAggregationRequests",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeRetentionConfigurations",
		},
		Permission: "DescribeRetentionConfigurations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetComplianceSummaryByConfigRule",
		},
		Permission: "GetComplianceSummaryByConfigRule",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetComplianceSummaryByResourceType",
		},
		Permission: "GetComplianceSummaryByResourceType",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetCustomRulePolicy",
		},
		Permission: "GetCustomRulePolicy",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetDiscoveredResourceCounts",
		},
		Permission: "GetDiscoveredResourceCounts",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.ListStoredQueries",
		},
		Permission: "ListStoredQueries",
	},
}
