package policy

import "github.com/securisec/cliam/shared"

var SageMakerPolicies = map[string]Service{
	"GetSagemakerServicecatalogPortfolioStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.GetSagemakerServicecatalogPortfolioStatus",
		},
		Permission: "GetSagemakerServicecatalogPortfolioStatus",
	},
	"ListActions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListActions",
		},
		Permission: "ListActions",
	},
	"ListAlgorithms": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListAlgorithms",
		},
		Permission: "ListAlgorithms",
	},
	"ListAppImageConfigs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListAppImageConfigs",
		},
		Permission: "ListAppImageConfigs",
	},
	"ListApps": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListApps",
		},
		Permission: "ListApps",
	},
	"ListArtifacts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListArtifacts",
		},
		Permission: "ListArtifacts",
	},
	"ListAssociations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListAssociations",
		},
		Permission: "ListAssociations",
	},
	"ListAutoMlJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListAutoMlJobs",
		},
		Permission: "ListAutoMlJobs",
	},
	"ListCodeRepositories": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListCodeRepositories",
		},
		Permission: "ListCodeRepositories",
	},
	"ListCompilationJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListCompilationJobs",
		},
		Permission: "ListCompilationJobs",
	},
	"ListContexts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListContexts",
		},
		Permission: "ListContexts",
	},
	"ListDataQualityJobDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListDataQualityJobDefinitions",
		},
		Permission: "ListDataQualityJobDefinitions",
	},
	"ListDeviceFleets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListDeviceFleets",
		},
		Permission: "ListDeviceFleets",
	},
	"ListDevices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListDevices",
		},
		Permission: "ListDevices",
	},
	"ListDomains": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListDomains",
		},
		Permission: "ListDomains",
	},
	"ListEdgePackagingJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListEdgePackagingJobs",
		},
		Permission: "ListEdgePackagingJobs",
	},
	"ListEndpointConfigs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListEndpointConfigs",
		},
		Permission: "ListEndpointConfigs",
	},
	"ListEndpoints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListEndpoints",
		},
		Permission: "ListEndpoints",
	},
	"ListExperiments": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListExperiments",
		},
		Permission: "ListExperiments",
	},
	"ListFeatureGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListFeatureGroups",
		},
		Permission: "ListFeatureGroups",
	},
	"ListFlowDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListFlowDefinitions",
		},
		Permission: "ListFlowDefinitions",
	},
	"ListHumanTaskUis": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListHumanTaskUis",
		},
		Permission: "ListHumanTaskUis",
	},
	"ListHyperParameterTuningJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListHyperParameterTuningJobs",
		},
		Permission: "ListHyperParameterTuningJobs",
	},
	"ListImages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListImages",
		},
		Permission: "ListImages",
	},
	"ListInferenceRecommendationsJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListInferenceRecommendationsJobs",
		},
		Permission: "ListInferenceRecommendationsJobs",
	},
	"ListLabelingJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListLabelingJobs",
		},
		Permission: "ListLabelingJobs",
	},
	"ListLineageGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListLineageGroups",
		},
		Permission: "ListLineageGroups",
	},
	"ListModelBiasJobDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListModelBiasJobDefinitions",
		},
		Permission: "ListModelBiasJobDefinitions",
	},
	"ListModelExplainabilityJobDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListModelExplainabilityJobDefinitions",
		},
		Permission: "ListModelExplainabilityJobDefinitions",
	},
	"ListModelPackageGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListModelPackageGroups",
		},
		Permission: "ListModelPackageGroups",
	},
	"ListModelPackages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListModelPackages",
		},
		Permission: "ListModelPackages",
	},
	"ListModelQualityJobDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListModelQualityJobDefinitions",
		},
		Permission: "ListModelQualityJobDefinitions",
	},
	"ListModels": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListModels",
		},
		Permission: "ListModels",
	},
	"ListMonitoringExecutions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListMonitoringExecutions",
		},
		Permission: "ListMonitoringExecutions",
	},
	"ListMonitoringSchedules": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListMonitoringSchedules",
		},
		Permission: "ListMonitoringSchedules",
	},
	"ListNotebookInstanceLifecycleConfigs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListNotebookInstanceLifecycleConfigs",
		},
		Permission: "ListNotebookInstanceLifecycleConfigs",
	},
	"ListNotebookInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListNotebookInstances",
		},
		Permission: "ListNotebookInstances",
	},
	"ListPipelineExecutionSteps": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListPipelineExecutionSteps",
		},
		Permission: "ListPipelineExecutionSteps",
	},
	"ListPipelines": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListPipelines",
		},
		Permission: "ListPipelines",
	},
	"ListProcessingJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListProcessingJobs",
		},
		Permission: "ListProcessingJobs",
	},
	"ListProjects": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListProjects",
		},
		Permission: "ListProjects",
	},
	"ListStudioLifecycleConfigs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListStudioLifecycleConfigs",
		},
		Permission: "ListStudioLifecycleConfigs",
	},
	"ListSubscribedWorkteams": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListSubscribedWorkteams",
		},
		Permission: "ListSubscribedWorkteams",
	},
	"ListTrainingJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListTrainingJobs",
		},
		Permission: "ListTrainingJobs",
	},
	"ListTransformJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListTransformJobs",
		},
		Permission: "ListTransformJobs",
	},
	"ListTrialComponents": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListTrialComponents",
		},
		Permission: "ListTrialComponents",
	},
	"ListTrials": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListTrials",
		},
		Permission: "ListTrials",
	},
	"ListUserProfiles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListUserProfiles",
		},
		Permission: "ListUserProfiles",
	},
	"ListWorkforces": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListWorkforces",
		},
		Permission: "ListWorkforces",
	},
	"ListWorkteams": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListWorkteams",
		},
		Permission: "ListWorkteams",
	},

	// extra
	"DescribeAction": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeAction",
		},
		Permission:             "DescribeAction",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ActionName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "action_name",
	},
	"DescribeAlgorithm": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeAlgorithm",
		},
		Permission:             "DescribeAlgorithm",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AlgorithmName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "algorithm_name",
	},
	"DescribeAppImageConfig": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeAppImageConfig",
		},
		Permission:             "DescribeAppImageConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AppImageConfigName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "app_image_config_name",
	},
	"DescribeArtifact": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeArtifact",
		},
		Permission:             "DescribeArtifact",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ArtifactArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "artifact_arn",
	},
	"DescribeAutoMLJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeAutoMLJob",
		},
		Permission:             "DescribeAutoMLJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoMLJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "auto_m_l_job_name",
	},
	"DescribeCodeRepository": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeCodeRepository",
		},
		Permission:             "DescribeCodeRepository",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CodeRepositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "code_repository_name",
	},
	"DescribeCompilationJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeCompilationJob",
		},
		Permission:             "DescribeCompilationJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CompilationJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "compilation_job_name",
	},
	"DescribeContext": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeContext",
		},
		Permission:             "DescribeContext",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ContextName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "context_name",
	},
	"DescribeDataQualityJobDefinition": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeDataQualityJobDefinition",
		},
		Permission:             "DescribeDataQualityJobDefinition",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobDefinitionName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_definition_name",
	},
	"DescribeDeviceFleet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeDeviceFleet",
		},
		Permission:             "DescribeDeviceFleet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DeviceFleetName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "device_fleet_name",
	},
	"DescribeDomain": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeDomain",
		},
		Permission:             "DescribeDomain",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_id",
	},
	"DescribeEdgePackagingJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeEdgePackagingJob",
		},
		Permission:             "DescribeEdgePackagingJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EdgePackagingJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "edge_packaging_job_name",
	},
	"DescribeEndpoint": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeEndpoint",
		},
		Permission:             "DescribeEndpoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_name",
	},
	"DescribeEndpointConfig": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeEndpointConfig",
		},
		Permission:             "DescribeEndpointConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointConfigName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_config_name",
	},
	"DescribeExperiment": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeExperiment",
		},
		Permission:             "DescribeExperiment",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ExperimentName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "experiment_name",
	},
	"DescribeFeatureGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeFeatureGroup",
		},
		Permission:             "DescribeFeatureGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FeatureGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "feature_group_name",
	},
	"DescribeFlowDefinition": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeFlowDefinition",
		},
		Permission:             "DescribeFlowDefinition",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FlowDefinitionName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "flow_definition_name",
	},
	"DescribeHumanTaskUi": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeHumanTaskUi",
		},
		Permission:             "DescribeHumanTaskUi",
		IsExtra:                true,
		ExtraComponentBodyKey:  "HumanTaskUiName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "human_task_ui_name",
	},
	"DescribeHyperParameterTuningJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeHyperParameterTuningJob",
		},
		Permission:             "DescribeHyperParameterTuningJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "HyperParameterTuningJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "hyper_parameter_tuning_job_name",
	},
	"DescribeImage": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeImage",
		},
		Permission:             "DescribeImage",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ImageName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "image_name",
	},
	"DescribeImageVersion": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeImageVersion",
		},
		Permission:             "DescribeImageVersion",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ImageName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "image_name",
	},
	"DescribeInferenceRecommendationsJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeInferenceRecommendationsJob",
		},
		Permission:             "DescribeInferenceRecommendationsJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_name",
	},
	"DescribeLabelingJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeLabelingJob",
		},
		Permission:             "DescribeLabelingJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LabelingJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "labeling_job_name",
	},
	"DescribeLineageGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeLineageGroup",
		},
		Permission:             "DescribeLineageGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LineageGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "lineage_group_name",
	},
	"DescribeModel": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeModel",
		},
		Permission:             "DescribeModel",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ModelName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "model_name",
	},
	"DescribeModelBiasJobDefinition": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeModelBiasJobDefinition",
		},
		Permission:             "DescribeModelBiasJobDefinition",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobDefinitionName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_definition_name",
	},
	"DescribeModelExplainabilityJobDefinition": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeModelExplainabilityJobDefinition",
		},
		Permission:             "DescribeModelExplainabilityJobDefinition",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobDefinitionName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_definition_name",
	},
	"DescribeModelPackage": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeModelPackage",
		},
		Permission:             "DescribeModelPackage",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ModelPackageName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "model_package_name",
	},
	"DescribeModelPackageGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeModelPackageGroup",
		},
		Permission:             "DescribeModelPackageGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ModelPackageGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "model_package_group_name",
	},
	"DescribeModelQualityJobDefinition": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeModelQualityJobDefinition",
		},
		Permission:             "DescribeModelQualityJobDefinition",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobDefinitionName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_definition_name",
	},
	"DescribeMonitoringSchedule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeMonitoringSchedule",
		},
		Permission:             "DescribeMonitoringSchedule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "MonitoringScheduleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "monitoring_schedule_name",
	},
	"DescribeNotebookInstance": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeNotebookInstance",
		},
		Permission:             "DescribeNotebookInstance",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NotebookInstanceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "notebook_instance_name",
	},
	"DescribeNotebookInstanceLifecycleConfig": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeNotebookInstanceLifecycleConfig",
		},
		Permission:             "DescribeNotebookInstanceLifecycleConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NotebookInstanceLifecycleConfigName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "notebook_instance_lifecycle_config_name",
	},
	"DescribePipeline": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribePipeline",
		},
		Permission:             "DescribePipeline",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PipelineName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pipeline_name",
	},
	"DescribePipelineDefinitionForExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribePipelineDefinitionForExecution",
		},
		Permission:             "DescribePipelineDefinitionForExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PipelineExecutionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pipeline_execution_arn",
	},
	"DescribePipelineExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribePipelineExecution",
		},
		Permission:             "DescribePipelineExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PipelineExecutionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pipeline_execution_arn",
	},
	"DescribeProcessingJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeProcessingJob",
		},
		Permission:             "DescribeProcessingJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProcessingJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "processing_job_name",
	},
	"DescribeProject": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeProject",
		},
		Permission:             "DescribeProject",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProjectName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "project_name",
	},
	"DescribeStudioLifecycleConfig": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeStudioLifecycleConfig",
		},
		Permission:             "DescribeStudioLifecycleConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StudioLifecycleConfigName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "studio_lifecycle_config_name",
	},
	"DescribeSubscribedWorkteam": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeSubscribedWorkteam",
		},
		Permission:             "DescribeSubscribedWorkteam",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkteamArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workteam_arn",
	},
	"DescribeTrainingJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeTrainingJob",
		},
		Permission:             "DescribeTrainingJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TrainingJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "training_job_name",
	},
	"DescribeTransformJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeTransformJob",
		},
		Permission:             "DescribeTransformJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransformJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "transform_job_name",
	},
	"DescribeTrial": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeTrial",
		},
		Permission:             "DescribeTrial",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TrialName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "trial_name",
	},
	"DescribeTrialComponent": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeTrialComponent",
		},
		Permission:             "DescribeTrialComponent",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TrialComponentName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "trial_component_name",
	},
	"DescribeWorkforce": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeWorkforce",
		},
		Permission:             "DescribeWorkforce",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkforceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workforce_name",
	},
	"DescribeWorkteam": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.DescribeWorkteam",
		},
		Permission:             "DescribeWorkteam",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkteamName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workteam_name",
	},
	"GetDeviceFleetReport": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.GetDeviceFleetReport",
		},
		Permission:             "GetDeviceFleetReport",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DeviceFleetName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "device_fleet_name",
	},
	"GetLineageGroupPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.GetLineageGroupPolicy",
		},
		Permission:             "GetLineageGroupPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LineageGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "lineage_group_name",
	},
	"GetModelPackageGroupPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.GetModelPackageGroupPolicy",
		},
		Permission:             "GetModelPackageGroupPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ModelPackageGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "model_package_group_name",
	},
	"GetSearchSuggestions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.GetSearchSuggestions",
		},
		Permission:             "GetSearchSuggestions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Resource",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource",
	},
	"ListCandidatesForAutoMLJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListCandidatesForAutoMLJob",
		},
		Permission:             "ListCandidatesForAutoMLJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoMLJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "auto_m_l_job_name",
	},
	"ListImageVersions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListImageVersions",
		},
		Permission:             "ListImageVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ImageName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "image_name",
	},
	"ListLabelingJobsForWorkteam": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListLabelingJobsForWorkteam",
		},
		Permission:             "ListLabelingJobsForWorkteam",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkteamArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workteam_arn",
	},
	"ListPipelineExecutions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListPipelineExecutions",
		},
		Permission:             "ListPipelineExecutions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PipelineName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pipeline_name",
	},
	"ListPipelineParametersForExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListPipelineParametersForExecution",
		},
		Permission:             "ListPipelineParametersForExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PipelineExecutionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pipeline_execution_arn",
	},
	"ListTags": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListTags",
		},
		Permission:             "ListTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTrainingJobsForHyperParameterTuningJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SageMaker.ListTrainingJobsForHyperParameterTuningJob",
		},
		Permission:             "ListTrainingJobsForHyperParameterTuningJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "HyperParameterTuningJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "hyper_parameter_tuning_job_name",
	},
}
