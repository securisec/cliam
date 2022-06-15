package policy

var Microsoft_ContainerRegistry_containerregistry = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/importImage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_ImportImage",
		Resource:    "Microsoft.ContainerRegistry",
	},
	{
		Path:   "/providers/Microsoft.ContainerRegistry/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerRegistry/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_CheckNameAvailability",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerRegistry/registries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_List",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_ListByResourceGroup",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_Get",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/listUsages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_ListUsages",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_ListPrivateLinkResources",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_GetPrivateLinkResource",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/listCredentials",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_ListCredentials",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/regenerateCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Registries_RegenerateCredential",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/replications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Replications_List",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/replications/{{.replicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Replications_Get",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/webhooks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Webhooks_List",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/webhooks/{{.webhookName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Webhooks_Get",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/webhooks/{{.webhookName}}/ping",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Webhooks_Ping",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/webhooks/{{.webhookName}}/listEvents",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Webhooks_ListEvents",
		Resource:    "Microsoft.ContainerRegistry",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerRegistry/registries/{{.registryName}}/webhooks/{{.webhookName}}/getCallbackConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Webhooks_GetCallbackConfig",
		Resource:    "Microsoft.ContainerRegistry",
	},
}
