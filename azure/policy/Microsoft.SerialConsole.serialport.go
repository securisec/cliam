package policy

// Microsoft_SerialConsole_serialport policy
var Microsoft_SerialConsole_serialport = map[string]Policy{
	"SerialPorts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourceType}}/{{.parentResource}}/providers/Microsoft.SerialConsole/serialPorts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "SerialPorts_List",
		Resource:    "Microsoft.SerialConsole",
	},
	"SerialPorts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourceType}}/{{.parentResource}}/providers/Microsoft.SerialConsole/serialPorts/{{.serialPort}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "SerialPorts_Get",
		Resource:    "Microsoft.SerialConsole",
	},
	"SerialPorts_ListBySubscriptions": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SerialConsole/serialPorts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "SerialPorts_ListBySubscriptions",
		Resource:    "Microsoft.SerialConsole",
	},
	"SerialPorts_Connect": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.resourceProviderNamespace}}/{{.parentResourceType}}/{{.parentResource}}/providers/Microsoft.SerialConsole/serialPorts/{{.serialPort}}/connect",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2018-05-01",
		},
		OperationID: "SerialPorts_Connect",
		Resource:    "Microsoft.SerialConsole",
	},
}
