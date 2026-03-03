package photos

import "encoding/json"

// Photo represents a photo record in a job (v3).
type Photo struct {
	ID                    string                    `json:"id,omitempty"`
	AssociatedLocations   map[string]string         `json:"associated_locations,omitempty"` // Map of entity ID to type (e.g., "node")
	CameraID              string                    `json:"camera_id,omitempty"`
	CameraMake            string                    `json:"camera_make,omitempty"`
	CameraModel           string                    `json:"camera_model,omitempty"`
	CameraSerialNumber    string                    `json:"camera_serial_number,omitempty"`
	DateTaken             int64                     `json:"date_taken,omitempty"`  // Unix timestamp
	FNumber               float64                   `json:"f_number,omitempty"`    // Aperture
	Filename              string                    `json:"filename,omitempty"`    // Original filename
	FocalLength           float64                   `json:"focal_length,omitempty"`
	FolderID              string                    `json:"folder_id,omitempty"`
	ImageHeight           int                       `json:"image_height,omitempty"`
	ImageWidth            int                       `json:"image_width,omitempty"`
	ISO                   int                       `json:"iso,omitempty"`
	Lens                  string                    `json:"lens,omitempty"`
	NameDate              string                    `json:"name_date,omitempty"` // "filename|timestamp" format
	Orientation           int                       `json:"orientation,omitempty"`
	OriginalSize          int64                     `json:"original_size,omitempty"` // Bytes
	PhotofirstData        *PhotofirstData           `json:"photofirst_data,omitempty"`
	ShutterSpeed          float64                   `json:"shutter_speed,omitempty"`
	Status                string                    `json:"status,omitempty"` // e.g., "upload_complete"
	StickAlign            *StickAlign               `json:"stick_align,omitempty"`
	SyncedTime            int64                     `json:"synced_time,omitempty"`
	UploadDate            int64                     `json:"upload_date,omitempty"`
	UploadedBy            string                    `json:"uploaded_by,omitempty"`
	VignettingCorrection  string                    `json:"vignetting_correction,omitempty"`
	VignettingCorrection2 string                    `json:"vignetting_correction_2,omitempty"`
}

// PhotofirstData contains measurement and calibration data for a photo.
type PhotofirstData struct {
	Editors           map[string]int64               `json:"_editors,omitempty"`           // Map of user ID to timestamp
	AnchorCalibration map[string]AnchorCalibration   `json:"anchor_calibration,omitempty"` // Keyed by anchor ID
	PoleHeight        map[string]bool                `json:"poleHeight,omitempty"`
	PoleTop           map[string]PoleTopMeasurement  `json:"pole_top,omitempty"`
	Wire              map[string]WireMeasurement     `json:"wire,omitempty"`
}

// AnchorCalibration represents a calibration anchor measurement.
type AnchorCalibration struct {
	RoutineInstanceID string            `json:"_routine_instance_id,omitempty"`
	Score             float64           `json:"_score,omitempty"`
	Height            float64           `json:"height,omitempty"` // Decimal feet
	PixelSelection    []PixelSelection  `json:"pixel_selection,omitempty"`
}

// PoleTopMeasurement represents a pole top measurement on a photo.
type PoleTopMeasurement struct {
	MeasuredHeight    float64           `json:"_measured_height,omitempty"` // Inches
	RoutineInstanceID string            `json:"_routine_instance_id,omitempty"`
	PixelSelection    []PixelSelection  `json:"pixel_selection,omitempty"`
	PoleTopExtension  bool              `json:"pole_top_extension,omitempty"`
}

// WireMeasurement represents a wire measurement on a photo.
type WireMeasurement struct {
	MeasuredHeight float64          `json:"_measured_height,omitempty"` // Inches
	Trace          string           `json:"_trace,omitempty"`           // Trace ID
	PixelSelection []PixelSelection `json:"pixel_selection,omitempty"`
}

// StickAlign contains stick alignment calibration data.
type StickAlign struct {
	A            float64             `json:"A,omitempty"`
	B            float64             `json:"B,omitempty"`
	C            float64             `json:"C,omitempty"`
	AD           float64             `json:"aD,omitempty"`
	BD           float64             `json:"bD,omitempty"`
	CD           float64             `json:"cD,omitempty"`
	DD           float64             `json:"dD,omitempty"`
	Error        float64             `json:"error,omitempty"`
	Height       int                 `json:"height,omitempty"`
	Width        int                 `json:"width,omitempty"`
	StickBottomX float64             `json:"stickBottomX,omitempty"`
	StickBottomY float64             `json:"stickBottomY,omitempty"`
	StickTopX    float64             `json:"stickTopX,omitempty"`
	StickTopY    float64             `json:"stickTopY,omitempty"`
	Type         string              `json:"type,omitempty"` // e.g., "flat"
	ErrorMessages []StickAlignError  `json:"errorMessages,omitempty"`
	OtherAligns   []StickAlignAlt    `json:"otherAligns,omitempty"`
}

// StickAlignError represents an error message in stick alignment.
type StickAlignError struct {
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"` // Can be object or number
}

// StickAlignAlt represents an alternative stick alignment calculation.
type StickAlignAlt struct {
	A            float64 `json:"A,omitempty"`
	B            float64 `json:"B,omitempty"`
	C            float64 `json:"C,omitempty"`
	AD           float64 `json:"aD,omitempty"`
	BD           float64 `json:"bD,omitempty"`
	CD           float64 `json:"cD,omitempty"`
	DD           float64 `json:"dD,omitempty"`
	Error        float64 `json:"error,omitempty"`
	Height       int     `json:"height,omitempty"`
	Width        int     `json:"width,omitempty"`
	StickBottomX float64 `json:"stickBottomX,omitempty"`
	StickBottomY float64 `json:"stickBottomY,omitempty"`
	StickTopX    float64 `json:"stickTopX,omitempty"`
	StickTopY    float64 `json:"stickTopY,omitempty"`
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
// The API returns this as an array in responses but accepts a single object in requests.
// This type handles both formats transparently.
// See: https://github.com/KatapultDevelopment/katapult-pro-api-documentation/blob/main/v3/README.md
type PixelSelection struct {
	PercentX float64 `json:"percentX"`
	PercentY float64 `json:"percentY"`
}

// UnmarshalJSON handles both single object and array formats from the API.
func (p *PixelSelection) UnmarshalJSON(data []byte) error {
	// Try single object first
	type plain PixelSelection
	if err := json.Unmarshal(data, (*plain)(p)); err == nil {
		return nil
	}
	// Try array format (API returns [{percentX, percentY}])
	var arr []plain
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	if len(arr) > 0 {
		*p = PixelSelection(arr[0])
	}
	return nil
}

// PhotoElementContext provides context about a photo element's location within the photo data.
type PhotoElementContext struct {
	ID   string `json:"id,omitempty"`
	Path string `json:"path,omitempty"` // e.g., "wire/-Ol6zELQobvpcR0XDIAM"
	Type string `json:"type,omitempty"` // e.g., "wire", "pole_top", "poleHeight"
}

// PhotoElement represents an element on a photo (v3). Attributes are a flat map and may also appear at root on the response.
type PhotoElement struct {
	ID                string                 `json:"id,omitempty"`
	ElementType       string                 `json:"element_type,omitempty"`
	PixelSelection    []PixelSelection       `json:"pixel_selection,omitempty"`
	ManualHeight      string                 `json:"manual_height,omitempty"`    // Feet-inches notation, e.g. "16-6".
	MeasuredHeight    float64                `json:"_measured_height,omitempty"` // Calculated height in inches
	RoutineInstanceID string                 `json:"_routine_instance_id,omitempty"`
	Trace             string                 `json:"_trace,omitempty"` // Associated trace ID
	PoleTopExtension  bool                   `json:"pole_top_extension,omitempty"`
	Context           *PhotoElementContext   `json:"_context,omitempty"`
	Attributes        map[string]interface{} `json:"attributes,omitempty"` // Flat map per API.
	ParentID          string                 `json:"parent_id,omitempty"`
	TraceID           string                 `json:"trace_id,omitempty"`
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
	ID                string           `json:"id,omitempty"`
	PixelSelection    []PixelSelection `json:"pixel_selection,omitempty"`
	Height            float64          `json:"height,omitempty"`               // Decimal feet.
	RoutineInstanceID string           `json:"_routine_instance_id,omitempty"` // Set by API when auto-generated.
	Score             float64          `json:"_score,omitempty"`               // Confidence score from auto-detection.
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
