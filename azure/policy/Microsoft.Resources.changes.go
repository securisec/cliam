package policy

var Microsoft_Resources_changes = map[string]Policy{
	"Changes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Resources/changes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Changes_List",
		Resource:    "Microsoft.Resources",
	},
	"Changes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Resources/changes/{{.changeResourceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Changes_Get",
		Resource:    "Microsoft.Resources",
	},
}
