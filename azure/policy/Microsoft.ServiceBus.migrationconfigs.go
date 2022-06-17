package policy

    var Microsoft_ServiceBus_migrationconfigs = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/migrationConfigurations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "MigrationConfigs_List",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/migrationConfigurations/{{.configName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "MigrationConfigs_Get",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/migrationConfigurations/{{.configName}}/upgrade",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "MigrationConfigs_CompleteMigration",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/migrationConfigurations/{{.configName}}/revert",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "MigrationConfigs_Revert",
    Resource:       "Microsoft.ServiceBus",
},
    }
    