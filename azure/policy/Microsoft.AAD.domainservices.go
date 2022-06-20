package policy

var Microsoft_AAD_domainservices = map[string]Policy{
	"DomainServiceOperations_List": {
		Path:   "/providers/Microsoft.AAD/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "DomainServiceOperations_List",
		Resource:    "Microsoft.AAD",
	},
	"DomainServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AAD/domainServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "DomainServices_List",
		Resource:    "Microsoft.AAD",
	},
	"DomainServices_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AAD/domainServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "DomainServices_ListByResourceGroup",
		Resource:    "Microsoft.AAD",
	},
	"DomainServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AAD/domainServices/{{.domainServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "DomainServices_Get",
		Resource:    "Microsoft.AAD",
	},
}
