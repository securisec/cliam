package policy

// Dynatrace_Observability_dynatrace policy
var Dynatrace_Observability_dynatrace = map[string]Policy{
	"Monitors_GetAccountCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/getAccountCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_GetAccountCredentials",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_ListMonitoredResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/listMonitoredResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_ListMonitoredResources",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_GetVMHostPayload": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/getVMHostPayload",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_GetVMHostPayload",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_Get",
		Resource:    "Dynatrace.Observability",
	},
	"Operations_List": {
		Path:   "/providers/Dynatrace.Observability/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_ListBySubscriptionId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Dynatrace.Observability/monitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_ListBySubscriptionId",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_ListByResourceGroup",
		Resource:    "Dynatrace.Observability",
	},
	"TagRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/tagRules/{{.ruleSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "TagRules_Get",
		Resource:    "Dynatrace.Observability",
	},
	"TagRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/tagRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "TagRules_List",
		Resource:    "Dynatrace.Observability",
	},
	"SingleSignOn_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/singleSignOnConfigurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "SingleSignOn_Get",
		Resource:    "Dynatrace.Observability",
	},
	"SingleSignOn_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/singleSignOnConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "SingleSignOn_List",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_ListHosts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/listHosts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_ListHosts",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_ListAppServices": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/listAppServices",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_ListAppServices",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_GetSSODetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/getSSODetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_GetSSODetails",
		Resource:    "Dynatrace.Observability",
	},
	"Monitors_ListLinkableEnvironments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Dynatrace.Observability/monitors/{{.monitorName}}/listLinkableEnvironments",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Monitors_ListLinkableEnvironments",
		Resource:    "Dynatrace.Observability",
	},
}
