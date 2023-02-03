package policy

// Microsoft_IoTCentral_iotcentral policy
var Microsoft_IoTCentral_iotcentral = map[string]Policy{
	"Apps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.IoTCentral/iotApps/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Apps_Get",
		Resource:    "Microsoft.IoTCentral",
	},
	"Apps_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.IoTCentral/iotApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Apps_ListBySubscription",
		Resource:    "Microsoft.IoTCentral",
	},
	"Apps_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.IoTCentral/iotApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Apps_ListByResourceGroup",
		Resource:    "Microsoft.IoTCentral",
	},
	"Apps_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.IoTCentral/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Apps_CheckNameAvailability",
		Resource:    "Microsoft.IoTCentral",
	},
	"Apps_CheckSubdomainAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.IoTCentral/checkSubdomainAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Apps_CheckSubdomainAvailability",
		Resource:    "Microsoft.IoTCentral",
	},
	"Apps_ListTemplates": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.IoTCentral/appTemplates",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Apps_ListTemplates",
		Resource:    "Microsoft.IoTCentral",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.IoTCentral/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.IoTCentral",
	},
}
