package policy

var Microsoft_ApiManagement_apimgroups = map[string]Policy{
	"Group_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/groups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Group_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Group_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/groups/{{.groupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Group_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"GroupUser_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/groups/{{.groupId}}/users",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GroupUser_List",
		Resource:    "Microsoft.ApiManagement",
	},
}
