package policy

import "github.com/securisec/cliam/shared"

var ServiceCatalogPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProduct",
		},
		Permission: "DescribeProduct",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProductAsAdmin",
		},
		Permission: "DescribeProductAsAdmin",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisionedProduct",
		},
		Permission: "DescribeProvisionedProduct",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisioningArtifact",
		},
		Permission: "DescribeProvisioningArtifact",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisioningParameters",
		},
		Permission: "DescribeProvisioningParameters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DisableAWSOrganizationsAccess",
		},
		Permission: "DisableAWSOrganizationsAccess",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.EnableAWSOrganizationsAccess",
		},
		Permission: "EnableAWSOrganizationsAccess",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.GetAWSOrganizationsAccessStatus",
		},
		Permission: "GetAWSOrganizationsAccessStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.GetProvisionedProductOutputs",
		},
		Permission: "GetProvisionedProductOutputs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListAcceptedPortfolioShares",
		},
		Permission: "ListAcceptedPortfolioShares",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListPortfolios",
		},
		Permission: "ListPortfolios",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListProvisionedProductPlans",
		},
		Permission: "ListProvisionedProductPlans",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListRecordHistory",
		},
		Permission: "ListRecordHistory",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListServiceActions",
		},
		Permission: "ListServiceActions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListTagOptions",
		},
		Permission: "ListTagOptions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ScanProvisionedProducts",
		},
		Permission: "ScanProvisionedProducts",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.SearchProducts",
		},
		Permission: "SearchProducts",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.SearchProductsAsAdmin",
		},
		Permission: "SearchProductsAsAdmin",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.SearchProvisionedProducts",
		},
		Permission: "SearchProvisionedProducts",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeConstraint",
		},
		Permission:             "DescribeConstraint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeCopyProductStatus",
		},
		Permission:             "DescribeCopyProductStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CopyProductToken",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "copy_product_token",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribePortfolio",
		},
		Permission:             "DescribePortfolio",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribePortfolioShareStatus",
		},
		Permission:             "DescribePortfolioShareStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PortfolioShareToken",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "portfolio_share_token",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProductView",
		},
		Permission:             "DescribeProductView",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisionedProductPlan",
		},
		Permission:             "DescribeProvisionedProductPlan",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PlanId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "plan_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeRecord",
		},
		Permission:             "DescribeRecord",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeServiceAction",
		},
		Permission:             "DescribeServiceAction",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeTagOption",
		},
		Permission:             "DescribeTagOption",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListBudgetsForResource",
		},
		Permission:             "ListBudgetsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListConstraintsForPortfolio",
		},
		Permission:             "ListConstraintsForPortfolio",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PortfolioId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "portfolio_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListLaunchPaths",
		},
		Permission:             "ListLaunchPaths",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProductId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "product_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListPortfolioAccess",
		},
		Permission:             "ListPortfolioAccess",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PortfolioId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "portfolio_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListPortfoliosForProduct",
		},
		Permission:             "ListPortfoliosForProduct",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProductId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "product_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListPrincipalsForPortfolio",
		},
		Permission:             "ListPrincipalsForPortfolio",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PortfolioId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "portfolio_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListProvisioningArtifacts",
		},
		Permission:             "ListProvisioningArtifacts",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProductId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "product_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListProvisioningArtifactsForServiceAction",
		},
		Permission:             "ListProvisioningArtifactsForServiceAction",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceActionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_action_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListResourcesForTagOption",
		},
		Permission:             "ListResourcesForTagOption",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TagOptionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "tag_option_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListStackInstancesForProvisionedProduct",
		},
		Permission:             "ListStackInstancesForProvisionedProduct",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ProvisionedProductId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "provisioned_product_id",
	},
}
