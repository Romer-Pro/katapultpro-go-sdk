package sections

import (
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/shared"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
)

// Section represents a section on a connection (v3).
type Section struct {
	Key           string                 `json:"key,omitempty"`
	Latitude      float64                `json:"latitude,omitempty"`
	Longitude     float64                `json:"longitude,omitempty"`
	MakeMidpoint  bool                   `json:"make_midpoint,omitempty"`
	Attributes    shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes map[string]interface{} `json:"add_attributes,omitempty"`
}

// CreateSectionRequest is the body for POST /v3/jobs/:job_id/connections/:connection_id/sections.
type CreateSectionRequest struct {
	MakeMidpoint  bool                   `json:"make_midpoint,omitempty"`
	Latitude      float64                `json:"latitude,omitempty"`
	Longitude     float64                `json:"longitude,omitempty"`
	Attributes    shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes map[string]interface{} `json:"add_attributes,omitempty"`
}

// UpdateSectionRequest is the body for POST /v3/jobs/.../sections/:section_key.
type UpdateSectionRequest struct {
	Latitude         float64                `json:"latitude,omitempty"`
	Longitude        float64                `json:"longitude,omitempty"`
	RemoveAttributes []string               `json:"remove_attributes,omitempty"`
	Attributes       shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes    map[string]interface{} `json:"add_attributes,omitempty"`
}

// UpdateSectionOptions are optional query parameters for UpdateSection.
type UpdateSectionOptions struct {
	OnlyIfExists bool
}

// UploadSectionPhotoOptions are optional query parameters for UploadSectionPhoto.
type UploadSectionPhotoOptions struct {
	AssociationValue photos.PhotoAssociationQuery
}
