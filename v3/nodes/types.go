package nodes

import (
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/shared"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
)

// Node represents a node in a job (v3).
type Node struct {
	ID             string                 `json:"id,omitempty"`
	Latitude       float64                `json:"latitude,omitempty"`
	Longitude      float64                `json:"longitude,omitempty"`
	Attributes     shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes  map[string]interface{} `json:"add_attributes,omitempty"`
}

// CreateNodeRequest is the body for POST /v3/jobs/:job_id/nodes.
type CreateNodeRequest struct {
	Latitude       float64                `json:"latitude"`
	Longitude      float64                `json:"longitude"`
	Attributes     shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes  map[string]interface{} `json:"add_attributes,omitempty"`
}

// UpdateNodeRequest is the body for POST /v3/jobs/:job_id/nodes/:node_id.
type UpdateNodeRequest struct {
	Latitude         float64                `json:"latitude,omitempty"`
	Longitude        float64                `json:"longitude,omitempty"`
	RemoveAttributes []string               `json:"remove_attributes,omitempty"`
	Attributes       shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes    map[string]interface{} `json:"add_attributes,omitempty"`
}

// UpdateNodeOptions are optional query parameters for UpdateNode.
type UpdateNodeOptions struct {
	OnlyIfExists bool
}

// UploadNodePhotoOptions are optional query parameters for UploadNodePhoto.
type UploadNodePhotoOptions struct {
	AssociationValue photos.PhotoAssociationQuery
}
