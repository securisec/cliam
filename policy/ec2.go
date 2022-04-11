package policy

var EC2 = []Service{
	{
		ServiceSuffix: "?Action=DescribeAccountAttributes&Version=2016-11-15",
		Permission:    "DescribeAccountAttributes",
	},
	{
		ServiceSuffix: "?Action=DescribeAddresses&Version=2016-11-15",
		Permission:    "DescribeAddresses",
	},
	{
		ServiceSuffix: "?Action=DescribeAggregateIdFormat&Version=2016-11-15",
		Permission:    "DescribeAggregateIdFormat",
	},
	{
		ServiceSuffix: "?Action=DescribeAvailabilityZones&Version=2016-11-15",
		Permission:    "DescribeAvailabilityZones",
	},
	{
		ServiceSuffix: "?Action=DescribeBundleTasks&Version=2016-11-15",
		Permission:    "DescribeBundleTasks",
	},
	{
		ServiceSuffix: "?Action=DescribeCapacityReservations&Version=2016-11-15",
		Permission:    "DescribeCapacityReservations",
	},
	{
		ServiceSuffix: "?Action=DescribeClassicLinkInstances&Version=2016-11-15",
		Permission:    "DescribeClassicLinkInstances",
	},
	{
		ServiceSuffix: "?Action=DescribeClientVpnEndpoints&Version=2016-11-15",
		Permission:    "DescribeClientVpnEndpoints",
	},
	{
		ServiceSuffix: "?Action=DescribeConversionTasks&Version=2016-11-15",
		Permission:    "DescribeConversionTasks",
	},
	{
		ServiceSuffix: "?Action=DescribeCustomerGateways&Version=2016-11-15",
		Permission:    "DescribeCustomerGateways",
	},
	{
		ServiceSuffix: "?Action=DescribeDhcpOptions&Version=2016-11-15",
		Permission:    "DescribeDhcpOptions",
	},
	{
		ServiceSuffix: "?Action=DescribeEgressOnlyInternetGateways&Version=2016-11-15",
		Permission:    "DescribeEgressOnlyInternetGateways",
	},
	{
		ServiceSuffix: "?Action=DescribeElasticGpus&Version=2016-11-15",
		Permission:    "DescribeElasticGpus",
	},
	{
		ServiceSuffix: "?Action=DescribeExportTasks&Version=2016-11-15",
		Permission:    "DescribeExportTasks",
	},
	{
		ServiceSuffix: "?Action=DescribeFleets&Version=2016-11-15",
		Permission:    "DescribeFleets",
	},
	{
		ServiceSuffix: "?Action=DescribeFlowLogs&Version=2016-11-15",
		Permission:    "DescribeFlowLogs",
	},
	{
		ServiceSuffix: "?Action=DescribeFpgaImages&Version=2016-11-15",
		Permission:    "DescribeFpgaImages",
	},
	{
		ServiceSuffix: "?Action=DescribeHostReservationOfferings&Version=2016-11-15",
		Permission:    "DescribeHostReservationOfferings",
	},
	{
		ServiceSuffix: "?Action=DescribeHostReservations&Version=2016-11-15",
		Permission:    "DescribeHostReservations",
	},
	{
		ServiceSuffix: "?Action=DescribeHosts&Version=2016-11-15",
		Permission:    "DescribeHosts",
	},
	{
		ServiceSuffix: "?Action=DescribeIamInstanceProfileAssociations&Version=2016-11-15",
		Permission:    "DescribeIamInstanceProfileAssociations",
	},
	{
		ServiceSuffix: "?Action=DescribeIdFormat&Version=2016-11-15",
		Permission:    "DescribeIdFormat",
	},
	{
		ServiceSuffix: "?Action=DescribeImages&Version=2016-11-15",
		Permission:    "DescribeImages",
	},
	{
		ServiceSuffix: "?Action=DescribeImportImageTasks&Version=2016-11-15",
		Permission:    "DescribeImportImageTasks",
	},
	{
		ServiceSuffix: "?Action=DescribeImportSnapshotTasks&Version=2016-11-15",
		Permission:    "DescribeImportSnapshotTasks",
	},
	{
		ServiceSuffix: "?Action=DescribeInstanceCreditSpecifications&Version=2016-11-15",
		Permission:    "DescribeInstanceCreditSpecifications",
	},
	{
		ServiceSuffix: "?Action=DescribeInstanceStatus&Version=2016-11-15",
		Permission:    "DescribeInstanceStatus",
	},
	{
		ServiceSuffix: "?Action=DescribeInstances&Version=2016-11-15",
		Permission:    "DescribeInstances",
	},
	{
		ServiceSuffix: "?Action=DescribeInternetGateways&Version=2016-11-15",
		Permission:    "DescribeInternetGateways",
	},
	{
		ServiceSuffix: "?Action=DescribeKeyPairs&Version=2016-11-15",
		Permission:    "DescribeKeyPairs",
	},
	{
		ServiceSuffix: "?Action=DescribeLaunchTemplateVersions&Version=2016-11-15",
		Permission:    "DescribeLaunchTemplateVersions",
	},
	{
		ServiceSuffix: "?Action=DescribeLaunchTemplates&Version=2016-11-15",
		Permission:    "DescribeLaunchTemplates",
	},
	{
		ServiceSuffix: "?Action=DescribeMovingAddresses&Version=2016-11-15",
		Permission:    "DescribeMovingAddresses",
	},
	{
		ServiceSuffix: "?Action=DescribeNatGateways&Version=2016-11-15",
		Permission:    "DescribeNatGateways",
	},
	{
		ServiceSuffix: "?Action=DescribeNetworkAcls&Version=2016-11-15",
		Permission:    "DescribeNetworkAcls",
	},
	{
		ServiceSuffix: "?Action=DescribeNetworkInterfacePermissions&Version=2016-11-15",
		Permission:    "DescribeNetworkInterfacePermissions",
	},
	{
		ServiceSuffix: "?Action=DescribeNetworkInterfaces&Version=2016-11-15",
		Permission:    "DescribeNetworkInterfaces",
	},
	{
		ServiceSuffix: "?Action=DescribePlacementGroups&Version=2016-11-15",
		Permission:    "DescribePlacementGroups",
	},
	{
		ServiceSuffix: "?Action=DescribePrefixLists&Version=2016-11-15",
		Permission:    "DescribePrefixLists",
	},
	{
		ServiceSuffix: "?Action=DescribePrincipalIdFormat&Version=2016-11-15",
		Permission:    "DescribePrincipalIdFormat",
	},
	{
		ServiceSuffix: "?Action=DescribePublicIpv4Pools&Version=2016-11-15",
		Permission:    "DescribePublicIpv4Pools",
	},
	{
		ServiceSuffix: "?Action=DescribeRegions&Version=2016-11-15",
		Permission:    "DescribeRegions",
	},
	{
		ServiceSuffix: "?Action=DescribeReservedInstances&Version=2016-11-15",
		Permission:    "DescribeReservedInstances",
	},
	{
		ServiceSuffix: "?Action=DescribeReservedInstancesListings&Version=2016-11-15",
		Permission:    "DescribeReservedInstancesListings",
	},
	{
		ServiceSuffix: "?Action=DescribeReservedInstancesModifications&Version=2016-11-15",
		Permission:    "DescribeReservedInstancesModifications",
	},
	{
		ServiceSuffix: "?Action=DescribeReservedInstancesOfferings&Version=2016-11-15",
		Permission:    "DescribeReservedInstancesOfferings",
	},
	{
		ServiceSuffix: "?Action=DescribeRouteTables&Version=2016-11-15",
		Permission:    "DescribeRouteTables",
	},
	{
		ServiceSuffix: "?Action=DescribeScheduledInstances&Version=2016-11-15",
		Permission:    "DescribeScheduledInstances",
	},
	{
		ServiceSuffix: "?Action=DescribeSecurityGroups&Version=2016-11-15",
		Permission:    "DescribeSecurityGroups",
	},
	{
		ServiceSuffix: "?Action=DescribeSnapshots&Version=2016-11-15",
		Permission:    "DescribeSnapshots",
	},
	{
		ServiceSuffix: "?Action=DescribeSpotDatafeedSubscription&Version=2016-11-15",
		Permission:    "DescribeSpotDatafeedSubscription",
	},
	{
		ServiceSuffix: "?Action=DescribeSpotFleetRequests&Version=2016-11-15",
		Permission:    "DescribeSpotFleetRequests",
	},
	{
		ServiceSuffix: "?Action=DescribeSpotInstanceRequests&Version=2016-11-15",
		Permission:    "DescribeSpotInstanceRequests",
	},
	{
		ServiceSuffix: "?Action=DescribeSpotPriceHistory&Version=2016-11-15",
		Permission:    "DescribeSpotPriceHistory",
	},
	{
		ServiceSuffix: "?Action=DescribeSubnets&Version=2016-11-15",
		Permission:    "DescribeSubnets",
	},
	{
		ServiceSuffix: "?Action=DescribeTags&Version=2016-11-15",
		Permission:    "DescribeTags",
	},
	{
		ServiceSuffix: "?Action=DescribeTransitGatewayAttachments&Version=2016-11-15",
		Permission:    "DescribeTransitGatewayAttachments",
	},
	{
		ServiceSuffix: "?Action=DescribeTransitGatewayRouteTables&Version=2016-11-15",
		Permission:    "DescribeTransitGatewayRouteTables",
	},
	{
		ServiceSuffix: "?Action=DescribeTransitGatewayVpcAttachments&Version=2016-11-15",
		Permission:    "DescribeTransitGatewayVpcAttachments",
	},
	{
		ServiceSuffix: "?Action=DescribeTransitGateways&Version=2016-11-15",
		Permission:    "DescribeTransitGateways",
	},
	{
		ServiceSuffix: "?Action=DescribeVolumeStatus&Version=2016-11-15",
		Permission:    "DescribeVolumeStatus",
	},
	{
		ServiceSuffix: "?Action=DescribeVolumes&Version=2016-11-15",
		Permission:    "DescribeVolumes",
	},
	{
		ServiceSuffix: "?Action=DescribeVolumesModifications&Version=2016-11-15",
		Permission:    "DescribeVolumesModifications",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcClassicLink&Version=2016-11-15",
		Permission:    "DescribeVpcClassicLink",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcClassicLinkDnsSupport&Version=2016-11-15",
		Permission:    "DescribeVpcClassicLinkDnsSupport",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcEndpointConnectionNotifications&Version=2016-11-15",
		Permission:    "DescribeVpcEndpointConnectionNotifications",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcEndpointConnections&Version=2016-11-15",
		Permission:    "DescribeVpcEndpointConnections",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcEndpointServiceConfigurations&Version=2016-11-15",
		Permission:    "DescribeVpcEndpointServiceConfigurations",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcEndpointServices&Version=2016-11-15",
		Permission:    "DescribeVpcEndpointServices",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcEndpoints&Version=2016-11-15",
		Permission:    "DescribeVpcEndpoints",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcPeeringConnections&Version=2016-11-15",
		Permission:    "DescribeVpcPeeringConnections",
	},
	{
		ServiceSuffix: "?Action=DescribeVpcs&Version=2016-11-15",
		Permission:    "DescribeVpcs",
	},
	{
		ServiceSuffix: "?Action=DescribeVpnConnections&Version=2016-11-15",
		Permission:    "DescribeVpnConnections",
	},
	{
		ServiceSuffix: "?Action=DescribeVpnGateways&Version=2016-11-15",
		Permission:    "DescribeVpnGateways",
	},
}
