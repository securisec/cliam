package policy

import "github.com/securisec/cliam/shared"

var RDSPolicies = map[string]Service{
	"DescribeDBEngineVersions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBEngineVersions&Version=2014-10-31",
		Permission:    "DescribeDBEngineVersions",
	},
	"DescribeDBInstances": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBInstances&Version=2014-10-31",
		Permission:    "DescribeDBInstances",
	},
	"DescribeDBParameterGroups": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBParameterGroups&Version=2014-10-31",
		Permission:    "DescribeDBParameterGroups",
	},
	"DescribeDBSecurityGroups": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBSecurityGroups&Version=2014-10-31",
		Permission:    "DescribeDBSecurityGroups",
	},
	"DescribeDBSnapshots": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBSnapshots&Version=2014-10-31",
		Permission:    "DescribeDBSnapshots",
	},
	"DescribeDBSubnetGroups": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBSubnetGroups&Version=2014-10-31",
		Permission:    "DescribeDBSubnetGroups",
	},
	"DescribeEventCategories": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeEventCategories&Version=2014-10-31",
		Permission:    "DescribeEventCategories",
	},
	"DescribeEventSubscriptions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeEventSubscriptions&Version=2014-10-31",
		Permission:    "DescribeEventSubscriptions",
	},
	"DescribeOptionGroups": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeOptionGroups&Version=2014-10-31",
		Permission:    "DescribeOptionGroups",
	},
	"DescribeReservedDBInstances": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeReservedDBInstances&Version=2014-10-31",
		Permission:    "DescribeReservedDBInstances",
	},
	"DescribeReservedDBInstancesOfferings": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeReservedDBInstancesOfferings&Version=2014-10-31",
		Permission:    "DescribeReservedDBInstancesOfferings",
	},
	"DescribeAccountAttributes": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountAttributes&Version=2014-10-31",
		Permission:    "DescribeAccountAttributes",
	},
	"DescribeCertificates": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeCertificates&Version=2014-10-31",
		Permission:    "DescribeCertificates",
	},
	"DescribeCustomAvailabilityZones": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeCustomAvailabilityZones&Version=2014-10-31",
		Permission:    "DescribeCustomAvailabilityZones",
	},
	"DescribeDBClusterEndpoints": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBClusterEndpoints&Version=2014-10-31",
		Permission:    "DescribeDBClusterEndpoints",
	},
	"DescribeDBClusterParameterGroups": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBClusterParameterGroups&Version=2014-10-31",
		Permission:    "DescribeDBClusterParameterGroups",
	},
	"DescribeDBClusterSnapshots": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBClusterSnapshots&Version=2014-10-31",
		Permission:    "DescribeDBClusterSnapshots",
	},
	"DescribeDBClusters": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBClusters&Version=2014-10-31",
		Permission:    "DescribeDBClusters",
	},
	"DescribeDBInstanceAutomatedBackups": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBInstanceAutomatedBackups&Version=2014-10-31",
		Permission:    "DescribeDBInstanceAutomatedBackups",
	},
	"DescribeDBProxies": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBProxies&Version=2014-10-31",
		Permission:    "DescribeDBProxies",
	},
	"DescribeDBProxyEndpoints": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeDBProxyEndpoints&Version=2014-10-31",
		Permission:    "DescribeDBProxyEndpoints",
	},
	"DescribeExportTasks": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeExportTasks&Version=2014-10-31",
		Permission:    "DescribeExportTasks",
	},
	"DescribeGlobalClusters": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeGlobalClusters&Version=2014-10-31",
		Permission:    "DescribeGlobalClusters",
	},
	"DescribeInstallationMedia": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeInstallationMedia&Version=2014-10-31",
		Permission:    "DescribeInstallationMedia",
	},
	"DescribePendingMaintenanceActions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribePendingMaintenanceActions&Version=2014-10-31",
		Permission:    "DescribePendingMaintenanceActions",
	},
	"DescribeSourceRegions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeSourceRegions&Version=2014-10-31",
		Permission:    "DescribeSourceRegions",
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
		ExtraCommandLineFlag:   "db_cluster_identifier",
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
		ExtraCommandLineFlag:   "db_cluster_parameter_group_name",
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
		ExtraCommandLineFlag:   "db_cluster_snapshot_identifier",
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
		ExtraCommandLineFlag:   "db_instance_identifier",
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
		ExtraCommandLineFlag:   "db_parameter_group_name",
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
		ExtraCommandLineFlag:   "db_proxy_name",
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
		ExtraCommandLineFlag:   "db_proxy_name",
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
		ExtraCommandLineFlag:   "db_snapshot_identifier",
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
		ExtraCommandLineFlag:   "db_parameter_group_family",
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
		ExtraCommandLineFlag:   "db_parameter_group_family",
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
		ExtraCommandLineFlag:   "db_instance_identifier",
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
