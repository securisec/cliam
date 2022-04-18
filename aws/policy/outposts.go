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
}
