package policy

var Microsoft_Compute_skus = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ResourceSkus_List",
		Resource:    "Microsoft.Compute",
	},
}
