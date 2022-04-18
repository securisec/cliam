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
}
