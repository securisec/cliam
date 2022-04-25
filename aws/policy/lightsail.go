package policy

import "github.com/securisec/cliam/shared"

var LightsailPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetActiveNames",
		},
		Permission: "GetActiveNames",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetAlarms",
		},
		Permission: "GetAlarms",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBlueprints",
		},
		Permission: "GetBlueprints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBucketBundles",
		},
		Permission: "GetBucketBundles",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBuckets",
		},
		Permission: "GetBuckets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBundles",
		},
		Permission: "GetBundles",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetCertificates",
		},
		Permission: "GetCertificates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetCloudFormationStackRecords",
		},
		Permission: "GetCloudFormationStackRecords",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContactMethods",
		},
		Permission: "GetContactMethods",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerApiMetadata",
		},
		Permission: "GetContainerApiMetadata",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerServicePowers",
		},
		Permission: "GetContainerServicePowers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerServices",
		},
		Permission: "GetContainerServices",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDiskSnapshots",
		},
		Permission: "GetDiskSnapshots",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDisks",
		},
		Permission: "GetDisks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDistributionBundles",
		},
		Permission: "GetDistributionBundles",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDistributionLatestCacheReset",
		},
		Permission: "GetDistributionLatestCacheReset",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDistributions",
		},
		Permission: "GetDistributions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDomains",
		},
		Permission: "GetDomains",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetExportSnapshotRecords",
		},
		Permission: "GetExportSnapshotRecords",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstanceSnapshots",
		},
		Permission: "GetInstanceSnapshots",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstances",
		},
		Permission: "GetInstances",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetKeyPairs",
		},
		Permission: "GetKeyPairs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetLoadBalancers",
		},
		Permission: "GetLoadBalancers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetOperations",
		},
		Permission: "GetOperations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRegions",
		},
		Permission: "GetRegions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseBlueprints",
		},
		Permission: "GetRelationalDatabaseBlueprints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseBundles",
		},
		Permission: "GetRelationalDatabaseBundles",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseSnapshots",
		},
		Permission: "GetRelationalDatabaseSnapshots",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabases",
		},
		Permission: "GetRelationalDatabases",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetStaticIps",
		},
		Permission: "GetStaticIps",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetAutoSnapshots",
		},
		Permission:             "GetAutoSnapshots",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBucketAccessKeys",
		},
		Permission:             "GetBucketAccessKeys",
		IsExtra:                true,
		ExtraComponentBodyKey:  "bucketName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "bucket_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerImages",
		},
		Permission:             "GetContainerImages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "serviceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerServiceDeployments",
		},
		Permission:             "GetContainerServiceDeployments",
		IsExtra:                true,
		ExtraComponentBodyKey:  "serviceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDisk",
		},
		Permission:             "GetDisk",
		IsExtra:                true,
		ExtraComponentBodyKey:  "diskName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "disk_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDiskSnapshot",
		},
		Permission:             "GetDiskSnapshot",
		IsExtra:                true,
		ExtraComponentBodyKey:  "diskSnapshotName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "disk_snapshot_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDomain",
		},
		Permission:             "GetDomain",
		IsExtra:                true,
		ExtraComponentBodyKey:  "domainName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstance",
		},
		Permission:             "GetInstance",
		IsExtra:                true,
		ExtraComponentBodyKey:  "instanceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstanceAccessDetails",
		},
		Permission:             "GetInstanceAccessDetails",
		IsExtra:                true,
		ExtraComponentBodyKey:  "instanceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstancePortStates",
		},
		Permission:             "GetInstancePortStates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "instanceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstanceSnapshot",
		},
		Permission:             "GetInstanceSnapshot",
		IsExtra:                true,
		ExtraComponentBodyKey:  "instanceSnapshotName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_snapshot_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstanceState",
		},
		Permission:             "GetInstanceState",
		IsExtra:                true,
		ExtraComponentBodyKey:  "instanceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetKeyPair",
		},
		Permission:             "GetKeyPair",
		IsExtra:                true,
		ExtraComponentBodyKey:  "keyPairName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_pair_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetLoadBalancer",
		},
		Permission:             "GetLoadBalancer",
		IsExtra:                true,
		ExtraComponentBodyKey:  "loadBalancerName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "load_balancer_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetLoadBalancerTlsCertificates",
		},
		Permission:             "GetLoadBalancerTlsCertificates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "loadBalancerName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "load_balancer_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetOperation",
		},
		Permission:             "GetOperation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "operationId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "operation_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetOperationsForResource",
		},
		Permission:             "GetOperationsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabase",
		},
		Permission:             "GetRelationalDatabase",
		IsExtra:                true,
		ExtraComponentBodyKey:  "relationalDatabaseName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "relational_database_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseEvents",
		},
		Permission:             "GetRelationalDatabaseEvents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "relationalDatabaseName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "relational_database_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseLogStreams",
		},
		Permission:             "GetRelationalDatabaseLogStreams",
		IsExtra:                true,
		ExtraComponentBodyKey:  "relationalDatabaseName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "relational_database_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseMasterUserPassword",
		},
		Permission:             "GetRelationalDatabaseMasterUserPassword",
		IsExtra:                true,
		ExtraComponentBodyKey:  "relationalDatabaseName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "relational_database_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseParameters",
		},
		Permission:             "GetRelationalDatabaseParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "relationalDatabaseName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "relational_database_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseSnapshot",
		},
		Permission:             "GetRelationalDatabaseSnapshot",
		IsExtra:                true,
		ExtraComponentBodyKey:  "relationalDatabaseSnapshotName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "relational_database_snapshot_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetStaticIp",
		},
		Permission:             "GetStaticIp",
		IsExtra:                true,
		ExtraComponentBodyKey:  "staticIpName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "static_ip_name",
	},
}
