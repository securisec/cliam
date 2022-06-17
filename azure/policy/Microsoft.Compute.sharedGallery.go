package policy

var Microsoft_Compute_sharedGallery = map[string]Policy{

	"SharedGalleries_List": {Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleries_List",
		Resource:    "Microsoft.Compute",
	},
	"SharedGalleryImages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImages_List",
		Resource:    "Microsoft.Compute",
	},
	"SharedGalleryImageVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images/{{.galleryImageName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImageVersions_List",
		Resource:    "Microsoft.Compute",
	},
	"SharedGalleries_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleries_Get",
		Resource:    "Microsoft.Compute",
	},
	"SharedGalleryImages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images/{{.galleryImageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImages_Get",
		Resource:    "Microsoft.Compute",
	},
	"SharedGalleryImageVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/sharedGalleries/{{.galleryUniqueName}}/images/{{.galleryImageName}}/versions/{{.galleryImageVersionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SharedGalleryImageVersions_Get",
		Resource:    "Microsoft.Compute",
	},
}
