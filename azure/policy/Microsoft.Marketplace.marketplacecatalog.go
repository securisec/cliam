package policy

// Microsoft_Marketplace_marketplacecatalog policy
var Microsoft_Marketplace_marketplacecatalog = map[string]Policy{
	"EdgeZonesProducts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Marketplace/locations/{{.location}}/edgeZones/{{.edgeZone}}/products",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-31",
		},
		OperationID: "EdgeZonesProducts_List",
		Resource:    "Microsoft.Marketplace",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Marketplace/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-31",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Marketplace",
	},
}
