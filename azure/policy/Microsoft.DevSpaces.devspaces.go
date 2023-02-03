package policy

// Microsoft_DevSpaces_devspaces policy
var Microsoft_DevSpaces_devspaces = map[string]Policy{
	"ContainerHostMappings_GetContainerHostMapping": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/locations/{{.location}}/checkContainerHostMapping",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "ContainerHostMappings_GetContainerHostMapping",
		Resource:    "Microsoft.DevSpaces",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DevSpaces/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DevSpaces",
	},
	"Controllers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/controllers/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Controllers_Get",
		Resource:    "Microsoft.DevSpaces",
	},
	"Controllers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/controllers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Controllers_ListByResourceGroup",
		Resource:    "Microsoft.DevSpaces",
	},
	"Controllers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DevSpaces/controllers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Controllers_List",
		Resource:    "Microsoft.DevSpaces",
	},
	"Controllers_ListConnectionDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DevSpaces/controllers/{{.name}}/listConnectionDetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-04-01",
		},
		OperationID: "Controllers_ListConnectionDetails",
		Resource:    "Microsoft.DevSpaces",
	},
}
