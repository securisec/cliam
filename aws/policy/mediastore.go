package policy

import "github.com/securisec/cliam/shared"

// MediaStorePolicies policy
var MediaStorePolicies = map[string]Service{
	"DescribeContainer": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.DescribeContainer",
		},
		Permission: "DescribeContainer",
	},
	"ListContainers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.ListContainers",
		},
		Permission: "ListContainers",
	},

	// extra
	"GetContainerPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.GetContainerPolicy",
		},
		Permission:             "GetContainerPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ContainerName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "container_name",
	},
	"GetCorsPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.GetCorsPolicy",
		},
		Permission:             "GetCorsPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ContainerName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "container_name",
	},
	"GetLifecyclePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.GetLifecyclePolicy",
		},
		Permission:             "GetLifecyclePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ContainerName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "container_name",
	},
	"GetMetricPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.GetMetricPolicy",
		},
		Permission:             "GetMetricPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ContainerName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "container_name",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "MediaStore_20170901.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Resource",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource",
	},
}
