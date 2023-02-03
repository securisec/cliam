package policy

// Microsoft_Resources_resources policy
var Microsoft_Resources_resources = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Resources/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_GetAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_GetAtScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_CancelAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_CancelAtScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ValidateAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ValidateAtScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ExportTemplateAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/exportTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ExportTemplateAtScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ListAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ListAtScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_GetAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_GetAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_CancelAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_CancelAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ValidateAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ValidateAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_WhatIfAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}/whatIf",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_WhatIfAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ExportTemplateAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}/exportTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ExportTemplateAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ListAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ListAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_GetAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_GetAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_CancelAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_CancelAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ValidateAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ValidateAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_WhatIfAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/whatIf",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_WhatIfAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ExportTemplateAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/exportTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ExportTemplateAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ListAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ListAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_GetAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_GetAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_CancelAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_CancelAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ValidateAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ValidateAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_WhatIfAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/whatIf",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_WhatIfAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ExportTemplateAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/exportTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ExportTemplateAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ListAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ListAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_Get",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_Cancel",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_Validate",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_WhatIf": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/whatIf",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_WhatIf",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ExportTemplate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/exportTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ExportTemplate",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Resources/deployments/",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_ListByResourceGroup",
		Resource:    "Microsoft.Resources",
	},
	"Providers_Unregister": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/{{.resourceProviderNamespace}}/unregister",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_Unregister",
		Resource:    "Microsoft.Resources",
	},
	"Providers_RegisterAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/{{.resourceProviderNamespace}}/register",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_RegisterAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"Providers_ProviderPermissions": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/{{.resourceProviderNamespace}}/providerPermissions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_ProviderPermissions",
		Resource:    "Microsoft.Resources",
	},
	"Providers_Register": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/{{.resourceProviderNamespace}}/register",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_Register",
		Resource:    "Microsoft.Resources",
	},
	"Providers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_List",
		Resource:    "Microsoft.Resources",
	},
	"Providers_ListAtTenantScope": {
		Path:   "/providers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_ListAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Providers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/{{.resourceProviderNamespace}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_Get",
		Resource:    "Microsoft.Resources",
	},
	"ProviderResourceTypes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/{{.resourceProviderNamespace}}/resourceTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ProviderResourceTypes_List",
		Resource:    "Microsoft.Resources",
	},
	"Providers_GetAtTenantScope": {
		Path:   "/providers/{{.resourceProviderNamespace}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Providers_GetAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"Resources_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/resources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Resources_ListByResourceGroup",
		Resource:    "Microsoft.Resources",
	},
	"ResourceGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ResourceGroups_Get",
		Resource:    "Microsoft.Resources",
	},
	"ResourceGroups_ExportTemplate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/exportTemplate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ResourceGroups_ExportTemplate",
		Resource:    "Microsoft.Resources",
	},
	"ResourceGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "ResourceGroups_List",
		Resource:    "Microsoft.Resources",
	},
	"Resources_MoveResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.sourceResourceGroupName}}/moveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Resources_MoveResources",
		Resource:    "Microsoft.Resources",
	},
	"Resources_ValidateMoveResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.sourceResourceGroupName}}/validateMoveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Resources_ValidateMoveResources",
		Resource:    "Microsoft.Resources",
	},
	"Resources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Resources_List",
		Resource:    "Microsoft.Resources",
	},
	"Resources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourcePath}}/{{.resourceType}}/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Resources_Get",
		Resource:    "Microsoft.Resources",
	},
	"Resources_GetById": {
		Path:   "/{{.resourceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Resources_GetById",
		Resource:    "Microsoft.Resources",
	},
	"Tags_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/tagNames",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Tags_List",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_GetAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_GetAtScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_ListAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_ListAtScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_GetAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_GetAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_ListAtTenantScope": {
		Path:   "/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_ListAtTenantScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_GetAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_GetAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_ListAtManagementGroupScope": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.groupId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_ListAtManagementGroupScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_GetAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_GetAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_ListAtSubscriptionScope": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Resources/deployments/{{.deploymentName}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_ListAtSubscriptionScope",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/deployments/{{.deploymentName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_Get",
		Resource:    "Microsoft.Resources",
	},
	"DeploymentOperations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/deployments/{{.deploymentName}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "DeploymentOperations_List",
		Resource:    "Microsoft.Resources",
	},
	"Deployments_CalculateTemplateHash": {
		Path:   "/providers/Microsoft.Resources/calculateTemplateHash",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Deployments_CalculateTemplateHash",
		Resource:    "Microsoft.Resources",
	},
	"Tags_GetAtScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Resources/tags/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "Tags_GetAtScope",
		Resource:    "Microsoft.Resources",
	},
}
