package policy

// Microsoft_Storage_queue policy
var Microsoft_Storage_queue = map[string]Policy{
	"QueueServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/queueServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "QueueServices_List",
		Resource:    "Microsoft.Storage",
	},
	"QueueServices_GetServiceProperties": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/queueServices/{{.queueServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "QueueServices_GetServiceProperties",
		Resource:    "Microsoft.Storage",
	},
	"Queue_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/queueServices/default/queues/{{.queueName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Queue_Get",
		Resource:    "Microsoft.Storage",
	},
	"Queue_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/queueServices/default/queues",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Queue_List",
		Resource:    "Microsoft.Storage",
	},
}
