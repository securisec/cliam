package policy

var Microsoft_Web_ResourceHealthMetadata = map[string]Policy{
	"ResourceHealthMetadata_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/resourceHealthMetadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ResourceHealthMetadata_List",
		Resource:    "Microsoft.Web",
	},
	"ResourceHealthMetadata_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/resourceHealthMetadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ResourceHealthMetadata_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"ResourceHealthMetadata_ListBySite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/resourceHealthMetadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ResourceHealthMetadata_ListBySite",
		Resource:    "Microsoft.Web",
	},
	"ResourceHealthMetadata_GetBySite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/resourceHealthMetadata/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ResourceHealthMetadata_GetBySite",
		Resource:    "Microsoft.Web",
	},
	"ResourceHealthMetadata_ListBySiteSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/resourceHealthMetadata",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ResourceHealthMetadata_ListBySiteSlot",
		Resource:    "Microsoft.Web",
	},
	"ResourceHealthMetadata_GetBySiteSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/resourceHealthMetadata/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ResourceHealthMetadata_GetBySiteSlot",
		Resource:    "Microsoft.Web",
	},
}
