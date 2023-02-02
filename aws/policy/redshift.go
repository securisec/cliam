package policy

import "github.com/securisec/cliam/shared"

// RedshiftPolicies policy
var RedshiftPolicies = map[string]Service{
	"DescribeAccountAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountAttributes",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountAttributes",
	},
	"DescribeAuthenticationProfiles": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAuthenticationProfiles",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAuthenticationProfiles",
	},
	"DescribeClusterDbRevisions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterDbRevisions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterDbRevisions",
	},
	"DescribeClusterParameterGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterParameterGroups",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterParameterGroups",
	},
	"DescribeClusterSecurityGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterSecurityGroups",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterSecurityGroups",
	},
	"DescribeClusterSnapshots": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterSnapshots",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterSnapshots",
	},
	"DescribeClusterSubnetGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterSubnetGroups",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterSubnetGroups",
	},
	"DescribeClusterTracks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterTracks",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterTracks",
	},
	"DescribeClusterVersions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterVersions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusterVersions",
	},
	"DescribeClusters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusters",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeClusters",
	},
	"DescribeDataShares": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDataShares",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDataShares",
	},
	"DescribeDataSharesForConsumer": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDataSharesForConsumer",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDataSharesForConsumer",
	},
	"DescribeDataSharesForProducer": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDataSharesForProducer",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDataSharesForProducer",
	},
	"DescribeEndpointAccess": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEndpointAccess",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEndpointAccess",
	},
	"DescribeEndpointAuthorization": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEndpointAuthorization",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEndpointAuthorization",
	},
	"DescribeEventCategories": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEventCategories",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEventCategories",
	},
	"DescribeEventSubscriptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEventSubscriptions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEventSubscriptions",
	},
	"DescribeEvents": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEvents",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEvents",
	},
	"DescribeHsmClientCertificates": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeHsmClientCertificates",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeHsmClientCertificates",
	},
	"DescribeHsmConfigurations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeHsmConfigurations",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeHsmConfigurations",
	},
	"DescribeOrderableClusterOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeOrderableClusterOptions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeOrderableClusterOptions",
	},
	"DescribeReservedNodeExchangeStatus": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedNodeExchangeStatus",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedNodeExchangeStatus",
	},
	"DescribeReservedNodeOfferings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedNodeOfferings",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedNodeOfferings",
	},
	"DescribeReservedNodes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedNodes",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedNodes",
	},
	"DescribeScheduledActions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeScheduledActions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeScheduledActions",
	},
	"DescribeSnapshotCopyGrants": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSnapshotCopyGrants",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSnapshotCopyGrants",
	},
	"DescribeSnapshotSchedules": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSnapshotSchedules",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSnapshotSchedules",
	},
	"DescribeStorage": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeStorage",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeStorage",
	},
	"DescribeTableRestoreStatus": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTableRestoreStatus",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTableRestoreStatus",
	},
	"DescribeTags": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTags",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeTags",
	},
	"DescribeUsageLimits": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeUsageLimits",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeUsageLimits",
	},
	"RevokeEndpointAccess": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RevokeEndpointAccess",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RevokeEndpointAccess",
	},

	// extra
	"DescribeClusterParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeClusterParameters",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeClusterParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ParameterGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "parameter_group_name",
	},
	"DescribeDefaultClusterParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDefaultClusterParameters",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDefaultClusterParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ParameterGroupFamily",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "parameter_group_family",
	},
	"DescribeLoggingStatus": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoggingStatus",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLoggingStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClusterIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "cluster_identifier",
	},
	"DescribeNodeConfigurationOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeNodeConfigurationOptions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeNodeConfigurationOptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ActionType",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "action_type",
	},
	"DescribeResize": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeResize",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeResize",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClusterIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "cluster_identifier",
	},
	"GetClusterCredentialsWithIAM": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetClusterCredentialsWithIAM",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetClusterCredentialsWithIAM",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClusterIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "cluster_identifier",
	},
	"GetReservedNodeExchangeConfigurationOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetReservedNodeExchangeConfigurationOptions",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetReservedNodeExchangeConfigurationOptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ActionType",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "action_type",
	},
	"GetReservedNodeExchangeOfferings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetReservedNodeExchangeOfferings",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetReservedNodeExchangeOfferings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ReservedNodeId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "reserved_node_id",
	},
}
