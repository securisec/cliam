package policy

import "github.com/securisec/cliam/shared"

var IOTPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "keys-and-certificate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateKeysAndCertificate",
	},
	{
		Method:        "GET",
		ServiceSuffix: "audit/configuration",
		Permission:    "DescribeAccountAuditConfiguration",
	},
	{
		Method:        "GET",
		ServiceSuffix: "default-authorizer",
		Permission:    "DescribeDefaultAuthorizer",
	},
	{
		Method:        "GET",
		ServiceSuffix: "endpoint",
		Permission:    "DescribeEndpoint",
	},
	{
		Method:        "GET",
		ServiceSuffix: "event-configurations",
		Permission:    "DescribeEventConfigurations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "behavior-model-training/summaries",
		Permission:    "GetBehaviorModelTrainingSummaries",
	},
	{
		Method:        "POST",
		ServiceSuffix: "effective-policies",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetEffectivePolicies",
	},
	{
		Method:        "GET",
		ServiceSuffix: "indexing/config",
		Permission:    "GetIndexingConfiguration",
	},
	{
		Method:        "GET",
		ServiceSuffix: "loggingOptions",
		Permission:    "GetLoggingOptions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "registrationcode",
		Permission:    "GetRegistrationCode",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v2LoggingOptions",
		Permission:    "GetV2LoggingOptions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "active-violations",
		Permission:    "ListActiveViolations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "audit/findings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListAuditFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "audit/suppressions/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListAuditSuppressions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "authorizers/",
		Permission:    "ListAuthorizers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "billing-groups",
		Permission:    "ListBillingGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "cacertificates",
		Permission:    "ListCACertificates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "certificates",
		Permission:    "ListCertificates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "custom-metrics",
		Permission:    "ListCustomMetrics",
	},
	{
		Method:        "GET",
		ServiceSuffix: "detect/mitigationactions/executions",
		Permission:    "ListDetectMitigationActionsExecutions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "dimensions",
		Permission:    "ListDimensions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "domainConfigurations",
		Permission:    "ListDomainConfigurations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "fleet-metrics",
		Permission:    "ListFleetMetrics",
	},
	{
		Method:        "GET",
		ServiceSuffix: "indices",
		Permission:    "ListIndices",
	},
	{
		Method:        "GET",
		ServiceSuffix: "job-templates",
		Permission:    "ListJobTemplates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "jobs",
		Permission:    "ListJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "managed-job-templates",
		Permission:    "ListManagedJobTemplates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "mitigationactions/actions",
		Permission:    "ListMitigationActions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "otaUpdates",
		Permission:    "ListOTAUpdates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "certificates-out-going",
		Permission:    "ListOutgoingCertificates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "policies",
		Permission:    "ListPolicies",
	},
	{
		Method:        "GET",
		ServiceSuffix: "provisioning-templates",
		Permission:    "ListProvisioningTemplates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "role-aliases",
		Permission:    "ListRoleAliases",
	},
	{
		Method:        "GET",
		ServiceSuffix: "audit/scheduledaudits",
		Permission:    "ListScheduledAudits",
	},
	{
		Method:        "GET",
		ServiceSuffix: "security-profiles",
		Permission:    "ListSecurityProfiles",
	},
	{
		Method:        "GET",
		ServiceSuffix: "streams",
		Permission:    "ListStreams",
	},
	{
		Method:        "GET",
		ServiceSuffix: "thing-groups",
		Permission:    "ListThingGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "thing-registration-tasks",
		Permission:    "ListThingRegistrationTasks",
	},
	{
		Method:        "GET",
		ServiceSuffix: "thing-types",
		Permission:    "ListThingTypes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "things",
		Permission:    "ListThings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "destinations",
		Permission:    "ListTopicRuleDestinations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "rules",
		Permission:    "ListTopicRules",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v2LoggingLevel",
		Permission:    "ListV2LoggingLevels",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v2LoggingOptions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "SetV2LoggingOptions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "indexing/config",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "UpdateIndexingConfiguration",
	},
}
