package policy

import "github.com/securisec/cliam/shared"

// ECSPolicies policy
var ECSPolicies = map[string]Service{
	"DescribeCapacityProviders": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DescribeCapacityProviders",
		},
		Permission: "DescribeCapacityProviders",
	},
	"DescribeClusters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DescribeClusters",
		},
		Permission: "DescribeClusters",
	},
	"DiscoverPollEndpoint": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DiscoverPollEndpoint",
		},
		Permission: "DiscoverPollEndpoint",
	},
	"ListAccountSettings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListAccountSettings",
		},
		Permission: "ListAccountSettings",
	},
	"ListClusters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListClusters",
		},
		Permission: "ListClusters",
	},
	"ListContainerInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListContainerInstances",
		},
		Permission: "ListContainerInstances",
	},
	"ListServices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListServices",
		},
		Permission: "ListServices",
	},
	"ListTaskDefinitionFamilies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListTaskDefinitionFamilies",
		},
		Permission: "ListTaskDefinitionFamilies",
	},
	"ListTaskDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListTaskDefinitions",
		},
		Permission: "ListTaskDefinitions",
	},
	"ListTasks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListTasks",
		},
		Permission: "ListTasks",
	},
	"RegisterContainerInstance": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.RegisterContainerInstance",
		},
		Permission: "RegisterContainerInstance",
	},
	"SubmitContainerStateChange": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.SubmitContainerStateChange",
		},
		Permission: "SubmitContainerStateChange",
	},
	"SubmitTaskStateChange": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.SubmitTaskStateChange",
		},
		Permission: "SubmitTaskStateChange",
	},

	// extra
	"DescribeContainerInstances": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DescribeContainerInstances",
		},
		Permission:             "DescribeContainerInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "containerInstances",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "container_instances",
	},
	"DescribeServices": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DescribeServices",
		},
		Permission:             "DescribeServices",
		IsExtra:                true,
		ExtraComponentBodyKey:  "services",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "services",
	},
	"DescribeTaskDefinition": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DescribeTaskDefinition",
		},
		Permission:             "DescribeTaskDefinition",
		IsExtra:                true,
		ExtraComponentBodyKey:  "taskDefinition",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "task_definition",
	},
	"DescribeTasks": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.DescribeTasks",
		},
		Permission:             "DescribeTasks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "tasks",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "tasks",
	},
	"GetTaskProtection": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.GetTaskProtection",
		},
		Permission:             "GetTaskProtection",
		IsExtra:                true,
		ExtraComponentBodyKey:  "cluster",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "cluster",
	},
	"ListAttributes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListAttributes",
		},
		Permission:             "ListAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "targetType",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "target_type",
	},
	"ListServicesByNamespace": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListServicesByNamespace",
		},
		Permission:             "ListServicesByNamespace",
		IsExtra:                true,
		ExtraComponentBodyKey:  "namespace",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "namespace",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerServiceV20141113.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
