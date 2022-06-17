package policy

var Microsoft_DocumentDB_restorable = map[string]Policy{
	"RestorableDatabaseAccounts_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableDatabaseAccounts_ListByLocation",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableDatabaseAccounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableDatabaseAccounts_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableDatabaseAccounts_GetByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableDatabaseAccounts_GetByLocation",
		Resource:    "Microsoft.DocumentDB",
	},
	"SqlResources_RetrieveContinuousBackupInformation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/retrieveContinuousBackupInformation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_RetrieveContinuousBackupInformation",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableSqlDatabases_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}/restorableSqlDatabases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableSqlDatabases_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableSqlContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}/restorableSqlContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableSqlContainers_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableSqlResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}/restorableSqlResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableSqlResources_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"MongoDBResources_RetrieveContinuousBackupInformation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/collections/{{.collectionName}}/retrieveContinuousBackupInformation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_RetrieveContinuousBackupInformation",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableMongodbDatabases_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}/restorableMongodbDatabases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableMongodbDatabases_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableMongodbCollections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}/restorableMongodbCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableMongodbCollections_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"RestorableMongodbResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}/restorableDatabaseAccounts/{{.instanceId}}/restorableMongodbResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "RestorableMongodbResources_List",
		Resource:    "Microsoft.DocumentDB",
	},
}
