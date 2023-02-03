package policy

// Microsoft_CostManagement_costmanagement policy
var Microsoft_CostManagement_costmanagement = map[string]Policy{
	"Views_List": {
		Path:   "/providers/Microsoft.CostManagement/views",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Views_List",
		Resource:    "Microsoft.CostManagement",
	},
	"Views_ListByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/views",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Views_ListByScope",
		Resource:    "Microsoft.CostManagement",
	},
	"Views_Get": {
		Path:   "/providers/Microsoft.CostManagement/views/{{.viewName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Views_Get",
		Resource:    "Microsoft.CostManagement",
	},
	"Views_GetByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/views/{{.viewName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Views_GetByScope",
		Resource:    "Microsoft.CostManagement",
	},
	"Alerts_List": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Alerts_List",
		Resource:    "Microsoft.CostManagement",
	},
	"Alerts_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/alerts/{{.alertId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Alerts_Get",
		Resource:    "Microsoft.CostManagement",
	},
	"Alerts_ListExternal": {
		Path:   "/providers/Microsoft.CostManagement/{{.externalCloudProviderType}}/{{.externalCloudProviderId}}/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Alerts_ListExternal",
		Resource:    "Microsoft.CostManagement",
	},
	"Forecast_Usage": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/forecast",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Forecast_Usage",
		Resource:    "Microsoft.CostManagement",
	},
	"Forecast_ExternalCloudProviderUsage": {
		Path:   "/providers/Microsoft.CostManagement/{{.externalCloudProviderType}}/{{.externalCloudProviderId}}/forecast",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Forecast_ExternalCloudProviderUsage",
		Resource:    "Microsoft.CostManagement",
	},
	"Dimensions_List": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/dimensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Dimensions_List",
		Resource:    "Microsoft.CostManagement",
	},
	"Dimensions_ByExternalCloudProviderType": {
		Path:   "/providers/Microsoft.CostManagement/{{.externalCloudProviderType}}/{{.externalCloudProviderId}}/dimensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Dimensions_ByExternalCloudProviderType",
		Resource:    "Microsoft.CostManagement",
	},
	"Query_Usage": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/query",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Query_Usage",
		Resource:    "Microsoft.CostManagement",
	},
	"Query_UsageByExternalCloudProviderType": {
		Path:   "/providers/Microsoft.CostManagement/{{.externalCloudProviderType}}/{{.externalCloudProviderId}}/query",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Query_UsageByExternalCloudProviderType",
		Resource:    "Microsoft.CostManagement",
	},
	"GenerateReservationDetailsReport_ByBillingAccountId": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.CostManagement/generateReservationDetailsReport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateReservationDetailsReport_ByBillingAccountId",
		Resource:    "Microsoft.CostManagement",
	},
	"GenerateReservationDetailsReport_ByBillingProfileId": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingProfiles/{{.billingProfileId}}/providers/Microsoft.CostManagement/generateReservationDetailsReport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateReservationDetailsReport_ByBillingProfileId",
		Resource:    "Microsoft.CostManagement",
	},
}
