package policy

// Microsoft_Datadog_datadog policy
var Microsoft_Datadog_datadog = map[string]Policy{
	"MarketplaceAgreements_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Datadog/agreements",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "MarketplaceAgreements_List",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_ListApiKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/listApiKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_ListApiKeys",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_GetDefaultKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/getDefaultKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_GetDefaultKey",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_SetDefaultKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/setDefaultKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_SetDefaultKey",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_ListHosts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/listHosts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_ListHosts",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_ListLinkedResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/listLinkedResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_ListLinkedResources",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_ListMonitoredResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/listMonitoredResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_ListMonitoredResources",
		Resource:    "Microsoft.Datadog",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Datadog/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Datadog/monitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_List",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_ListByResourceGroup",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_Get",
		Resource:    "Microsoft.Datadog",
	},
	"Monitors_RefreshSetPasswordLink": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/refreshSetPasswordLink",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Monitors_RefreshSetPasswordLink",
		Resource:    "Microsoft.Datadog",
	},
	"TagRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/tagRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "TagRules_List",
		Resource:    "Microsoft.Datadog",
	},
	"TagRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/tagRules/{{.ruleSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "TagRules_Get",
		Resource:    "Microsoft.Datadog",
	},
	"SingleSignOnConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/singleSignOnConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "SingleSignOnConfigurations_List",
		Resource:    "Microsoft.Datadog",
	},
	"SingleSignOnConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Datadog/monitors/{{.monitorName}}/singleSignOnConfigurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "SingleSignOnConfigurations_Get",
		Resource:    "Microsoft.Datadog",
	},
}
