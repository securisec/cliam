package policy

// IOTAnalyticsPolicies policy
var IOTAnalyticsPolicies = map[string]Service{
	"DescribeLoggingOptions": {
		Method:        "GET",
		ServiceSuffix: "logging",
		Permission:    "DescribeLoggingOptions",
	},
	"ListChannels": {
		Method:        "GET",
		ServiceSuffix: "channels",
		Permission:    "ListChannels",
	},
	"ListDatasets": {
		Method:        "GET",
		ServiceSuffix: "datasets",
		Permission:    "ListDatasets",
	},
	"ListDatastores": {
		Method:        "GET",
		ServiceSuffix: "datastores",
		Permission:    "ListDatastores",
	},
	"ListPipelines": {
		Method:        "GET",
		ServiceSuffix: "pipelines",
		Permission:    "ListPipelines",
	},

	// extra
	"DescribeChannel": {
		ServiceSuffix:          "/channels/{{.channel_name}}",
		Permission:             "DescribeChannel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "channel_name",
	},
	"DescribeDataset": {
		ServiceSuffix:          "/datasets/{{.dataset_name}}",
		Permission:             "DescribeDataset",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "dataset_name",
	},
	"DescribeDatastore": {
		ServiceSuffix:          "/datastores/{{.datastore_name}}",
		Permission:             "DescribeDatastore",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "datastore_name",
	},
	"DescribePipeline": {
		ServiceSuffix:          "/pipelines/{{.pipeline_name}}",
		Permission:             "DescribePipeline",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "pipeline_name",
	},
	"GetDatasetContent": {
		ServiceSuffix:          "/datasets/{{.dataset_name}}/content",
		Permission:             "GetDatasetContent",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "dataset_name",
	},
	"ListDatasetContents": {
		ServiceSuffix:          "/datasets/{{.dataset_name}}/contents",
		Permission:             "ListDatasetContents",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "dataset_name",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
