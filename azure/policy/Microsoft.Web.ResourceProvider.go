package policy

// Microsoft_Web_ResourceProvider policy
var Microsoft_Web_ResourceProvider = map[string]Policy{
	"GetPublishingUser": {
		Path:   "/providers/Microsoft.Web/publishingUsers/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "GetPublishingUser",
		Resource:    "Microsoft.Web",
	},
	"ListSourceControls": {
		Path:   "/providers/Microsoft.Web/sourcecontrols",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListSourceControls",
		Resource:    "Microsoft.Web",
	},
	"GetSourceControl": {
		Path:   "/providers/Microsoft.Web/sourcecontrols/{{.sourceControlType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "GetSourceControl",
		Resource:    "Microsoft.Web",
	},
	"ListBillingMeters": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/billingMeters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListBillingMeters",
		Resource:    "Microsoft.Web",
	},
	"CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/checknameavailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "CheckNameAvailability",
		Resource:    "Microsoft.Web",
	},
	"ListCustomHostNameSites": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/customhostnameSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListCustomHostNameSites",
		Resource:    "Microsoft.Web",
	},
	"GetSubscriptionDeploymentLocations": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/deploymentLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "GetSubscriptionDeploymentLocations",
		Resource:    "Microsoft.Web",
	},
	"ListGeoRegions": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/geoRegions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListGeoRegions",
		Resource:    "Microsoft.Web",
	},
	"ListSiteIdentifiersAssignedToHostName": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/listSitesAssignedToHostName",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListSiteIdentifiersAssignedToHostName",
		Resource:    "Microsoft.Web",
	},
	"ListPremierAddOnOffers": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/premieraddonoffers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListPremierAddOnOffers",
		Resource:    "Microsoft.Web",
	},
	"ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ListSkus",
		Resource:    "Microsoft.Web",
	},
	"VerifyHostingEnvironmentVnet": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/verifyHostingEnvironmentVnet",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "VerifyHostingEnvironmentVnet",
		Resource:    "Microsoft.Web",
	},
	"Move": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/moveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Move",
		Resource:    "Microsoft.Web",
	},
	"Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Validate",
		Resource:    "Microsoft.Web",
	},
	"ValidateMove": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/validateMoveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ValidateMove",
		Resource:    "Microsoft.Web",
	},
}
