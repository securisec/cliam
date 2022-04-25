package policy

var OutpostsPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "catalog/items",
		Permission:    "ListCatalogItems",
	},
	{
		Method:        "GET",
		ServiceSuffix: "list-orders",
		Permission:    "ListOrders",
	},
	{
		Method:        "GET",
		ServiceSuffix: "outposts",
		Permission:    "ListOutposts",
	},
	{
		Method:        "GET",
		ServiceSuffix: "sites",
		Permission:    "ListSites",
	},

	// extra
	{
		ServiceSuffix:          "/catalog/item/{{.catalog_item_id}}",
		Permission:             "GetCatalogItem",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "catalog_item_id",
	},
	{
		ServiceSuffix:          "/orders/{{.order_id}}",
		Permission:             "GetOrder",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "order_id",
	},
	{
		ServiceSuffix:          "/outposts/{{.outpost_id}}",
		Permission:             "GetOutpost",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "outpost_id",
	},
	{
		ServiceSuffix:          "/outposts/{{.outpost_id}}/instanceTypes",
		Permission:             "GetOutpostInstanceTypes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "outpost_id",
	},
	{
		ServiceSuffix:          "/sites/{{.site_id}}",
		Permission:             "GetSite",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "site_id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
