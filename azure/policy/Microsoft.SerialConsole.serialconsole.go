package policy

    var Microsoft_SerialConsole_serialconsole = []Policy{
        {
    Path: "/providers/Microsoft.SerialConsole/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "ListOperations",
    Resource:       "Microsoft.SerialConsole",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SerialConsole/consoleServices/{{.default}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "GetConsoleStatus",
    Resource:       "Microsoft.SerialConsole",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SerialConsole/consoleServices/{{.default}}/disableConsole",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "DisableConsole",
    Resource:       "Microsoft.SerialConsole",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SerialConsole/consoleServices/{{.default}}/enableConsole",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2018-05-01",
    },
	OperationID:    "EnableConsole",
    Resource:       "Microsoft.SerialConsole",
},
    }
    