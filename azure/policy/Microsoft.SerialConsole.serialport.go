package policy

    var Microsoft_SerialConsole_serialport = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourceType}}/{{.parentResource}}/providers/Microsoft.SerialConsole/serialPorts",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "SerialPorts_List",
    Resource:       "Microsoft.SerialConsole",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourceType}}/{{.parentResource}}/providers/Microsoft.SerialConsole/serialPorts/{{.serialPort}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "SerialPorts_Get",
    Resource:       "Microsoft.SerialConsole",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SerialConsole/serialPorts",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "SerialPorts_ListBySubscriptions",
    Resource:       "Microsoft.SerialConsole",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourceType}}/{{.parentResource}}/providers/Microsoft.SerialConsole/serialPorts/{{.serialPort}}/connect",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "SerialPorts_Connect",
    Resource:       "Microsoft.SerialConsole",
},
    }
    