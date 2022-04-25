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

	// extra
	{
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}",
		Permission:             "DescribeApplicationInstance",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	{
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}/details",
		Permission:             "DescribeApplicationInstanceDetails",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	{
		ServiceSuffix:          "/devices/{{.device_id}}",
		Permission:             "DescribeDevice",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "device_id",
	},
	{
		ServiceSuffix:          "/jobs/{{.job_id}}",
		Permission:             "DescribeDeviceJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	{
		ServiceSuffix:          "/nodes/{{.node_id}}",
		Permission:             "DescribeNode",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "node_id",
	},
	{
		ServiceSuffix:          "/packages/template-job/{{.job_id}}",
		Permission:             "DescribeNodeFromTemplateJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	{
		ServiceSuffix:          "/packages/metadata/{{.package_id}}",
		Permission:             "DescribePackage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "package_id",
	},
	{
		ServiceSuffix:          "/packages/import-jobs/{{.job_id}}",
		Permission:             "DescribePackageImportJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	{
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}/package-dependencies",
		Permission:             "ListApplicationInstanceDependencies",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	{
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}/node-instances",
		Permission:             "ListApplicationInstanceNodeInstances",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
