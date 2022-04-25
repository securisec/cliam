package policy

import "github.com/securisec/cliam/shared"

var CodecommitPolicies = []Service{
	{
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "ListApprovalRuleTemplates",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListApprovalRuleTemplates",
		},
	},
	{
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "ListRepositories",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListRepositories",
		},
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.DescribePullRequestEvents",
		},
		Permission:             "DescribePullRequestEvents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "pullRequestId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pull_request_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetApprovalRuleTemplate",
		},
		Permission:             "GetApprovalRuleTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "approvalRuleTemplateName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "approval_rule_template_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetComment",
		},
		Permission:             "GetComment",
		IsExtra:                true,
		ExtraComponentBodyKey:  "commentId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "comment_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetCommentReactions",
		},
		Permission:             "GetCommentReactions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "commentId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "comment_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetCommentsForPullRequest",
		},
		Permission:             "GetCommentsForPullRequest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "pullRequestId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pull_request_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetPullRequest",
		},
		Permission:             "GetPullRequest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "pullRequestId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pull_request_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetRepository",
		},
		Permission:             "GetRepository",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.GetRepositoryTriggers",
		},
		Permission:             "GetRepositoryTriggers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListAssociatedApprovalRuleTemplatesForRepository",
		},
		Permission:             "ListAssociatedApprovalRuleTemplatesForRepository",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListBranches",
		},
		Permission:             "ListBranches",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListPullRequests",
		},
		Permission:             "ListPullRequests",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListRepositoriesForApprovalRuleTemplate",
		},
		Permission:             "ListRepositoriesForApprovalRuleTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "approvalRuleTemplateName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "approval_rule_template_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeCommit_20150413.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
