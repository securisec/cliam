package policy

import "github.com/securisec/cliam/shared"

var PersonalizePolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListBatchInferenceJobs",
		},
		Permission: "ListBatchInferenceJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListBatchSegmentJobs",
		},
		Permission: "ListBatchSegmentJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListCampaigns",
		},
		Permission: "ListCampaigns",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasetExportJobs",
		},
		Permission: "ListDatasetExportJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasetGroups",
		},
		Permission: "ListDatasetGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasetImportJobs",
		},
		Permission: "ListDatasetImportJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListDatasets",
		},
		Permission: "ListDatasets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListEventTrackers",
		},
		Permission: "ListEventTrackers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListFilters",
		},
		Permission: "ListFilters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListRecipes",
		},
		Permission: "ListRecipes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListRecommenders",
		},
		Permission: "ListRecommenders",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListSchemas",
		},
		Permission: "ListSchemas",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListSolutionVersions",
		},
		Permission: "ListSolutionVersions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonPersonalize.ListSolutions",
		},
		Permission: "ListSolutions",
	},

	// extra
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
