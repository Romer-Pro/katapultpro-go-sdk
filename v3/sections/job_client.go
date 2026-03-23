package sections

import (
	"context"
	"net/http"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
)

// JobClient performs section API operations at the job level (without requiring connectionID).
// Create with NewJobClient(do, jobID).
// Use this when you have a sectionID but not the connectionID.
type JobClient struct {
	do    request.Doer
	jobID string
}

// NewJobClient returns a sections client scoped to a job for direct section access.
func NewJobClient(do request.Doer, jobID string) *JobClient {
	return &JobClient{do: do, jobID: jobID}
}

// List returns all sections in the job across all connections (v3).
// GET /v3/jobs/{jobID}/sections
func (c *JobClient) List(ctx context.Context) ([]Section, error) {
	path := "v3/jobs/" + c.jobID + "/sections"
	var out []Section
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns a section by ID directly without requiring connectionID (v3).
// GET /v3/jobs/{jobID}/sections/{sectionID}
func (c *JobClient) Get(ctx context.Context, sectionID string) (*Section, error) {
	path := "v3/jobs/" + c.jobID + "/sections/" + sectionID
	var section Section
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &section); err != nil {
		return nil, err
	}
	return &section, nil
}

// Section returns a JobSectionScope for the given section ID, enabling drill-down operations.
// Example: client.Job(jobID).Sections().Section(sectionID).Get(ctx)
func (c *JobClient) Section(sectionID string) *JobSectionScope {
	return &JobSectionScope{do: c.do, jobID: c.jobID, sectionID: sectionID}
}

// JobSectionScope scopes operations to a single section accessed directly by ID.
type JobSectionScope struct {
	do        request.Doer
	jobID     string
	sectionID string
}

// SectionID returns the scoped section ID.
func (s *JobSectionScope) SectionID() string { return s.sectionID }

// Get returns the section.
func (s *JobSectionScope) Get(ctx context.Context) (*Section, error) {
	path := "v3/jobs/" + s.jobID + "/sections/" + s.sectionID
	var section Section
	if err := s.do.Do(ctx, http.MethodGet, path, nil, nil, &section); err != nil {
		return nil, err
	}
	return &section, nil
}
