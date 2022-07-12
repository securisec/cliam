package policy

    var Microsoft_StreamAnalytics_subscriptions = map[string]Policy{
        "Subscriptions_ListQuotas": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StreamAnalytics/locations/{{.location}}/quotas",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "Subscriptions_ListQuotas",
    Resource:       "Microsoft.StreamAnalytics",
},
    }
    