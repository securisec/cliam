package policy

var Microsoft_AppPlatform_appplatform = map[string]Policy{
	"Services_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_ListTestKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/listTestKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_ListTestKeys",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_RegenerateTestKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/regenerateTestKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_RegenerateTestKey",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_DisableTestEndpoint": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/disableTestEndpoint",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_DisableTestEndpoint",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_EnableTestEndpoint": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/enableTestEndpoint",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_EnableTestEndpoint",
		Resource:    "Microsoft.AppPlatform",
	},
	"ConfigServers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/configServers/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ConfigServers_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"ConfigServers_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/configServers/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ConfigServers_Validate",
		Resource:    "Microsoft.AppPlatform",
	},
	"ConfigurationServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/configurationServices/{{.configurationServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ConfigurationServices_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"ConfigurationServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/configurationServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ConfigurationServices_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"ConfigurationServices_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/configurationServices/{{.configurationServiceName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ConfigurationServices_Validate",
		Resource:    "Microsoft.AppPlatform",
	},
	"ServiceRegistries_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/serviceRegistries/{{.serviceRegistryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ServiceRegistries_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"ServiceRegistries_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/serviceRegistries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ServiceRegistries_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_ListBuildServices": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_ListBuildServices",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetBuildService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetBuildService",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_ListBuilds": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builds",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_ListBuilds",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetBuild": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builds/{{.buildName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetBuild",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildpackBinding_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builders/{{.builderName}}/buildpackBindings/{{.buildpackBindingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildpackBinding_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildpackBinding_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builders/{{.builderName}}/buildpackBindings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildpackBinding_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_ListBuildResults": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builds/{{.buildName}}/results",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_ListBuildResults",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetBuildResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builds/{{.buildName}}/results/{{.buildResultName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetBuildResult",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetBuildResultLog": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builds/{{.buildName}}/results/{{.buildResultName}}/getLogFileUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetBuildResultLog",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildServiceBuilder_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builders/{{.builderName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildServiceBuilder_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildServiceBuilder_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/builders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildServiceBuilder_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetResourceUploadUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/getResourceUploadUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetResourceUploadUrl",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_ListSupportedBuildpacks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/supportedBuildpacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_ListSupportedBuildpacks",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetSupportedBuildpack": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/supportedBuildpacks/{{.buildpackName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetSupportedBuildpack",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_ListSupportedStacks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/supportedStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_ListSupportedStacks",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildService_GetSupportedStack": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/supportedStacks/{{.stackName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildService_GetSupportedStack",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildServiceAgentPool_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/agentPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildServiceAgentPool_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"BuildServiceAgentPool_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/buildServices/{{.buildServiceName}}/agentPools/{{.agentPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "BuildServiceAgentPool_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"MonitoringSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/monitoringSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "MonitoringSettings_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"Apps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Apps_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"Apps_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Apps_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"Apps_GetResourceUploadUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/getResourceUploadUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Apps_GetResourceUploadUrl",
		Resource:    "Microsoft.AppPlatform",
	},
	"Apps_SetActiveDeployments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/setActiveDeployments",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Apps_SetActiveDeployments",
		Resource:    "Microsoft.AppPlatform",
	},
	"Bindings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/bindings/{{.bindingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Bindings_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"Bindings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/bindings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Bindings_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"Certificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Certificates_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"Certificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Certificates_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppPlatform/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_CheckNameAvailability",
		Resource:    "Microsoft.AppPlatform",
	},
	"CustomDomains_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/domains/{{.domainName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "CustomDomains_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"CustomDomains_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/domains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "CustomDomains_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"Apps_ValidateDomain": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/validateDomain",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Apps_ValidateDomain",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_Get",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_ListForCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_ListForCluster",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_Start",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_Stop",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_Restart",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_GetLogFileUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/getLogFileUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_GetLogFileUrl",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_GenerateHeapDump": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/generateHeapDump",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_GenerateHeapDump",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_GenerateThreadDump": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/generateThreadDump",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_GenerateThreadDump",
		Resource:    "Microsoft.AppPlatform",
	},
	"Deployments_StartJFR": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring/{{.serviceName}}/apps/{{.appName}}/deployments/{{.deploymentName}}/startJFR",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Deployments_StartJFR",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppPlatform/Spring",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_ListBySubscription",
		Resource:    "Microsoft.AppPlatform",
	},
	"Services_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppPlatform/Spring",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Services_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.AppPlatform/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AppPlatform",
	},
	"RuntimeVersions_ListRuntimeVersions": {
		Path:   "/providers/Microsoft.AppPlatform/runtimeVersions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "RuntimeVersions_ListRuntimeVersions",
		Resource:    "Microsoft.AppPlatform",
	},
	"Skus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppPlatform/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Skus_List",
		Resource:    "Microsoft.AppPlatform",
	},
}
