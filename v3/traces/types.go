package traces

// Trace represents a trace in a job (v3).
type Trace struct {
	ID         string                 `json:"id,omitempty"`
	TraceType  string                 `json:"trace_type,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

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
