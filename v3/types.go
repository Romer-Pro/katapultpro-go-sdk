package katapultpro

import (
	"github.com/romer-pro/katapultpro-go-sdk/v3/connections"
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/shared"
	"github.com/romer-pro/katapultpro-go-sdk/v3/jobs"
	"github.com/romer-pro/katapultpro-go-sdk/v3/nodes"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
	"github.com/romer-pro/katapultpro-go-sdk/v3/sections"
	"github.com/romer-pro/katapultpro-go-sdk/v3/traces"
)

// Meta is included in every API response with token bucket state for rate limiting.
// See the Katapult Pro API v3 docs for rate limits and refill behavior.
type Meta struct {
	TokenCount     int64 `json:"token_count"`
	LastRefillTime int64 `json:"last_refill_time"`
}

// EntityAttributeList is the attribute structure for nodes, connections, and sections.
// Re-exported from internal for use in request/response types.
type EntityAttributeList = shared.EntityAttributeList

// Jobs domain types (re-exported).
type (
	Job                  = jobs.Job
	CreateJobRequest     = jobs.CreateJobRequest
	UpdateJobRequest     = jobs.UpdateJobRequest
	UpdateJobStatusRequest = jobs.UpdateJobStatusRequest
	ListJobsOptions      = jobs.ListJobsOptions
	GetJobOptions        = jobs.GetJobOptions
	JobStatusResponse    = jobs.JobStatusResponse
)

// Nodes domain types (re-exported).
type (
	Node                    = nodes.Node
	CreateNodeRequest       = nodes.CreateNodeRequest
	UpdateNodeRequest       = nodes.UpdateNodeRequest
	UpdateNodeOptions       = nodes.UpdateNodeOptions
	UploadNodePhotoOptions  = nodes.UploadNodePhotoOptions
)

// Connections domain types (re-exported).
type (
	Connection               = connections.Connection
	CreateConnectionRequest   = connections.CreateConnectionRequest
	UpdateConnectionRequest  = connections.UpdateConnectionRequest
	UpdateConnectionOptions  = connections.UpdateConnectionOptions
)

// Sections domain types (re-exported).
type (
	Section                    = sections.Section
	CreateSectionRequest       = sections.CreateSectionRequest
	UpdateSectionRequest       = sections.UpdateSectionRequest
	UpdateSectionOptions       = sections.UpdateSectionOptions
	UploadSectionPhotoOptions  = sections.UploadSectionPhotoOptions
)

// Photos domain types (re-exported).
type (
	Photo                 = photos.Photo
	AssociatePhotoRequest = photos.AssociatePhotoRequest
)

// Traces domain types (re-exported).
type (
	Trace                 = traces.Trace
	CreateTraceRequest    = traces.CreateTraceRequest
	UpdateTraceRequest    = traces.UpdateTraceRequest
	UpdateTraceOptions    = traces.UpdateTraceOptions
)
