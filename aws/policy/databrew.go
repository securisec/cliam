package policy

// DataBrewPolicies policy
var DataBrewPolicies = map[string]Service{
	"ListDatasets": {
		Method:        "GET",
		ServiceSuffix: "datasets",
		Permission:    "ListDatasets",
	},
	"ListJobs": {
		Method:        "GET",
		ServiceSuffix: "jobs",
		Permission:    "ListJobs",
	},
	"ListProjects": {
		Method:        "GET",
		ServiceSuffix: "projects",
		Permission:    "ListProjects",
	},
	"ListRecipes": {
		Method:        "GET",
		ServiceSuffix: "recipes",
		Permission:    "ListRecipes",
	},
	"ListRulesets": {
		Method:        "GET",
		ServiceSuffix: "rulesets",
		Permission:    "ListRulesets",
	},
	"ListSchedules": {
		Method:        "GET",
		ServiceSuffix: "schedules",
		Permission:    "ListSchedules",
	},

	// extra
	"DescribeDataset": {
		ServiceSuffix:          "/datasets/{{.name}}",
		Permission:             "DescribeDataset",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"DescribeJob": {
		ServiceSuffix:          "/jobs/{{.name}}",
		Permission:             "DescribeJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"DescribeProject": {
		ServiceSuffix:          "/projects/{{.name}}",
		Permission:             "DescribeProject",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"DescribeRecipe": {
		ServiceSuffix:          "/recipes/{{.name}}",
		Permission:             "DescribeRecipe",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"DescribeRuleset": {
		ServiceSuffix:          "/rulesets/{{.name}}",
		Permission:             "DescribeRuleset",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"DescribeSchedule": {
		ServiceSuffix:          "/schedules/{{.name}}",
		Permission:             "DescribeSchedule",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"ListJobRuns": {
		ServiceSuffix:          "/jobs/{{.name}}/jobRuns",
		Permission:             "ListJobRuns",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"ListRecipeVersions": {
		ServiceSuffix:          "/recipeVersions",
		Permission:             "ListRecipeVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
