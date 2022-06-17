package policy

var Microsoft_DomainRegistration_Domains = map[string]Policy{
	"Domains_CheckAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/checkDomainAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_CheckAvailability",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/domains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_List",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_GetControlCenterSsoRequest": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/generateSsoRequest",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_GetControlCenterSsoRequest",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_ListRecommendations": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DomainRegistration/listDomainRecommendations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_ListRecommendations",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DomainRegistration/domains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_ListByResourceGroup",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DomainRegistration/domains/{{.domainName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_Get",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_ListOwnershipIdentifiers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DomainRegistration/domains/{{.domainName}}/domainOwnershipIdentifiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_ListOwnershipIdentifiers",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_GetOwnershipIdentifier": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DomainRegistration/domains/{{.domainName}}/domainOwnershipIdentifiers/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_GetOwnershipIdentifier",
		Resource:    "Microsoft.DomainRegistration",
	},
	"Domains_Renew": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DomainRegistration/domains/{{.domainName}}/renew",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Domains_Renew",
		Resource:    "Microsoft.DomainRegistration",
	},
}
