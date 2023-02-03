package policy

// Microsoft_ApiManagement_apimusers policy
var Microsoft_ApiManagement_apimusers = map[string]Policy{
	"User_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "User_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"User_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "User_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"User_GenerateSsoUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/generateSsoUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "User_GenerateSsoUrl",
		Resource:    "Microsoft.ApiManagement",
	},
	"UserGroup_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/groups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "UserGroup_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"UserSubscription_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/subscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "UserSubscription_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"UserSubscription_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/subscriptions/{{.sid}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "UserSubscription_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"UserIdentities_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/identities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "UserIdentities_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"User_GetSharedAccessToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/token",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "User_GetSharedAccessToken",
		Resource:    "Microsoft.ApiManagement",
	},
	"UserConfirmationPassword_Send": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/users/{{.userId}}/confirmations/password/send",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "UserConfirmationPassword_Send",
		Resource:    "Microsoft.ApiManagement",
	},
}
