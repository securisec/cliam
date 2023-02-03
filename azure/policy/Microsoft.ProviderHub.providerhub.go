package policy

// Microsoft_ProviderHub_providerhub policy
var Microsoft_ProviderHub_providerhub = map[string]Policy{
	"CustomRollouts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/customRollouts/{{.rolloutName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "CustomRollouts_Get",
		Resource:    "Microsoft.ProviderHub",
	},
	"CustomRollouts_ListByProviderRegistration": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/customRollouts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "CustomRollouts_ListByProviderRegistration",
		Resource:    "Microsoft.ProviderHub",
	},
	"DefaultRollouts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/defaultRollouts/{{.rolloutName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "DefaultRollouts_Get",
		Resource:    "Microsoft.ProviderHub",
	},
	"DefaultRollouts_ListByProviderRegistration": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/defaultRollouts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "DefaultRollouts_ListByProviderRegistration",
		Resource:    "Microsoft.ProviderHub",
	},
	"DefaultRollouts_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/defaultRollouts/{{.rolloutName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "DefaultRollouts_Stop",
		Resource:    "Microsoft.ProviderHub",
	},
	"GenerateManifest": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/generateManifest",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "GenerateManifest",
		Resource:    "Microsoft.ProviderHub",
	},
	"CheckinManifest": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/checkinManifest",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "CheckinManifest",
		Resource:    "Microsoft.ProviderHub",
	},
	"NotificationRegistrations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/notificationRegistrations/{{.notificationRegistrationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "NotificationRegistrations_Get",
		Resource:    "Microsoft.ProviderHub",
	},
	"NotificationRegistrations_ListByProviderRegistration": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/notificationRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "NotificationRegistrations_ListByProviderRegistration",
		Resource:    "Microsoft.ProviderHub",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ProviderHub/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ProviderHub",
	},
	"Operations_ListByProviderRegistration": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/operations/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Operations_ListByProviderRegistration",
		Resource:    "Microsoft.ProviderHub",
	},
	"ProviderRegistrations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "ProviderRegistrations_Get",
		Resource:    "Microsoft.ProviderHub",
	},
	"ProviderRegistrations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "ProviderRegistrations_List",
		Resource:    "Microsoft.ProviderHub",
	},
	"ProviderRegistrations_GenerateOperations": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/generateOperations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "ProviderRegistrations_GenerateOperations",
		Resource:    "Microsoft.ProviderHub",
	},
	"ResourceTypeRegistrations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "ResourceTypeRegistrations_Get",
		Resource:    "Microsoft.ProviderHub",
	},
	"ResourceTypeRegistrations_ListByProviderRegistration": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "ResourceTypeRegistrations_ListByProviderRegistration",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/skus/{{.sku}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_Get",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_GetNestedResourceTypeFirst": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/resourcetypeRegistrations/{{.nestedResourceTypeFirst}}/skus/{{.sku}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_GetNestedResourceTypeFirst",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_GetNestedResourceTypeSecond": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/resourcetypeRegistrations/{{.nestedResourceTypeFirst}}/resourcetypeRegistrations/{{.nestedResourceTypeSecond}}/skus/{{.sku}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_GetNestedResourceTypeSecond",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_GetNestedResourceTypeThird": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/resourcetypeRegistrations/{{.nestedResourceTypeFirst}}/resourcetypeRegistrations/{{.nestedResourceTypeSecond}}/resourcetypeRegistrations/{{.nestedResourceTypeThird}}/skus/{{.sku}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_GetNestedResourceTypeThird",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_ListByResourceTypeRegistrations": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_ListByResourceTypeRegistrations",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/resourcetypeRegistrations/{{.nestedResourceTypeFirst}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/resourcetypeRegistrations/{{.nestedResourceTypeFirst}}/resourcetypeRegistrations/{{.nestedResourceTypeSecond}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond",
		Resource:    "Microsoft.ProviderHub",
	},
	"Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ProviderHub/providerRegistrations/{{.providerNamespace}}/resourcetypeRegistrations/{{.resourceType}}/resourcetypeRegistrations/{{.nestedResourceTypeFirst}}/resourcetypeRegistrations/{{.nestedResourceTypeSecond}}/resourcetypeRegistrations/{{.nestedResourceTypeThird}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-20",
		},
		OperationID: "Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird",
		Resource:    "Microsoft.ProviderHub",
	},
}
