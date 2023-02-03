package policy

// Microsoft_Advisor_advisor policy
var Microsoft_Advisor_advisor = map[string]Policy{
	"RecommendationMetadata_Get": {
		Path:   "/providers/Microsoft.Advisor/metadata/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "RecommendationMetadata_Get",
		Resource:    "Microsoft.Advisor",
	},
	"RecommendationMetadata_List": {
		Path:   "/providers/Microsoft.Advisor/metadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "RecommendationMetadata_List",
		Resource:    "Microsoft.Advisor",
	},
	"Configurations_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Advisor/configurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Configurations_ListBySubscription",
		Resource:    "Microsoft.Advisor",
	},
	"Configurations_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.Advisor/configurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Configurations_ListByResourceGroup",
		Resource:    "Microsoft.Advisor",
	},
	"Recommendations_Generate": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Advisor/generateRecommendations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Recommendations_Generate",
		Resource:    "Microsoft.Advisor",
	},
	"Recommendations_GetGenerateStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Advisor/generateRecommendations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Recommendations_GetGenerateStatus",
		Resource:    "Microsoft.Advisor",
	},
	"Recommendations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Advisor/recommendations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Recommendations_List",
		Resource:    "Microsoft.Advisor",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Advisor/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Advisor",
	},
	"Recommendations_Get": {
		Path:   "/{{.resourceUri}}/providers/Microsoft.Advisor/recommendations/{{.recommendationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Recommendations_Get",
		Resource:    "Microsoft.Advisor",
	},
	"Suppressions_Get": {
		Path:   "/{{.resourceUri}}/providers/Microsoft.Advisor/recommendations/{{.recommendationId}}/suppressions/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Suppressions_Get",
		Resource:    "Microsoft.Advisor",
	},
	"Suppressions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Advisor/suppressions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Suppressions_List",
		Resource:    "Microsoft.Advisor",
	},
	"Predict": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Advisor/predict",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Predict",
		Resource:    "Microsoft.Advisor",
	},
}
