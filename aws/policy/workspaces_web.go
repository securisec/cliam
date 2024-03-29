package policy

// WorkspacesWebPolicies policy
var WorkspacesWebPolicies = map[string]Service{
	"ListBrowserSettings": {
		Method:        "GET",
		ServiceSuffix: "browserSettings",
		Permission:    "ListBrowserSettings",
	},
	"ListNetworkSettings": {
		Method:        "GET",
		ServiceSuffix: "networkSettings",
		Permission:    "ListNetworkSettings",
	},
	"ListPortals": {
		Method:        "GET",
		ServiceSuffix: "portals",
		Permission:    "ListPortals",
	},
	"ListTrustStores": {
		Method:        "GET",
		ServiceSuffix: "trustStores",
		Permission:    "ListTrustStores",
	},
	"ListUserAccessLoggingSettings": {
		Method:        "GET",
		ServiceSuffix: "userAccessLoggingSettings",
		Permission:    "ListUserAccessLoggingSettings",
	},
	"ListUserSettings": {
		Method:        "GET",
		ServiceSuffix: "userSettings",
		Permission:    "ListUserSettings",
	},

	// extra
	"GetBrowserSettings": {
		ServiceSuffix:          "/browserSettings/{{.browser_settings_arn}}",
		Permission:             "GetBrowserSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "browser_settings_arn",
	},
	"GetIdentityProvider": {
		ServiceSuffix:          "/identityProviders/{{.identity_provider_arn}}",
		Permission:             "GetIdentityProvider",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "identity_provider_arn",
	},
	"GetNetworkSettings": {
		ServiceSuffix:          "/networkSettings/{{.network_settings_arn}}",
		Permission:             "GetNetworkSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "network_settings_arn",
	},
	"GetPortal": {
		ServiceSuffix:          "/portals/{{.portal_arn}}",
		Permission:             "GetPortal",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "portal_arn",
	},
	"GetPortalServiceProviderMetadata": {
		ServiceSuffix:          "/portalIdp/{{.portal_arn}}",
		Permission:             "GetPortalServiceProviderMetadata",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "portal_arn",
	},
	"GetTrustStore": {
		ServiceSuffix:          "/trustStores/{{.trust_store_arn}}",
		Permission:             "GetTrustStore",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "trust_store_arn",
	},
	"GetUserAccessLoggingSettings": {
		ServiceSuffix:          "/userAccessLoggingSettings/{{.user_access_logging_settings_arn}}",
		Permission:             "GetUserAccessLoggingSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "user_access_logging_settings_arn",
	},
	"GetUserSettings": {
		ServiceSuffix:          "/userSettings/{{.user_settings_arn}}",
		Permission:             "GetUserSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "user_settings_arn",
	},
	"ListIdentityProviders": {
		ServiceSuffix:          "/portals/{{.portal_arn}}/identityProviders",
		Permission:             "ListIdentityProviders",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "portal_arn",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTrustStoreCertificates": {
		ServiceSuffix:          "/trustStores/{{.trust_store_arn}}/certificates",
		Permission:             "ListTrustStoreCertificates",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "trust_store_arn",
	},
}
