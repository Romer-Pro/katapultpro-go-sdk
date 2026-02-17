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
