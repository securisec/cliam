package policy

// Microsoft_Quota_quota policy
var Microsoft_Quota_quota = map[string]Policy{
	"Quota_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Quota/quotaLimits/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "Quota_Get",
		Resource:    "Microsoft.Quota",
	},
	"Quota_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Quota/quotaLimits",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "Quota_List",
		Resource:    "Microsoft.Quota",
	},
	"QuotaRequestStatus_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Quota/quotaLimitsRequests/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "QuotaRequestStatus_Get",
		Resource:    "Microsoft.Quota",
	},
	"QuotaRequestStatus_List": {
		Path:   "/{{.scope}}/providers/Microsoft.Quota/quotaLimitsRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "QuotaRequestStatus_List",
		Resource:    "Microsoft.Quota",
	},
	"QuotaResourceProviders_List": {
		Path:   "/providers/Microsoft.Quota/quotaLimitProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "QuotaResourceProviders_List",
		Resource:    "Microsoft.Quota",
	},
	"Operation_List": {
		Path:   "/providers/Microsoft.Quota/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "Operation_List",
		Resource:    "Microsoft.Quota",
	},
}
