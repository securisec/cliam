package policy

var Microsoft_Web_StaticSites = map[string]Policy{
	"StaticSites_PreviewWorkflow": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/locations/{{.location}}/previewStaticSiteWorkflowFile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_PreviewWorkflow",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/staticSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_List",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetStaticSitesByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSitesByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetStaticSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSite",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteUsers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/authproviders/{{.authprovider}}/listUsers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteUsers",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetStaticSiteBuilds": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSiteBuilds",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetStaticSiteBuild": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteBuildFunctions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteBuildFunctions",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteBuildAppSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/listAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteBuildAppSettings",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteBuildFunctionAppSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/listFunctionAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteBuildFunctionAppSettings",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetUserProvidedFunctionAppsForStaticSiteBuild": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/userProvidedFunctionApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppsForStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetUserProvidedFunctionAppForStaticSiteBuild": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/userProvidedFunctionApps/{{.functionAppName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppForStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_CreateZipDeploymentForStaticSiteBuild": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/builds/{{.environmentName}}/zipdeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_CreateZipDeploymentForStaticSiteBuild",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_CreateUserRolesInvitationLink": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/createUserInvitation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_CreateUserRolesInvitationLink",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteCustomDomains": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/customDomains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteCustomDomains",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetStaticSiteCustomDomain": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/customDomains/{{.domainName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetStaticSiteCustomDomain",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ValidateCustomDomainCanBeAddedToStaticSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/customDomains/{{.domainName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ValidateCustomDomainCanBeAddedToStaticSite",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_DetachStaticSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/detach",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_DetachStaticSite",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteFunctions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteFunctions",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteAppSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteAppSettings",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteConfiguredRoles": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listConfiguredRoles",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteConfiguredRoles",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteFunctionAppSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listFunctionAppSettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteFunctionAppSettings",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ListStaticSiteSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ListStaticSiteSecrets",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetPrivateEndpointConnectionList": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetPrivateEndpointConnectionList",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetPrivateLinkResources",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_ResetStaticSiteApiKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/resetapikey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_ResetStaticSiteApiKey",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetUserProvidedFunctionAppsForStaticSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/userProvidedFunctionApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppsForStaticSite",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_GetUserProvidedFunctionAppForStaticSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/userProvidedFunctionApps/{{.functionAppName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_GetUserProvidedFunctionAppForStaticSite",
		Resource:    "Microsoft.Web",
	},
	"StaticSites_CreateZipDeploymentForStaticSite": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/staticSites/{{.name}}/zipdeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "StaticSites_CreateZipDeploymentForStaticSite",
		Resource:    "Microsoft.Web",
	},
}
