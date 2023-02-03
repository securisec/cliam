package policy

// Microsoft_Capacity_reservations policy
var Microsoft_Capacity_reservations = map[string]Policy{
	"Reservation_AvailableScopes": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/reservations/{{.reservationId}}/availableScopes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_AvailableScopes",
		Resource:    "Microsoft.Capacity",
	},
	"GetCatalog": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Capacity/catalogs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "GetCatalog",
		Resource:    "Microsoft.Capacity",
	},
	"GetAppliedReservationList": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Capacity/appliedReservations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "GetAppliedReservationList",
		Resource:    "Microsoft.Capacity",
	},
	"ReservationOrder_Calculate": {
		Path:   "/providers/Microsoft.Capacity/calculatePrice",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "ReservationOrder_Calculate",
		Resource:    "Microsoft.Capacity",
	},
	"ReservationOrder_List": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "ReservationOrder_List",
		Resource:    "Microsoft.Capacity",
	},
	"ReservationOrder_Get": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "ReservationOrder_Get",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_Split": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/split",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_Split",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_Merge": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/merge",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_Merge",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_List": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/reservations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_List",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_Get": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/reservations/{{.reservationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_Get",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_Archive": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/reservations/{{.reservationId}}/archive",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_Archive",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_Unarchive": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/reservations/{{.reservationId}}/unarchive",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_Unarchive",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_ListRevisions": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/reservations/{{.reservationId}}/revisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_ListRevisions",
		Resource:    "Microsoft.Capacity",
	},
	"Operation_List": {
		Path:   "/providers/Microsoft.Capacity/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Operation_List",
		Resource:    "Microsoft.Capacity",
	},
	"CalculateRefund_Post": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/calculateRefund",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "CalculateRefund_Post",
		Resource:    "Microsoft.Capacity",
	},
	"Return_Post": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/return",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Return_Post",
		Resource:    "Microsoft.Capacity",
	},
	"CalculateExchange_Post": {
		Path:   "/providers/Microsoft.Capacity/calculateExchange",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "CalculateExchange_Post",
		Resource:    "Microsoft.Capacity",
	},
	"Exchange_Post": {
		Path:   "/providers/Microsoft.Capacity/exchange",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Exchange_Post",
		Resource:    "Microsoft.Capacity",
	},
	"Reservation_ListAll": {
		Path:   "/providers/Microsoft.Capacity/reservations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Reservation_ListAll",
		Resource:    "Microsoft.Capacity",
	},
	"ReservationOrder_ChangeDirectory": {
		Path:   "/providers/Microsoft.Capacity/reservationOrders/{{.reservationOrderId}}/changeDirectory",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "ReservationOrder_ChangeDirectory",
		Resource:    "Microsoft.Capacity",
	},
}
