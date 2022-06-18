package policy

import "github.com/securisec/cliam/shared"

var ElasticFileSystemPolicies = map[string]Service{
	"DescribeFileSystems": {
		ServiceSuffix: "2015-02-01/file-systems",
		Permission:    "DescribeFileSystems",
	},

	// extra
	"DescribeAcceleratorOfferings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAcceleratorOfferings",
			"Version": "2017-07-25",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeAcceleratorOfferings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "locationType",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "location_type",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
