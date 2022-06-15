package policy

var Microsoft_DocumentDB_cosmos_db = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_Get",
		Resource:    "Microsoft.DocumentDB",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/failoverPriorityChange",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_FailoverPriorityChange",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/databaseAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_List",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListByResourceGroup",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListKeys",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/listConnectionStrings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListConnectionStrings",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/offlineRegion",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_OfflineRegion",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/onlineRegion",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_OnlineRegion",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/readonlykeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_GetReadOnlyKeys",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/readonlykeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListReadOnlyKeys",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_RegenerateKey",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/providers/Microsoft.DocumentDB/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Database_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Collection_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/region/{{.region}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CollectionRegion_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/region/{{.region}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccountRegion_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sourceRegion/{{.sourceRegion}}/targetRegion/{{.targetRegion}}/percentile/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PercentileSourceTarget_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/targetRegion/{{.targetRegion}}/percentile/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PercentileTarget_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/percentile/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Percentile_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/region/{{.region}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/partitions/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CollectionPartitionRegion_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/partitions/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CollectionPartition_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/partitionKeyRangeId/{{.partitionKeyRangeId}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PartitionKeyRangeId_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/region/{{.region}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/partitionKeyRangeId/{{.partitionKeyRangeId}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PartitionKeyRangeIdRegion_ListMetrics",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListUsages",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Database_ListUsages",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Collection_ListUsages",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/partitions/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CollectionPartition_ListUsages",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/metricDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Database_ListMetricDefinitions",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/databases/{{.databaseRid}}/collections/{{.collectionRid}}/metricDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Collection_ListMetricDefinitions",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/metricDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "DatabaseAccounts_ListMetricDefinitions",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlDatabases",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlDatabase",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlDatabaseThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_MigrateSqlDatabaseToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_MigrateSqlDatabaseToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlContainers",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlContainer",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlContainerThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_MigrateSqlContainerToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_MigrateSqlContainerToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/storedProcedures",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlStoredProcedures",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/storedProcedures/{{.storedProcedureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlStoredProcedure",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/userDefinedFunctions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlUserDefinedFunctions",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/userDefinedFunctions/{{.userDefinedFunctionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlUserDefinedFunction",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/triggers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlTriggers",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlDatabases/{{.databaseName}}/containers/{{.containerName}}/triggers/{{.triggerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlTrigger",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_ListMongoDBDatabases",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_GetMongoDBDatabase",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_GetMongoDBDatabaseThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_MigrateMongoDBDatabaseToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_MigrateMongoDBDatabaseToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/collections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_ListMongoDBCollections",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/collections/{{.collectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_GetMongoDBCollection",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/collections/{{.collectionName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_GetMongoDBCollectionThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/collections/{{.collectionName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_MigrateMongoDBCollectionToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/mongodbDatabases/{{.databaseName}}/collections/{{.collectionName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "MongoDBResources_MigrateMongoDBCollectionToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/tables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "TableResources_ListTables",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/tables/{{.tableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "TableResources_GetTable",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/tables/{{.tableName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "TableResources_GetTableThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/tables/{{.tableName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "TableResources_MigrateTableToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/tables/{{.tableName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "TableResources_MigrateTableToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_ListCassandraKeyspaces",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_GetCassandraKeyspace",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_GetCassandraKeyspaceThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_MigrateCassandraKeyspaceToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_MigrateCassandraKeyspaceToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/tables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_ListCassandraTables",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/tables/{{.tableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_GetCassandraTable",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/tables/{{.tableName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_GetCassandraTableThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/tables/{{.tableName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_MigrateCassandraTableToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/cassandraKeyspaces/{{.keyspaceName}}/tables/{{.tableName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraResources_MigrateCassandraTableToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_ListGremlinDatabases",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_GetGremlinDatabase",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_GetGremlinDatabaseThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_MigrateGremlinDatabaseToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_MigrateGremlinDatabaseToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/graphs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_ListGremlinGraphs",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/graphs/{{.graphName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_GetGremlinGraph",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/graphs/{{.graphName}}/throughputSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_GetGremlinGraphThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/graphs/{{.graphName}}/throughputSettings/default/migrateToAutoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_MigrateGremlinGraphToAutoscale",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/gremlinDatabases/{{.databaseName}}/graphs/{{.graphName}}/throughputSettings/default/migrateToManualThroughput",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "GremlinResources_MigrateGremlinGraphToManualThroughput",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Locations_List",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/locations/{{.location}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "Locations_Get",
		Resource:    "Microsoft.DocumentDB",
	},
}
