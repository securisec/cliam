package policy

// Microsoft_ResourceGraph_resourcegraph policy
var Microsoft_ResourceGraph_resourcegraph = map[string]Policy{
	"Resources": {
		Path:   "/providers/Microsoft.ResourceGraph/resources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Resources",
		Resource:    "Microsoft.ResourceGraph",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ResourceGraph/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ResourceGraph",
	},
}
