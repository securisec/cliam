package policy

// Microsoft_Web_Certificates policy
var Microsoft_Web_Certificates = map[string]Policy{
	"Certificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Certificates_List",
		Resource:    "Microsoft.Web",
	},
	"Certificates_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Certificates_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"Certificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/certificates/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Certificates_Get",
		Resource:    "Microsoft.Web",
	},
}
