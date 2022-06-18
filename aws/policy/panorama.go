package policy

var PanoramaPolicies = map[string]Service{
	"ListApplicationInstances": {
		Method:        "GET",
		ServiceSuffix: "application-instances",
		Permission:    "ListApplicationInstances",
	},
	"ListDevices": {
		Method:        "GET",
		ServiceSuffix: "devices",
		Permission:    "ListDevices",
	},
	"ListDevicesJobs": {
		Method:        "GET",
		ServiceSuffix: "jobs",
		Permission:    "ListDevicesJobs",
	},
	"ListNodeFromTemplateJobs": {
		Method:        "GET",
		ServiceSuffix: "packages/template-job",
		Permission:    "ListNodeFromTemplateJobs",
	},
	"ListNodes": {
		Method:        "GET",
		ServiceSuffix: "nodes",
		Permission:    "ListNodes",
	},
	"ListPackageImportJobs": {
		Method:        "GET",
		ServiceSuffix: "packages/import-jobs",
		Permission:    "ListPackageImportJobs",
	},
	"ListPackages": {
		Method:        "GET",
		ServiceSuffix: "packages",
		Permission:    "ListPackages",
	},

	// extra
	"DescribeApplicationInstance": {
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}",
		Permission:             "DescribeApplicationInstance",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	"DescribeApplicationInstanceDetails": {
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}/details",
		Permission:             "DescribeApplicationInstanceDetails",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	"DescribeDevice": {
		ServiceSuffix:          "/devices/{{.device_id}}",
		Permission:             "DescribeDevice",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "device_id",
	},
	"DescribeDeviceJob": {
		ServiceSuffix:          "/jobs/{{.job_id}}",
		Permission:             "DescribeDeviceJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	"DescribeNode": {
		ServiceSuffix:          "/nodes/{{.node_id}}",
		Permission:             "DescribeNode",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "node_id",
	},
	"DescribeNodeFromTemplateJob": {
		ServiceSuffix:          "/packages/template-job/{{.job_id}}",
		Permission:             "DescribeNodeFromTemplateJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	"DescribePackage": {
		ServiceSuffix:          "/packages/metadata/{{.package_id}}",
		Permission:             "DescribePackage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "package_id",
	},
	"DescribePackageImportJob": {
		ServiceSuffix:          "/packages/import-jobs/{{.job_id}}",
		Permission:             "DescribePackageImportJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	"ListApplicationInstanceDependencies": {
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}/package-dependencies",
		Permission:             "ListApplicationInstanceDependencies",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	"ListApplicationInstanceNodeInstances": {
		ServiceSuffix:          "/application-instances/{{.application_instance_id}}/node-instances",
		Permission:             "ListApplicationInstanceNodeInstances",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_instance_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
