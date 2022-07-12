package policy

    var Microsoft_ServiceLinker_servicelinker = map[string]Policy{
        "Linker_List": {
    Path: "/{{.resourceUri}}/providers/Microsoft.ServiceLinker/linkers",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Linker_List",
    Resource:       "Microsoft.ServiceLinker",
},
"Linker_Get": {
    Path: "/{{.resourceUri}}/providers/Microsoft.ServiceLinker/linkers/{{.linkerName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Linker_Get",
    Resource:       "Microsoft.ServiceLinker",
},
"Linker_Validate": {
    Path: "/{{.resourceUri}}/providers/Microsoft.ServiceLinker/linkers/{{.linkerName}}/validateLinker",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Linker_Validate",
    Resource:       "Microsoft.ServiceLinker",
},
"Linker_ListConfigurations": {
    Path: "/{{.resourceUri}}/providers/Microsoft.ServiceLinker/linkers/{{.linkerName}}/listConfigurations",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Linker_ListConfigurations",
    Resource:       "Microsoft.ServiceLinker",
},
"Operations_List": {
    Path: "/providers/Microsoft.ServiceLinker/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-05-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.ServiceLinker",
},
    }
    