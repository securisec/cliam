package policy

import "github.com/securisec/cliam/shared"

var WellarchitectedPolicies = map[string]Service{
	"ListLenses": {
		Method:        "GET",
		ServiceSuffix: "lenses",
		Permission:    "ListLenses",
	},
	"ListNotifications": {
		Method:        "POST",
		ServiceSuffix: "notifications",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListNotifications",
	},
	"ListShareInvitations": {
		Method:        "GET",
		ServiceSuffix: "shareInvitations",
		Permission:    "ListShareInvitations",
	},
	"ListWorkloads": {
		Method:        "POST",
		ServiceSuffix: "workloadsSummaries",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorkloads",
	},

	// extra
	"GetLens": {
		ServiceSuffix:          "/lenses/{{.lens_alias}}",
		Permission:             "GetLens",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "lens_alias",
	},
	"GetLensVersionDifference": {
		ServiceSuffix:          "/lenses/{{.lens_alias}}/versionDifference",
		Permission:             "GetLensVersionDifference",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "lens_alias",
	},
	"GetWorkload": {
		ServiceSuffix:          "/workloads/{{.workload_id}}",
		Permission:             "GetWorkload",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_id",
	},
	"ListLensReviews": {
		ServiceSuffix:          "/workloads/{{.workload_id}}/lensReviews",
		Permission:             "ListLensReviews",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_id",
	},
	"ListLensShares": {
		ServiceSuffix:          "/lenses/{{.lens_alias}}/shares",
		Permission:             "ListLensShares",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "lens_alias",
	},
	"ListMilestones": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListMilestones",
			"Version": "2020-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListMilestones",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkloadId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "workload_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.workload_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_arn",
	},
	"ListWorkloadShares": {
		ServiceSuffix:          "/workloads/{{.workload_id}}/shares",
		Permission:             "ListWorkloadShares",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_id",
	},
}
