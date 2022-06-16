package policy

var Microsoft_Compute_communityGallery = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/communityGalleries/{{.publicGalleryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CommunityGalleries_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/communityGalleries/{{.publicGalleryName}}/images/{{.galleryImageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CommunityGalleryImages_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/communityGalleries/{{.publicGalleryName}}/images/{{.galleryImageName}}/versions/{{.galleryImageVersionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "CommunityGalleryImageVersions_Get",
		Resource:    "Microsoft.Compute",
	},
}
