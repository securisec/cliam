package policy

    var Microsoft_HybridNetwork_device = map[string]Policy{
        "Devices_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridNetwork/devices/{{.deviceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "Devices_Get",
    Resource:       "Microsoft.HybridNetwork",
},
"Devices_ListBySubscription": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/devices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "Devices_ListBySubscription",
    Resource:       "Microsoft.HybridNetwork",
},
"Devices_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridNetwork/devices",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "Devices_ListByResourceGroup",
    Resource:       "Microsoft.HybridNetwork",
},
"Devices_ListRegistrationKey": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridNetwork/devices/{{.deviceName}}/listRegistrationKey",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "Devices_ListRegistrationKey",
    Resource:       "Microsoft.HybridNetwork",
},
    }
    