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

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeAggregateComplianceByConfigRules",
		},
		Permission:             "DescribeAggregateComplianceByConfigRules",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationAggregatorName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "configuration_aggregator_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeAggregateComplianceByConformancePacks",
		},
		Permission:             "DescribeAggregateComplianceByConformancePacks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationAggregatorName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "configuration_aggregator_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationAggregatorSourcesStatus",
		},
		Permission:             "DescribeConfigurationAggregatorSourcesStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationAggregatorName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "configuration_aggregator_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConformancePackCompliance",
		},
		Permission:             "DescribeConformancePackCompliance",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConformancePackName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "conformance_pack_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeRemediationConfigurations",
		},
		Permission:             "DescribeRemediationConfigurations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigRuleNames",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "config_rule_names",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeRemediationExceptions",
		},
		Permission:             "DescribeRemediationExceptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigRuleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "config_rule_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeRemediationExecutionStatus",
		},
		Permission:             "DescribeRemediationExecutionStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigRuleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "config_rule_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetAggregateConfigRuleComplianceSummary",
		},
		Permission:             "GetAggregateConfigRuleComplianceSummary",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationAggregatorName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "configuration_aggregator_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetAggregateConformancePackComplianceSummary",
		},
		Permission:             "GetAggregateConformancePackComplianceSummary",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationAggregatorName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "configuration_aggregator_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetAggregateDiscoveredResourceCounts",
		},
		Permission:             "GetAggregateDiscoveredResourceCounts",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationAggregatorName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "configuration_aggregator_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetComplianceDetailsByConfigRule",
		},
		Permission:             "GetComplianceDetailsByConfigRule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigRuleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "config_rule_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetConformancePackComplianceDetails",
		},
		Permission:             "GetConformancePackComplianceDetails",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConformancePackName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "conformance_pack_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetConformancePackComplianceSummary",
		},
		Permission:             "GetConformancePackComplianceSummary",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConformancePackNames",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "conformance_pack_names",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetOrganizationConfigRuleDetailedStatus",
		},
		Permission:             "GetOrganizationConfigRuleDetailedStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OrganizationConfigRuleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "organization_config_rule_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetOrganizationConformancePackDetailedStatus",
		},
		Permission:             "GetOrganizationConformancePackDetailedStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OrganizationConformancePackName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "organization_conformance_pack_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetOrganizationCustomRulePolicy",
		},
		Permission:             "GetOrganizationCustomRulePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OrganizationConfigRuleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "organization_config_rule_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetStoredQuery",
		},
		Permission:             "GetStoredQuery",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.ListDiscoveredResources",
		},
		Permission:             "ListDiscoveredResources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceType",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_type",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
