package policy

var MediaconnectPolicies = []Service{
	{
		ServiceSuffix: "v1/flows",
		Permission:    "ListFlows",
	},
	{
		ServiceSuffix: "v1/entitlements",
		Permission:    "ListEntitlements",
	},
	{
		ServiceSuffix: "v1/offerings",
		Permission:    "ListOfferings",
	},
	{
		ServiceSuffix: "v1/reservations",
		Permission:    "ListReservations",
	},
}
