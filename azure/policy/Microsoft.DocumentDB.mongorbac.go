package policy

// Microsoft_DocumentDB_mongorbac policy
var Microsoft_DocumentDB_mongorbac = map[string]Policy{
	"MongoDBResources_GetMongoRoleDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbRoleDefinitions/{{.mongoRoleDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "MongoDBResources_GetMongoRoleDefinition",
		Resource:    "Microsoft.DocumentDB",
	},
	"MongoDBResources_ListMongoRoleDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbRoleDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "MongoDBResources_ListMongoRoleDefinitions",
		Resource:    "Microsoft.DocumentDB",
	},
	"MongoDBResources_GetMongoUserDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbUserDefinitions/{{.mongoUserDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "MongoDBResources_GetMongoUserDefinition",
		Resource:    "Microsoft.DocumentDB",
	},
	"MongoDBResources_ListMongoUserDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbUserDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "MongoDBResources_ListMongoUserDefinitions",
		Resource:    "Microsoft.DocumentDB",
	},
}
