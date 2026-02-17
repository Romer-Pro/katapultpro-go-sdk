# User guide

Quick reference for using the Katapult Pro Go SDK. For API details see the [API v3 documentation](https://github.com/KatapultDevelopment/katapult-pro-api-documentation/blob/main/v3/README.md). For package reference see [pkg.go.dev](https://pkg.go.dev/github.com/romer-pro/katapultpro-go-sdk/v3) (after publishing).

## Install (v3)

```bash
go get github.com/romer-pro/katapultpro-go-sdk/v3
```

## Basic usage

```go
import (
    "context"
    "log"
    "github.com/romer-pro/katapultpro-go-sdk/v3"
)

func main() {
    client, err := katapultpro.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }
    ctx := context.Background()

    jobs, err := client.ListJobs(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    job, err := client.GetJob(ctx, "job-id", nil)
    // ...
}
```

## Scopes

Use `Client.Job(jobID)` to scope calls so you don’t repeat the job ID:

```go
job := client.Job("job-123")
nodes, _ := job.Nodes().List(ctx)
sections, _ := job.Connections().Sections("connection-id").List(ctx)
status, _ := job.Status(ctx)
```

## Rate limiting

The API allows about 1 call per 50ms. Enable client-side throttling:

```go
client, _ := katapultpro.NewClient("api-key",
    katapultpro.WithRateLimit(katapultpro.DefaultRateLimitInterval), // 50ms
)
```

## Options

```go
client, _ := katapultpro.NewClient("api-key",
    katapultpro.WithBaseURL("https://custom.api.example.com"),
    katapultpro.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)
```

## Errors

Use `errors.As` to detect API errors:

```go
var apiErr *katapultpro.APIError
if errors.As(err, &apiErr) {
    log.Printf("API error %d (%s): %s", apiErr.StatusCode, apiErr.Type, apiErr.Message)
}
```

## Testing

The client implements `katapultpro.Interface`. Accept the interface in your code to mock the client in tests.

## Using a different API version

- **v3:** `import "github.com/romer-pro/katapultpro-go-sdk/v3"`
- **v4 (when available):** `import "github.com/romer-pro/katapultpro-go-sdk/v4"`

Each version is a separate module; you can have both in the same project if needed.
