package policy

// OutpostsPolicies policy
var OutpostsPolicies = map[string]Service{
	"ListCatalogItems": {
		Method:        "GET",
		ServiceSuffix: "catalog/items",
		Permission:    "ListCatalogItems",
	},
	"ListOrders": {
		Method:        "GET",
		ServiceSuffix: "list-orders",
		Permission:    "ListOrders",
	},
	"ListOutposts": {
		Method:        "GET",
		ServiceSuffix: "outposts",
		Permission:    "ListOutposts",
	},
	"ListSites": {
		Method:        "GET",
		ServiceSuffix: "sites",
		Permission:    "ListSites",
	},

	// extra
	"GetCatalogItem": {
		ServiceSuffix:          "/catalog/item/{{.catalog_item_id}}",
		Permission:             "GetCatalogItem",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "catalog_item_id",
	},
	"GetConnection": {
		ServiceSuffix:          "/connections/{{.connection_id}}",
		Permission:             "GetConnection",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "connection_id",
	},
	"GetOrder": {
		ServiceSuffix:          "/orders/{{.order_id}}",
		Permission:             "GetOrder",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "order_id",
	},
	"GetOutpost": {
		ServiceSuffix:          "/outposts/{{.outpost_id}}",
		Permission:             "GetOutpost",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "outpost_id",
	},
	"GetOutpostInstanceTypes": {
		ServiceSuffix:          "/outposts/{{.outpost_id}}/instanceTypes",
		Permission:             "GetOutpostInstanceTypes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "outpost_id",
	},
	"GetSite": {
		ServiceSuffix:          "/sites/{{.site_id}}",
		Permission:             "GetSite",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "site_id",
	},
	"ListAssets": {
		ServiceSuffix:          "/outposts/{{.outpost_identifier}}/assets",
		Permission:             "ListAssets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "outpost_identifier",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
