package policy

import "github.com/securisec/cliam/shared"

var ConfigServicePolicies = map[string]Service{
	"DescribeAggregationAuthorizations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeAggregationAuthorizations",
		},
		Permission: "DescribeAggregationAuthorizations",
	},
	"DescribeComplianceByConfigRule": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeComplianceByConfigRule",
		},
		Permission: "DescribeComplianceByConfigRule",
	},
	"DescribeComplianceByResource": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeComplianceByResource",
		},
		Permission: "DescribeComplianceByResource",
	},
	"DescribeConfigRuleEvaluationStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigRuleEvaluationStatus",
		},
		Permission: "DescribeConfigRuleEvaluationStatus",
	},
	"DescribeConfigRules": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigRules",
		},
		Permission: "DescribeConfigRules",
	},
	"DescribeConfigurationAggregators": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationAggregators",
		},
		Permission: "DescribeConfigurationAggregators",
	},
	"DescribeConfigurationRecorderStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationRecorderStatus",
		},
		Permission: "DescribeConfigurationRecorderStatus",
	},
	"DescribeConfigurationRecorders": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConfigurationRecorders",
		},
		Permission: "DescribeConfigurationRecorders",
	},
	"DescribeConformancePackStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConformancePackStatus",
		},
		Permission: "DescribeConformancePackStatus",
	},
	"DescribeConformancePacks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeConformancePacks",
		},
		Permission: "DescribeConformancePacks",
	},
	"DescribeDeliveryChannelStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeDeliveryChannelStatus",
		},
		Permission: "DescribeDeliveryChannelStatus",
	},
	"DescribeDeliveryChannels": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeDeliveryChannels",
		},
		Permission: "DescribeDeliveryChannels",
	},
	"DescribeOrganizationConfigRuleStatuses": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConfigRuleStatuses",
		},
		Permission: "DescribeOrganizationConfigRuleStatuses",
	},
	"DescribeOrganizationConfigRules": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConfigRules",
		},
		Permission: "DescribeOrganizationConfigRules",
	},
	"DescribeOrganizationConformancePackStatuses": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConformancePackStatuses",
		},
		Permission: "DescribeOrganizationConformancePackStatuses",
	},
	"DescribeOrganizationConformancePacks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeOrganizationConformancePacks",
		},
		Permission: "DescribeOrganizationConformancePacks",
	},
	"DescribePendingAggregationRequests": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribePendingAggregationRequests",
		},
		Permission: "DescribePendingAggregationRequests",
	},
	"DescribeRetentionConfigurations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.DescribeRetentionConfigurations",
		},
		Permission: "DescribeRetentionConfigurations",
	},
	"GetComplianceSummaryByConfigRule": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetComplianceSummaryByConfigRule",
		},
		Permission: "GetComplianceSummaryByConfigRule",
	},
	"GetComplianceSummaryByResourceType": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetComplianceSummaryByResourceType",
		},
		Permission: "GetComplianceSummaryByResourceType",
	},
	"GetCustomRulePolicy": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetCustomRulePolicy",
		},
		Permission: "GetCustomRulePolicy",
	},
	"GetDiscoveredResourceCounts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.GetDiscoveredResourceCounts",
		},
		Permission: "GetDiscoveredResourceCounts",
	},
	"ListStoredQueries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StarlingDoveService.ListStoredQueries",
		},
		Permission: "ListStoredQueries",
	},

	// extra
	"DescribeAggregateComplianceByConfigRules": {
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
	"DescribeAggregateComplianceByConformancePacks": {
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
	"DescribeConfigurationAggregatorSourcesStatus": {
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
	"DescribeConformancePackCompliance": {
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
	"DescribeRemediationConfigurations": {
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
	"DescribeRemediationExceptions": {
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
	"DescribeRemediationExecutionStatus": {
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
	"GetAggregateConfigRuleComplianceSummary": {
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
	"GetAggregateConformancePackComplianceSummary": {
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
	"GetAggregateDiscoveredResourceCounts": {
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
	"GetComplianceDetailsByConfigRule": {
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
	"GetConformancePackComplianceDetails": {
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
	"GetConformancePackComplianceSummary": {
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
	"GetOrganizationConfigRuleDetailedStatus": {
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
	"GetOrganizationConformancePackDetailedStatus": {
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
	"GetOrganizationCustomRulePolicy": {
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
	"GetStoredQuery": {
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
	"ListDiscoveredResources": {
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
	"ListTagsForResource": {
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
