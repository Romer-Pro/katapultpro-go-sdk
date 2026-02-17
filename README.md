# katapultpro-go-sdk

Unofficial Go SDK for the [Katapult Pro](https://katapult.pro) **API v3**. Lightweight, stdlib-only, safe for concurrent use.

- **API reference:** [Katapult Pro API v3 (GitHub)](https://github.com/KatapultDevelopment/katapult-pro-api-documentation/blob/main/v3/README.md)
- **Documentation:** [pkg.go.dev/github.com/romer-pro/katapultpro-go-sdk/v3](https://pkg.go.dev/github.com/romer-pro/katapultpro-go-sdk/v3) (after publishing)
- **Layout:** Versioned by API (v3, future v4). Design and user docs in [docs/](docs/). See [docs/design.md](docs/design.md) for versioning and domain layout.

## Install (v3)

**Latest version (tracking main):**

```bash
go get github.com/romer-pro/katapultpro-go-sdk/v3
```

**Specific release (recommended for production):**

```bash
# Replace v3.0.0 with the release tag (e.g. v3.1.0)
go get github.com/romer-pro/katapultpro-go-sdk/v3@v3.0.0
```

Then in your code:

```go
import "github.com/romer-pro/katapultpro-go-sdk/v3"
```

## Releases and security

Releases are published as [GitHub Releases](https://github.com/romer-pro/katapultpro-go-sdk/releases) when a version tag (e.g. `v3.0.0`) is pushed. Each release includes:

- **SBOM (Software Bill of Materials)** in CycloneDX and SPDX format for dependency and supply-chain review.
- **Checksums** for the SBOM files (SHA-256 in `checksums.txt`).

Use a tagged version in `go get ...@vx.x.x` for reproducible builds and to verify artifacts from the release page if needed.

## Usage

```go
package main

import (
    "context"
    "errors"
    "log"

    "github.com/romer-pro/katapultpro-go-sdk/v3"
)

func main() {
    client, err := katapultpro.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }
    ctx := context.Background()

    // List jobs (high-level method)
    jobs, err := client.ListJobs(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    // Get a single job
    job, err := client.GetJob(ctx, "job-id", nil)
    if err != nil {
        var apiErr *katapultpro.APIError
        if errors.As(err, &apiErr) {
            log.Printf("API error %d (%s): %s", apiErr.StatusCode, apiErr.Type, apiErr.Message)
            return
        }
        log.Fatal(err)
    }
    _ = job
    // client.LastMeta has token_count and last_refill_time for rate limits
}
```

### Scopes (simplified API)

Use `Client.Job(jobID)` to get a scope so you don’t repeat the job ID on every call:

```go
job := client.Job("job-123")
nodes, _ := job.Nodes().List(ctx)
sections, _ := job.Connections().Sections("connection-id").List(ctx)
status, _ := job.Status(ctx)
```

### Rate limiting

The API allows 1 call per 50ms. Enable client-side throttling with `WithRateLimit`:

```go
client, _ := katapultpro.NewClient("api-key",
    katapultpro.WithRateLimit(katapultpro.DefaultRateLimitInterval), // 50ms
)
// Or a custom interval:
client, _ = katapultpro.NewClient("api-key", katapultpro.WithRateLimit(100*time.Millisecond))
```

### Options

```go
client, _ := katapultpro.NewClient("api-key",
    katapultpro.WithBaseURL("https://custom.api.example.com"),
    katapultpro.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)
```

### Testing

The client implements the `Interface` type. Accept `katapultpro.Interface` in your code to allow mocking:

```go
func DoSomething(ctx context.Context, c katapultpro.Interface) error {
    return c.Get(ctx, "/v3/jobs/123", &out)
}
```

## Project layout

- **docs/** — Design and user docs: [design.md](docs/design.md) (versioning, layout), [user-guide.md](docs/user-guide.md).
- **v3/** — Go module for API v3 (`import github.com/romer-pro/katapultpro-go-sdk/v3`). Root package holds `Client`, resource methods, scopes, and re-exports.
  - **Domain subpackages:** Each has a **Client** that holds the doer (and jobID/connectionID when scoped). You get a client once and call methods without passing it again: `client.Jobs().List(ctx, opts)`, `client.Job("id").Nodes().List(ctx)`, `client.Job("id").Connections().Sections("conn-id").List(ctx)`. Root re-exports types and delegates; tests can build a domain client with a mock Doer and keep tests colocated in the subpackage.
  - **v3/internal/** — `envelope`, `transport`, `ratelimit`, `shared`, `request` (Doer interface for domain packages).
  - Root: `client.go` (Client, Do/DoWithBody implement Doer), `option.go`, `errors.go`, `enums.go`, `types.go`, `jobs.go`… (delegation only), `scopes.go`, `ratelimit.go`, `*_test.go`.

Future **v4/** will live alongside v3 when the API adds a new major version. Low-level Get/Post/Put/Delete remain for custom paths.

## Development (Justfile)

The project uses [just](https://github.com/casey/just) for common tasks. Install with `brew install just` or `cargo install just`, then:

| Command | Description |
|--------|-------------|
| `just` | List available recipes |
| `just build` | Build the module |
| `just test` | Run tests |
| `just test-verbose` | Run tests with `-v` |
| `just test-race` | Run tests with race detector |
| `just test-cover` | Run tests with coverage |
| `just vet` | Run `go vet` |
| `just fmt` | Format code |
| `just check` | Build, vet, and test |
| `just ci` | Tidy, build, vet, test with race (CI). |

## Contributing

Contributions are welcome. All changes go through pull requests: tests must pass and a maintainer must approve. See [CONTRIBUTING.md](CONTRIBUTING.md) for the full process and branch protection details.

## License

See [LICENSE](LICENSE).
