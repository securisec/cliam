package policy

import "github.com/securisec/cliam/shared"

// Route53DomainsPolicies policy
var Route53DomainsPolicies = map[string]Service{
	"GetContactReachabilityStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.GetContactReachabilityStatus",
		},
		Permission: "GetContactReachabilityStatus",
	},
	"ListDomains": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.ListDomains",
		},
		Permission: "ListDomains",
	},
	"ListOperations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.ListOperations",
		},
		Permission: "ListOperations",
	},
	"ListPrices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.ListPrices",
		},
		Permission: "ListPrices",
	},
	"ResendContactReachabilityEmail": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.ResendContactReachabilityEmail",
		},
		Permission: "ResendContactReachabilityEmail",
	},
	"ViewBilling": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.ViewBilling",
		},
		Permission: "ViewBilling",
	},

	// extra
	"GetDomainDetail": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.GetDomainDetail",
		},
		Permission:             "GetDomainDetail",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetOperationDetail": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.GetOperationDetail",
		},
		Permission:             "GetOperationDetail",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OperationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "operation_id",
	},
	"ListTagsForDomain": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Route53Domains_v20140515.ListTagsForDomain",
		},
		Permission:             "ListTagsForDomain",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_name",
	},
}
