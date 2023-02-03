package policy

// Microsoft_PowerBIdedicated_powerbidedicated policy
var Microsoft_PowerBIdedicated_powerbidedicated = map[string]Policy{
	"Capacities_GetDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/capacities/{{.dedicatedCapacityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_GetDetails",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_Suspend": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/capacities/{{.dedicatedCapacityName}}/suspend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_Suspend",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/capacities/{{.dedicatedCapacityName}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_Resume",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/capacities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_ListByResourceGroup",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBIDedicated/capacities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_List",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBIDedicated/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_ListSkus",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_ListSkusForCapacity": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBIDedicated/capacities/{{.dedicatedCapacityName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_ListSkusForCapacity",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.PowerBIDedicated/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.PowerBIdedicated",
	},
	"Capacities_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBIDedicated/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Capacities_CheckNameAvailability",
		Resource:    "Microsoft.PowerBIdedicated",
	},
}
