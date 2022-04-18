package policy

var ChimePolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "endpoints/messaging-session",
		Permission:    "GetMessagingSessionEndpoint",
	},
	{
		Method:        "GET",
		ServiceSuffix: "accounts",
		Permission:    "ListAccounts",
	},
	{
		Method:        "GET",
		ServiceSuffix: "app-instances",
		Permission:    "ListAppInstances",
	},
	{
		Method:        "GET",
		ServiceSuffix: "channels?scope=app-instance-user-memberships",
		Permission:    "ListChannelMembershipsForAppInstanceUser",
	},
	{
		Method:        "GET",
		ServiceSuffix: "channels?scope=app-instance-user-moderated-channels",
		Permission:    "ListChannelsModeratedByAppInstanceUser",
	},
	{
		Method:        "GET",
		ServiceSuffix: "media-capture-pipelines",
		Permission:    "ListMediaCapturePipelines",
	},
	{
		Method:        "GET",
		ServiceSuffix: "meetings",
		Permission:    "ListMeetings",
	},
	{
		Method:        "GET",
		ServiceSuffix: "phone-number-orders",
		Permission:    "ListPhoneNumberOrders",
	},
	{
		Method:        "GET",
		ServiceSuffix: "phone-numbers",
		Permission:    "ListPhoneNumbers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "sip-media-applications",
		Permission:    "ListSipMediaApplications",
	},
	{
		Method:        "GET",
		ServiceSuffix: "sip-rules",
		Permission:    "ListSipRules",
	},
	{
		Method:        "GET",
		ServiceSuffix: "voice-connector-groups",
		Permission:    "ListVoiceConnectorGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "voice-connectors",
		Permission:    "ListVoiceConnectors",
	},
	{
		Method:        "GET",
		ServiceSuffix: "search?type=phone-numbers",
		Permission:    "SearchAvailablePhoneNumbers",
	},
}
