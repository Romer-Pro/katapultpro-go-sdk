package katapultpro

import (
	"context"

	"github.com/romer-pro/katapultpro-go-sdk/v3/jobs"
)

// Jobs returns a jobs client. Its methods do not take the client; use it as client.Jobs().List(ctx, opts).
func (c *Client) Jobs() *jobs.Client {
	return jobs.NewClient(c)
}

// ListJobs returns all jobs accessible to the requester (v3).
func (c *Client) ListJobs(ctx context.Context, opts *ListJobsOptions) ([]Job, error) {
	return c.Jobs().List(ctx, opts)
}

// GetJob returns partial or full job data for the given job ID (v3).
func (c *Client) GetJob(ctx context.Context, jobID string, opts *GetJobOptions) (*Job, error) {
	return c.Jobs().Get(ctx, jobID, opts)
}

// CreateJob creates a new job (v3).
func (c *Client) CreateJob(ctx context.Context, req *CreateJobRequest) (*Job, error) {
	return c.Jobs().Create(ctx, req)
}

// UpdateJob updates the specified job (v3).
func (c *Client) UpdateJob(ctx context.Context, jobID string, req *UpdateJobRequest) (*Job, error) {
	return c.Jobs().Update(ctx, jobID, req)
}

// GetJobStatus returns the status of the job (v3).
func (c *Client) GetJobStatus(ctx context.Context, jobID string) (JobStatus, error) {
	return c.Jobs().GetStatus(ctx, jobID)
}

// UpdateJobStatus sets the job status (v3).
func (c *Client) UpdateJobStatus(ctx context.Context, jobID string, status JobStatus) error {
	return c.Jobs().UpdateStatus(ctx, jobID, status)
}
