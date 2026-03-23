package connections

import (
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/shared"
)

// Connection represents a connection between two nodes (v3).
type Connection struct {
	ID            string                     `json:"id,omitempty"`
	NodeID1       string                     `json:"node_id_1,omitempty"`
	NodeID2       string                     `json:"node_id_2,omitempty"`
	Button        string                     `json:"button,omitempty"` // e.g., "aerial"
	Created       *shared.CreatedInfo        `json:"_created,omitempty"`
	Attributes    shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes map[string]interface{}     `json:"add_attributes,omitempty"`
	Sections      map[string]EmbeddedSection `json:"sections,omitempty"` // Sections embedded in connection response
	Photos        shared.PhotoAssociationMap `json:"photos,omitempty"`   // Photos associated with the connection
}

// EmbeddedSection represents a section embedded within a connection response.
// This differs from the standalone Section type as sections are keyed by ID in the map.
type EmbeddedSection struct {
	Created         *shared.CreatedInfo        `json:"_created,omitempty"`
	Latitude        float64                    `json:"latitude,omitempty"`
	Longitude       float64                    `json:"longitude,omitempty"`
	MultiAttributes shared.EntityAttributeList `json:"multi_attributes,omitempty"`
	Photos          shared.PhotoAssociationMap `json:"photos,omitempty"`
}

// CreateConnectionRequest is the body for POST /v3/jobs/:job_id/connections.
type CreateConnectionRequest struct {
	NodeID1       string                 `json:"node_id_1"`
	NodeID2       string                 `json:"node_id_2"`
	Attributes    shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes map[string]interface{} `json:"add_attributes,omitempty"`
}

// UpdateConnectionRequest is the body for POST /v3/jobs/:job_id/connections/:connection_id.
type UpdateConnectionRequest struct {
	NodeID1          string                 `json:"node_id_1,omitempty"`
	NodeID2          string                 `json:"node_id_2,omitempty"`
	RemoveAttributes []string               `json:"remove_attributes,omitempty"`
	Attributes       shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes    map[string]interface{} `json:"add_attributes,omitempty"`
}

// UpdateConnectionOptions are optional query parameters for UpdateConnection.
type UpdateConnectionOptions struct {
	OnlyIfExists bool
}
