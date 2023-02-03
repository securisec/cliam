package policy

// Microsoft_Insights_autoscale_API policy
var Microsoft_Insights_autoscale_API = map[string]Policy{
	"AutoscaleSettings_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Insights/autoscalesettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "AutoscaleSettings_ListByResourceGroup",
		Resource:    "Microsoft.Insights",
	},
	"AutoscaleSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Insights/autoscalesettings/{{.autoscaleSettingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "AutoscaleSettings_Get",
		Resource:    "Microsoft.Insights",
	},
	"PredictiveMetric_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Insights/autoscalesettings/{{.autoscaleSettingName}}/predictiveMetrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PredictiveMetric_Get",
		Resource:    "Microsoft.Insights",
	},
	"AutoscaleSettings_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Insights/autoscalesettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "AutoscaleSettings_ListBySubscription",
		Resource:    "Microsoft.Insights",
	},
}
