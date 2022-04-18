package policy

var PanoramaPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "application-instances",
		Permission:    "ListApplicationInstances",
	},
	{
		Method:        "GET",
		ServiceSuffix: "devices",
		Permission:    "ListDevices",
	},
	{
		Method:        "GET",
		ServiceSuffix: "jobs",
		Permission:    "ListDevicesJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "packages/template-job",
		Permission:    "ListNodeFromTemplateJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "nodes",
		Permission:    "ListNodes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "packages/import-jobs",
		Permission:    "ListPackageImportJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "packages",
		Permission:    "ListPackages",
	},
}
