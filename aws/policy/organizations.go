package policy

import "github.com/securisec/cliam/shared"

var OrganizationsPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeOrganization",
		},
		Permission: "DescribeOrganization",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAccounts",
		},
		Permission: "ListAccounts",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAwsServiceAccessForOrganization",
		},
		Permission: "ListAwsServiceAccessForOrganization",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListCreateAccountStatus",
		},
		Permission: "ListCreateAccountStatus",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListDelegatedAdministrators",
		},
		Permission: "ListDelegatedAdministrators",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListHandshakesForAccount",
		},
		Permission: "ListHandshakesForAccount",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListHandshakesForOrganization",
		},
		Permission: "ListHandshakesForOrganization",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListRoots",
		},
		Permission: "ListRoots",
	},
}
