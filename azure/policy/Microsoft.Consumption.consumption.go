package policy

// Microsoft_Consumption_consumption policy
var Microsoft_Consumption_consumption = map[string]Policy{
	"UsageDetails_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Consumption/usageDetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "UsageDetails_List",
		Resource:    "Microsoft.Consumption",
	},
	"Marketplaces_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Consumption/marketplaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Marketplaces_List",
		Resource:    "Microsoft.Consumption",
	},
	"Budgets_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Consumption/budgets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Budgets_List",
		Resource:    "Microsoft.Consumption",
	},
	"Budgets_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Consumption/budgets/{{.budgetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Budgets_Get",
		Resource:    "Microsoft.Consumption",
	},
	"Tags_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Consumption/tags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Tags_Get",
		Resource:    "Microsoft.Consumption",
	},
	"Charges_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Consumption/charges",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Charges_List",
		Resource:    "Microsoft.Consumption",
	},
	"Balances_GetByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.Consumption/balances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Balances_GetByBillingAccount",
		Resource:    "Microsoft.Consumption",
	},
	"Balances_GetForBillingPeriodByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingPeriods/{{.billingPeriodName}}/providers/Microsoft.Consumption/balances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Balances_GetForBillingPeriodByBillingAccount",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationsSummaries_ListByReservationOrder": {
		Path:   "/providers/Microsoft.Capacity/reservationorders/{{.reservationOrderId}}/providers/Microsoft.Consumption/reservationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationsSummaries_ListByReservationOrder",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationsSummaries_ListByReservationOrderAndReservation": {
		Path:   "/providers/Microsoft.Capacity/reservationorders/{{.reservationOrderId}}/reservations/{{.reservationId}}/providers/Microsoft.Consumption/reservationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationsSummaries_ListByReservationOrderAndReservation",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationsSummaries_List": {
		Path:   "/{{.resourceScope}}/providers/Microsoft.Consumption/reservationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationsSummaries_List",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationsDetails_ListByReservationOrder": {
		Path:   "/providers/Microsoft.Capacity/reservationorders/{{.reservationOrderId}}/providers/Microsoft.Consumption/reservationDetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationsDetails_ListByReservationOrder",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationsDetails_ListByReservationOrderAndReservation": {
		Path:   "/providers/Microsoft.Capacity/reservationorders/{{.reservationOrderId}}/reservations/{{.reservationId}}/providers/Microsoft.Consumption/reservationDetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationsDetails_ListByReservationOrderAndReservation",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationsDetails_List": {
		Path:   "/{{.resourceScope}}/providers/Microsoft.Consumption/reservationDetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationsDetails_List",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationRecommendations_List": {
		Path:   "/{{.resourceScope}}/providers/Microsoft.Consumption/reservationRecommendations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationRecommendations_List",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationRecommendationDetails_Get": {
		Path:   "/{{.resourceScope}}/providers/Microsoft.Consumption/reservationRecommendationDetails",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationRecommendationDetails_Get",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationTransactions_List": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.Consumption/reservationTransactions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationTransactions_List",
		Resource:    "Microsoft.Consumption",
	},
	"ReservationTransactions_ListByBillingProfile": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingProfiles/{{.billingProfileId}}/providers/Microsoft.Consumption/reservationTransactions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ReservationTransactions_ListByBillingProfile",
		Resource:    "Microsoft.Consumption",
	},
	"PriceSheet_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Consumption/pricesheets/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PriceSheet_Get",
		Resource:    "Microsoft.Consumption",
	},
	"PriceSheet_GetByBillingPeriod": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Billing/billingPeriods/{{.billingPeriodName}}/providers/Microsoft.Consumption/pricesheets/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PriceSheet_GetByBillingPeriod",
		Resource:    "Microsoft.Consumption",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Consumption/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Consumption",
	},
	"AggregatedCost_GetByManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Consumption/aggregatedcost",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "AggregatedCost_GetByManagementGroup",
		Resource:    "Microsoft.Consumption",
	},
	"AggregatedCost_GetForBillingPeriodByManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Billing/billingPeriods/{{.billingPeriodName}}/providers/Microsoft.Consumption/aggregatedCost",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "AggregatedCost_GetForBillingPeriodByManagementGroup",
		Resource:    "Microsoft.Consumption",
	},
	"Events_ListByBillingProfile": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingProfiles/{{.billingProfileId}}/providers/Microsoft.Consumption/events",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Events_ListByBillingProfile",
		Resource:    "Microsoft.Consumption",
	},
	"Events_ListByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.Consumption/events",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Events_ListByBillingAccount",
		Resource:    "Microsoft.Consumption",
	},
	"Lots_ListByBillingProfile": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingProfiles/{{.billingProfileId}}/providers/Microsoft.Consumption/lots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Lots_ListByBillingProfile",
		Resource:    "Microsoft.Consumption",
	},
	"Lots_ListByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.Consumption/lots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Lots_ListByBillingAccount",
		Resource:    "Microsoft.Consumption",
	},
	"Lots_ListByCustomer": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/customers/{{.customerId}}/providers/Microsoft.Consumption/lots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Lots_ListByCustomer",
		Resource:    "Microsoft.Consumption",
	},
	"Credits_Get": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingProfiles/{{.billingProfileId}}/providers/Microsoft.Consumption/credits/balanceSummary",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Credits_Get",
		Resource:    "Microsoft.Consumption",
	},
}
