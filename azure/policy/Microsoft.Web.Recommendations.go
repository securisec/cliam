package policy

var Microsoft_Web_Recommendations = map[string]Policy{
	"Recommendations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/recommendations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_List",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ResetAllFilters": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/recommendations/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ResetAllFilters",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_DisableRecommendationForSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/recommendations/{{.name}}/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_DisableRecommendationForSubscription",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ListHistoryForHostingEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.hostingEnvironmentName}}/recommendationHistory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ListHistoryForHostingEnvironment",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ListRecommendedRulesForHostingEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.hostingEnvironmentName}}/recommendations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ListRecommendedRulesForHostingEnvironment",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_DisableAllForHostingEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.hostingEnvironmentName}}/recommendations/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_DisableAllForHostingEnvironment",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ResetAllFiltersForHostingEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.hostingEnvironmentName}}/recommendations/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ResetAllFiltersForHostingEnvironment",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_GetRuleDetailsByHostingEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.hostingEnvironmentName}}/recommendations/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_GetRuleDetailsByHostingEnvironment",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_DisableRecommendationForHostingEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/hostingEnvironments/{{.hostingEnvironmentName}}/recommendations/{{.name}}/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_DisableRecommendationForHostingEnvironment",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ListHistoryForWebApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/recommendationHistory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ListHistoryForWebApp",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ListRecommendedRulesForWebApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/recommendations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ListRecommendedRulesForWebApp",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_DisableAllForWebApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/recommendations/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_DisableAllForWebApp",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_ResetAllFiltersForWebApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/recommendations/reset",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_ResetAllFiltersForWebApp",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_GetRuleDetailsByWebApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/recommendations/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_GetRuleDetailsByWebApp",
		Resource:    "Microsoft.Web",
	},
	"Recommendations_DisableRecommendationForSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.siteName}}/recommendations/{{.name}}/disable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Recommendations_DisableRecommendationForSite",
		Resource:    "Microsoft.Web",
	},
}
