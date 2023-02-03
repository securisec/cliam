package policy

// Microsoft_Network_ddosProtectionPlan policy
var Microsoft_Network_ddosProtectionPlan = map[string]Policy{
	"DdosProtectionPlans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ddosProtectionPlans/{{.ddosProtectionPlanName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DdosProtectionPlans_Get",
		Resource:    "Microsoft.Network",
	},
	"DdosProtectionPlans_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ddosProtectionPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DdosProtectionPlans_List",
		Resource:    "Microsoft.Network",
	},
	"DdosProtectionPlans_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ddosProtectionPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DdosProtectionPlans_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
}
