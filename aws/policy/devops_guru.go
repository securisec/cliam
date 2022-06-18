package policy

import "github.com/securisec/cliam/shared"

var DevopsGuruPolicies = map[string]Service{
	"DescribeAccountHealth": {
		Method:        "GET",
		ServiceSuffix: "accounts/health",
		Permission:    "DescribeAccountHealth",
	},
	"DescribeEventSourcesConfig": {
		Method:        "POST",
		ServiceSuffix: "event-sources",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeEventSourcesConfig",
	},
	"DescribeFeedback": {
		Method:        "POST",
		ServiceSuffix: "feedback",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeFeedback",
	},
	"DescribeOrganizationHealth": {
		Method:        "POST",
		ServiceSuffix: "organization/health",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOrganizationHealth",
	},
	"DescribeServiceIntegration": {
		Method:        "GET",
		ServiceSuffix: "service-integrations",
		Permission:    "DescribeServiceIntegration",
	},
	"GetCostEstimation": {
		Method:        "GET",
		ServiceSuffix: "cost-estimation",
		Permission:    "GetCostEstimation",
	},
	"ListNotificationChannels": {
		Method:        "POST",
		ServiceSuffix: "channels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListNotificationChannels",
	},

	// extra
	"DescribeAccountOverview": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountOverview",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeAccountOverview",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FromTime",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "from_time",
	},
	"DescribeAnomaly": {
		ServiceSuffix:          "/anomalies/{{.id}}",
		Permission:             "DescribeAnomaly",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"DescribeInsight": {
		ServiceSuffix:          "/insights/{{.id}}",
		Permission:             "DescribeInsight",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"DescribeOrganizationOverview": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeOrganizationOverview",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeOrganizationOverview",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FromTime",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "from_time",
	},
	"DescribeOrganizationResourceCollectionHealth": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeOrganizationResourceCollectionHealth",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeOrganizationResourceCollectionHealth",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OrganizationResourceCollectionType",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "organization_resource_collection_type",
	},
	"DescribeResourceCollectionHealth": {
		ServiceSuffix:          "/accounts/health/resource-collection/{{.resource_collection_type}}",
		Permission:             "DescribeResourceCollectionHealth",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_collection_type",
	},
	"GetResourceCollection": {
		ServiceSuffix:          "/resource-collections/{{.resource_collection_type}}",
		Permission:             "GetResourceCollection",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_collection_type",
	},
	"ListAnomaliesForInsight": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListAnomaliesForInsight",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListAnomaliesForInsight",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InsightId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "insight_id",
	},
	"ListEvents": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListEvents",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListEvents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Filters",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "filters",
	},
	"ListInsights": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListInsights",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListInsights",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StatusFilter",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "status_filter",
	},
	"ListOrganizationInsights": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListOrganizationInsights",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListOrganizationInsights",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StatusFilter",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "status_filter",
	},
	"ListRecommendations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListRecommendations",
			"Version": "2020-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListRecommendations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InsightId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "insight_id",
	},
}
