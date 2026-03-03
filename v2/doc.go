// Package katapultpro provides a Go client for the Katapult Pro API v2.
//
// The v2 API provides read-only access to job data and photo URLs.
// For full CRUD operations, use the v3 SDK.
//
// # Quick Start
//
//	client, err := katapultpro.NewClient("your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Get a photo download URL
//	photoURL, err := client.GetPhotoURL(ctx, "job-id", "photo-id", katapultpro.PhotoSizeExtraLarge)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(photoURL.URL)
//
// # Private Servers
//
// For private Katapult Pro servers, use WithBaseURL:
//
//	client, err := katapultpro.NewClient("api-key",
//	    katapultpro.WithBaseURL("https://yourcompany.katapultpro.com/api"),
//	)
package katapultpro
