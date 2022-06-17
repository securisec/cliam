package policy

var Microsoft_Compute_gallery = map[string]Policy{
	"Galleries_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Galleries_Get",
		Resource:    "Microsoft.Compute",
	},
	"GalleryImages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/images/{{.galleryImageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryImages_Get",
		Resource:    "Microsoft.Compute",
	},
	"GalleryImageVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/images/{{.galleryImageName}}/versions/{{.galleryImageVersionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryImageVersions_Get",
		Resource:    "Microsoft.Compute",
	}, "GalleryApplications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/applications/{{.galleryApplicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryApplications_Get",
		Resource:    "Microsoft.Compute",
	}, "GalleryApplicationVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/applications/{{.galleryApplicationName}}/versions/{{.galleryApplicationVersionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryApplicationVersions_Get",
		Resource:    "Microsoft.Compute",
	}, "Galleries_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Galleries_ListByResourceGroup",
		Resource:    "Microsoft.Compute",
	}, "Galleries_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/galleries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Galleries_List",
		Resource:    "Microsoft.Compute",
	}, "GalleryImages_ListByGallery": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/images",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryImages_ListByGallery",
		Resource:    "Microsoft.Compute",
	}, "GalleryImageVersions_ListByGalleryImage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/images/{{.galleryImageName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryImageVersions_ListByGalleryImage",
		Resource:    "Microsoft.Compute",
	}, "GalleryApplications_ListByGallery": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryApplications_ListByGallery",
		Resource:    "Microsoft.Compute",
	}, "GalleryApplicationVersions_ListByGalleryApplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/applications/{{.galleryApplicationName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GalleryApplicationVersions_ListByGalleryApplication",
		Resource:    "Microsoft.Compute",
	}, "GallerySharingProfile_Update": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/galleries/{{.galleryName}}/share",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "GallerySharingProfile_Update",
		Resource:    "Microsoft.Compute",
	},
}
