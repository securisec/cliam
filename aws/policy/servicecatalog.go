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
}
