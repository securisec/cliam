package policy

var AccessAnalyzersPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "analyzer",
		Permission:    "ListAnalyzers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "policy/generation",
		Permission:    "ListPolicyGenerations",
	},
}
