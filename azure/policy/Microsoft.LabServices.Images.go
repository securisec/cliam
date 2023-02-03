package policy

// Microsoft_LabServices_Images policy
var Microsoft_LabServices_Images = map[string]Policy{
	"Images_ListByLabPlan": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labPlans/{{.labPlanName}}/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Images_ListByLabPlan",
		Resource:    "Microsoft.LabServices",
	},
	"Images_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labPlans/{{.labPlanName}}/images/{{.imageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Images_Get",
		Resource:    "Microsoft.LabServices",
	},
}
