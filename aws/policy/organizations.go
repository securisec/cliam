package policy

import "github.com/securisec/cliam/shared"

var OrganizationsPolicies = map[string]Service{
	"DescribeOrganization": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeOrganization",
		},
		Permission: "DescribeOrganization",
	},
	"ListAccounts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAccounts",
		},
		Permission: "ListAccounts",
	},
	"ListAwsServiceAccessForOrganization": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAwsServiceAccessForOrganization",
		},
		Permission: "ListAwsServiceAccessForOrganization",
	},
	"ListCreateAccountStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListCreateAccountStatus",
		},
		Permission: "ListCreateAccountStatus",
	},
	"ListDelegatedAdministrators": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListDelegatedAdministrators",
		},
		Permission: "ListDelegatedAdministrators",
	},
	"ListHandshakesForAccount": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListHandshakesForAccount",
		},
		Permission: "ListHandshakesForAccount",
	},
	"ListHandshakesForOrganization": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListHandshakesForOrganization",
		},
		Permission: "ListHandshakesForOrganization",
	},
	"ListRoots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListRoots",
		},
		Permission: "ListRoots",
	},

	// extra
	"DescribeAccount": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeAccount",
		},
		Permission:             "DescribeAccount",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccountId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "account_id",
	},
	"DescribeCreateAccountStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeCreateAccountStatus",
		},
		Permission:             "DescribeCreateAccountStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CreateAccountRequestId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "create_account_request_id",
	},
	"DescribeEffectivePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeEffectivePolicy",
		},
		Permission:             "DescribeEffectivePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyType",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "policy_type",
	},
	"DescribeHandshake": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeHandshake",
		},
		Permission:             "DescribeHandshake",
		IsExtra:                true,
		ExtraComponentBodyKey:  "HandshakeId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "handshake_id",
	},
	"DescribeOrganizationalUnit": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribeOrganizationalUnit",
		},
		Permission:             "DescribeOrganizationalUnit",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OrganizationalUnitId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "organizational_unit_id",
	},
	"DescribePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.DescribePolicy",
		},
		Permission:             "DescribePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "policy_id",
	},
	"ListAccountsForParent": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListAccountsForParent",
		},
		Permission:             "ListAccountsForParent",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ParentId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "parent_id",
	},
	"ListDelegatedServicesForAccount": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListDelegatedServicesForAccount",
		},
		Permission:             "ListDelegatedServicesForAccount",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccountId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "account_id",
	},
	"ListOrganizationalUnitsForParent": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListOrganizationalUnitsForParent",
		},
		Permission:             "ListOrganizationalUnitsForParent",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ParentId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "parent_id",
	},
	"ListParents": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListParents",
		},
		Permission:             "ListParents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ChildId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "child_id",
	},
	"ListPolicies": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListPolicies",
		},
		Permission:             "ListPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Filter",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "filter",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	"ListTargetsForPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSOrganizationsV20161128.ListTargetsForPolicy",
		},
		Permission:             "ListTargetsForPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "policy_id",
	},
}
