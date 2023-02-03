package policy

// Microsoft_TimeSeriesInsights_timeseriesinsights policy
var Microsoft_TimeSeriesInsights_timeseriesinsights = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.TimeSeriesInsights/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"Environments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "Environments_Get",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"Environments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "Environments_ListByResourceGroup",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"Environments_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.TimeSeriesInsights/environments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "Environments_ListBySubscription",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"EventSources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}/eventSources/{{.eventSourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "EventSources_Get",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"EventSources_ListByEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}/eventSources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "EventSources_ListByEnvironment",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"ReferenceDataSets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}/referenceDataSets/{{.referenceDataSetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "ReferenceDataSets_Get",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"ReferenceDataSets_ListByEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}/referenceDataSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "ReferenceDataSets_ListByEnvironment",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"AccessPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}/accessPolicies/{{.accessPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "AccessPolicies_Get",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
	"AccessPolicies_ListByEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.TimeSeriesInsights/environments/{{.environmentName}}/accessPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-15",
		},
		OperationID: "AccessPolicies_ListByEnvironment",
		Resource:    "Microsoft.TimeSeriesInsights",
	},
}
