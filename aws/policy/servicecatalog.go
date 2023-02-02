package policy

import "github.com/securisec/cliam/shared"

// ServiceCatalogPolicies policy
var ServiceCatalogPolicies = map[string]Service{
	"DescribeProduct": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProduct",
		},
		Permission: "DescribeProduct",
	},
	"DescribeProductAsAdmin": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProductAsAdmin",
		},
		Permission: "DescribeProductAsAdmin",
	},
	"DescribeProvisionedProduct": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisionedProduct",
		},
		Permission: "DescribeProvisionedProduct",
	},
	"DescribeProvisioningArtifact": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisioningArtifact",
		},
		Permission: "DescribeProvisioningArtifact",
	},
	"DescribeProvisioningParameters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DescribeProvisioningParameters",
		},
		Permission: "DescribeProvisioningParameters",
	},
	"DisableAWSOrganizationsAccess": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.DisableAWSOrganizationsAccess",
		},
		Permission: "DisableAWSOrganizationsAccess",
	},
	"EnableAWSOrganizationsAccess": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.EnableAWSOrganizationsAccess",
		},
		Permission: "EnableAWSOrganizationsAccess",
	},
	"GetAWSOrganizationsAccessStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.GetAWSOrganizationsAccessStatus",
		},
		Permission: "GetAWSOrganizationsAccessStatus",
	},
	"GetProvisionedProductOutputs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.GetProvisionedProductOutputs",
		},
		Permission: "GetProvisionedProductOutputs",
	},
	"ListAcceptedPortfolioShares": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListAcceptedPortfolioShares",
		},
		Permission: "ListAcceptedPortfolioShares",
	},
	"ListPortfolios": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListPortfolios",
		},
		Permission: "ListPortfolios",
	},
	"ListProvisionedProductPlans": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListProvisionedProductPlans",
		},
		Permission: "ListProvisionedProductPlans",
	},
	"ListRecordHistory": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListRecordHistory",
		},
		Permission: "ListRecordHistory",
	},
	"ListServiceActions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListServiceActions",
		},
		Permission: "ListServiceActions",
	},
	"ListTagOptions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ListTagOptions",
		},
		Permission: "ListTagOptions",
	},
	"ScanProvisionedProducts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.ScanProvisionedProducts",
		},
		Permission: "ScanProvisionedProducts",
	},
	"SearchProducts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.SearchProducts",
		},
		Permission: "SearchProducts",
	},
	"SearchProductsAsAdmin": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.SearchProductsAsAdmin",
		},
		Permission: "SearchProductsAsAdmin",
	},
	"SearchProvisionedProducts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWS242ServiceCatalogService.SearchProvisionedProducts",
		},
		Permission: "SearchProvisionedProducts",
	},

	// extra
	"DescribeConstraint": {
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
	"DescribeCopyProductStatus": {
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
	"DescribePortfolio": {
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
	"DescribePortfolioShareStatus": {
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
	"DescribeProductView": {
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
	"DescribeProvisionedProductPlan": {
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
	"DescribeRecord": {
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
	"DescribeServiceAction": {
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
	"DescribeTagOption": {
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
	"ListBudgetsForResource": {
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
	"ListConstraintsForPortfolio": {
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
	"ListLaunchPaths": {
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
	"ListPortfolioAccess": {
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
	"ListPortfoliosForProduct": {
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
	"ListPrincipalsForPortfolio": {
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
	"ListProvisioningArtifacts": {
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
	"ListProvisioningArtifactsForServiceAction": {
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
	"ListResourcesForTagOption": {
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
	"ListStackInstancesForProvisionedProduct": {
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
