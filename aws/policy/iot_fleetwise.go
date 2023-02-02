package policy

import "github.com/securisec/cliam/shared"

// IOTFleetWisePolicies policy
var IOTFleetWisePolicies = map[string]Service{
	"GetLoggingOptions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetLoggingOptions",
		},
		Permission: "GetLoggingOptions",
	},
	"GetRegisterAccountStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetRegisterAccountStatus",
		},
		Permission: "GetRegisterAccountStatus",
	},
	"ListCampaigns": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListCampaigns",
		},
		Permission: "ListCampaigns",
	},
	"ListDecoderManifests": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListDecoderManifests",
		},
		Permission: "ListDecoderManifests",
	},
	"ListFleets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListFleets",
		},
		Permission: "ListFleets",
	},
	"ListModelManifests": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListModelManifests",
		},
		Permission: "ListModelManifests",
	},
	"ListSignalCatalogs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListSignalCatalogs",
		},
		Permission: "ListSignalCatalogs",
	},
	"ListVehicles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListVehicles",
		},
		Permission: "ListVehicles",
	},

	// extra
	"GetCampaign": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetCampaign",
		},
		Permission:             "GetCampaign",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetDecoderManifest": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetDecoderManifest",
		},
		Permission:             "GetDecoderManifest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetFleet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetFleet",
		},
		Permission:             "GetFleet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "fleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	"GetModelManifest": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetModelManifest",
		},
		Permission:             "GetModelManifest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetSignalCatalog": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetSignalCatalog",
		},
		Permission:             "GetSignalCatalog",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetVehicle": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetVehicle",
		},
		Permission:             "GetVehicle",
		IsExtra:                true,
		ExtraComponentBodyKey:  "vehicleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vehicle_name",
	},
	"GetVehicleStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.GetVehicleStatus",
		},
		Permission:             "GetVehicleStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "vehicleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vehicle_name",
	},
	"ListDecoderManifestNetworkInterfaces": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListDecoderManifestNetworkInterfaces",
		},
		Permission:             "ListDecoderManifestNetworkInterfaces",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListDecoderManifestSignals": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListDecoderManifestSignals",
		},
		Permission:             "ListDecoderManifestSignals",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListFleetsForVehicle": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListFleetsForVehicle",
		},
		Permission:             "ListFleetsForVehicle",
		IsExtra:                true,
		ExtraComponentBodyKey:  "vehicleName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vehicle_name",
	},
	"ListModelManifestNodes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListModelManifestNodes",
		},
		Permission:             "ListModelManifestNodes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListSignalCatalogNodes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListSignalCatalogNodes",
		},
		Permission:             "ListSignalCatalogNodes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListVehiclesInFleet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "IoTAutobahnControlPlane.ListVehiclesInFleet",
		},
		Permission:             "ListVehiclesInFleet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "fleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
}
