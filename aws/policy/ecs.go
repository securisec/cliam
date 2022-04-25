package policy

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
}
