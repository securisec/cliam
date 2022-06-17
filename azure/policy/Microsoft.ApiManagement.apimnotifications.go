package policy

var Microsoft_ApiManagement_apimnotifications = map[string]Policy{
	"Notification_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/notifications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Notification_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Notification_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/notifications/{{.notificationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Notification_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"NotificationRecipientUser_ListByNotification": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/notifications/{{.notificationName}}/recipientUsers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NotificationRecipientUser_ListByNotification",
		Resource:    "Microsoft.ApiManagement",
	},
	"NotificationRecipientEmail_ListByNotification": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/notifications/{{.notificationName}}/recipientEmails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "NotificationRecipientEmail_ListByNotification",
		Resource:    "Microsoft.ApiManagement",
	},
}
