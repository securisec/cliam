package policy

// Microsoft_PolicyInsights_attestations policy
var Microsoft_PolicyInsights_attestations = map[string]Policy{
	"Attestations_ListForSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PolicyInsights/attestations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Attestations_ListForSubscription",
		Resource:    "Microsoft.PolicyInsights",
	},
	"Attestations_GetAtSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PolicyInsights/attestations/{{.attestationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Attestations_GetAtSubscription",
		Resource:    "Microsoft.PolicyInsights",
	},
	"Attestations_ListForResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PolicyInsights/attestations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Attestations_ListForResourceGroup",
		Resource:    "Microsoft.PolicyInsights",
	},
	"Attestations_GetAtResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PolicyInsights/attestations/{{.attestationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Attestations_GetAtResourceGroup",
		Resource:    "Microsoft.PolicyInsights",
	},
	"Attestations_ListForResource": {
		Path:   "/{{.resourceId}}/providers/Microsoft.PolicyInsights/attestations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Attestations_ListForResource",
		Resource:    "Microsoft.PolicyInsights",
	},
	"Attestations_GetAtResource": {
		Path:   "/{{.resourceId}}/providers/Microsoft.PolicyInsights/attestations/{{.attestationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Attestations_GetAtResource",
		Resource:    "Microsoft.PolicyInsights",
	},
}
