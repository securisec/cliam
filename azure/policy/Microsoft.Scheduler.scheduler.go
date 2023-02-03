package policy

// Microsoft_Scheduler_scheduler policy
var Microsoft_Scheduler_scheduler = map[string]Policy{
	"JobCollections_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Scheduler/jobCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "JobCollections_ListBySubscription",
		Resource:    "Microsoft.Scheduler",
	},
	"JobCollections_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "JobCollections_ListByResourceGroup",
		Resource:    "Microsoft.Scheduler",
	},
	"JobCollections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "JobCollections_Get",
		Resource:    "Microsoft.Scheduler",
	},
	"JobCollections_Enable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}/enable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "JobCollections_Enable",
		Resource:    "Microsoft.Scheduler",
	},
	"JobCollections_Disable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "JobCollections_Disable",
		Resource:    "Microsoft.Scheduler",
	},
	"Jobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}/jobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "Jobs_Get",
		Resource:    "Microsoft.Scheduler",
	},
	"Jobs_Run": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}/jobs/{{.jobName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "Jobs_Run",
		Resource:    "Microsoft.Scheduler",
	},
	"Jobs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "Jobs_List",
		Resource:    "Microsoft.Scheduler",
	},
	"Jobs_ListJobHistory": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Scheduler/jobCollections/{{.jobCollectionName}}/jobs/{{.jobName}}/history",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-03-01",
		},
		OperationID: "Jobs_ListJobHistory",
		Resource:    "Microsoft.Scheduler",
	},
}
