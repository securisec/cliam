package policy

import "github.com/securisec/cliam/shared"

var ECSPolicies = []Service{
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.DescribeClusters",
		},
		Permission: "DescribeClusters",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListAccountSettings",
		},
		Permission: "ListAccountSettings",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListClusters",
		},
		Permission: "ListClusters",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListContainerInstances",
		},
		Permission: "ListContainerInstances",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListServices",
		},
		Permission: "ListServices",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListTaskDefinitionFamilies",
		},
		Permission: "ListTaskDefinitionFamilies",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListTaskDefinitions",
		},
		Permission: "ListTaskDefinitions",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerServiceV20141113.ListTasks",
		},
		Permission: "ListTasks",
	},

	// extra
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
