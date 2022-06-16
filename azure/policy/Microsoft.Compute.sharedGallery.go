package policy

var Microsoft_Compute_sharedGallery = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleries_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImages_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images/{{.galleryImageName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImageVersions_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleries_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images/{{.galleryImageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImages_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images/{{.galleryImageName}}/versions/{{.galleryImageVersionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImageVersions_Get",
		Resource:    "Microsoft.Compute",
	},
}
