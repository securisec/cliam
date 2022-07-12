package policy

    var Microsoft_DataFactory_datafactory = map[string]Policy{
        "Operations_List": {
    Path: "/providers/Microsoft.DataFactory/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.DataFactory",
},
"Factories_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataFactory/factories",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Factories_List",
    Resource:       "Microsoft.DataFactory",
},
"Factories_ConfigureFactoryRepo": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataFactory/locations/{{.locationId}}/configureFactoryRepo",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Factories_ConfigureFactoryRepo",
    Resource:       "Microsoft.DataFactory",
},
"ExposureControl_GetFeatureValue": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataFactory/locations/{{.locationId}}/getFeatureValue",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ExposureControl_GetFeatureValue",
    Resource:       "Microsoft.DataFactory",
},
"ExposureControl_GetFeatureValueByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/getFeatureValue",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ExposureControl_GetFeatureValueByFactory",
    Resource:       "Microsoft.DataFactory",
},
"ExposureControl_QueryFeatureValuesByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/queryFeaturesValue",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ExposureControl_QueryFeatureValuesByFactory",
    Resource:       "Microsoft.DataFactory",
},
"Factories_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Factories_ListByResourceGroup",
    Resource:       "Microsoft.DataFactory",
},
"Factories_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Factories_Get",
    Resource:       "Microsoft.DataFactory",
},
"Factories_GetGitHubAccessToken": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/getGitHubAccessToken",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Factories_GetGitHubAccessToken",
    Resource:       "Microsoft.DataFactory",
},
"Factories_GetDataPlaneAccess": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/getDataPlaneAccess",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Factories_GetDataPlaneAccess",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_Get",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_GetStatus": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/getStatus",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_GetStatus",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/outboundNetworkDependenciesEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_GetConnectionInfo": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/getConnectionInfo",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_GetConnectionInfo",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_RegenerateAuthKey": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/regenerateAuthKey",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_RegenerateAuthKey",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_ListAuthKeys": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/listAuthKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_ListAuthKeys",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_Start": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_Start",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_Stop": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_Stop",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_SyncCredentials": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/syncCredentials",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_SyncCredentials",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_GetMonitoringData": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/monitoringData",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_GetMonitoringData",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_Upgrade": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/upgrade",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_Upgrade",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_RemoveLinks": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/removeLinks",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_RemoveLinks",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimes_CreateLinkedIntegrationRuntime": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/linkedIntegrationRuntime",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimes_CreateLinkedIntegrationRuntime",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimeObjectMetadata_Refresh": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/refreshObjectMetadata",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimeObjectMetadata_Refresh",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimeObjectMetadata_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/getObjectMetadata",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimeObjectMetadata_Get",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimeNodes_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/nodes/{{.nodeName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimeNodes_Get",
    Resource:       "Microsoft.DataFactory",
},
"IntegrationRuntimeNodes_GetIpAddress": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/integrationRuntimes/{{.integrationRuntimeName}}/nodes/{{.nodeName}}/ipAddress",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "IntegrationRuntimeNodes_GetIpAddress",
    Resource:       "Microsoft.DataFactory",
},
"LinkedServices_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/linkedservices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "LinkedServices_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"LinkedServices_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/linkedservices/{{.linkedServiceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "LinkedServices_Get",
    Resource:       "Microsoft.DataFactory",
},
"Datasets_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/datasets",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Datasets_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"Datasets_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/datasets/{{.datasetName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Datasets_Get",
    Resource:       "Microsoft.DataFactory",
},
"Pipelines_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/pipelines",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Pipelines_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"Pipelines_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/pipelines/{{.pipelineName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Pipelines_Get",
    Resource:       "Microsoft.DataFactory",
},
"Pipelines_CreateRun": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/pipelines/{{.pipelineName}}/createRun",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Pipelines_CreateRun",
    Resource:       "Microsoft.DataFactory",
},
"PipelineRuns_QueryByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/queryPipelineRuns",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "PipelineRuns_QueryByFactory",
    Resource:       "Microsoft.DataFactory",
},
"PipelineRuns_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/pipelineruns/{{.runId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "PipelineRuns_Get",
    Resource:       "Microsoft.DataFactory",
},
"ActivityRuns_QueryByPipelineRun": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/pipelineruns/{{.runId}}/queryActivityruns",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ActivityRuns_QueryByPipelineRun",
    Resource:       "Microsoft.DataFactory",
},
"PipelineRuns_Cancel": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/pipelineruns/{{.runId}}/cancel",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "PipelineRuns_Cancel",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_QueryByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/querytriggers",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_QueryByFactory",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_Get",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_SubscribeToEvents": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/subscribeToEvents",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_SubscribeToEvents",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_GetEventSubscriptionStatus": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/getEventSubscriptionStatus",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_GetEventSubscriptionStatus",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_UnsubscribeFromEvents": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/unsubscribeFromEvents",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_UnsubscribeFromEvents",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_Start": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_Start",
    Resource:       "Microsoft.DataFactory",
},
"Triggers_Stop": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "Triggers_Stop",
    Resource:       "Microsoft.DataFactory",
},
"TriggerRuns_Rerun": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/triggerRuns/{{.runId}}/rerun",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "TriggerRuns_Rerun",
    Resource:       "Microsoft.DataFactory",
},
"TriggerRuns_Cancel": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/triggers/{{.triggerName}}/triggerRuns/{{.runId}}/cancel",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "TriggerRuns_Cancel",
    Resource:       "Microsoft.DataFactory",
},
"TriggerRuns_QueryByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/queryTriggerRuns",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "TriggerRuns_QueryByFactory",
    Resource:       "Microsoft.DataFactory",
},
"DataFlows_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/dataflows/{{.dataFlowName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlows_Get",
    Resource:       "Microsoft.DataFactory",
},
"DataFlows_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/dataflows",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlows_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"DataFlowDebugSession_Create": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/createDataFlowDebugSession",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlowDebugSession_Create",
    Resource:       "Microsoft.DataFactory",
},
"DataFlowDebugSession_QueryByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/queryDataFlowDebugSessions",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlowDebugSession_QueryByFactory",
    Resource:       "Microsoft.DataFactory",
},
"DataFlowDebugSession_AddDataFlow": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/addDataFlowToDebugSession",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlowDebugSession_AddDataFlow",
    Resource:       "Microsoft.DataFactory",
},
"DataFlowDebugSession_Delete": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/deleteDataFlowDebugSession",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlowDebugSession_Delete",
    Resource:       "Microsoft.DataFactory",
},
"DataFlowDebugSession_ExecuteCommand": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/executeDataFlowDebugCommand",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "DataFlowDebugSession_ExecuteCommand",
    Resource:       "Microsoft.DataFactory",
},
"ManagedVirtualNetworks_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/managedVirtualNetworks",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ManagedVirtualNetworks_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"ManagedVirtualNetworks_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/managedVirtualNetworks/{{.managedVirtualNetworkName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ManagedVirtualNetworks_Get",
    Resource:       "Microsoft.DataFactory",
},
"ManagedPrivateEndpoints_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/managedVirtualNetworks/{{.managedVirtualNetworkName}}/managedPrivateEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ManagedPrivateEndpoints_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"ManagedPrivateEndpoints_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/managedVirtualNetworks/{{.managedVirtualNetworkName}}/managedPrivateEndpoints/{{.managedPrivateEndpointName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "ManagedPrivateEndpoints_Get",
    Resource:       "Microsoft.DataFactory",
},
"privateEndPointConnections_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/privateEndPointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "privateEndPointConnections_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"PrivateEndpointConnection_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "PrivateEndpointConnection_Get",
    Resource:       "Microsoft.DataFactory",
},
"privateLinkResources_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "privateLinkResources_Get",
    Resource:       "Microsoft.DataFactory",
},
"GlobalParameters_ListByFactory": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/globalParameters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "GlobalParameters_ListByFactory",
    Resource:       "Microsoft.DataFactory",
},
"GlobalParameters_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataFactory/factories/{{.factoryName}}/globalParameters/{{.globalParameterName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-06-01",
    },
	OperationID:    "GlobalParameters_Get",
    Resource:       "Microsoft.DataFactory",
},
    }
    