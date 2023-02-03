package policy

// MicrosoftMixedRealityspatialanchors policy
var MicrosoftMixedRealityspatialanchors = map[string]Policy{
	"SpatialAnchorsAccounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MixedReality/spatialAnchorsAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "SpatialAnchorsAccounts_ListBySubscription",
		Resource:    "Microsoft.MixedReality",
	},
	"SpatialAnchorsAccounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/spatialAnchorsAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "SpatialAnchorsAccounts_ListByResourceGroup",
		Resource:    "Microsoft.MixedReality",
	},
	"SpatialAnchorsAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "SpatialAnchorsAccounts_Get",
		Resource:    "Microsoft.MixedReality",
	},
	"SpatialAnchorsAccounts_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "SpatialAnchorsAccounts_ListKeys",
		Resource:    "Microsoft.MixedReality",
	},
	"SpatialAnchorsAccounts_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{{.accountName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "SpatialAnchorsAccounts_RegenerateKeys",
		Resource:    "Microsoft.MixedReality",
	},
}
