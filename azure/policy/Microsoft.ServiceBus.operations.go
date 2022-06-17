package policy

var Microsoft_ServiceBus_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.ServiceBus/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ServiceBus",
	},
}
