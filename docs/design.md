# Design and layout

This document describes the repository layout, versioning strategy, and how the SDK is organized for long-term maintenance and future API versions.

## Directory layout

```
katapultpro-go-sdk/
  docs/           # Design and user documentation (this file, user-guide, etc.)
  v3/             # Go module for API v3 (import path .../v3)
  v4/             # (Future) Go module for API v4 (import path .../v4)
  go.work         # Workspace for local development (use ./v3, ./v4, …)
  justfile        # Build, test, lint
  README.md       # Entry point and install/usage
```

- **docs/** — Design and user docs in addition to godoc. No Go code.
- **internal/** — (Optional) Shared internal code used by multiple versioned modules (e.g. v3 and v4). Not part of the public API. If we add v4 and share envelope or transport logic, it can live here.
- **v3/** — Full SDK for Katapult Pro API v3. Single importable module.
- **v4/** — Reserved for a future API v4 SDK when the backend adds a new major version.

There is no `pkg/` at repo root: the public API surface is the versioned modules (`v3`, and later `v4`). This keeps the import path explicit about which API version you use.

## Versioning and future-proofing

- **Import paths:**  
  - v3: `github.com/romer-pro/katapultpro-go-sdk/v3`  
  - v4 (future): `github.com/romer-pro/katapultpro-go-sdk/v4`

- **Go modules:** Each major API version is a separate Go module:
  - `v3/go.mod` → `module github.com/romer-pro/katapultpro-go-sdk/v3`
  - `v4/go.mod` (later) → `module github.com/romer-pro/katapultpro-go-sdk/v4`

- **Stability:** When v4 is introduced, v3 remains unchanged. Consumers can stay on v3 or migrate to v4 by changing the import path and adapting to any API differences.

- **Install:**
  - v3: `go get github.com/romer-pro/katapultpro-go-sdk/v3@latest`
  - v4: (when available) `go get github.com/romer-pro/katapultpro-go-sdk/v4@latest`

## Domain separation within v3

The v3 module has **domain subpackages**; the root package re-exports types and keeps the `Client` and all resource methods.

- **v3/jobs** — Types and a **Client** that holds a request.Doer. `jobs.NewClient(do)`; then `c.List(ctx, opts)`, `c.Get(ctx, jobID, opts)`, etc. No client passed into each method.
- **v3/nodes** — **Client** holds doer and jobID. `nodes.NewClient(do, jobID)`; then `c.List(ctx)`, `c.Get(ctx, nodeID)`, etc.
- **v3/connections** — **Client** holds doer and jobID. `.Sections(connectionID)` returns `*sections.Client` for the fluent chain.
- **v3/sections** — **Client** holds doer, jobID, connectionID.
- **v3/photos** — **Client** holds doer and jobID.
- **v3/traces** — **Client** holds doer and jobID.

Shared types live in **v3/internal/shared**. **v3/internal/request** defines the Doer interface. Root exposes `client.Jobs()` (returns `*jobs.Client`) and `client.Job(jobID).Nodes()` / `.Connections()` / `.Photos()` / `.Traces()` (return domain clients). So callers get the same experience as root (`client.Job("id").Nodes().List(ctx)`) with no client passed into domain methods; types and logic stay colocated in subpackages, and tests can construct a domain client with a mock Doer.

## Internal code

- **v3/internal:** Code used only inside the v3 module; not importable by external projects.
  - **v3/internal/envelope** — Parses the Katapult Pro v3 response envelope (`Parse(body) -> Envelope`). Caller unmarshals `Envelope.Meta` into their own type.
  - **v3/internal/transport** — HTTP execution: `Do(...)` returns status code and response body. No envelope parsing.
  - **v3/internal/ratelimit** — `NewTransport(base, interval)` returns an `http.RoundTripper` that enforces a minimum interval between requests. Used by the public `WithRateLimit` option.
  - **v3/internal/request** — `Doer` interface (`Do`, `DoWithBody`). Domain packages take a `Doer` to perform requests; *Client implements it.
  - **v3/internal/shared** — `EntityAttributeList` and other types shared by nodes, connections, sections.
- **Repo root:** An `internal/` at repo root can hold code shared by v3 and v4 (e.g. common envelope types or transport helpers) if we add v4 and want to share logic. That code would not be importable by external projects.

## Summary

| Concern            | Location / approach |
|--------------------|----------------------|
| User & design docs | `docs/`              |
| Public API v3      | `v3/` (one module, one package) |
| Public API v4      | `v4/` (future)       |
| Shared internals   | `internal/` (optional) or `v3/internal` |
| Version in import  | `/v3`, `/v4` in path |
| Domain separation | Files today; optional subpackages later |
