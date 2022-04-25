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

	// extra
	{
		ServiceSuffix:          "/v1/flows/{{.flow_arn}}",
		Permission:             "DescribeFlow",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "flow_arn",
	},
	{
		ServiceSuffix:          "/v1/offerings/{{.offering_arn}}",
		Permission:             "DescribeOffering",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "offering_arn",
	},
	{
		ServiceSuffix:          "/v1/reservations/{{.reservation_arn}}",
		Permission:             "DescribeReservation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "reservation_arn",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
