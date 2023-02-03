package policy

// Microsoft_ExtendedLocation_customlocations policy
var Microsoft_ExtendedLocation_customlocations = map[string]Policy{
	"CustomLocations_ListOperations": {
		Path:   "/providers/Microsoft.ExtendedLocation/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-15",
		},
		OperationID: "CustomLocations_ListOperations",
		Resource:    "Microsoft.ExtendedLocation",
	},
	"CustomLocations_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ExtendedLocation/customLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-15",
		},
		OperationID: "CustomLocations_ListBySubscription",
		Resource:    "Microsoft.ExtendedLocation",
	},
	"CustomLocations_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ExtendedLocation/customLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-15",
		},
		OperationID: "CustomLocations_ListByResourceGroup",
		Resource:    "Microsoft.ExtendedLocation",
	},
	"CustomLocations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ExtendedLocation/customLocations/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-15",
		},
		OperationID: "CustomLocations_Get",
		Resource:    "Microsoft.ExtendedLocation",
	},
	"CustomLocations_ListEnabledResourceTypes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ExtendedLocation/customLocations/{{.resourceName}}/enabledResourceTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-15",
		},
		OperationID: "CustomLocations_ListEnabledResourceTypes",
		Resource:    "Microsoft.ExtendedLocation",
	},
}
