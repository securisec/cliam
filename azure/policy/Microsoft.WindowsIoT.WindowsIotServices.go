package policy

    var Microsoft_WindowsIoT_WindowsIotServices = map[string]Policy{
        "Operations_List": {
    Path: "/providers/Microsoft.WindowsIoT/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-06-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.WindowsIoT",
},
"Services_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.WindowsIoT/deviceServices/{{.deviceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-06-01",
    },
	OperationID:    "Services_Get",
    Resource:       "Microsoft.WindowsIoT",
},
"Services_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.WindowsIoT/deviceServices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-06-01",
    },
	OperationID:    "Services_ListByResourceGroup",
    Resource:       "Microsoft.WindowsIoT",
},
"Services_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.WindowsIoT/deviceServices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-06-01",
    },
	OperationID:    "Services_List",
    Resource:       "Microsoft.WindowsIoT",
},
"Services_CheckDeviceServiceNameAvailability": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.WindowsIoT/checkDeviceServiceNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-06-01",
    },
	OperationID:    "Services_CheckDeviceServiceNameAvailability",
    Resource:       "Microsoft.WindowsIoT",
},
    }
    