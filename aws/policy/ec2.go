package policy

import "github.com/securisec/cliam/shared"

var EC2Policies = []Service{
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AcceptTransitGatewayMulticastDomainAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AcceptTransitGatewayMulticastDomainAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AcceptVpcPeeringConnection",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AcceptVpcPeeringConnection",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AllocateAddress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AllocateAddress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AssociateAddress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AssociateAddress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AssociateEnclaveCertificateIamRole",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AssociateEnclaveCertificateIamRole",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AssociateTransitGatewayMulticastDomain",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AssociateTransitGatewayMulticastDomain",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "AuthorizeSecurityGroupIngress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "AuthorizeSecurityGroupIngress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CancelImportTask",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CancelImportTask",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateDefaultVpc",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateDefaultVpc",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateInstanceEventWindow",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateInstanceEventWindow",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateInternetGateway",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateInternetGateway",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateIpam",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateIpam",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreatePlacementGroup",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreatePlacementGroup",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreatePublicIpv4Pool",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreatePublicIpv4Pool",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateTrafficMirrorFilter",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateTrafficMirrorFilter",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateTrafficMirrorTarget",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateTrafficMirrorTarget",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateTransitGateway",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateTransitGateway",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateVpc",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateVpc",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateVpcEndpointServiceConfiguration",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateVpcEndpointServiceConfiguration",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateVpcPeeringConnection",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateVpcPeeringConnection",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeleteKeyPair",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeleteKeyPair",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeleteLaunchTemplate",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeleteLaunchTemplate",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeleteSecurityGroup",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeleteSecurityGroup",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeleteSpotDatafeedSubscription",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeleteSpotDatafeedSubscription",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeregisterInstanceEventNotificationAttributes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeregisterInstanceEventNotificationAttributes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeregisterTransitGatewayMulticastGroupMembers",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeregisterTransitGatewayMulticastGroupMembers",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DeregisterTransitGatewayMulticastGroupSources",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DeregisterTransitGatewayMulticastGroupSources",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountAttributes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountAttributes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAddresses",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAddresses",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAddressesAttribute",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAddressesAttribute",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAggregateIdFormat",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAggregateIdFormat",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAvailabilityZones",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAvailabilityZones",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeBundleTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeBundleTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCapacityReservationFleets",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCapacityReservationFleets",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCapacityReservations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCapacityReservations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCarrierGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCarrierGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClassicLinkInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClassicLinkInstances",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClientVpnEndpoints",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClientVpnEndpoints",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCoipPools",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCoipPools",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConversionTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeConversionTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCustomerGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCustomerGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDhcpOptions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDhcpOptions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEgressOnlyInternetGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEgressOnlyInternetGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeElasticGpus",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeElasticGpus",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeExportImageTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeExportImageTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeExportTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeExportTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFastLaunchImages",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeFastLaunchImages",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFastSnapshotRestores",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeFastSnapshotRestores",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFleets",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeFleets",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFlowLogs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeFlowLogs",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFpgaImages",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeFpgaImages",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeHostReservationOfferings",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeHostReservationOfferings",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeHostReservations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeHostReservations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeHosts",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeHosts",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIamInstanceProfileAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIamInstanceProfileAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIdFormat",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIdFormat",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeImages",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeImages",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeImportImageTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeImportImageTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeImportSnapshotTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeImportSnapshotTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceCreditSpecifications",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstanceCreditSpecifications",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceEventNotificationAttributes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstanceEventNotificationAttributes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceEventWindows",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstanceEventWindows",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceStatus",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstanceStatus",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceTypeOfferings",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstanceTypeOfferings",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceTypes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstanceTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInstances",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInternetGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeInternetGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIpamPools",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIpamPools",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIpamScopes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIpamScopes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIpams",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIpams",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIpv6Pools",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIpv6Pools",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeKeyPairs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeKeyPairs",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLaunchTemplateVersions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLaunchTemplateVersions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLaunchTemplates",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLaunchTemplates",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLocalGatewayRouteTableVpcAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLocalGatewayRouteTableVpcAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLocalGatewayRouteTables",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLocalGatewayRouteTables",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLocalGatewayVirtualInterfaceGroups",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLocalGatewayVirtualInterfaceGroups",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLocalGatewayVirtualInterfaces",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLocalGatewayVirtualInterfaces",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLocalGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeLocalGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeManagedPrefixLists",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeManagedPrefixLists",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeMovingAddresses",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeMovingAddresses",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNatGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNatGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkAcls",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkAcls",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInsightsAccessScopeAnalyses",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkInsightsAccessScopeAnalyses",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInsightsAccessScopes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkInsightsAccessScopes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInsightsAnalyses",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkInsightsAnalyses",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInsightsPaths",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkInsightsPaths",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInterfacePermissions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkInterfacePermissions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInterfaces",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeNetworkInterfaces",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePlacementGroups",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePlacementGroups",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePrefixLists",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePrefixLists",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePrincipalIdFormat",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePrincipalIdFormat",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePublicIpv4Pools",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePublicIpv4Pools",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeRegions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeRegions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReplaceRootVolumeTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReplaceRootVolumeTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedInstances",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedInstancesListings",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedInstancesListings",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedInstancesModifications",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedInstancesModifications",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedInstancesOfferings",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedInstancesOfferings",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeRouteTables",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeRouteTables",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeScheduledInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeScheduledInstances",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSecurityGroupRules",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSecurityGroupRules",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSecurityGroups",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSecurityGroups",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSnapshotTierStatus",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSnapshotTierStatus",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSnapshots",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSnapshots",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSpotDatafeedSubscription",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSpotDatafeedSubscription",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSpotFleetRequests",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSpotFleetRequests",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSpotInstanceRequests",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSpotInstanceRequests",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSpotPriceHistory",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSpotPriceHistory",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStoreImageTasks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeStoreImageTasks",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSubnets",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSubnets",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTags",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTags",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTrafficMirrorFilters",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTrafficMirrorFilters",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTrafficMirrorSessions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTrafficMirrorSessions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTrafficMirrorTargets",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTrafficMirrorTargets",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayAttachments",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayAttachments",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayConnectPeers",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayConnectPeers",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayConnects",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayConnects",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayMulticastDomains",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayMulticastDomains",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayPeeringAttachments",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayPeeringAttachments",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayRouteTables",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayRouteTables",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayVpcAttachments",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayVpcAttachments",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTrunkInterfaceAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTrunkInterfaceAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVolumeStatus",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVolumeStatus",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVolumes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVolumes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVolumesModifications",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVolumesModifications",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcClassicLink",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcClassicLink",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcClassicLinkDnsSupport",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcClassicLinkDnsSupport",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpointConnectionNotifications",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcEndpointConnectionNotifications",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpointConnections",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcEndpointConnections",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpointServiceConfigurations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcEndpointServiceConfigurations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpointServices",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcEndpointServices",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpoints",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcEndpoints",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcPeeringConnections",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcPeeringConnections",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpcs",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpnConnections",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpnConnections",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpnGateways",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVpnGateways",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisableEbsEncryptionByDefault",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisableEbsEncryptionByDefault",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisableSerialConsoleAccess",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisableSerialConsoleAccess",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisableVpcClassicLinkDnsSupport",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisableVpcClassicLinkDnsSupport",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisassociateAddress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisassociateAddress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisassociateEnclaveCertificateIamRole",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisassociateEnclaveCertificateIamRole",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisassociateTransitGatewayMulticastDomain",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisassociateTransitGatewayMulticastDomain",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "EnableEbsEncryptionByDefault",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "EnableEbsEncryptionByDefault",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "EnableSerialConsoleAccess",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "EnableSerialConsoleAccess",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "EnableVpcClassicLinkDnsSupport",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "EnableVpcClassicLinkDnsSupport",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAssociatedEnclaveCertificateIamRoles",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetAssociatedEnclaveCertificateIamRoles",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetEbsDefaultKmsKeyId",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetEbsDefaultKmsKeyId",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetEbsEncryptionByDefault",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetEbsEncryptionByDefault",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSerialConsoleAccessStatus",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetSerialConsoleAccessStatus",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayMulticastDomainAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetTransitGatewayMulticastDomainAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetVpnConnectionDeviceTypes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetVpnConnectionDeviceTypes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ImportImage",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ImportImage",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ImportSnapshot",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ImportSnapshot",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListImagesInRecycleBin",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListImagesInRecycleBin",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListSnapshotsInRecycleBin",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListSnapshotsInRecycleBin",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ModifyLaunchTemplate",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ModifyLaunchTemplate",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ModifyPrivateDnsNameOptions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ModifyPrivateDnsNameOptions",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RegisterInstanceEventNotificationAttributes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RegisterInstanceEventNotificationAttributes",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RegisterTransitGatewayMulticastGroupMembers",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RegisterTransitGatewayMulticastGroupMembers",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RegisterTransitGatewayMulticastGroupSources",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RegisterTransitGatewayMulticastGroupSources",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RejectTransitGatewayMulticastDomainAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RejectTransitGatewayMulticastDomainAssociations",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ReleaseAddress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ReleaseAddress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RequestSpotInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RequestSpotInstances",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ResetEbsDefaultKmsKeyId",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ResetEbsDefaultKmsKeyId",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RevokeSecurityGroupIngress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RevokeSecurityGroupIngress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "SearchTransitGatewayMulticastGroups",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "SearchTransitGatewayMulticastGroups",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "UpdateSecurityGroupRuleDescriptionsEgress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "UpdateSecurityGroupRuleDescriptionsEgress",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "UpdateSecurityGroupRuleDescriptionsIngress",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "UpdateSecurityGroupRuleDescriptionsIngress",
	},
}
