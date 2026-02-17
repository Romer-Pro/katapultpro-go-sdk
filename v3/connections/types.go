package connections

import (
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/shared"
)

// Connection represents a connection between two nodes (v3).
type Connection struct {
	ID            string                 `json:"id,omitempty"`
	NodeID1       string                 `json:"node_id_1,omitempty"`
	NodeID2       string                 `json:"node_id_2,omitempty"`
	Attributes    shared.EntityAttributeList `json:"attributes,omitempty"`
	AddAttributes map[string]interface{} `json:"add_attributes,omitempty"`
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
