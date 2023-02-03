package policy

// Microsoft_LabServices_Users policy
var Microsoft_LabServices_Users = map[string]Policy{
	"Users_ListByLab": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/users",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Users_ListByLab",
		Resource:    "Microsoft.LabServices",
	},
	"Users_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/users/{{.userName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Users_Get",
		Resource:    "Microsoft.LabServices",
	},
	"Users_Invite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/users/{{.userName}}/invite",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Users_Invite",
		Resource:    "Microsoft.LabServices",
	},
}
