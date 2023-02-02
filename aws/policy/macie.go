package policy

import "github.com/securisec/cliam/shared"

// MaciePolicies policy
var MaciePolicies = map[string]Service{
	"ListMemberAccounts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MacieService.ListMemberAccounts",
		},
		Permission: "ListMemberAccounts",
	},
	"ListS3Resources": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MacieService.ListS3Resources",
		},
		Permission: "ListS3Resources",
	},

	// extra
}
