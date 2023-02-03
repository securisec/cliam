package policy

// Microsoft_LabServices_Schedules policy
var Microsoft_LabServices_Schedules = map[string]Policy{
	"Schedules_ListByLab": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Schedules_ListByLab",
		Resource:    "Microsoft.LabServices",
	},
	"Schedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/schedules/{{.scheduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Schedules_Get",
		Resource:    "Microsoft.LabServices",
	},
}
