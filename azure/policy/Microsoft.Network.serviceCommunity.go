package policy

// Microsoft_Network_serviceCommunity policy
var Microsoft_Network_serviceCommunity = map[string]Policy{
	"BgpServiceCommunities_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/bgpServiceCommunities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "BgpServiceCommunities_List",
		Resource:    "Microsoft.Network",
	},
}
