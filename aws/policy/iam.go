package policy

import "github.com/securisec/cliam/shared"

// IAMPolicies iam permissions
var IAMPolicies = map[string]Service{
	"ListRoles": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListRoles&Version=2010-05-08",
		Permission:    "ListRoles",
		ResponseParser: &ResponseParser{
			ResponseFormat: "json",
			KeysToExtract: []CommandLineFlagMap{
				{Flag: "role_name", ResponseKey: "RoleName"},
				{Flag: "role_name", ResponseKey: "RoleName"},
			},
			ObjectPath: []string{"ListRolesResponse", "ListRolesResult", "Roles"},
		},
	},
	"GetAccountAuthorizationDetails": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountAuthorizationDetails&Version=2010-05-08",
		Permission:    "GetAccountAuthorizationDetails",
	},
	"GetAccountPasswordPolicy": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountPasswordPolicy&Version=2010-05-08",
		Permission:    "GetAccountPasswordPolicy",
	},
	"GetAccountSummary": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountSummary&Version=2010-05-08",
		Permission:    "GetAccountSummary",
	},
	"GetCredentialReport": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetCredentialReport&Version=2010-05-08",
		Permission:    "GetCredentialReport",
	},
	"GetUser": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetUser&Version=2010-05-08",
		Permission:    "GetUser",
	},
	"ListAccessKeys": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListAccessKeys&Version=2010-05-08",
		Permission:    "ListAccessKeys",
	},
	"ListAccountAliases": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListAccountAliases&Version=2010-05-08",
		Permission:    "ListAccountAliases",
	},
	"ListGroups": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListGroups&Version=2010-05-08",
		Permission:    "ListGroups",
	},
	"ListInstanceProfiles": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListInstanceProfiles&Version=2010-05-08",
		Permission:    "ListInstanceProfiles",
	},
	"ListMFADevices": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListMFADevices&Version=2010-05-08",
		Permission:    "ListMFADevices",
	},
	"ListOpenIDConnectProviders": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListOpenIDConnectProviders&Version=2010-05-08",
		Permission:    "ListOpenIDConnectProviders",
	},
	"ListPolicies": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListPolicies&Version=2010-05-08",
		Permission:    "ListPolicies",
	},
	"ListSamlProviders": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListSamlProviders&Version=2010-05-08",
		Permission:    "ListSamlProviders",
	},
	"ListServerCertificates": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListServerCertificates&Version=2010-05-08",
		Permission:    "ListServerCertificates",
	},
	"ListServiceSpecificCredentials": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListServiceSpecificCredentials&Version=2010-05-08",
		Permission:    "ListServiceSpecificCredentials",
	},
	"ListSigningCertificates": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListSigningCertificates&Version=2010-05-08",
		Permission:    "ListSigningCertificates",
	},
	"ListSshPublicKeys": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListSshPublicKeys&Version=2010-05-08",
		Permission:    "ListSshPublicKeys",
	},
	"ListUsers": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListUsers&Version=2010-05-08",
		Permission:    "ListUsers",
	},
	"ListVirtualMfaDevices": {
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListVirtualMfaDevices&Version=2010-05-08",
		Permission:    "ListVirtualMfaDevices",
	},

	// extra
	"GetAccessKeyLastUsed": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetAccessKeyLastUsed",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetAccessKeyLastUsed",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccessKeyId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "access_key_id",
	},
	"GetContextKeysForCustomPolicy": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetContextKeysForCustomPolicy",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetContextKeysForCustomPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyInputList",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "policy_input_list",
	},
	"GetContextKeysForPrincipalPolicy": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetContextKeysForPrincipalPolicy",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetContextKeysForPrincipalPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicySourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "policy_source_arn",
	},
	"GetGroup": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetGroup",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "group_name",
	},
	"GetInstanceProfile": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetInstanceProfile",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetInstanceProfile",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceProfileName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_profile_name",
	},
	"GetLoginProfile": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetLoginProfile",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetLoginProfile",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "user_name",
	},
	"GetOpenIDConnectProvider": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetOpenIDConnectProvider",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetOpenIDConnectProvider",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OpenIDConnectProviderArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "open_i_d_connect_provider_arn",
	},
	"GetOrganizationsAccessReport": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetOrganizationsAccessReport",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetOrganizationsAccessReport",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "job_id",
	},
	"GetPolicy": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetPolicy",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "policy_arn",
	},
	"GetRole": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetRole",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetRole",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RoleName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "role_name",
	},
	"GetServerCertificate": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetServerCertificate",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetServerCertificate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServerCertificateName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "server_certificate_name",
	},
	"GetServiceLastAccessedDetails": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetServiceLastAccessedDetails",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetServiceLastAccessedDetails",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "job_id",
	},
	"GetServiceLinkedRoleDeletionStatus": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "GetServiceLinkedRoleDeletionStatus",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetServiceLinkedRoleDeletionStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DeletionTaskId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "deletion_task_id",
	},
	"ListAttachedGroupPolicies": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListAttachedGroupPolicies",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListAttachedGroupPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "group_name",
	},
	"ListAttachedRolePolicies": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListAttachedRolePolicies",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListAttachedRolePolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RoleName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "role_name",
		ResponseParser: &ResponseParser{
			ResponseFormat: "json",
			KeysToExtract: []CommandLineFlagMap{
				{Flag: "policy_arn", ResponseKey: "PolicyArn"},
			},
			ObjectPath: []string{"ListAttachedRolePoliciesResponse", "ListAttachedRolePoliciesResult", "AttachedPolicies"},
		},
	},
	"ListAttachedUserPolicies": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListAttachedUserPolicies",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListAttachedUserPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "user_name",
	},
	"ListEntitiesForPolicy": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListEntitiesForPolicy",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListEntitiesForPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "policy_arn",
	},
	"ListGroupPolicies": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListGroupPolicies",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListGroupPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "group_name",
	},
	"ListGroupsForUser": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListGroupsForUser",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListGroupsForUser",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "user_name",
	},
	"ListInstanceProfileTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListInstanceProfileTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListInstanceProfileTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceProfileName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_profile_name",
	},
	"ListInstanceProfilesForRole": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListInstanceProfilesForRole",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListInstanceProfilesForRole",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RoleName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "role_name",
	},
	"ListMFADeviceTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListMFADeviceTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListMFADeviceTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SerialNumber",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "serial_number",
	},
	"ListOpenIDConnectProviderTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListOpenIDConnectProviderTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListOpenIDConnectProviderTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OpenIDConnectProviderArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "open_i_d_connect_provider_arn",
	},
	"ListPolicyTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListPolicyTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListPolicyTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "policy_arn",
	},
	"ListPolicyVersions": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListPolicyVersions",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListPolicyVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PolicyArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "policy_arn",
	},
	"ListRolePolicies": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListRolePolicies",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListRolePolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RoleName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "role_name",
	},
	"ListRoleTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListRoleTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListRoleTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RoleName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "role_name",
	},
	"ListSAMLProviderTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListSAMLProviderTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListSAMLProviderTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SAMLProviderArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "s_a_m_l_provider_arn",
	},
	"ListServerCertificateTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListServerCertificateTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListServerCertificateTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServerCertificateName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "server_certificate_name",
	},
	"ListUserPolicies": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListUserPolicies",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListUserPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "user_name",
	},
	"ListUserTags": {
		Method:       "POST",
		IgnoreRegion: true,
		FormData: map[string]string{
			"Action":  "ListUserTags",
			"Version": "2010-05-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListUserTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "user_name",
	},
}
