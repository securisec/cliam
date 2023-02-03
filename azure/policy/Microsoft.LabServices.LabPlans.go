package policy

// Microsoft_LabServices_LabPlans policy
var Microsoft_LabServices_LabPlans = map[string]Policy{
	"LabPlans_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LabServices/labPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "LabPlans_ListBySubscription",
		Resource:    "Microsoft.LabServices",
	},
	"LabPlans_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "LabPlans_ListByResourceGroup",
		Resource:    "Microsoft.LabServices",
	},
	"LabPlans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labPlans/{{.labPlanName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "LabPlans_Get",
		Resource:    "Microsoft.LabServices",
	},
	"LabPlans_SaveImage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labPlans/{{.labPlanName}}/saveImage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "LabPlans_SaveImage",
		Resource:    "Microsoft.LabServices",
	},
}
