package policy

import "github.com/securisec/cliam/shared"

var RedshiftPolicies = []Service{
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "CreateSnapshotSchedule",
			"Version": "2012-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "CreateSnapshotSchedule",
	},
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
}
