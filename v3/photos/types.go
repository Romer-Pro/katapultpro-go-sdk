package photos

import "encoding/json"

// Photo represents a photo record in a job (v3).
type Photo struct {
	ID string `json:"id,omitempty"`
}

// PhotoAssociationQuery is the association_value query param for node/section photo uploads.
// Valid values are Main ("main") or True ("true"); the API default is "true".
type PhotoAssociationQuery string

const (
	PhotoAssociationQueryMain PhotoAssociationQuery = "main"
	PhotoAssociationQueryTrue PhotoAssociationQuery = "true"
)

// String returns the API value.
func (q PhotoAssociationQuery) String() string { return string(q) }

// IsValid reports whether q is a defined value.
func (q PhotoAssociationQuery) IsValid() bool {
	return q == PhotoAssociationQueryMain || q == PhotoAssociationQueryTrue
}

// PhotoAssociationValue is the association_value body for AssociatePhoto.
// Valid values: PhotoAssociationMain ("main"), PhotoAssociationTrue (boolean true), or nil (unassociate).
// Use PtrPhotoAssociationMain(), PtrPhotoAssociationTrue(), or nil when building AssociatePhotoRequest.
type PhotoAssociationValue interface {
	json.Marshaler
	sealedPhotoAssociation()
}

type photoAssociationMain struct{}

func (photoAssociationMain) sealedPhotoAssociation() {}

func (photoAssociationMain) MarshalJSON() ([]byte, error) {
	return []byte(`"main"`), nil
}

type photoAssociationTrue struct{}

func (photoAssociationTrue) sealedPhotoAssociation() {}

func (photoAssociationTrue) MarshalJSON() ([]byte, error) {
	return []byte(`true`), nil
}

var (
	// PhotoAssociationMain is the "main" association value for AssociatePhoto.
	PhotoAssociationMain PhotoAssociationValue = photoAssociationMain{}
	// PhotoAssociationTrue is the true association value for AssociatePhoto.
	PhotoAssociationTrue PhotoAssociationValue = photoAssociationTrue{}
)

// PtrPhotoAssociationMain returns a pointer suitable for AssociatePhotoRequest.AssociationValue ("main").
func PtrPhotoAssociationMain() *PhotoAssociationValue {
	v := PhotoAssociationMain
	return &v
}

// PtrPhotoAssociationTrue returns a pointer suitable for AssociatePhotoRequest.AssociationValue (true).
func PtrPhotoAssociationTrue() *PhotoAssociationValue {
	v := PhotoAssociationTrue
	return &v
}

// AssociatePhotoRequest is the body for POST /v3/jobs/:job_id/photos/:photo_id/associate.
type AssociatePhotoRequest struct {
	NodeID           string                 `json:"node_id,omitempty"`
	SectionID        string                 `json:"section_id,omitempty"`
	ConnectionID     string                 `json:"connection_id,omitempty"`
	AssociationValue *PhotoAssociationValue  `json:"association_value"`
}

// PixelSelection is the API shape for photo element and calibration anchor location.
// Uses percent coordinates: percentX and percentY (0–100), not raw pixel values.
// See: https://github.com/KatapultDevelopment/katapult-pro-api-documentation/blob/main/v3/README.md
type PixelSelection struct {
	PercentX float64 `json:"percentX"`
	PercentY float64 `json:"percentY"`
}

// PhotoElement represents an element on a photo (v3). Attributes are a flat map and may also appear at root on the response.
type PhotoElement struct {
	ID              string                 `json:"id,omitempty"`
	ElementType     string                 `json:"element_type,omitempty"`
	PixelSelection  *PixelSelection       `json:"pixel_selection,omitempty"`
	ManualHeight    string                 `json:"manual_height,omitempty"` // Feet-inches notation, e.g. "16-6".
	Attributes      map[string]interface{} `json:"attributes,omitempty"`    // Flat map per API.
	ParentID        string                 `json:"parent_id,omitempty"`
	TraceID         string                 `json:"trace_id,omitempty"`
}

// CreatePhotoElementRequest is the body for POST /v3/jobs/:job_id/photos/:photo_id/photo_elements.
type CreatePhotoElementRequest struct {
	ElementType     string                 `json:"element_type"`              // Required.
	PixelSelection  *PixelSelection       `json:"pixel_selection,omitempty"`
	ManualHeight    string                 `json:"manual_height,omitempty"`  // Feet-inches, e.g. "16-6".
	Attributes      map[string]interface{} `json:"attributes,omitempty"`     // Flat map.
	ParentID        string                 `json:"parent_id,omitempty"`     // Child element; omit to not nest.
	TraceID         string                 `json:"trace_id,omitempty"`
}

// UpdatePhotoElementRequest is the body for POST /v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id.
// ElementType can only be set if the element does not already exist.
type UpdatePhotoElementRequest struct {
	ElementType     string                 `json:"element_type,omitempty"`
	PixelSelection  *PixelSelection       `json:"pixel_selection,omitempty"`
	ManualHeight    string                 `json:"manual_height,omitempty"`
	Attributes      map[string]interface{} `json:"attributes,omitempty"`
	ParentID        *string                `json:"parent_id,omitempty"`     // Set to nil to de-nest.
	TraceID         string                 `json:"trace_id,omitempty"`
}

// UpdatePhotoElementOptions are optional query parameters for UpdateElement.
type UpdatePhotoElementOptions struct {
	OnlyIfExists bool // If true, update only when the element exists (no create with specified id).
}

// PhotoCalibrationAnchor represents a calibration anchor on a photo (v3). Height in decimal feet.
type PhotoCalibrationAnchor struct {
	ID              string           `json:"id,omitempty"`
	PixelSelection  *PixelSelection `json:"pixel_selection,omitempty"`
	Height          float64          `json:"height,omitempty"` // Decimal feet.
}

// CreatePhotoCalibrationAnchorRequest is the body for POST /v3/jobs/:job_id/photos/:photo_id/calibration_anchors.
// Both pixel_selection and height are required by the API.
type CreatePhotoCalibrationAnchorRequest struct {
	PixelSelection PixelSelection `json:"pixel_selection"` // Required.
	Height         float64        `json:"height"`          // Required; decimal feet.
}

// UpdatePhotoCalibrationAnchorRequest is the body for POST /v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id.
type UpdatePhotoCalibrationAnchorRequest struct {
	PixelSelection *PixelSelection `json:"pixel_selection,omitempty"`
	Height         *float64        `json:"height,omitempty"`
}

// UpdatePhotoCalibrationAnchorOptions are optional query parameters for UpdateCalibrationAnchor.
type UpdatePhotoCalibrationAnchorOptions struct {
	OnlyIfExists bool
}
