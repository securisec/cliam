package policy

// Microsoft_DesktopVirtualization_desktopvirtualization policy
var Microsoft_DesktopVirtualization_desktopvirtualization = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.DesktopVirtualization/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Workspaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/workspaces/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Workspaces_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Workspaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Workspaces_ListByResourceGroup",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Workspaces_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DesktopVirtualization/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Workspaces_ListBySubscription",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ScalingPlans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/scalingPlans/{{.scalingPlanName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ScalingPlans_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ScalingPlans_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/scalingPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ScalingPlans_ListByResourceGroup",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ScalingPlanPooledSchedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/scalingPlans/{{.scalingPlanName}}/pooledSchedules/{{.scalingPlanScheduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ScalingPlanPooledSchedules_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ScalingPlanPooledSchedules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/scalingPlans/{{.scalingPlanName}}/pooledSchedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ScalingPlanPooledSchedules_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ScalingPlans_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DesktopVirtualization/scalingPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ScalingPlans_ListBySubscription",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ApplicationGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups/{{.applicationGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ApplicationGroups_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ApplicationGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ApplicationGroups_ListByResourceGroup",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ApplicationGroups_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DesktopVirtualization/applicationGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ApplicationGroups_ListBySubscription",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"StartMenuItems_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups/{{.applicationGroupName}}/startMenuItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "StartMenuItems_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Applications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups/{{.applicationGroupName}}/applications/{{.applicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Applications_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Applications_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups/{{.applicationGroupName}}/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Applications_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Desktops_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups/{{.applicationGroupName}}/desktops/{{.desktopName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Desktops_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"Desktops_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/applicationGroups/{{.applicationGroupName}}/desktops",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "Desktops_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"HostPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "HostPools_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"HostPools_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "HostPools_ListByResourceGroup",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"HostPools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DesktopVirtualization/hostPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "HostPools_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"HostPools_RetrieveRegistrationToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/retrieveRegistrationToken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "HostPools_RetrieveRegistrationToken",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"UserSessions_ListByHostPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/userSessions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "UserSessions_ListByHostPool",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"SessionHosts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/sessionHosts/{{.sessionHostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "SessionHosts_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"SessionHosts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/sessionHosts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "SessionHosts_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"MSIXPackages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/msixPackages/{{.msixPackageFullName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "MSIXPackages_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"MSIXPackages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/msixPackages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "MSIXPackages_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"MsixImages_Expand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/expandMsixImage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "MsixImages_Expand",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"UserSessions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/sessionHosts/{{.sessionHostName}}/userSessions/{{.userSessionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "UserSessions_Get",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"UserSessions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/sessionHosts/{{.sessionHostName}}/userSessions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "UserSessions_List",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"UserSessions_Disconnect": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/sessionHosts/{{.sessionHostName}}/userSessions/{{.userSessionId}}/disconnect",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "UserSessions_Disconnect",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"UserSessions_SendMessage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/sessionHosts/{{.sessionHostName}}/userSessions/{{.userSessionId}}/sendMessage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "UserSessions_SendMessage",
		Resource:    "Microsoft.DesktopVirtualization",
	},
	"ScalingPlans_ListByHostPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DesktopVirtualization/hostPools/{{.hostPoolName}}/scalingPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-09",
		},
		OperationID: "ScalingPlans_ListByHostPool",
		Resource:    "Microsoft.DesktopVirtualization",
	},
}
