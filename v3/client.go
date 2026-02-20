package katapultpro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/envelope"
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/transport"
)

const defaultBaseURL = "https://katapultpro.com/api"

// Interface defines the Katapult Pro API surface. *Client implements Interface.
// Use this interface in your code to allow mocking or swapping implementations in tests.
type Interface interface {
	Get(ctx context.Context, path string, out any) error
	Post(ctx context.Context, path string, body, out any) error
	Put(ctx context.Context, path string, body, out any) error
	Delete(ctx context.Context, path string) error
}

// Client is the Katapult Pro API client. A Client is safe for concurrent use by multiple goroutines.
// LastMeta is set after each successful request and holds the API's token_count and last_refill_time for rate-limit awareness.
type Client struct {
	baseURL    *url.URL
	apiKey     string
	httpClient *http.Client
	LastMeta   *Meta // Set after each request; nil before the first call or if the response had no meta.
}

// Ensure Client implements Interface and request.Doer at compile time.
var _ Interface = (*Client)(nil)
var _ request.Doer = (*Client)(nil)

// NewClient returns a new Client for the Katapult Pro API.
// apiKey is sent as a Bearer token on all requests and is required.
// Options (e.g. WithBaseURL, WithHTTPClient) customize the client behavior.
// Returns ErrMissingAPIKey if apiKey is empty.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

// Do performs an HTTP request and decodes the v3 JSON response envelope into out.
// Domain packages use this via the internal request.Doer interface.
// query is optional; when non-nil it is set as the request URL's RawQuery.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body any, out any) error {
	c.LastMeta = nil
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("encode request: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}
	statusCode, slurp, err := transport.Do(ctx, method, c.baseURL, path, query, bodyReader, "application/json", c.apiKey, c.httpClient)
	if err != nil {
		return err
	}
	return c.handleResponse(statusCode, slurp, out)
}

// do is an alias for Do for use in root package resource methods that delegate to domain packages.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, out any) error {
	return c.Do(ctx, method, path, query, body, out)
}

// doWithBody is an alias for DoWithBody.
func (c *Client) doWithBody(ctx context.Context, method, path string, query url.Values, contentType string, body io.Reader, out any) error {
	return c.DoWithBody(ctx, method, path, query, contentType, body, out)
}

// DoWithBody performs a request with a raw body and optional content type (e.g. image/jpeg for photo upload).
// Domain packages use this via the internal request.Doer interface.
func (c *Client) DoWithBody(ctx context.Context, method, path string, query url.Values, contentType string, body io.Reader, out any) error {
	c.LastMeta = nil
	if contentType == "" {
		contentType = "application/json"
	}
	statusCode, slurp, err := transport.Do(ctx, method, c.baseURL, path, query, body, contentType, c.apiKey, c.httpClient)
	if err != nil {
		return err
	}
	return c.handleResponse(statusCode, slurp, out)
}

// handleResponse parses the v3 envelope, sets LastMeta, and returns an APIError or decodes data into out.
func (c *Client) handleResponse(statusCode int, slurp []byte, out any) error {
	env, err := envelope.Parse(slurp)
	if err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	var meta *Meta
	if len(env.Meta) > 0 {
		meta = &Meta{}
		_ = json.Unmarshal(env.Meta, meta)
	}
	if statusCode < 200 || statusCode >= 300 {
		apiErr := &APIError{StatusCode: statusCode, Message: env.Message, Type: env.Type, Meta: meta}
		if apiErr.Message == "" {
			apiErr.Message = string(slurp)
		}
		return apiErr
	}
	if env.Status == "error" {
		return &APIError{StatusCode: statusCode, Message: env.Message, Type: env.Type, Meta: meta}
	}
	c.LastMeta = meta
	if out != nil && len(env.Data) > 0 {
		if err := json.Unmarshal(env.Data, out); err != nil {
			return fmt.Errorf("decode response data: %w", err)
		}
	}
	return nil
}

// Get sends a GET request to path and decodes the response body into out.
// out may be nil to discard the body.
func (c *Client) Get(ctx context.Context, path string, out any) error {
	return c.do(ctx, http.MethodGet, path, nil, nil, out)
}

// Post sends a POST request to path with body and decodes the response into out.
// body may be nil; out may be nil to discard the response body.
func (c *Client) Post(ctx context.Context, path string, body, out any) error {
	return c.do(ctx, http.MethodPost, path, nil, body, out)
}

// Put sends a PUT request to path with body and decodes the response into out.
// body may be nil; out may be nil to discard the response body.
func (c *Client) Put(ctx context.Context, path string, body, out any) error {
	return c.do(ctx, http.MethodPut, path, nil, body, out)
}

// Delete sends a DELETE request to path.
func (c *Client) Delete(ctx context.Context, path string) error {
	return c.do(ctx, http.MethodDelete, path, nil, nil, nil)
}
