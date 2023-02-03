package policy

// Microsoft_ApiManagement_apimsubscriptions policy
var Microsoft_ApiManagement_apimsubscriptions = map[string]Policy{
	"Subscription_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/subscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Subscription_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"Subscription_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/subscriptions/{{.sid}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Subscription_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"Subscription_RegeneratePrimaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/subscriptions/{{.sid}}/regeneratePrimaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Subscription_RegeneratePrimaryKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"Subscription_RegenerateSecondaryKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/subscriptions/{{.sid}}/regenerateSecondaryKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Subscription_RegenerateSecondaryKey",
		Resource:    "Microsoft.ApiManagement",
	},
	"Subscription_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/subscriptions/{{.sid}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Subscription_ListSecrets",
		Resource:    "Microsoft.ApiManagement",
	},
}
