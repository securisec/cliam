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

	// extra
	{
		ServiceSuffix:          "/datasets/{{.name}}",
		Permission:             "DescribeDataset",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/jobs/{{.name}}",
		Permission:             "DescribeJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/projects/{{.name}}",
		Permission:             "DescribeProject",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/recipes/{{.name}}",
		Permission:             "DescribeRecipe",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/rulesets/{{.name}}",
		Permission:             "DescribeRuleset",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/schedules/{{.name}}",
		Permission:             "DescribeSchedule",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/jobs/{{.name}}/jobRuns",
		Permission:             "ListJobRuns",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/recipeVersions",
		Permission:             "ListRecipeVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
