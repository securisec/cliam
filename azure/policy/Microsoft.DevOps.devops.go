package policy

var Microsoft_DevOps_devops = []Policy{
	{
		Path:   "/providers/Microsoft.DevOps/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-13-preview",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DevOps",
	}, {
		Path:   "/providers/Microsoft.DevOps/pipelineTemplateDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-13-preview",
		},
		OperationID: "PipelineTemplateDefinitions_List",
		Resource:    "Microsoft.DevOps",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevOps/pipelines/{{.pipelineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-13-preview",
		},
		OperationID: "Pipelines_Get",
		Resource:    "Microsoft.DevOps",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevOps/pipelines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-13-preview",
		},
		OperationID: "Pipelines_ListByResourceGroup",
		Resource:    "Microsoft.DevOps",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DevOps/pipelines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-07-13-preview",
		},
		OperationID: "Pipelines_ListBySubscription",
		Resource:    "Microsoft.DevOps",
	},
}
