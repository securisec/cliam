package policy

import "github.com/securisec/cliam/shared"

var Route53ResolverPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallConfigs",
		},
		Permission: "ListFirewallConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallDomainLists",
		},
		Permission: "ListFirewallDomainLists",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallRuleGroupAssociations",
		},
		Permission: "ListFirewallRuleGroupAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallRuleGroups",
		},
		Permission: "ListFirewallRuleGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverConfigs",
		},
		Permission: "ListResolverConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverDnssecConfigs",
		},
		Permission: "ListResolverDnssecConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverEndpoints",
		},
		Permission: "ListResolverEndpoints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverQueryLogConfigAssociations",
		},
		Permission: "ListResolverQueryLogConfigAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverQueryLogConfigs",
		},
		Permission: "ListResolverQueryLogConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverRuleAssociations",
		},
		Permission: "ListResolverRuleAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverRules",
		},
		Permission: "ListResolverRules",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetFirewallConfig",
		},
		Permission:             "GetFirewallConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetFirewallDomainList",
		},
		Permission:             "GetFirewallDomainList",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FirewallDomainListId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "firewall_domain_list_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetFirewallRuleGroup",
		},
		Permission:             "GetFirewallRuleGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FirewallRuleGroupId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "firewall_rule_group_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetFirewallRuleGroupAssociation",
		},
		Permission:             "GetFirewallRuleGroupAssociation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FirewallRuleGroupAssociationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "firewall_rule_group_association_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetFirewallRuleGroupPolicy",
		},
		Permission:             "GetFirewallRuleGroupPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverConfig",
		},
		Permission:             "GetResolverConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverDnssecConfig",
		},
		Permission:             "GetResolverDnssecConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverEndpoint",
		},
		Permission:             "GetResolverEndpoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResolverEndpointId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resolver_endpoint_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverQueryLogConfig",
		},
		Permission:             "GetResolverQueryLogConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResolverQueryLogConfigId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resolver_query_log_config_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverQueryLogConfigAssociation",
		},
		Permission:             "GetResolverQueryLogConfigAssociation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResolverQueryLogConfigAssociationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resolver_query_log_config_association_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverQueryLogConfigPolicy",
		},
		Permission:             "GetResolverQueryLogConfigPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverRule",
		},
		Permission:             "GetResolverRule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResolverRuleId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resolver_rule_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverRuleAssociation",
		},
		Permission:             "GetResolverRuleAssociation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResolverRuleAssociationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resolver_rule_association_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.GetResolverRulePolicy",
		},
		Permission:             "GetResolverRulePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallDomains",
		},
		Permission:             "ListFirewallDomains",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FirewallDomainListId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "firewall_domain_list_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallRules",
		},
		Permission:             "ListFirewallRules",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FirewallRuleGroupId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "firewall_rule_group_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverEndpointIpAddresses",
		},
		Permission:             "ListResolverEndpointIpAddresses",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResolverEndpointId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resolver_endpoint_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
