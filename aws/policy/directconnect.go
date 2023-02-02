package policy

import "github.com/securisec/cliam/shared"

// DirectConnectPolicies policy
var DirectConnectPolicies = map[string]Service{
	"ConfirmCustomerAgreement": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.ConfirmCustomerAgreement",
		},
		Permission: "ConfirmCustomerAgreement",
	},
	"DeleteBGPPeer": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DeleteBGPPeer",
		},
		Permission: "DeleteBGPPeer",
	},
	"DeleteDirectConnectGatewayAssociation": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DeleteDirectConnectGatewayAssociation",
		},
		Permission: "DeleteDirectConnectGatewayAssociation",
	},
	"DescribeConnections": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeConnections",
		},
		Permission: "DescribeConnections",
	},
	"DescribeCustomerMetadata": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeCustomerMetadata",
		},
		Permission: "DescribeCustomerMetadata",
	},
	"DescribeDirectConnectGatewayAssociationProposals": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAssociationProposals",
		},
		Permission: "DescribeDirectConnectGatewayAssociationProposals",
	},
	"DescribeDirectConnectGatewayAssociations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAssociations",
		},
		Permission: "DescribeDirectConnectGatewayAssociations",
	},
	"DescribeDirectConnectGatewayAttachments": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGatewayAttachments",
		},
		Permission: "DescribeDirectConnectGatewayAttachments",
	},
	"DescribeDirectConnectGateways": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeDirectConnectGateways",
		},
		Permission: "DescribeDirectConnectGateways",
	},
	"DescribeInterconnects": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeInterconnects",
		},
		Permission: "DescribeInterconnects",
	},
	"DescribeLags": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLags",
		},
		Permission: "DescribeLags",
	},
	"DescribeLocations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeLocations",
		},
		Permission: "DescribeLocations",
	},
	"DescribeVirtualGateways": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeVirtualGateways",
		},
		Permission: "DescribeVirtualGateways",
	},
	"DescribeVirtualInterfaces": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.DescribeVirtualInterfaces",
		},
		Permission: "DescribeVirtualInterfaces",
	},
	"ListVirtualInterfaceTestHistory": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.ListVirtualInterfaceTestHistory",
		},
		Permission: "ListVirtualInterfaceTestHistory",
	},
	"UpdateDirectConnectGatewayAssociation": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "OvertureService.UpdateDirectConnectGatewayAssociation",
		},
		Permission: "UpdateDirectConnectGatewayAssociation",
	},

	// extra
	"DescribeConnectionLoa": {
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
	"DescribeConnectionsOnInterconnect": {
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
	"DescribeHostedConnections": {
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
	"DescribeInterconnectLoa": {
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
	"DescribeLoa": {
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
	"DescribeRouterConfiguration": {
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
	"DescribeTags": {
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
