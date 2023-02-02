package policy

import "github.com/securisec/cliam/shared"

// IOTWinMakerPolicies policy
var IOTWinMakerPolicies = map[string]Service{
	"GetPricingPlan": {
		Method:        "GET",
		ServiceSuffix: "pricingplan",
		Permission:    "GetPricingPlan",
	},
	"ListWorkspaces": {
		Method:        "POST",
		ServiceSuffix: "workspaces-list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorkspaces",
	},

	// extra
	"GetSyncJob": {
		ServiceSuffix:          "/sync-jobs/{{.sync_source}}",
		Permission:             "GetSyncJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "sync_source",
	},
	"GetWorkspace": {
		ServiceSuffix:          "/workspaces/{{.workspace_id}}",
		Permission:             "GetWorkspace",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListComponentTypes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListComponentTypes",
			"Version": "2021-11-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListComponentTypes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "workspaceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListEntities": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListEntities",
			"Version": "2021-11-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListEntities",
		IsExtra:                true,
		ExtraComponentBodyKey:  "workspaceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListScenes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListScenes",
			"Version": "2021-11-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListScenes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "workspaceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListSyncJobs": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListSyncJobs",
			"Version": "2021-11-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListSyncJobs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "workspaceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2021-11-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceARN",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
