package policy

    var Microsoft_GuestConfiguration_guestconfiguration = map[string]Policy{
        "GuestConfigurationAssignments_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.guestConfigurationAssignmentName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignments_Get",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignments_SubscriptionList": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignments_SubscriptionList",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignments_RGList": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignments_RGList",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignments_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignments_List",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignmentReports_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.guestConfigurationAssignmentName}}/reports",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignmentReports_List",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignmentReports_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.guestConfigurationAssignmentName}}/reports/{{.reportId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignmentReports_Get",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationHCRPAssignments_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.guestConfigurationAssignmentName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationHCRPAssignments_Get",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationHCRPAssignments_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationHCRPAssignments_List",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationHCRPAssignmentReports_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.guestConfigurationAssignmentName}}/reports",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationHCRPAssignmentReports_List",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationHCRPAssignmentReports_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.guestConfigurationAssignmentName}}/reports/{{.reportId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationHCRPAssignmentReports_Get",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignmentsVMSS_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmssName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.name}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignmentsVMSS_Get",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignmentsVMSS_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmssName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignmentsVMSS_List",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignmentReportsVMSS_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmssName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.name}}/reports",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignmentReportsVMSS_List",
    Resource:       "Microsoft.GuestConfiguration",
},
"GuestConfigurationAssignmentReportsVMSS_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmssName}}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{{.name}}/reports/{{.id}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "GuestConfigurationAssignmentReportsVMSS_Get",
    Resource:       "Microsoft.GuestConfiguration",
},
"Operations_List": {
    Path: "/providers/Microsoft.GuestConfiguration/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-01-25",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.GuestConfiguration",
},
    }
    