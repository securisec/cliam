package policy

import "github.com/securisec/cliam/shared"

var WellarchitectedPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "lenses",
		Permission:    "ListLenses",
	},
	{
		Method:        "POST",
		ServiceSuffix: "notifications",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListNotifications",
	},
	{
		Method:        "GET",
		ServiceSuffix: "shareInvitations",
		Permission:    "ListShareInvitations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "workloadsSummaries",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorkloads",
	},

	// extra
	{
		ServiceSuffix:          "/lenses/{{.lens_alias}}",
		Permission:             "GetLens",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "lens_alias",
	},
	{
		ServiceSuffix:          "/lenses/{{.lens_alias}}/versionDifference",
		Permission:             "GetLensVersionDifference",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "lens_alias",
	},
	{
		ServiceSuffix:          "/workloads/{{.workload_id}}",
		Permission:             "GetWorkload",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_id",
	},
	{
		ServiceSuffix:          "/workloads/{{.workload_id}}/lensReviews",
		Permission:             "ListLensReviews",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_id",
	},
	{
		ServiceSuffix:          "/lenses/{{.lens_alias}}/shares",
		Permission:             "ListLensShares",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "lens_alias",
	},
	{
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
	{
		ServiceSuffix:          "/tags/{{.workload_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_arn",
	},
	{
		ServiceSuffix:          "/workloads/{{.workload_id}}/shares",
		Permission:             "ListWorkloadShares",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workload_id",
	},
}
