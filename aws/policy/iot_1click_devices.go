package policy

// IOT1ClickDevicesPolicies policy
var IOT1ClickDevicesPolicies = map[string]Service{
	"ListDevices": {
		Method:        "GET",
		ServiceSuffix: "devices",
		Permission:    "ListDevices",
	},

	// extra
	"DescribeDevice": {
		ServiceSuffix:          "/devices/{{.device_id}}",
		Permission:             "DescribeDevice",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "device_id",
	},
	"GetDeviceMethods": {
		ServiceSuffix:          "/devices/{{.device_id}}/methods",
		Permission:             "GetDeviceMethods",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "device_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
