package policy

import "github.com/securisec/cliam/shared"

// RDSPolicies policy
var RDSPolicies = map[string]Service{
	"DescribeAccountAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAccountAttributes",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeAccountAttributes",
	},
	"DescribeBlueGreenDeployments": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeBlueGreenDeployments",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeBlueGreenDeployments",
	},
	"DescribeCertificates": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCertificates",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCertificates",
	},
	"DescribeDBClusterEndpoints": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusterEndpoints",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBClusterEndpoints",
	},
	"DescribeDBClusterParameterGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusterParameterGroups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBClusterParameterGroups",
	},
	"DescribeDBClusterSnapshots": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusterSnapshots",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBClusterSnapshots",
	},
	"DescribeDBClusters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusters",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBClusters",
	},
	"DescribeDBEngineVersions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBEngineVersions",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBEngineVersions",
	},
	"DescribeDBInstanceAutomatedBackups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBInstanceAutomatedBackups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBInstanceAutomatedBackups",
	},
	"DescribeDBInstances": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBInstances",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBInstances",
	},
	"DescribeDBParameterGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBParameterGroups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBParameterGroups",
	},
	"DescribeDBProxies": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBProxies",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBProxies",
	},
	"DescribeDBProxyEndpoints": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBProxyEndpoints",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBProxyEndpoints",
	},
	"DescribeDBSecurityGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBSecurityGroups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBSecurityGroups",
	},
	"DescribeDBSnapshots": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBSnapshots",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBSnapshots",
	},
	"DescribeDBSubnetGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBSubnetGroups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDBSubnetGroups",
	},
	"DescribeEventCategories": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEventCategories",
			"Version": "2014-10-31",
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
			"Version": "2014-10-31",
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
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEvents",
	},
	"DescribeExportTasks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeExportTasks",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeExportTasks",
	},
	"DescribeGlobalClusters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeGlobalClusters",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeGlobalClusters",
	},
	"DescribeOptionGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeOptionGroups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeOptionGroups",
	},
	"DescribePendingMaintenanceActions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribePendingMaintenanceActions",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribePendingMaintenanceActions",
	},
	"DescribeReservedDBInstances": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedDBInstances",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedDBInstances",
	},
	"DescribeReservedDBInstancesOfferings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedDBInstancesOfferings",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedDBInstancesOfferings",
	},
	"DescribeSourceRegions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSourceRegions",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSourceRegions",
	},
	"ModifyActivityStream": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ModifyActivityStream",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ModifyActivityStream",
	},
	"ModifyCertificates": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ModifyCertificates",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ModifyCertificates",
	},
	"ModifyGlobalCluster": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ModifyGlobalCluster",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ModifyGlobalCluster",
	},
	"RemoveFromGlobalCluster": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "RemoveFromGlobalCluster",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "RemoveFromGlobalCluster",
	},

	// extra
	"DescribeDBClusterBacktracks": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusterBacktracks",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBClusterBacktracks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBClusterIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_cluster_identifier",
	},
	"DescribeDBClusterParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusterParameters",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBClusterParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBClusterParameterGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_cluster_parameter_group_name",
	},
	"DescribeDBClusterSnapshotAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBClusterSnapshotAttributes",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBClusterSnapshotAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBClusterSnapshotIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_cluster_snapshot_identifier",
	},
	"DescribeDBLogFiles": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBLogFiles",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBLogFiles",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBInstanceIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_instance_identifier",
	},
	"DescribeDBParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBParameters",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBParameterGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_parameter_group_name",
	},
	"DescribeDBProxyTargetGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBProxyTargetGroups",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBProxyTargetGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBProxyName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_proxy_name",
	},
	"DescribeDBProxyTargets": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBProxyTargets",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBProxyTargets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBProxyName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_proxy_name",
	},
	"DescribeDBSnapshotAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDBSnapshotAttributes",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDBSnapshotAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBSnapshotIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_snapshot_identifier",
	},
	"DescribeEngineDefaultClusterParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEngineDefaultClusterParameters",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeEngineDefaultClusterParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBParameterGroupFamily",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_parameter_group_family",
	},
	"DescribeEngineDefaultParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEngineDefaultParameters",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeEngineDefaultParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBParameterGroupFamily",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_parameter_group_family",
	},
	"DescribeOptionGroupOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeOptionGroupOptions",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeOptionGroupOptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EngineName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "engine_name",
	},
	"DescribeOrderableDBInstanceOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeOrderableDBInstanceOptions",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeOrderableDBInstanceOptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Engine",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "engine",
	},
	"DescribeValidDBInstanceModifications": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeValidDBInstanceModifications",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeValidDBInstanceModifications",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DBInstanceIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "d_b_instance_identifier",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2014-10-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_name",
	},
}
