package policy

var GroundStationPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "config",
		Permission:    "ListConfigs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "dataflowEndpointGroup",
		Permission:    "ListDataflowEndpointGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "groundstation",
		Permission:    "ListGroundStations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "missionprofile",
		Permission:    "ListMissionProfiles",
	},
	{
		Method:        "GET",
		ServiceSuffix: "satellite",
		Permission:    "ListSatellites",
	},

	// extra
	{
		ServiceSuffix:          "/contact/{{.contact_id}}",
		Permission:             "DescribeContact",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "contact_id",
	},
	{
		ServiceSuffix:          "/dataflowEndpointGroup/{{.dataflow_endpoint_group_id}}",
		Permission:             "GetDataflowEndpointGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "dataflow_endpoint_group_id",
	},
	{
		ServiceSuffix:          "/missionprofile/{{.mission_profile_id}}",
		Permission:             "GetMissionProfile",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "mission_profile_id",
	},
	{
		ServiceSuffix:          "/satellite/{{.satellite_id}}",
		Permission:             "GetSatellite",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "satellite_id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
