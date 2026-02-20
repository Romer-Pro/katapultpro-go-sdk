// Package katapultpro provides a lightweight Go client for the Katapult Pro API v3.
// Import as: github.com/romer-pro/katapultpro-go-sdk/v3
// It depends only on the standard library and is safe for concurrent use.
//
// API reference: https://github.com/KatapultDevelopment/katapult-pro-api-documentation/blob/main/v3/README.md
//
// # Overview
//
// Create a client with NewClient, then call Get, Post, Put, or Delete using v3 paths
// (e.g. /v3/jobs, /v3/jobs/:job_id/nodes). The client unwraps the API's response envelope
// ({ status, data, meta }) so your out value receives the contents of data. LastMeta on
// the client is updated after each request with token_count and last_refill_time for
// rate-limit awareness.
//
// # Usage
//
// NewClient requires a valid API key; passing an empty string returns ErrMissingAPIKey.
//
//	client, err := katapultpro.NewClient("your-api-key")
//	if err != nil {
//	    log.Fatal(err) // e.g., ErrMissingAPIKey if API key is empty
//	}
//	ctx := context.Background()
//
//	var job katapultpro.Job
//	if err := client.Get(ctx, "/v3/jobs/some-job-id", &job); err != nil {
//	    var apiErr *katapultpro.APIError
//	    if errors.As(err, &apiErr) {
//	        log.Printf("API error %d (%s): %s", apiErr.StatusCode, apiErr.Type, apiErr.Message)
//	        return
//	    }
//	    log.Fatal(err)
//	}
//	// Check client.LastMeta for token_count and last_refill_time
//
// # Scopes
//
// Use Client.Job(jobID) to scope operations to one job and avoid repeating the job ID:
//
//	nodes, _ := client.Job("job-123").Nodes().List(ctx)
//	sections, _ := client.Job("job-123").Connections().Sections("conn-id").List(ctx)
//
// # Rate limiting
//
// The API allows 1 call per 50ms. Use WithRateLimit(DefaultRateLimitInterval) to throttle requests.
//
// # Configuration
//
// Use options to customize the client:
//
//	client, _ := katapultpro.NewClient("api-key",
//	    katapultpro.WithBaseURL("https://katapultpro.com/api"),
//	    katapultpro.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
//	)
//
// # Testing
//
// The Interface type defines the client's public methods. Accept Interface in your
// code so you can pass a mock implementation in tests:
//
//	func DoSomething(c katapultpro.Interface) error {
//	    return c.Get(ctx, "/v3/jobs/123", &out)
//	}
package katapultpro
