package traces

// Trace represents a trace in a job (v3).
type Trace struct {
	ID         string                 `json:"id,omitempty"`
	TraceType  string                 `json:"_trace_type,omitempty"` // e.g., "cable"
	CableType  string                 `json:"cable_type,omitempty"`  // e.g., "Telco Com"
	Company    string                 `json:"company,omitempty"`     // e.g., "Brightspeed"
	Label      string                 `json:"label,omitempty"`
	Items      map[string]TraceItem   `json:"items,omitempty"` // Map of photo IDs to trace items
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// TraceItem represents a photo's trace elements within a trace.
// The keys are element type names (e.g., "wire"), values map element IDs to true.
type TraceItem map[string]map[string]bool

// CreateTraceRequest is the body for POST /v3/jobs/:job_id/traces.
type CreateTraceRequest struct {
	TraceType  string                 `json:"trace_type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// UpdateTraceRequest is the body for POST /v3/jobs/:job_id/traces/:trace_id.
type UpdateTraceRequest struct {
	TraceType  string                 `json:"trace_type,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// UpdateTraceOptions are optional query parameters for UpdateTrace.
type UpdateTraceOptions struct {
	OnlyIfExists bool
}
