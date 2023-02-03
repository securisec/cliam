package policy

// Microsoft_ChangeAnalysis_changeanalysis policy
var Microsoft_ChangeAnalysis_changeanalysis = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.ChangeAnalysis/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ChangeAnalysis",
	},
	"ResourceChanges_List": {
		Path:   "/{{.resourceId}}/providers/Microsoft.ChangeAnalysis/resourceChanges",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "ResourceChanges_List",
		Resource:    "Microsoft.ChangeAnalysis",
	},
	"Changes_ListChangesByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ChangeAnalysis/changes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "Changes_ListChangesByResourceGroup",
		Resource:    "Microsoft.ChangeAnalysis",
	},
	"Changes_ListChangesBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ChangeAnalysis/changes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-04-01",
		},
		OperationID: "Changes_ListChangesBySubscription",
		Resource:    "Microsoft.ChangeAnalysis",
	},
}
