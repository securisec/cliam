package policy

var DataBrewPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "datasets",
		Permission:    "ListDatasets",
	},
	{
		Method:        "GET",
		ServiceSuffix: "jobs",
		Permission:    "ListJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "projects",
		Permission:    "ListProjects",
	},
	{
		Method:        "GET",
		ServiceSuffix: "recipes",
		Permission:    "ListRecipes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "rulesets",
		Permission:    "ListRulesets",
	},
	{
		Method:        "GET",
		ServiceSuffix: "schedules",
		Permission:    "ListSchedules",
	},
}
