package policy

import "github.com/securisec/cliam/shared"

var DirectconnectPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeConnections",
		},
		Permission: "DescribeConnections",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeCustomerMetadata",
		},
		Permission: "DescribeCustomerMetadata",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAssociationProposals",
		},
		Permission: "DescribeDirectConnectGatewayAssociationProposals",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAssociations",
		},
		Permission: "DescribeDirectConnectGatewayAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAttachments",
		},
		Permission: "DescribeDirectConnectGatewayAttachments",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGateways",
		},
		Permission: "DescribeDirectConnectGateways",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeInterconnects",
		},
		Permission: "DescribeInterconnects",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLags",
		},
		Permission: "DescribeLags",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLocations",
		},
		Permission: "DescribeLocations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeVirtualGateways",
		},
		Permission: "DescribeVirtualGateways",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeVirtualInterfaces",
		},
		Permission: "DescribeVirtualInterfaces",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.ListVirtualInterfaceTestHistory",
		},
		Permission: "ListVirtualInterfaceTestHistory",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeConnectionLoa",
		},
		Permission:             "DescribeConnectionLoa",
		IsExtra:                true,
		ExtraComponentBodyKey:  "connectionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "connection_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeConnectionsOnInterconnect",
		},
		Permission:             "DescribeConnectionsOnInterconnect",
		IsExtra:                true,
		ExtraComponentBodyKey:  "interconnectId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "interconnect_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeHostedConnections",
		},
		Permission:             "DescribeHostedConnections",
		IsExtra:                true,
		ExtraComponentBodyKey:  "connectionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "connection_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeInterconnectLoa",
		},
		Permission:             "DescribeInterconnectLoa",
		IsExtra:                true,
		ExtraComponentBodyKey:  "interconnectId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "interconnect_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLoa",
		},
		Permission:             "DescribeLoa",
		IsExtra:                true,
		ExtraComponentBodyKey:  "connectionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "connection_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeRouterConfiguration",
		},
		Permission:             "DescribeRouterConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "virtualInterfaceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "virtual_interface_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeTags",
		},
		Permission:             "DescribeTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arns",
	},
}
