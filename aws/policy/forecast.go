package policy

import "github.com/securisec/cliam/shared"

var ForecastPolicies = map[string]Service{
	"ListDatasetGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListDatasetGroups",
		},
		Permission: "ListDatasetGroups",
	},
	"ListDatasetImportJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListDatasetImportJobs",
		},
		Permission: "ListDatasetImportJobs",
	},
	"ListDatasets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListDatasets",
		},
		Permission: "ListDatasets",
	},
	"ListExplainabilities": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListExplainabilities",
		},
		Permission: "ListExplainabilities",
	},
	"ListExplainabilityExports": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListExplainabilityExports",
		},
		Permission: "ListExplainabilityExports",
	},
	"ListForecastExportJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListForecastExportJobs",
		},
		Permission: "ListForecastExportJobs",
	},
	"ListForecasts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListForecasts",
		},
		Permission: "ListForecasts",
	},
	"ListPredictorBacktestExportJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListPredictorBacktestExportJobs",
		},
		Permission: "ListPredictorBacktestExportJobs",
	},
	"ListPredictors": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListPredictors",
		},
		Permission: "ListPredictors",
	},

	// extra
	"DescribeAutoPredictor": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeAutoPredictor",
		},
		Permission:             "DescribeAutoPredictor",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PredictorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "predictor_arn",
	},
	"DescribeDataset": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeDataset",
		},
		Permission:             "DescribeDataset",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DatasetArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_arn",
	},
	"DescribeDatasetGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeDatasetGroup",
		},
		Permission:             "DescribeDatasetGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DatasetGroupArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_group_arn",
	},
	"DescribeDatasetImportJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeDatasetImportJob",
		},
		Permission:             "DescribeDatasetImportJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DatasetImportJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "dataset_import_job_arn",
	},
	"DescribeExplainability": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeExplainability",
		},
		Permission:             "DescribeExplainability",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ExplainabilityArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "explainability_arn",
	},
	"DescribeExplainabilityExport": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeExplainabilityExport",
		},
		Permission:             "DescribeExplainabilityExport",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ExplainabilityExportArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "explainability_export_arn",
	},
	"DescribeForecast": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeForecast",
		},
		Permission:             "DescribeForecast",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ForecastArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "forecast_arn",
	},
	"DescribeForecastExportJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribeForecastExportJob",
		},
		Permission:             "DescribeForecastExportJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ForecastExportJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "forecast_export_job_arn",
	},
	"DescribePredictor": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribePredictor",
		},
		Permission:             "DescribePredictor",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PredictorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "predictor_arn",
	},
	"DescribePredictorBacktestExportJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.DescribePredictorBacktestExportJob",
		},
		Permission:             "DescribePredictorBacktestExportJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PredictorBacktestExportJobArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "predictor_backtest_export_job_arn",
	},
	"GetAccuracyMetrics": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.GetAccuracyMetrics",
		},
		Permission:             "GetAccuracyMetrics",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PredictorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "predictor_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonForecast.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
