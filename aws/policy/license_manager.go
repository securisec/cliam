package policy

import "github.com/securisec/cliam/shared"

var LicenseManagerPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetServiceSettings",
		},
		Permission: "GetServiceSettings",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListDistributedGrants",
		},
		Permission: "ListDistributedGrants",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListLicenseConfigurations",
		},
		Permission: "ListLicenseConfigurations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListLicenseConversionTasks",
		},
		Permission: "ListLicenseConversionTasks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListLicenseManagerReportGenerators",
		},
		Permission: "ListLicenseManagerReportGenerators",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListLicenses",
		},
		Permission: "ListLicenses",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListReceivedGrants",
		},
		Permission: "ListReceivedGrants",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListReceivedLicenses",
		},
		Permission: "ListReceivedLicenses",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListResourceInventory",
		},
		Permission: "ListResourceInventory",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListTokens",
		},
		Permission: "ListTokens",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.UpdateServiceSettings",
		},
		Permission: "UpdateServiceSettings",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetAccessToken",
		},
		Permission:             "GetAccessToken",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Token",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "token",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetGrant",
		},
		Permission:             "GetGrant",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GrantArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "grant_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetLicense",
		},
		Permission:             "GetLicense",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetLicenseConfiguration",
		},
		Permission:             "GetLicenseConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseConfigurationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_configuration_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetLicenseConversionTask",
		},
		Permission:             "GetLicenseConversionTask",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseConversionTaskId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_conversion_task_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetLicenseManagerReportGenerator",
		},
		Permission:             "GetLicenseManagerReportGenerator",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseManagerReportGeneratorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_manager_report_generator_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.GetLicenseUsage",
		},
		Permission:             "GetLicenseUsage",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListAssociationsForLicenseConfiguration",
		},
		Permission:             "ListAssociationsForLicenseConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseConfigurationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_configuration_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListFailuresForLicenseConfigurationOperations",
		},
		Permission:             "ListFailuresForLicenseConfigurationOperations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseConfigurationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_configuration_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListLicenseSpecificationsForResource",
		},
		Permission:             "ListLicenseSpecificationsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListLicenseVersions",
		},
		Permission:             "ListLicenseVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSLicenseManager.ListUsageForLicenseConfiguration",
		},
		Permission:             "ListUsageForLicenseConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LicenseConfigurationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "license_configuration_arn",
	},
}
