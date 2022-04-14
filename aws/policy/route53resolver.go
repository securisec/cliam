package policy

import "github.com/securisec/cliam/shared"

var Route53ResolverPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallConfigs",
		},
		Permission: "ListFirewallConfigs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallDomainLists",
		},
		Permission: "ListFirewallDomainLists",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallRuleGroupAssociations",
		},
		Permission: "ListFirewallRuleGroupAssociations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListFirewallRuleGroups",
		},
		Permission: "ListFirewallRuleGroups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverConfigs",
		},
		Permission: "ListResolverConfigs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverDnssecConfigs",
		},
		Permission: "ListResolverDnssecConfigs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverEndpoints",
		},
		Permission: "ListResolverEndpoints",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverQueryLogConfigAssociations",
		},
		Permission: "ListResolverQueryLogConfigAssociations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverQueryLogConfigs",
		},
		Permission: "ListResolverQueryLogConfigs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverRuleAssociations",
		},
		Permission: "ListResolverRuleAssociations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Route53Resolver.ListResolverRules",
		},
		Permission: "ListResolverRules",
	},
}
