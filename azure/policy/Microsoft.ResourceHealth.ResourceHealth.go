package policy

// Microsoft_ResourceHealth_ResourceHealth policy
var Microsoft_ResourceHealth_ResourceHealth = map[string]Policy{
	"AvailabilityStatuses_ListBySubscriptionId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceHealth/availabilityStatuses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "AvailabilityStatuses_ListBySubscriptionId",
		Resource:    "Microsoft.ResourceHealth",
	},
	"AvailabilityStatuses_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ResourceHealth/availabilityStatuses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "AvailabilityStatuses_ListByResourceGroup",
		Resource:    "Microsoft.ResourceHealth",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ResourceHealth/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ResourceHealth",
	},
	"AvailabilityStatuses_GetByResource": {
		Path:   "/{{.resourceUri}}/providers/Microsoft.ResourceHealth/availabilityStatuses/current",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "AvailabilityStatuses_GetByResource",
		Resource:    "Microsoft.ResourceHealth",
	},
	"AvailabilityStatuses_List": {
		Path:   "/{{.resourceUri}}/providers/Microsoft.ResourceHealth/availabilityStatuses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "AvailabilityStatuses_List",
		Resource:    "Microsoft.ResourceHealth",
	},
	"ImpactedResources_ListBySubscriptionIdAndEventId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceHealth/events/{{.eventTrackingId}}/impactedResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ImpactedResources_ListBySubscriptionIdAndEventId",
		Resource:    "Microsoft.ResourceHealth",
	},
	"ImpactedResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceHealth/events/{{.eventTrackingId}}/impactedResources/{{.impactedResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ImpactedResources_Get",
		Resource:    "Microsoft.ResourceHealth",
	},
	"Events_ListBySubscriptionId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceHealth/events",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Events_ListBySubscriptionId",
		Resource:    "Microsoft.ResourceHealth",
	},
	"Event_GetBySubscriptionIdAndTrackingId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceHealth/events/{{.eventTrackingId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Event_GetBySubscriptionIdAndTrackingId",
		Resource:    "Microsoft.ResourceHealth",
	},
	"Events_ListByTenantId": {
		Path:   "/providers/Microsoft.ResourceHealth/events",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Events_ListByTenantId",
		Resource:    "Microsoft.ResourceHealth",
	},
	"Event_GetByTenantIdAndTrackingId": {
		Path:   "/providers/Microsoft.ResourceHealth/events/{{.eventTrackingId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Event_GetByTenantIdAndTrackingId",
		Resource:    "Microsoft.ResourceHealth",
	},
	"Events_ListBySingleResource": {
		Path:   "/{{.resourceUri}}/providers/Microsoft.ResourceHealth/events",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Events_ListBySingleResource",
		Resource:    "Microsoft.ResourceHealth",
	},
}
