package policy

import "github.com/securisec/cliam/shared"

var OrganizationsPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeOrganization",
		},
		Permission: "DescribeOrganization",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAccounts",
		},
		Permission: "ListAccounts",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAwsServiceAccessForOrganization",
		},
		Permission: "ListAwsServiceAccessForOrganization",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListCreateAccountStatus",
		},
		Permission: "ListCreateAccountStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListDelegatedAdministrators",
		},
		Permission: "ListDelegatedAdministrators",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListHandshakesForAccount",
		},
		Permission: "ListHandshakesForAccount",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListHandshakesForOrganization",
		},
		Permission: "ListHandshakesForOrganization",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListRoots",
		},
		Permission: "ListRoots",
	},
}
