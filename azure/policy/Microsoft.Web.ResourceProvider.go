package policy

var Microsoft_Web_ResourceProvider = []Policy{
	{
		Path:   "/providers/Microsoft.Web/publishingUsers/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "GetPublishingUser",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/providers/Microsoft.Web/sourcecontrols",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListSourceControls",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/providers/Microsoft.Web/sourcecontrols/{{.sourceControlType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID:   "GetSourceControl",
		Resource:      "Microsoft.Web",
		IsExtra:       true,
		ExtraLocation: "path",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/billingMeters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListBillingMeters",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/checknameavailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "CheckNameAvailability",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/customhostnameSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListCustomHostNameSites",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deploymentLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "GetSubscriptionDeploymentLocations",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/geoRegions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListGeoRegions",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/listSitesAssignedToHostName",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListSiteIdentifiersAssignedToHostName",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/premieraddonoffers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListPremierAddOnOffers",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ListSkus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/verifyHostingEnvironmentVnet",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "VerifyHostingEnvironmentVnet",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/moveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Move",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Validate",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/validateMoveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ValidateMove",
		Resource:    "Microsoft.Web",
	},
}
