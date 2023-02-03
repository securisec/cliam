package policy

// Microsoft_LabServices_Labs policy
var Microsoft_LabServices_Labs = map[string]Policy{
	"Labs_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LabServices/labs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Labs_ListBySubscription",
		Resource:    "Microsoft.LabServices",
	},
	"Labs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Labs_ListByResourceGroup",
		Resource:    "Microsoft.LabServices",
	},
	"Labs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Labs_Get",
		Resource:    "Microsoft.LabServices",
	},
	"Labs_Publish": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/publish",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Labs_Publish",
		Resource:    "Microsoft.LabServices",
	},
	"Labs_SyncGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/syncGroup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Labs_SyncGroup",
		Resource:    "Microsoft.LabServices",
	},
}
