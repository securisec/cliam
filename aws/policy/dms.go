package policy

import "github.com/securisec/cliam/shared"

var DMSPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeAccountAttributes",
		},
		Permission: "DescribeAccountAttributes",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeApplicableIndividualAssessments",
		},
		Permission: "DescribeApplicableIndividualAssessments",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeCertificates",
		},
		Permission: "DescribeCertificates",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeConnections",
		},
		Permission: "DescribeConnections",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpointTypes",
		},
		Permission: "DescribeEndpointTypes",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpoints",
		},
		Permission: "DescribeEndpoints",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEventCategories",
		},
		Permission: "DescribeEventCategories",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEventSubscriptions",
		},
		Permission: "DescribeEventSubscriptions",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeOrderableReplicationInstances",
		},
		Permission: "DescribeOrderableReplicationInstances",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribePendingMaintenanceActions",
		},
		Permission: "DescribePendingMaintenanceActions",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationInstances",
		},
		Permission: "DescribeReplicationInstances",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationSubnetGroups",
		},
		Permission: "DescribeReplicationSubnetGroups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskAssessmentResults",
		},
		Permission: "DescribeReplicationTaskAssessmentResults",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskAssessmentRuns",
		},
		Permission: "DescribeReplicationTaskAssessmentRuns",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskIndividualAssessments",
		},
		Permission: "DescribeReplicationTaskIndividualAssessments",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTasks",
		},
		Permission: "DescribeReplicationTasks",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.ListTagsForResource",
		},
		Permission: "ListTagsForResource",
	},
}
