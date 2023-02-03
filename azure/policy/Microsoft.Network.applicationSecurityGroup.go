package policy

// Microsoft_Network_applicationSecurityGroup policy
var Microsoft_Network_applicationSecurityGroup = map[string]Policy{
	"ApplicationSecurityGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationSecurityGroups/{{.applicationSecurityGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationSecurityGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"ApplicationSecurityGroups_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/applicationSecurityGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationSecurityGroups_ListAll",
		Resource:    "Microsoft.Network",
	},
	"ApplicationSecurityGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/applicationSecurityGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationSecurityGroups_List",
		Resource:    "Microsoft.Network",
	},
}
