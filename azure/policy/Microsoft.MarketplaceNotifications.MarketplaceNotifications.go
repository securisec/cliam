package policy

// Microsoft_MarketplaceNotifications_MarketplaceNotifications policy
var Microsoft_MarketplaceNotifications_MarketplaceNotifications = map[string]Policy{
	"Notifications_List": {
		Path:   "/subscriptions/{{.subscription}}/providers/Microsoft.MarketplaceNotifications/reviewsNotifications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-03",
		},
		OperationID: "Notifications_List",
		Resource:    "Microsoft.MarketplaceNotifications",
	},
	"Notification_Get": {
		Path:   "/subscriptions/{{.subscription}}/providers/Microsoft.MarketplaceNotifications/reviewsNotification/{{.notification}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-03",
		},
		OperationID: "Notification_Get",
		Resource:    "Microsoft.MarketplaceNotifications",
	},
	"Notification_GetOperations": {
		Path:   "/providers/Microsoft.MarketplaceNotifications/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-03",
		},
		OperationID: "Notification_GetOperations",
		Resource:    "Microsoft.MarketplaceNotifications",
	},
}
