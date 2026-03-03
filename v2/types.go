package katapultpro

// PhotoSize specifies the image size variant to retrieve.
type PhotoSize string

const (
	// PhotoSizeFull is the original full-resolution image.
	PhotoSizeFull PhotoSize = "full"
	// PhotoSizeExtraLarge is the size typically used when viewing photos in Katapult Pro.
	PhotoSizeExtraLarge PhotoSize = "extra_large"
	// PhotoSizeLarge is a large thumbnail.
	PhotoSizeLarge PhotoSize = "large"
	// PhotoSizeSmall is a small thumbnail.
	PhotoSizeSmall PhotoSize = "small"
	// PhotoSizeTiny is the smallest thumbnail.
	PhotoSizeTiny PhotoSize = "tiny"
)

// String returns the API value.
func (s PhotoSize) String() string { return string(s) }

// IsValid reports whether s is a defined photo size.
func (s PhotoSize) IsValid() bool {
	switch s {
	case PhotoSizeFull, PhotoSizeExtraLarge, PhotoSizeLarge, PhotoSizeSmall, PhotoSizeTiny:
		return true
	}
	return false
}

// PhotoURLResponse is the response from GetPhotoURL.
type PhotoURLResponse struct {
	URL string `json:"url"`
}
