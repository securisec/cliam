package policy

var Microsoft_DomainRegistration_TopLevelDomains = map[string]Policy{
	"TopLevelDomains_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/topLevelDomains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "TopLevelDomains_List",
		Resource:    "Microsoft.DomainRegistration",
	},
	"TopLevelDomains_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/topLevelDomains/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "TopLevelDomains_Get",
		Resource:    "Microsoft.DomainRegistration",
	},
	"TopLevelDomains_ListAgreements": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/topLevelDomains/{{.name}}/listAgreements",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "TopLevelDomains_ListAgreements",
		Resource:    "Microsoft.DomainRegistration",
	},
}
