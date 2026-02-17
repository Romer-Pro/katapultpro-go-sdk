package katapultpro

import (
	"context"

	"github.com/romer-pro/katapultpro-go-sdk/v3/connections"
	"github.com/romer-pro/katapultpro-go-sdk/v3/nodes"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
	"github.com/romer-pro/katapultpro-go-sdk/v3/traces"
)

// JobScope scopes all operations to a single job. Use Client.Job(jobID) to create one.
// Nodes(), Connections(), Photos(), and Traces() return domain clients that hold the scope;
// call their methods without passing the client again, e.g. client.Job("id").Nodes().List(ctx).
type JobScope struct {
	c     *Client
	jobID string
}

// Job returns a scope for the given job.
func (c *Client) Job(jobID string) *JobScope {
	return &JobScope{c: c, jobID: jobID}
}

// JobID returns the scoped job ID.
func (s *JobScope) JobID() string { return s.jobID }

// Get returns partial or full job data. Use opts.Paths to request specific fields.
func (s *JobScope) Get(ctx context.Context, opts *GetJobOptions) (*Job, error) {
	return s.c.GetJob(ctx, s.jobID, opts)
}

// Update updates the job.
func (s *JobScope) Update(ctx context.Context, req *UpdateJobRequest) (*Job, error) {
	return s.c.UpdateJob(ctx, s.jobID, req)
}

// Status returns the job status (active or archived).
func (s *JobScope) Status(ctx context.Context) (JobStatus, error) {
	return s.c.GetJobStatus(ctx, s.jobID)
}

// SetStatus sets the job status.
func (s *JobScope) SetStatus(ctx context.Context, status JobStatus) error {
	return s.c.UpdateJobStatus(ctx, s.jobID, status)
}

// Nodes returns a nodes client for this job. No need to pass the client into its methods.
func (s *JobScope) Nodes() *nodes.Client {
	return nodes.NewClient(s.c, s.jobID)
}

// Connections returns a connections client for this job. Use .Sections(connectionID) for sections.
func (s *JobScope) Connections() *connections.Client {
	return connections.NewClient(s.c, s.jobID)
}

// Photos returns a photos client for this job.
func (s *JobScope) Photos() *photos.Client {
	return photos.NewClient(s.c, s.jobID)
}

// Traces returns a traces client for this job.
func (s *JobScope) Traces() *traces.Client {
	return traces.NewClient(s.c, s.jobID)
}
