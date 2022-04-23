package policy

import "github.com/securisec/cliam/shared"

var CodecommitPolicies = []Service{
	{
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "ListApprovalRuleTemplates",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListApprovalRuleTemplates",
		},
	},
	{
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "ListRepositories",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListRepositories",
		},
	},
}
