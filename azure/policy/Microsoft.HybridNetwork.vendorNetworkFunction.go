package policy

    var Microsoft_HybridNetwork_vendorNetworkFunction = map[string]Policy{
        "VendorNetworkFunctions_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions/{{.serviceKey}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "VendorNetworkFunctions_Get",
    Resource:       "Microsoft.HybridNetwork",
},
"RoleInstances_Start": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions/{{.serviceKey}}/roleInstances/{{.roleInstanceName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "RoleInstances_Start",
    Resource:       "Microsoft.HybridNetwork",
},
"RoleInstances_Stop": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions/{{.serviceKey}}/roleInstances/{{.roleInstanceName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "RoleInstances_Stop",
    Resource:       "Microsoft.HybridNetwork",
},
"RoleInstances_Restart": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions/{{.serviceKey}}/roleInstances/{{.roleInstanceName}}/restart",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "RoleInstances_Restart",
    Resource:       "Microsoft.HybridNetwork",
},
"RoleInstances_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions/{{.serviceKey}}/roleInstances/{{.roleInstanceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "RoleInstances_Get",
    Resource:       "Microsoft.HybridNetwork",
},
"VendorNetworkFunctions_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "VendorNetworkFunctions_List",
    Resource:       "Microsoft.HybridNetwork",
},
"RoleInstances_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/locations/{{.locationName}}/vendors/{{.vendorName}}/networkFunctions/{{.serviceKey}}/roleInstances",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "RoleInstances_List",
    Resource:       "Microsoft.HybridNetwork",
},
    }
    