package policy

// Microsoft_MarketplaceOrdering_Agreements policy
var Microsoft_MarketplaceOrdering_Agreements = map[string]Policy{
	"MarketplaceAgreements_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MarketplaceOrdering/offerTypes/{{.offerType}}/publishers/{{.publisherId}}/offers/{{.offerId}}/plans/{{.planId}}/agreements/current",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "MarketplaceAgreements_Get",
		Resource:    "Microsoft.MarketplaceOrdering",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.MarketplaceOrdering/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.MarketplaceOrdering",
	},
	"MarketplaceAgreements_Sign": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MarketplaceOrdering/agreements/{{.publisherId}}/offers/{{.offerId}}/plans/{{.planId}}/sign",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "MarketplaceAgreements_Sign",
		Resource:    "Microsoft.MarketplaceOrdering",
	},
	"MarketplaceAgreements_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MarketplaceOrdering/agreements/{{.publisherId}}/offers/{{.offerId}}/plans/{{.planId}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "MarketplaceAgreements_Cancel",
		Resource:    "Microsoft.MarketplaceOrdering",
	},
	"MarketplaceAgreements_GetAgreement": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MarketplaceOrdering/agreements/{{.publisherId}}/offers/{{.offerId}}/plans/{{.planId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "MarketplaceAgreements_GetAgreement",
		Resource:    "Microsoft.MarketplaceOrdering",
	},
	"MarketplaceAgreements_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MarketplaceOrdering/agreements",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-01-01",
		},
		OperationID: "MarketplaceAgreements_List",
		Resource:    "Microsoft.MarketplaceOrdering",
	},
}
