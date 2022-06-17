package policy

    var Microsoft_Logz_logz = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/listMonitoredResources",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitors_ListMonitoredResources",
    Resource:       "Microsoft.Logz",
},{
    Path: "/providers/Microsoft.Logz/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Logz/monitors",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitors_ListBySubscription",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitors_ListByResourceGroup",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitors_Get",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/tagRules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "TagRules_List",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/tagRules/{{.ruleSetName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "TagRules_Get",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/singleSignOnConfigurations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SingleSignOn_List",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/singleSignOnConfigurations/{{.configurationName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SingleSignOn_Get",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccount_List",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccount_Get",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}/tagRules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccountTagRules_List",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}/tagRules/{{.ruleSetName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccountTagRules_Get",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}/listMonitoredResources",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccount_ListMonitoredResources",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}/vmHostPayload",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccount_VMHostPayload",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}/vmHostUpdate",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccount_ListVmHostUpdate",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/accounts/{{.subAccountName}}/listVMHosts",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "SubAccount_ListVMHosts",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/vmHostPayload",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitor_VMHostPayload",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/vmHostUpdate",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitor_ListVmHostUpdate",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/listVMHosts",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitor_ListVMHosts",
    Resource:       "Microsoft.Logz",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Logz/monitors/{{.monitorName}}/listUserRoles",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-10-01",
    },
	OperationID:    "Monitors_ListUserRoles",
    Resource:       "Microsoft.Logz",
},
    }
    