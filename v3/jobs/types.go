package jobs

// JobStatus is the status of a job in the API. Only Active and Archived are valid.
type JobStatus string

const (
	JobStatusActive   JobStatus = "active"
	JobStatusArchived JobStatus = "archived"
)

// String returns the API value (e.g. "active", "archived").
func (s JobStatus) String() string { return string(s) }

// IsValid reports whether s is a defined job status.
func (s JobStatus) IsValid() bool {
	return s == JobStatusActive || s == JobStatusArchived
}

// JobPath is a data path for GetJob partial responses. Use these constants in GetJobOptions.Paths.
type JobPath string

const (
	JobPathName          JobPath = "name"
	JobPathJobCreator    JobPath = "job_creator"
	JobPathJobOwner      JobPath = "job_owner"
	JobPathProjectFolder JobPath = "project_folder"
	JobPathProjectID     JobPath = "project_id"
	JobPathStatus        JobPath = "status"
	JobPathDone          JobPath = "done"
	JobPathMapStyles     JobPath = "map_styles"
	JobPathMetadata      JobPath = "metadata"
	JobPathSharing       JobPath = "sharing"
)

// Job represents a job in the Katapult Pro API (v3).
// Use paths like /v3/jobs and /v3/jobs/:job_id.
type Job struct {
	ID            string                 `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Model         string                 `json:"model,omitempty"`
	Status        JobStatus              `json:"status,omitempty"`
	MapStyles     string                 `json:"map_styles,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	Sharing       map[string]interface{} `json:"sharing,omitempty"`
	JobCreator    string                 `json:"job_creator,omitempty"`
	JobOwner      string                 `json:"job_owner,omitempty"`
	ProjectFolder string                 `json:"project_folder,omitempty"`
	ProjectID     string                 `json:"project_id,omitempty"`
	Done          bool                   `json:"done,omitempty"`
}

// CreateJobRequest is the body for POST /v3/jobs.
type CreateJobRequest struct {
	Name      string                 `json:"name"`
	Model     string                 `json:"model"`
	MapStyles string                 `json:"map_styles,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Sharing   map[string]interface{} `json:"sharing,omitempty"`
}

// UpdateJobRequest is the body for POST /v3/jobs/:job_id.
type UpdateJobRequest struct {
	Name      string                 `json:"name,omitempty"`
	Model     string                 `json:"model,omitempty"`
	MapStyles string                 `json:"map_styles,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Sharing   map[string]interface{} `json:"sharing,omitempty"`
}

// UpdateJobStatusRequest is the body for POST /v3/jobs/:job_id/status.
type UpdateJobStatusRequest struct {
	Status JobStatus `json:"status"`
}

// ListJobsOptions are optional query parameters for ListJobs.
type ListJobsOptions struct {
	IncludeArchived bool
	MetadataFilter  string
}

// GetJobOptions are optional query parameters for GetJob (partial data).
type GetJobOptions struct {
	Paths []JobPath
}

// JobStatusResponse is the response body for GetJobStatus.
type JobStatusResponse struct {
	Status JobStatus `json:"status"`
}
