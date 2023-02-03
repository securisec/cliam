package policy

// Microsoft_Marketplace_Marketplace policy
var Microsoft_Marketplace_Marketplace = map[string]Policy{
	"PrivateStore_List": {
		Path:   "/providers/Microsoft.Marketplace/privateStores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_List",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_Get": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_Get",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_AnyExistingOffersInTheCollections": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/anyExistingOffersInTheCollections",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_AnyExistingOffersInTheCollections",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_QueryOffers": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/queryOffers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_QueryOffers",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_QueryUserOffers": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/queryUserOffers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_QueryUserOffers",
		Resource:    "Microsoft.Marketplace",
	},
	"QueryUserRules": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/queryUserRules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "QueryUserRules",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_BillingAccounts": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/billingAccounts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_BillingAccounts",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_CollectionsToSubscriptionsMapping": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collectionsToSubscriptionsMapping",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_CollectionsToSubscriptionsMapping",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollection_List": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollection_List",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollection_Get": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollection_Get",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollection_Post": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollection_Post",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollection_TransferOffers": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/transferOffers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollection_TransferOffers",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_QueryApprovedPlans": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/queryApprovedPlans",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_QueryApprovedPlans",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_BulkCollectionsAction": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/bulkCollectionsAction",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_BulkCollectionsAction",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollection_ApproveAllItems": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/approveAllItems",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollection_ApproveAllItems",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollection_DisableApproveAllItems": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/disableApproveAllItems",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollection_DisableApproveAllItems",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollectionOffer_List": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollectionOffer_List",
		Resource:    "Microsoft.Marketplace",
	},
	"SetCollectionRules": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/setRules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "SetCollectionRules",
		Resource:    "Microsoft.Marketplace",
	},
	"QueryRules": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/queryRules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "QueryRules",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollectionOffer_Get": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/offers/{{.offerId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollectionOffer_Get",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollectionOffer_Post": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/offers/{{.offerId}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollectionOffer_Post",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStoreCollectionOffer_UpsertOfferWithMultiContext": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/collections/{{.collectionId}}/offers/{{.offerId}}/upsertOfferWithMultiContext",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStoreCollectionOffer_UpsertOfferWithMultiContext",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_GetApprovalRequestsList": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/requestApprovals",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_GetApprovalRequestsList",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_GetRequestApproval": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/requestApprovals/{{.requestApprovalId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_GetRequestApproval",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_QueryRequestApproval": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/requestApprovals/{{.requestApprovalId}}/query",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_QueryRequestApproval",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_AdminRequestApprovalsList": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/adminRequestApprovals",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_AdminRequestApprovalsList",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_GetAdminRequestApproval": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/adminRequestApprovals/{{.adminRequestApprovalId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_GetAdminRequestApproval",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_QueryNotificationsState": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/queryNotificationsState",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_QueryNotificationsState",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_AcknowledgeOfferNotification": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/offers/{{.offerId}}/acknowledgeNotification",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_AcknowledgeOfferNotification",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_WithdrawPlan": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/requestApprovals/{{.requestApprovalId}}/withdrawPlan",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_WithdrawPlan",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_FetchAllSubscriptionsInTenant": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/fetchAllSubscriptionsInTenant",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_FetchAllSubscriptionsInTenant",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_ListNewPlansNotifications": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/listNewPlansNotifications",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_ListNewPlansNotifications",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_ListStopSellOffersPlansNotifications": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/listStopSellOffersPlansNotifications",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_ListStopSellOffersPlansNotifications",
		Resource:    "Microsoft.Marketplace",
	},
	"PrivateStore_ListSubscriptionsContext": {
		Path:   "/providers/Microsoft.Marketplace/privateStores/{{.privateStoreId}}/listSubscriptionsContext",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "PrivateStore_ListSubscriptionsContext",
		Resource:    "Microsoft.Marketplace",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Marketplace/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Marketplace",
	},
}
