package policy

// Microsoft_Network_checkDnsAvailability policy
var Microsoft_Network_checkDnsAvailability = map[string]Policy{
	"CheckDnsNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/CheckDnsNameAvailability",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "CheckDnsNameAvailability",
		Resource:    "Microsoft.Network",
	},
}
