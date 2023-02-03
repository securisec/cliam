package policy

// Microsoft_ApiManagement_apimapis policy
var Microsoft_ApiManagement_apimapis = map[string]Policy{
	"Api_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Api_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Api_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Api_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiRevision_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/revisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiRevision_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiRelease_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/releases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiRelease_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiRelease_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/releases/{{.releaseId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiRelease_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiOperation_ListByApi": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiOperation_ListByApi",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiOperation_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiOperation_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiOperationPolicy_ListByOperation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operations/{{.operationId}}/policies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiOperationPolicy_ListByOperation",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiOperationPolicy_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operations/{{.operationId}}/policies/{{.policyId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiOperationPolicy_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_ListByOperation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operations/{{.operationId}}/tags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_ListByOperation",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_GetByOperation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operations/{{.operationId}}/tags/{{.tagId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_GetByOperation",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiProduct_ListByApis": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/products",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiProduct_ListByApis",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiPolicy_ListByApi": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/policies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiPolicy_ListByApi",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiPolicy_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/policies/{{.policyId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiPolicy_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiSchema_ListByApi": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/schemas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiSchema_ListByApi",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiSchema_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/schemas/{{.schemaId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiSchema_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiDiagnostic_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/diagnostics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiDiagnostic_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiDiagnostic_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/diagnostics/{{.diagnosticId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiDiagnostic_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiIssue_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/issues",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiIssue_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiIssue_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/issues/{{.issueId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiIssue_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiIssueComment_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/issues/{{.issueId}}/comments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiIssueComment_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiIssueComment_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/issues/{{.issueId}}/comments/{{.commentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiIssueComment_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiIssueAttachment_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/issues/{{.issueId}}/attachments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiIssueAttachment_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiIssueAttachment_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/issues/{{.issueId}}/attachments/{{.attachmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiIssueAttachment_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiTagDescription_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/tagDescriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiTagDescription_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiTagDescription_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/tagDescriptions/{{.tagDescriptionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiTagDescription_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_ListByApi": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/tags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_ListByApi",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_GetByApi": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/tags/{{.tagId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_GetByApi",
		Resource:    "Microsoft.ApiManagement",
	},
	"Operation_ListByTags": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apis/{{.apiId}}/operationsByTags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Operation_ListByTags",
		Resource:    "Microsoft.ApiManagement",
	},
}
