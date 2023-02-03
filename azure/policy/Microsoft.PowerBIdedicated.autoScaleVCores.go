package policy

// Microsoft_PowerBIdedicated_autoScaleVCores policy
var Microsoft_PowerBIdedicated_autoScaleVCores = map[string]Policy{
	"AutoScaleVCores_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/autoScaleVCores/{{.vcoreName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "AutoScaleVCores_Get",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"AutoScaleVCores_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/autoScaleVCores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "AutoScaleVCores_ListByResourceGroup",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"AutoScaleVCores_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBIDedicated/autoScaleVCores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "AutoScaleVCores_ListBySubscription",
		Resource:    "Microsoft.PowerBIdedicated",
	},
}
