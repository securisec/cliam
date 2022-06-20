package policy

    var Microsoft_Management_management = map[string]Policy{
        "ManagementGroups_List": {
    Path: "/providers/Microsoft.Management/managementGroups",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "ManagementGroups_List",
    Resource:       "Microsoft.Management",
},
"ManagementGroups_Get": {
    Path: "/providers/Microsoft.Management/managementGroups/{{.groupId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "ManagementGroups_Get",
    Resource:       "Microsoft.Management",
},
"ManagementGroups_GetDescendants": {
    Path: "/providers/Microsoft.Management/managementGroups/{{.groupId}}/descendants",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "ManagementGroups_GetDescendants",
    Resource:       "Microsoft.Management",
},
"ManagementGroupSubscriptions_GetSubscription": {
    Path: "/providers/Microsoft.Management/managementGroups/{{.groupId}}/subscriptions/{{.subscriptionId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "ManagementGroupSubscriptions_GetSubscription",
    Resource:       "Microsoft.Management",
},
"ManagementGroupSubscriptions_GetSubscriptionsUnderManagementGroup": {
    Path: "/providers/Microsoft.Management/managementGroups/{{.groupId}}/subscriptions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "ManagementGroupSubscriptions_GetSubscriptionsUnderManagementGroup",
    Resource:       "Microsoft.Management",
},
"HierarchySettings_List": {
    Path: "/providers/Microsoft.Management/managementGroups/{{.groupId}}/settings",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "HierarchySettings_List",
    Resource:       "Microsoft.Management",
},
"HierarchySettings_Get": {
    Path: "/providers/Microsoft.Management/managementGroups/{{.groupId}}/settings/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "HierarchySettings_Get",
    Resource:       "Microsoft.Management",
},
"Operations_List": {
    Path: "/providers/Microsoft.Management/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.Management",
},
"CheckNameAvailability": {
    Path: "/providers/Microsoft.Management/checkNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "CheckNameAvailability",
    Resource:       "Microsoft.Management",
},
"Entities_List": {
    Path: "/providers/Microsoft.Management/getEntities",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "Entities_List",
    Resource:       "Microsoft.Management",
},
"StartTenantBackfill": {
    Path: "/providers/Microsoft.Management/startTenantBackfill",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "StartTenantBackfill",
    Resource:       "Microsoft.Management",
},
"TenantBackfillStatus": {
    Path: "/providers/Microsoft.Management/tenantBackfillStatus",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-04-01",
    },
	OperationID:    "TenantBackfillStatus",
    Resource:       "Microsoft.Management",
},
    }
    