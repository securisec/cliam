package policy

    var Microsoft_App_Global = []Policy{
        {
    Path: "/providers/Microsoft.App/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.App",
},
    }
    