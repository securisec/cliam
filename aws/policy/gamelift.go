package policy

import "github.com/securisec/cliam/shared"

var GameLiftPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.CreateBuild",
		},
		Permission: "CreateBuild",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.CreateScript",
		},
		Permission: "CreateScript",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeEC2InstanceLimits",
		},
		Permission: "DescribeEC2InstanceLimits",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeFleetAttributes",
		},
		Permission: "DescribeFleetAttributes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeFleetCapacity",
		},
		Permission: "DescribeFleetCapacity",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeFleetUtilization",
		},
		Permission: "DescribeFleetUtilization",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeGameSessionDetails",
		},
		Permission: "DescribeGameSessionDetails",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeGameSessionQueues",
		},
		Permission: "DescribeGameSessionQueues",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeGameSessions",
		},
		Permission: "DescribeGameSessions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeMatchmakingConfigurations",
		},
		Permission: "DescribeMatchmakingConfigurations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeMatchmakingRuleSets",
		},
		Permission: "DescribeMatchmakingRuleSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribePlayerSessions",
		},
		Permission: "DescribePlayerSessions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeVpcPeeringAuthorizations",
		},
		Permission: "DescribeVpcPeeringAuthorizations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeVpcPeeringConnections",
		},
		Permission: "DescribeVpcPeeringConnections",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListAliases",
		},
		Permission: "ListAliases",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListBuilds",
		},
		Permission: "ListBuilds",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListFleets",
		},
		Permission: "ListFleets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListGameServerGroups",
		},
		Permission: "ListGameServerGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListScripts",
		},
		Permission: "ListScripts",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.SearchGameSessions",
		},
		Permission: "SearchGameSessions",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeAlias",
		},
		Permission:             "DescribeAlias",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AliasId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "alias_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeBuild",
		},
		Permission:             "DescribeBuild",
		IsExtra:                true,
		ExtraComponentBodyKey:  "BuildId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "build_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeFleetEvents",
		},
		Permission:             "DescribeFleetEvents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeFleetLocationAttributes",
		},
		Permission:             "DescribeFleetLocationAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeFleetPortSettings",
		},
		Permission:             "DescribeFleetPortSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeGameServerGroup",
		},
		Permission:             "DescribeGameServerGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GameServerGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "game_server_group_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeGameServerInstances",
		},
		Permission:             "DescribeGameServerInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GameServerGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "game_server_group_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeGameSessionPlacement",
		},
		Permission:             "DescribeGameSessionPlacement",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PlacementId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "placement_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeInstances",
		},
		Permission:             "DescribeInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeMatchmaking",
		},
		Permission:             "DescribeMatchmaking",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TicketIds",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "ticket_ids",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeRuntimeConfiguration",
		},
		Permission:             "DescribeRuntimeConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeScalingPolicies",
		},
		Permission:             "DescribeScalingPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.DescribeScript",
		},
		Permission:             "DescribeScript",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ScriptId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "script_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.GetGameSessionLogUrl",
		},
		Permission:             "GetGameSessionLogUrl",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GameSessionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "game_session_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListGameServers",
		},
		Permission:             "ListGameServers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GameServerGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "game_server_group_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "GameLift.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
