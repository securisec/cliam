package policy

import "github.com/securisec/cliam/shared"

var PersonalizePolicies = map[string]Service{
	"ListBatchInferenceJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListBatchInferenceJobs",
		},
		Permission: "ListBatchInferenceJobs",
	},
	"ListBatchSegmentJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListBatchSegmentJobs",
		},
		Permission: "ListBatchSegmentJobs",
	},
	"ListCampaigns": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListCampaigns",
		},
		Permission: "ListCampaigns",
	},
	"ListDatasetExportJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasetExportJobs",
		},
		Permission: "ListDatasetExportJobs",
	},
	"ListDatasetGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasetGroups",
		},
		Permission: "ListDatasetGroups",
	},
	"ListDatasetImportJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasetImportJobs",
		},
		Permission: "ListDatasetImportJobs",
	},
	"ListDatasets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasets",
		},
		Permission: "ListDatasets",
	},
	"ListEventTrackers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListEventTrackers",
		},
		Permission: "ListEventTrackers",
	},
	"ListFilters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListFilters",
		},
		Permission: "ListFilters",
	},
	"ListRecipes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListRecipes",
		},
		Permission: "ListRecipes",
	},
	"ListRecommenders": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListRecommenders",
		},
		Permission: "ListRecommenders",
	},
	"ListSchemas": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListSchemas",
		},
		Permission: "ListSchemas",
	},
	"ListSolutionVersions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListSolutionVersions",
		},
		Permission: "ListSolutionVersions",
	},
	"ListSolutions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListSolutions",
		},
		Permission: "ListSolutions",
	},

	// extra
	"DescribeAlgorithm": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeAlgorithm",
		},
		Permission:             "DescribeAlgorithm",
		IsExtra:                true,
		ExtraComponentBodyKey:  "algorithmArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "algorithm_arn",
	},
	"DescribeBatchInferenceJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeBatchInferenceJob",
		},
		Permission:             "DescribeBatchInferenceJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "batchInferenceJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "batch_inference_job_arn",
	},
	"DescribeBatchSegmentJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeBatchSegmentJob",
		},
		Permission:             "DescribeBatchSegmentJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "batchSegmentJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "batch_segment_job_arn",
	},
	"DescribeCampaign": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeCampaign",
		},
		Permission:             "DescribeCampaign",
		IsExtra:                true,
		ExtraComponentBodyKey:  "campaignArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "campaign_arn",
	},
	"DescribeDataset": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeDataset",
		},
		Permission:             "DescribeDataset",
		IsExtra:                true,
		ExtraComponentBodyKey:  "datasetArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_arn",
	},
	"DescribeDatasetExportJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeDatasetExportJob",
		},
		Permission:             "DescribeDatasetExportJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "datasetExportJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_export_job_arn",
	},
	"DescribeDatasetGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeDatasetGroup",
		},
		Permission:             "DescribeDatasetGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "datasetGroupArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_group_arn",
	},
	"DescribeDatasetImportJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeDatasetImportJob",
		},
		Permission:             "DescribeDatasetImportJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "datasetImportJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_import_job_arn",
	},
	"DescribeEventTracker": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeEventTracker",
		},
		Permission:             "DescribeEventTracker",
		IsExtra:                true,
		ExtraComponentBodyKey:  "eventTrackerArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "event_tracker_arn",
	},
	"DescribeFeatureTransformation": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeFeatureTransformation",
		},
		Permission:             "DescribeFeatureTransformation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "featureTransformationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "feature_transformation_arn",
	},
	"DescribeFilter": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeFilter",
		},
		Permission:             "DescribeFilter",
		IsExtra:                true,
		ExtraComponentBodyKey:  "filterArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "filter_arn",
	},
	"DescribeRecipe": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeRecipe",
		},
		Permission:             "DescribeRecipe",
		IsExtra:                true,
		ExtraComponentBodyKey:  "recipeArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "recipe_arn",
	},
	"DescribeRecommender": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeRecommender",
		},
		Permission:             "DescribeRecommender",
		IsExtra:                true,
		ExtraComponentBodyKey:  "recommenderArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "recommender_arn",
	},
	"DescribeSchema": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeSchema",
		},
		Permission:             "DescribeSchema",
		IsExtra:                true,
		ExtraComponentBodyKey:  "schemaArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "schema_arn",
	},
	"DescribeSolution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeSolution",
		},
		Permission:             "DescribeSolution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "solutionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "solution_arn",
	},
	"DescribeSolutionVersion": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.DescribeSolutionVersion",
		},
		Permission:             "DescribeSolutionVersion",
		IsExtra:                true,
		ExtraComponentBodyKey:  "solutionVersionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "solution_version_arn",
	},
	"GetSolutionMetrics": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.GetSolutionMetrics",
		},
		Permission:             "GetSolutionMetrics",
		IsExtra:                true,
		ExtraComponentBodyKey:  "solutionVersionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "solution_version_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
