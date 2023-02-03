package policy

// Microsoft_EventGrid_EventGrid policy
var Microsoft_EventGrid_EventGrid = map[string]Policy{
	"Channels_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces/{{.partnerNamespaceName}}/channels/{{.channelName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Channels_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"Channels_ListByPartnerNamespace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces/{{.partnerNamespaceName}}/channels",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Channels_ListByPartnerNamespace",
		Resource:    "Microsoft.EventGrid",
	},
	"Channels_GetFullUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces/{{.partnerNamespaceName}}/channels/{{.channelName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Channels_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"Domains_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Domains_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"Domains_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/domains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Domains_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"Domains_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Domains_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"Domains_ListSharedAccessKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Domains_ListSharedAccessKeys",
		Resource:    "Microsoft.EventGrid",
	},
	"Domains_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Domains_RegenerateKey",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainTopics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics/{{.domainTopicName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainTopics_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainTopics_ListByDomain": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainTopics_ListByDomain",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicEventSubscriptions_GetDeliveryAttributes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getDeliveryAttributes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicEventSubscriptions_GetDeliveryAttributes",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainEventSubscriptions_GetDeliveryAttributes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/eventSubscriptions/{{.eventSubscriptionName}}/getDeliveryAttributes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainEventSubscriptions_GetDeliveryAttributes",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.EventGrid/eventSubscriptions/{{.eventSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainTopicEventSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics/{{.topicName}}/eventSubscriptions/{{.eventSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainTopicEventSubscriptions_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicEventSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}/eventSubscriptions/{{.eventSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicEventSubscriptions_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainEventSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/eventSubscriptions/{{.eventSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainEventSubscriptions_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopicEventSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/systemTopics/{{.systemTopicName}}/eventSubscriptions/{{.eventSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopicEventSubscriptions_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopicEventSubscriptions_GetFullUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/systemTopics/{{.systemTopicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopicEventSubscriptions_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_GetFullUrl": {
		Path:   "/{{.scope}}/providers/Microsoft.EventGrid/eventSubscriptions/{{.eventSubscriptionName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopicEventSubscriptions_ListBySystemTopic": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/systemTopics/{{.systemTopicName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopicEventSubscriptions_ListBySystemTopic",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopicEventSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}/eventSubscriptions/{{.eventSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopicEventSubscriptions_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopicEventSubscriptions_GetFullUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopicEventSubscriptions_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopicEventSubscriptions_ListByPartnerTopic": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopicEventSubscriptions_ListByPartnerTopic",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainTopicEventSubscriptions_GetFullUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics/{{.topicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainTopicEventSubscriptions_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicEventSubscriptions_GetFullUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicEventSubscriptions_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainEventSubscriptions_GetFullUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/eventSubscriptions/{{.eventSubscriptionName}}/getFullUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainEventSubscriptions_GetFullUrl",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListGlobalBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListGlobalBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListGlobalBySubscriptionForTopicType": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/topicTypes/{{.topicTypeName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListGlobalBySubscriptionForTopicType",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListGlobalByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListGlobalByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListGlobalByResourceGroupForTopicType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topicTypes/{{.topicTypeName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListGlobalByResourceGroupForTopicType",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListRegionalBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/locations/{{.location}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListRegionalBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListRegionalByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/locations/{{.location}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListRegionalByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListRegionalBySubscriptionForTopicType": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/locations/{{.location}}/topicTypes/{{.topicTypeName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListRegionalBySubscriptionForTopicType",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListRegionalByResourceGroupForTopicType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/locations/{{.location}}/topicTypes/{{.topicTypeName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListRegionalByResourceGroupForTopicType",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.providerNamespace}}/{{.resourceTypeName}}/{{.resourceName}}/providers/Microsoft.EventGrid/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListByResource",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicEventSubscriptions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicEventSubscriptions_List",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainEventSubscriptions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainEventSubscriptions_List",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_ListByDomainTopic": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics/{{.topicName}}/providers/Microsoft.EventGrid/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_ListByDomainTopic",
		Resource:    "Microsoft.EventGrid",
	},
	"EventSubscriptions_GetDeliveryAttributes": {
		Path:   "/{{.scope}}/providers/Microsoft.EventGrid/eventSubscriptions/{{.eventSubscriptionName}}/getDeliveryAttributes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "EventSubscriptions_GetDeliveryAttributes",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopicEventSubscriptions_GetDeliveryAttributes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/systemTopics/{{.systemTopicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getDeliveryAttributes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopicEventSubscriptions_GetDeliveryAttributes",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainTopicEventSubscriptions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics/{{.topicName}}/eventSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainTopicEventSubscriptions_List",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopicEventSubscriptions_GetDeliveryAttributes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getDeliveryAttributes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopicEventSubscriptions_GetDeliveryAttributes",
		Resource:    "Microsoft.EventGrid",
	},
	"DomainTopicEventSubscriptions_GetDeliveryAttributes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/domains/{{.domainName}}/topics/{{.topicName}}/eventSubscriptions/{{.eventSubscriptionName}}/getDeliveryAttributes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "DomainTopicEventSubscriptions_GetDeliveryAttributes",
		Resource:    "Microsoft.EventGrid",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.EventGrid/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.EventGrid",
	},
	"Topics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Topics_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"Topics_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/topics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Topics_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"Topics_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Topics_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"Topics_ListSharedAccessKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Topics_ListSharedAccessKeys",
		Resource:    "Microsoft.EventGrid",
	},
	"Topics_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/topics/{{.topicName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Topics_RegenerateKey",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerConfigurations/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerConfigurations_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerConfigurations_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerConfigurations_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerConfigurations_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/partnerConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerConfigurations_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerConfigurations_AuthorizePartner": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerConfigurations/default/authorizePartner",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerConfigurations_AuthorizePartner",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerConfigurations_UnauthorizePartner": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerConfigurations/default/unauthorizePartner",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerConfigurations_UnauthorizePartner",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerNamespaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces/{{.partnerNamespaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerNamespaces_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerNamespaces_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/partnerNamespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerNamespaces_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerNamespaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerNamespaces_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerNamespaces_ListSharedAccessKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces/{{.partnerNamespaceName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerNamespaces_ListSharedAccessKeys",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerNamespaces_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerNamespaces/{{.partnerNamespaceName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerNamespaces_RegenerateKey",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerRegistrations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerRegistrations/{{.partnerRegistrationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerRegistrations_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerRegistrations_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/partnerRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerRegistrations_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerRegistrations_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerRegistrations_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopics_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopics_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/partnerTopics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopics_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopics_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopics_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopics_Activate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}/activate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopics_Activate",
		Resource:    "Microsoft.EventGrid",
	},
	"PartnerTopics_Deactivate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/partnerTopics/{{.partnerTopicName}}/deactivate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PartnerTopics_Deactivate",
		Resource:    "Microsoft.EventGrid",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/{{.parentType}}/{{.parentName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PrivateEndpointConnections_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/{{.parentType}}/{{.parentName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PrivateEndpointConnections_ListByResource",
		Resource:    "Microsoft.EventGrid",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/{{.parentType}}/{{.parentName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"PrivateLinkResources_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/{{.parentType}}/{{.parentName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "PrivateLinkResources_ListByResource",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/systemTopics/{{.systemTopicName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopics_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopics_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EventGrid/systemTopics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopics_ListBySubscription",
		Resource:    "Microsoft.EventGrid",
	},
	"SystemTopics_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EventGrid/systemTopics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "SystemTopics_ListByResourceGroup",
		Resource:    "Microsoft.EventGrid",
	},
	"Topics_ListEventTypes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/{{.providerNamespace}}/{{.resourceTypeName}}/{{.resourceName}}/providers/Microsoft.EventGrid/eventTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "Topics_ListEventTypes",
		Resource:    "Microsoft.EventGrid",
	},
	"ExtensionTopics_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.EventGrid/extensionTopics/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "ExtensionTopics_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicTypes_List": {
		Path:   "/providers/Microsoft.EventGrid/topicTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicTypes_List",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicTypes_Get": {
		Path:   "/providers/Microsoft.EventGrid/topicTypes/{{.topicTypeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicTypes_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"TopicTypes_ListEventTypes": {
		Path:   "/providers/Microsoft.EventGrid/topicTypes/{{.topicTypeName}}/eventTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "TopicTypes_ListEventTypes",
		Resource:    "Microsoft.EventGrid",
	},
	"VerifiedPartners_Get": {
		Path:   "/providers/Microsoft.EventGrid/verifiedPartners/{{.verifiedPartnerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "VerifiedPartners_Get",
		Resource:    "Microsoft.EventGrid",
	},
	"VerifiedPartners_List": {
		Path:   "/providers/Microsoft.EventGrid/verifiedPartners",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "VerifiedPartners_List",
		Resource:    "Microsoft.EventGrid",
	},
}
