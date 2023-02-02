package policy

// MediaconnectPolicies policy
var MediaconnectPolicies = map[string]Service{
	"ListEntitlements": {
		Method:        "GET",
		ServiceSuffix: "v1/entitlements",
		Permission:    "ListEntitlements",
	},
	"ListFlows": {
		Method:        "GET",
		ServiceSuffix: "v1/flows",
		Permission:    "ListFlows",
	},
	"ListOfferings": {
		Method:        "GET",
		ServiceSuffix: "v1/offerings",
		Permission:    "ListOfferings",
	},
	"ListReservations": {
		Method:        "GET",
		ServiceSuffix: "v1/reservations",
		Permission:    "ListReservations",
	},

	// extra
	"DescribeFlow": {
		ServiceSuffix:          "/v1/flows/{{.flow_arn}}",
		Permission:             "DescribeFlow",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "flow_arn",
	},
	"DescribeOffering": {
		ServiceSuffix:          "/v1/offerings/{{.offering_arn}}",
		Permission:             "DescribeOffering",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "offering_arn",
	},
	"DescribeReservation": {
		ServiceSuffix:          "/v1/reservations/{{.reservation_arn}}",
		Permission:             "DescribeReservation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "reservation_arn",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
