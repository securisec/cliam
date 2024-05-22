package policy

import "github.com/securisec/cliam/shared"

// WorkdocsPolicies policy
var WorkdocsPolicies = map[string]Service{
	"DescribeActivities": {
		Method:                 "GET",
		ServiceSuffix:          "api/v1/activities?organizationId={{.organization_id}}",
		Permission:             "DescribeActivities",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "organization_id",
	},
	"DescribeUsers": {
		Method:                 "GET",
		ServiceSuffix:          "api/v1/users?organizationId={{.organization_id}}",
		Permission:             "DescribeUsers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "organization_id",
	},
	"GetResources": {
		Method:        "GET",
		ServiceSuffix: "api/v1/resources",
		Permission:    "GetResources",
	},
	"InitiateDocumentVersionUpload": {
		Method:        "POST",
		ServiceSuffix: "api/v1/documents",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "InitiateDocumentVersionUpload",
	},
	"SearchResources": {
		Method:        "POST",
		ServiceSuffix: "api/v1/search",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "SearchResources",
	},

	// extra
	"DescribeDocumentVersions": {
		ServiceSuffix:          "/api/v1/documents/{{.document_id}}/versions",
		Permission:             "DescribeDocumentVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "document_id",
	},
	"DescribeFolderContents": {
		ServiceSuffix:          "/api/v1/folders/{{.folder_id}}/contents",
		Permission:             "DescribeFolderContents",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "folder_id",
	},
	"DescribeGroups": {
		ServiceSuffix:          "/api/v1/groups",
		Permission:             "DescribeGroups",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "search_query",
	},
	"DescribeNotificationSubscriptions": {
		ServiceSuffix:          "/api/v1/organizations/{{.organization_id}}/subscriptions",
		Permission:             "DescribeNotificationSubscriptions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "organization_id",
	},
	"DescribeResourcePermissions": {
		ServiceSuffix:          "/api/v1/resources/{{.resource_id}}/permissions",
		Permission:             "DescribeResourcePermissions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_id",
	},
	"DescribeRootFolders": {
		ServiceSuffix:          "/api/v1/me/root",
		Permission:             "DescribeRootFolders",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "authentication_token",
	},
	"GetCurrentUser": {
		ServiceSuffix:          "/api/v1/me",
		Permission:             "GetCurrentUser",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "authentication_token",
	},
	"GetDocument": {
		ServiceSuffix:          "/api/v1/documents/{{.document_id}}",
		Permission:             "GetDocument",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "document_id",
	},
	"GetDocumentPath": {
		ServiceSuffix:          "/api/v1/documents/{{.document_id}}/path",
		Permission:             "GetDocumentPath",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "document_id",
	},
	"GetFolder": {
		ServiceSuffix:          "/api/v1/folders/{{.folder_id}}",
		Permission:             "GetFolder",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "folder_id",
	},
	"GetFolderPath": {
		ServiceSuffix:          "/api/v1/folders/{{.folder_id}}/path",
		Permission:             "GetFolderPath",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "folder_id",
	},
}
