package policy

// Microsoft_NotificationHubs_notificationhubs policy
var Microsoft_NotificationHubs_notificationhubs = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.NotificationHubs/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_CheckAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NotificationHubs/checkNamespaceAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_CheckAvailability",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_Get",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_GetAuthorizationRule",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_List",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NotificationHubs/namespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_ListAll",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/AuthorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_ListAuthorizationRules",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_ListKeys",
		Resource:    "Microsoft.NotificationHubs",
	},
	"Namespaces_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "Namespaces_RegenerateKeys",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_CheckNotificationHubAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/checkNotificationHubAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_CheckNotificationHubAvailability",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_Get",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_DebugSend": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}/debugsend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_DebugSend",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}/AuthorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_GetAuthorizationRule",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_List",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}/AuthorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_ListAuthorizationRules",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}/AuthorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_ListKeys",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}/AuthorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_RegenerateKeys",
		Resource:    "Microsoft.NotificationHubs",
	},
	"NotificationHubs_GetPnsCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NotificationHubs/namespaces/{{.namespaceName}}/notificationHubs/{{.notificationHubName}}/pnsCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-04-01",
		},
		OperationID: "NotificationHubs_GetPnsCredentials",
		Resource:    "Microsoft.NotificationHubs",
	},
}
