package policy

    var Microsoft_HybridNetwork_operation = map[string]Policy{
        "Operations_List": {
    Path: "/providers/Microsoft.HybridNetwork/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.HybridNetwork",
},
    }
    