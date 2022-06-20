package policy

var Microsoft_Insights_workbooks_API = map[string]Policy{
	"Workbooks_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Insights/workbooks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Workbooks_ListBySubscription",
		Resource:    "Microsoft.Insights",
	},
	"Workbooks_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/workbooks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Workbooks_ListByResourceGroup",
		Resource:    "Microsoft.Insights",
	},
	"Workbooks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/workbooks/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Workbooks_Get",
		Resource:    "Microsoft.Insights",
	},
	"Workbooks_RevisionsList": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/workbooks/{{.resourceName}}/revisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Workbooks_RevisionsList",
		Resource:    "Microsoft.Insights",
	},
	"Workbooks_RevisionGet": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/workbooks/{{.resourceName}}/revisions/{{.revisionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Workbooks_RevisionGet",
		Resource:    "Microsoft.Insights",
	},
}
