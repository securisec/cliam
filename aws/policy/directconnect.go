package policy

import "github.com/securisec/cliam/shared"

var DirectconnectPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeConnections",
		},
		Permission: "DescribeConnections",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeCustomerMetadata",
		},
		Permission: "DescribeCustomerMetadata",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAssociationProposals",
		},
		Permission: "DescribeDirectConnectGatewayAssociationProposals",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAssociations",
		},
		Permission: "DescribeDirectConnectGatewayAssociations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAttachments",
		},
		Permission: "DescribeDirectConnectGatewayAttachments",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGateways",
		},
		Permission: "DescribeDirectConnectGateways",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeInterconnects",
		},
		Permission: "DescribeInterconnects",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLags",
		},
		Permission: "DescribeLags",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLocations",
		},
		Permission: "DescribeLocations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeVirtualGateways",
		},
		Permission: "DescribeVirtualGateways",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeVirtualInterfaces",
		},
		Permission: "DescribeVirtualInterfaces",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "OvertureService.ListVirtualInterfaceTestHistory",
		},
		Permission: "ListVirtualInterfaceTestHistory",
	},
}
