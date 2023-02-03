package policy

// Microsoft_AlertsManagement_AlertProcessingRules policy
var Microsoft_AlertsManagement_AlertProcessingRules = map[string]Policy{
	"AlertProcessingRules_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AlertsManagement/actionRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-08",
		},
		OperationID: "AlertProcessingRules_ListBySubscription",
		Resource:    "Microsoft.AlertsManagement",
	},
	"AlertProcessingRules_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AlertsManagement/actionRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-08",
		},
		OperationID: "AlertProcessingRules_ListByResourceGroup",
		Resource:    "Microsoft.AlertsManagement",
	},
	"AlertProcessingRules_GetByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AlertsManagement/actionRules/{{.alertProcessingRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-08",
		},
		OperationID: "AlertProcessingRules_GetByName",
		Resource:    "Microsoft.AlertsManagement",
	},
}
