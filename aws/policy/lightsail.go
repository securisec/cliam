package policy

import "github.com/securisec/cliam/shared"

var LightsailPolicies = map[string]Service{
	"GetActiveNames": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetActiveNames",
		},
		Permission: "GetActiveNames",
	},
	"GetAlarms": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetAlarms",
		},
		Permission: "GetAlarms",
	},
	"GetBlueprints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBlueprints",
		},
		Permission: "GetBlueprints",
	},
	"GetBucketBundles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBucketBundles",
		},
		Permission: "GetBucketBundles",
	},
	"GetBuckets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBuckets",
		},
		Permission: "GetBuckets",
	},
	"GetBundles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetBundles",
		},
		Permission: "GetBundles",
	},
	"GetCertificates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetCertificates",
		},
		Permission: "GetCertificates",
	},
	"GetCloudFormationStackRecords": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetCloudFormationStackRecords",
		},
		Permission: "GetCloudFormationStackRecords",
	},
	"GetContactMethods": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContactMethods",
		},
		Permission: "GetContactMethods",
	},
	"GetContainerApiMetadata": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerApiMetadata",
		},
		Permission: "GetContainerApiMetadata",
	},
	"GetContainerServicePowers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerServicePowers",
		},
		Permission: "GetContainerServicePowers",
	},
	"GetContainerServices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetContainerServices",
		},
		Permission: "GetContainerServices",
	},
	"GetDiskSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDiskSnapshots",
		},
		Permission: "GetDiskSnapshots",
	},
	"GetDisks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDisks",
		},
		Permission: "GetDisks",
	},
	"GetDistributionBundles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDistributionBundles",
		},
		Permission: "GetDistributionBundles",
	},
	"GetDistributionLatestCacheReset": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDistributionLatestCacheReset",
		},
		Permission: "GetDistributionLatestCacheReset",
	},
	"GetDistributions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDistributions",
		},
		Permission: "GetDistributions",
	},
	"GetDomains": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetDomains",
		},
		Permission: "GetDomains",
	},
	"GetExportSnapshotRecords": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetExportSnapshotRecords",
		},
		Permission: "GetExportSnapshotRecords",
	},
	"GetInstanceSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstanceSnapshots",
		},
		Permission: "GetInstanceSnapshots",
	},
	"GetInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetInstances",
		},
		Permission: "GetInstances",
	},
	"GetKeyPairs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetKeyPairs",
		},
		Permission: "GetKeyPairs",
	},
	"GetLoadBalancers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetLoadBalancers",
		},
		Permission: "GetLoadBalancers",
	},
	"GetOperations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetOperations",
		},
		Permission: "GetOperations",
	},
	"GetRegions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRegions",
		},
		Permission: "GetRegions",
	},
	"GetRelationalDatabaseBlueprints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseBlueprints",
		},
		Permission: "GetRelationalDatabaseBlueprints",
	},
	"GetRelationalDatabaseBundles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseBundles",
		},
		Permission: "GetRelationalDatabaseBundles",
	},
	"GetRelationalDatabaseSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabaseSnapshots",
		},
		Permission: "GetRelationalDatabaseSnapshots",
	},
	"GetRelationalDatabases": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetRelationalDatabases",
		},
		Permission: "GetRelationalDatabases",
	},
	"GetStaticIps": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Lightsail_20161128.GetStaticIps",
		},
		Permission: "GetStaticIps",
	},

	// extra
	"GetAutoSnapshots": {
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
	"GetBucketAccessKeys": {
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
	"GetContainerImages": {
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
	"GetContainerServiceDeployments": {
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
	"GetDisk": {
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
	"GetDiskSnapshot": {
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
	"GetDomain": {
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
	"GetInstance": {
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
	"GetInstanceAccessDetails": {
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
	"GetInstancePortStates": {
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
	"GetInstanceSnapshot": {
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
	"GetInstanceState": {
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
	"GetKeyPair": {
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
	"GetLoadBalancer": {
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
	"GetLoadBalancerTlsCertificates": {
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
	"GetOperation": {
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
	"GetOperationsForResource": {
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
	"GetRelationalDatabase": {
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
	"GetRelationalDatabaseEvents": {
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
	"GetRelationalDatabaseLogStreams": {
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
	"GetRelationalDatabaseMasterUserPassword": {
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
	"GetRelationalDatabaseParameters": {
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
	"GetRelationalDatabaseSnapshot": {
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
	"GetStaticIp": {
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
