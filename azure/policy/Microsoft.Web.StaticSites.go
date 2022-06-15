package policy

var Microsoft_Web_StaticSites = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/locations/{{.location}}/previewStaticSiteWorkflowFile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_PreviewWorkflow",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/staticSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_List",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSitesByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSite",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/authproviders/{{.authprovider}}/listUsers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteUsers",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSiteBuilds",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteBuildFunctions",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/listAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteBuildAppSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/listFunctionAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteBuildFunctionAppSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/userProvidedFunctionApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppsForStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/userProvidedFunctionApps/{{.functionAppName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppForStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/zipdeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_CreateZipDeploymentForStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/createUserInvitation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_CreateUserRolesInvitationLink",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/customDomains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteCustomDomains",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/customDomains/{{.domainName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSiteCustomDomain",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/customDomains/{{.domainName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ValidateCustomDomainCanBeAddedToStaticSite",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/detach",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_DetachStaticSite",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteFunctions",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteAppSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listConfiguredRoles",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteConfiguredRoles",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listFunctionAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteFunctionAppSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteSecrets",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetPrivateEndpointConnectionList",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetPrivateLinkResources",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/resetapikey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ResetStaticSiteApiKey",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/userProvidedFunctionApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppsForStaticSite",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/userProvidedFunctionApps/{{.functionAppName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppForStaticSite",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/zipdeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_CreateZipDeploymentForStaticSite",
		Resource:    "Microsoft.Web",
	},
}
