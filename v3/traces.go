package katapultpro

import (
	"context"

	"github.com/romer-pro/katapultpro-go-sdk/v3/traces"
)

// ListTraces returns all traces in the job (v3).
func (c *Client) ListTraces(ctx context.Context, jobID string) ([]Trace, error) {
	return traces.NewClient(c, jobID).List(ctx)
}

// GetTrace returns the specified trace (v3).
func (c *Client) GetTrace(ctx context.Context, jobID, traceID string) (*Trace, error) {
	return traces.NewClient(c, jobID).Get(ctx, traceID)
}

// CreateTrace creates a new trace in the job (v3).
func (c *Client) CreateTrace(ctx context.Context, jobID string, req *CreateTraceRequest) (*Trace, error) {
	return traces.NewClient(c, jobID).Create(ctx, req)
}

// UpdateTrace updates the specified trace, or creates it with the given ID if it does not exist (v3).
func (c *Client) UpdateTrace(ctx context.Context, jobID, traceID string, req *UpdateTraceRequest, opts *UpdateTraceOptions) (*Trace, error) {
	return traces.NewClient(c, jobID).Update(ctx, traceID, req, opts)
}

// DeleteTrace deletes the specified trace (v3).
func (c *Client) DeleteTrace(ctx context.Context, jobID, traceID string) error {
	return traces.NewClient(c, jobID).Delete(ctx, traceID)
}
