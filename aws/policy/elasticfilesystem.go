package policy

// ElasticFileSystemPolicies policy
var ElasticFileSystemPolicies = map[string]Service{
	"DescribeAccessPoints": {
		Method:        "GET",
		ServiceSuffix: "2015-02-01/access-points",
		Permission:    "DescribeAccessPoints",
	},
	"DescribeAccountPreferences": {
		Method:        "GET",
		ServiceSuffix: "2015-02-01/account-preferences",
		Permission:    "DescribeAccountPreferences",
	},
	"DescribeFileSystems": {
		Method:        "GET",
		ServiceSuffix: "2015-02-01/file-systems",
		Permission:    "DescribeFileSystems",
	},
	"DescribeMountTargets": {
		Method:        "GET",
		ServiceSuffix: "2015-02-01/mount-targets",
		Permission:    "DescribeMountTargets",
	},
	"DescribeReplicationConfigurations": {
		Method:        "GET",
		ServiceSuffix: "2015-02-01/file-systems/replication-configurations",
		Permission:    "DescribeReplicationConfigurations",
	},

	// extra
	"DescribeBackupPolicy": {
		ServiceSuffix:          "/2015-02-01/file-systems/{{.file_system_id}}/backup-policy",
		Permission:             "DescribeBackupPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "file_system_id",
	},
	"DescribeFileSystemPolicy": {
		ServiceSuffix:          "/2015-02-01/file-systems/{{.file_system_id}}/policy",
		Permission:             "DescribeFileSystemPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "file_system_id",
	},
	"DescribeLifecycleConfiguration": {
		ServiceSuffix:          "/2015-02-01/file-systems/{{.file_system_id}}/lifecycle-configuration",
		Permission:             "DescribeLifecycleConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "file_system_id",
	},
	"DescribeMountTargetSecurityGroups": {
		ServiceSuffix:          "/2015-02-01/mount-targets/{{.mount_target_id}}/security-groups",
		Permission:             "DescribeMountTargetSecurityGroups",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "mount_target_id",
	},
	"DescribeTags": {
		ServiceSuffix:          "/2015-02-01/tags/{{.file_system_id}}/",
		Permission:             "DescribeTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "file_system_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/2015-02-01/resource-tags/{{.resource_id}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_id",
	},
}
