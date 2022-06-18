package policy

var GroundStationPolicies = map[string]Service{
	"ListConfigs": {
		Method:        "GET",
		ServiceSuffix: "config",
		Permission:    "ListConfigs",
	},
	"ListDataflowEndpointGroups": {
		Method:        "GET",
		ServiceSuffix: "dataflowEndpointGroup",
		Permission:    "ListDataflowEndpointGroups",
	},
	"ListGroundStations": {
		Method:        "GET",
		ServiceSuffix: "groundstation",
		Permission:    "ListGroundStations",
	},
	"ListMissionProfiles": {
		Method:        "GET",
		ServiceSuffix: "missionprofile",
		Permission:    "ListMissionProfiles",
	},
	"ListSatellites": {
		Method:        "GET",
		ServiceSuffix: "satellite",
		Permission:    "ListSatellites",
	},

	// extra
	"DescribeContact": {
		ServiceSuffix:          "/contact/{{.contact_id}}",
		Permission:             "DescribeContact",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "contact_id",
	},
	"GetDataflowEndpointGroup": {
		ServiceSuffix:          "/dataflowEndpointGroup/{{.dataflow_endpoint_group_id}}",
		Permission:             "GetDataflowEndpointGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "dataflow_endpoint_group_id",
	},
	"GetMissionProfile": {
		ServiceSuffix:          "/missionprofile/{{.mission_profile_id}}",
		Permission:             "GetMissionProfile",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "mission_profile_id",
	},
	"GetSatellite": {
		ServiceSuffix:          "/satellite/{{.satellite_id}}",
		Permission:             "GetSatellite",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "satellite_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
