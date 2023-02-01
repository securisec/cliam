package policy

import "github.com/securisec/cliam/shared"

// EC2Policies policies
var EC2Policies = map[string]Service{
	"AcceptTransitGatewayMulticastDomainAssociations": {
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
	"AcceptVpcPeeringConnection": {
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
	"AllocateAddress": {
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
	"AssociateAddress": {
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
	"AssociateEnclaveCertificateIamRole": {
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
	"AssociateTransitGatewayMulticastDomain": {
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
	"AuthorizeSecurityGroupIngress": {
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
	"CancelImportTask": {
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
	"DeleteKeyPair": {
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
	"DeleteLaunchTemplate": {
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
	"DeleteSecurityGroup": {
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
	"DeleteSpotDatafeedSubscription": {
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
	"DeregisterInstanceEventNotificationAttributes": {
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
	"DeregisterTransitGatewayMulticastGroupMembers": {
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
	"DeregisterTransitGatewayMulticastGroupSources": {
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
	"DescribeAccountAttributes": {
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
	"DescribeAddressTransfers": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAddressTransfers",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAddressTransfers",
	},
	"DescribeAddresses": {
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
	"DescribeAddressesAttribute": {
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
	"DescribeAggregateIdFormat": {
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
	"DescribeAvailabilityZones": {
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
	"DescribeAwsNetworkPerformanceMetricSubscriptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAwsNetworkPerformanceMetricSubscriptions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAwsNetworkPerformanceMetricSubscriptions",
	},
	"DescribeBundleTasks": {
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
	"DescribeCapacityReservationFleets": {
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
	"DescribeCapacityReservations": {
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
	"DescribeCarrierGateways": {
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
	"DescribeClassicLinkInstances": {
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
	"DescribeClientVpnEndpoints": {
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
	"DescribeCoipPools": {
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
	"DescribeConversionTasks": {
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
	"DescribeCustomerGateways": {
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
	"DescribeDhcpOptions": {
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
	"DescribeEgressOnlyInternetGateways": {
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
	"DescribeElasticGpus": {
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
	"DescribeExportImageTasks": {
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
	"DescribeExportTasks": {
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
	"DescribeFastLaunchImages": {
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
	"DescribeFastSnapshotRestores": {
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
	"DescribeFleets": {
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
	"DescribeFlowLogs": {
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
	"DescribeFpgaImages": {
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
	"DescribeHostReservationOfferings": {
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
	"DescribeHostReservations": {
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
	"DescribeHosts": {
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
	"DescribeIamInstanceProfileAssociations": {
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
	"DescribeIdFormat": {
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
	"DescribeImages": {
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
	"DescribeImportImageTasks": {
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
	"DescribeImportSnapshotTasks": {
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
	"DescribeInstanceCreditSpecifications": {
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
	"DescribeInstanceEventNotificationAttributes": {
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
	"DescribeInstanceEventWindows": {
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
	"DescribeInstanceStatus": {
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
	"DescribeInstanceTypeOfferings": {
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
	"DescribeInstanceTypes": {
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
	"DescribeInstances": {
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
	"DescribeInternetGateways": {
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
	"DescribeIpamPools": {
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
	"DescribeIpamResourceDiscoveries": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIpamResourceDiscoveries",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIpamResourceDiscoveries",
	},
	"DescribeIpamResourceDiscoveryAssociations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIpamResourceDiscoveryAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeIpamResourceDiscoveryAssociations",
	},
	"DescribeIpamScopes": {
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
	"DescribeIpams": {
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
	"DescribeIpv6Pools": {
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
	"DescribeKeyPairs": {
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
	"DescribeLaunchTemplateVersions": {
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
	"DescribeLaunchTemplates": {
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
	"DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations": {
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
	"DescribeLocalGatewayRouteTableVpcAssociations": {
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
	"DescribeLocalGatewayRouteTables": {
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
	"DescribeLocalGatewayVirtualInterfaceGroups": {
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
	"DescribeLocalGatewayVirtualInterfaces": {
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
	"DescribeLocalGateways": {
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
	"DescribeManagedPrefixLists": {
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
	"DescribeMovingAddresses": {
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
	"DescribeNatGateways": {
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
	"DescribeNetworkAcls": {
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
	"DescribeNetworkInsightsAccessScopeAnalyses": {
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
	"DescribeNetworkInsightsAccessScopes": {
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
	"DescribeNetworkInsightsAnalyses": {
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
	"DescribeNetworkInsightsPaths": {
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
	"DescribeNetworkInterfacePermissions": {
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
	"DescribeNetworkInterfaces": {
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
	"DescribePlacementGroups": {
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
	"DescribePrefixLists": {
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
	"DescribePrincipalIdFormat": {
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
	"DescribePublicIpv4Pools": {
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
	"DescribeRegions": {
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
	"DescribeReplaceRootVolumeTasks": {
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
	"DescribeReservedInstances": {
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
	"DescribeReservedInstancesListings": {
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
	"DescribeReservedInstancesModifications": {
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
	"DescribeReservedInstancesOfferings": {
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
	"DescribeRouteTables": {
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
	"DescribeScheduledInstances": {
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
	"DescribeSecurityGroupRules": {
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
	"DescribeSecurityGroups": {
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
	"DescribeSnapshotTierStatus": {
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
	"DescribeSnapshots": {
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
	"DescribeSpotDatafeedSubscription": {
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
	"DescribeSpotFleetRequests": {
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
	"DescribeSpotInstanceRequests": {
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
	"DescribeSpotPriceHistory": {
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
	"DescribeStoreImageTasks": {
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
	"DescribeSubnets": {
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
	"DescribeTags": {
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
	"DescribeTrafficMirrorFilters": {
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
	"DescribeTrafficMirrorSessions": {
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
	"DescribeTrafficMirrorTargets": {
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
	"DescribeTransitGatewayAttachments": {
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
	"DescribeTransitGatewayConnectPeers": {
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
	"DescribeTransitGatewayConnects": {
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
	"DescribeTransitGatewayMulticastDomains": {
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
	"DescribeTransitGatewayPeeringAttachments": {
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
	"DescribeTransitGatewayPolicyTables": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayPolicyTables",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayPolicyTables",
	},
	"DescribeTransitGatewayRouteTableAnnouncements": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTransitGatewayRouteTableAnnouncements",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTransitGatewayRouteTableAnnouncements",
	},
	"DescribeTransitGatewayRouteTables": {
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
	"DescribeTransitGatewayVpcAttachments": {
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
	"DescribeTransitGateways": {
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
	"DescribeTrunkInterfaceAssociations": {
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
	"DescribeVerifiedAccessEndpoints": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVerifiedAccessEndpoints",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVerifiedAccessEndpoints",
	},
	"DescribeVerifiedAccessGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVerifiedAccessGroups",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVerifiedAccessGroups",
	},
	"DescribeVerifiedAccessInstanceLoggingConfigurations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVerifiedAccessInstanceLoggingConfigurations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVerifiedAccessInstanceLoggingConfigurations",
	},
	"DescribeVerifiedAccessInstances": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVerifiedAccessInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVerifiedAccessInstances",
	},
	"DescribeVerifiedAccessTrustProviders": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVerifiedAccessTrustProviders",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeVerifiedAccessTrustProviders",
	},
	"DescribeVolumeStatus": {
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
	"DescribeVolumes": {
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
	"DescribeVolumesModifications": {
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
	"DescribeVpcClassicLink": {
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
	"DescribeVpcClassicLinkDnsSupport": {
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
	"DescribeVpcEndpointConnectionNotifications": {
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
	"DescribeVpcEndpointConnections": {
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
	"DescribeVpcEndpointServiceConfigurations": {
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
	"DescribeVpcEndpointServices": {
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
	"DescribeVpcEndpoints": {
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
	"DescribeVpcPeeringConnections": {
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
	"DescribeVpcs": {
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
	"DescribeVpnConnections": {
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
	"DescribeVpnGateways": {
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
	"DisableAwsNetworkPerformanceMetricSubscription": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DisableAwsNetworkPerformanceMetricSubscription",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DisableAwsNetworkPerformanceMetricSubscription",
	},
	"DisableEbsEncryptionByDefault": {
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
	"DisableSerialConsoleAccess": {
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
	"DisableVpcClassicLinkDnsSupport": {
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
	"DisassociateAddress": {
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
	"DisassociateEnclaveCertificateIamRole": {
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
	"DisassociateTransitGatewayMulticastDomain": {
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
	"EnableAwsNetworkPerformanceMetricSubscription": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "EnableAwsNetworkPerformanceMetricSubscription",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "EnableAwsNetworkPerformanceMetricSubscription",
	},
	"EnableEbsEncryptionByDefault": {
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
	"EnableReachabilityAnalyzerOrganizationSharing": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "EnableReachabilityAnalyzerOrganizationSharing",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "EnableReachabilityAnalyzerOrganizationSharing",
	},
	"EnableSerialConsoleAccess": {
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
	"EnableVpcClassicLinkDnsSupport": {
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
	"GetAssociatedEnclaveCertificateIamRoles": {
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
	"GetAwsNetworkPerformanceData": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAwsNetworkPerformanceData",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetAwsNetworkPerformanceData",
	},
	"GetEbsDefaultKmsKeyId": {
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
	"GetEbsEncryptionByDefault": {
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
	"GetSerialConsoleAccessStatus": {
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
	"GetTransitGatewayMulticastDomainAssociations": {
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
	"GetVpnConnectionDeviceTypes": {
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
	"ImportImage": {
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
	"ImportSnapshot": {
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
	"ListImagesInRecycleBin": {
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
	"ListSnapshotsInRecycleBin": {
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
	"ModifyLaunchTemplate": {
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
	"ModifyPrivateDnsNameOptions": {
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
	"RegisterInstanceEventNotificationAttributes": {
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
	"RegisterTransitGatewayMulticastGroupMembers": {
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
	"RegisterTransitGatewayMulticastGroupSources": {
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
	"RejectTransitGatewayMulticastDomainAssociations": {
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
	"ReleaseAddress": {
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
	"RequestSpotInstances": {
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
	"ResetEbsDefaultKmsKeyId": {
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
	"RevokeSecurityGroupIngress": {
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
	"SearchTransitGatewayMulticastGroups": {
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
	"UpdateSecurityGroupRuleDescriptionsEgress": {
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
	"UpdateSecurityGroupRuleDescriptionsIngress": {
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

	// extra
	"DescribeByoipCidrs": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeByoipCidrs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeByoipCidrs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "MaxResults",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "max_results",
	},
	"DescribeClientVpnAuthorizationRules": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClientVpnAuthorizationRules",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeClientVpnAuthorizationRules",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClientVpnEndpointId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "client_vpn_endpoint_id",
	},
	"DescribeClientVpnConnections": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClientVpnConnections",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeClientVpnConnections",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClientVpnEndpointId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "client_vpn_endpoint_id",
	},
	"DescribeClientVpnRoutes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClientVpnRoutes",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeClientVpnRoutes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClientVpnEndpointId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "client_vpn_endpoint_id",
	},
	"DescribeClientVpnTargetNetworks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClientVpnTargetNetworks",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeClientVpnTargetNetworks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClientVpnEndpointId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "client_vpn_endpoint_id",
	},
	"DescribeFleetInstances": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFleetInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeFleetInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "fleet_id",
	},
	"DescribeIdentityIdFormat": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIdentityIdFormat",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeIdentityIdFormat",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PrincipalArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "principal_arn",
	},
	"DescribeNetworkInterfaceAttribute": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNetworkInterfaceAttribute",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeNetworkInterfaceAttribute",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NetworkInterfaceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "network_interface_id",
	},
	"DescribeSecurityGroupReferences": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSecurityGroupReferences",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSecurityGroupReferences",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GroupId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "group_id",
	},
	"DescribeSpotFleetInstances": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSpotFleetInstances",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSpotFleetInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SpotFleetRequestId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "spot_fleet_request_id",
	},
	"DescribeStaleSecurityGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStaleSecurityGroups",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeStaleSecurityGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VpcId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "vpc_id",
	},
	"DescribeVpcEndpointServicePermissions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpointServicePermissions",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeVpcEndpointServicePermissions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "service_id",
	},
	"GetAssociatedIpv6PoolCidrs": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAssociatedIpv6PoolCidrs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetAssociatedIpv6PoolCidrs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PoolId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "pool_id",
	},
	"GetCapacityReservationUsage": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetCapacityReservationUsage",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetCapacityReservationUsage",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CapacityReservationId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "capacity_reservation_id",
	},
	"GetCoipPoolUsage": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetCoipPoolUsage",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetCoipPoolUsage",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PoolId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "pool_id",
	},
	"GetConsoleOutput": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetConsoleOutput",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetConsoleOutput",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_id",
	},
	"GetConsoleScreenshot": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetConsoleScreenshot",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetConsoleScreenshot",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_id",
	},
	"GetDefaultCreditSpecification": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetDefaultCreditSpecification",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetDefaultCreditSpecification",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceFamily",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_family",
	},
	"GetGroupsForCapacityReservation": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetGroupsForCapacityReservation",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetGroupsForCapacityReservation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CapacityReservationId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "capacity_reservation_id",
	},
	"GetInstanceUefiData": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetInstanceUefiData",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetInstanceUefiData",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_id",
	},
	"GetIpamPoolAllocations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIpamPoolAllocations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIpamPoolAllocations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IpamPoolId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "ipam_pool_id",
	},
	"GetIpamPoolCidrs": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIpamPoolCidrs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIpamPoolCidrs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IpamPoolId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "ipam_pool_id",
	},
	"GetIpamResourceCidrs": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIpamResourceCidrs",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIpamResourceCidrs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IpamScopeId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "ipam_scope_id",
	},
	"GetLaunchTemplateData": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetLaunchTemplateData",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetLaunchTemplateData",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_id",
	},
	"GetManagedPrefixListAssociations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetManagedPrefixListAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetManagedPrefixListAssociations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PrefixListId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "prefix_list_id",
	},
	"GetManagedPrefixListEntries": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetManagedPrefixListEntries",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetManagedPrefixListEntries",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PrefixListId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "prefix_list_id",
	},
	"GetNetworkInsightsAccessScopeAnalysisFindings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetNetworkInsightsAccessScopeAnalysisFindings",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetNetworkInsightsAccessScopeAnalysisFindings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NetworkInsightsAccessScopeAnalysisId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "network_insights_access_scope_analysis_id",
	},
	"GetNetworkInsightsAccessScopeContent": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetNetworkInsightsAccessScopeContent",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetNetworkInsightsAccessScopeContent",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NetworkInsightsAccessScopeId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "network_insights_access_scope_id",
	},
	"GetPasswordData": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetPasswordData",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetPasswordData",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InstanceId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "instance_id",
	},
	"GetReservedInstancesExchangeQuote": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetReservedInstancesExchangeQuote",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetReservedInstancesExchangeQuote",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ReservedInstanceIds",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "reserved_instance_ids",
	},
	"GetSpotPlacementScores": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSpotPlacementScores",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetSpotPlacementScores",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TargetCapacity",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "target_capacity",
	},
	"GetSubnetCidrReservations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSubnetCidrReservations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetSubnetCidrReservations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SubnetId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "subnet_id",
	},
	"GetTransitGatewayAttachmentPropagations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayAttachmentPropagations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTransitGatewayAttachmentPropagations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransitGatewayAttachmentId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transit_gateway_attachment_id",
	},
	"GetTransitGatewayPolicyTableAssociations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayPolicyTableAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTransitGatewayPolicyTableAssociations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransitGatewayPolicyTableId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transit_gateway_policy_table_id",
	},
	"GetTransitGatewayPolicyTableEntries": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayPolicyTableEntries",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTransitGatewayPolicyTableEntries",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransitGatewayPolicyTableId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transit_gateway_policy_table_id",
	},
	"GetTransitGatewayPrefixListReferences": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayPrefixListReferences",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTransitGatewayPrefixListReferences",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransitGatewayRouteTableId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transit_gateway_route_table_id",
	},
	"GetTransitGatewayRouteTableAssociations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayRouteTableAssociations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTransitGatewayRouteTableAssociations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransitGatewayRouteTableId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transit_gateway_route_table_id",
	},
	"GetTransitGatewayRouteTablePropagations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTransitGatewayRouteTablePropagations",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTransitGatewayRouteTablePropagations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransitGatewayRouteTableId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "transit_gateway_route_table_id",
	},
	"GetVerifiedAccessEndpointPolicy": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetVerifiedAccessEndpointPolicy",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetVerifiedAccessEndpointPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VerifiedAccessEndpointId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "verified_access_endpoint_id",
	},
	"GetVerifiedAccessGroupPolicy": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetVerifiedAccessGroupPolicy",
			"Version": "2016-11-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetVerifiedAccessGroupPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VerifiedAccessGroupId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "verified_access_group_id",
	},
}
