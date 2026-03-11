package katapultpro

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://katapultpro.com/api/"

// Client is the Katapult Pro API v2 client.
// A Client is safe for concurrent use by multiple goroutines.
type Client struct {
	baseURL    *url.URL
	apiKey     string
	httpClient *http.Client
}

// NewClient returns a new Client for the Katapult Pro API v2.
// apiKey is sent as a query parameter on all requests and is required.
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

// GetPhotoURL retrieves a downloadable URL for a photo in a job.
// The returned URL is a signed Firebase Storage URL with a temporary token.
//
// Parameters:
//   - jobID: The job containing the photo
//   - photoID: The photo ID (obtained from job data)
//   - size: The image size variant (use PhotoSizeFull for original, PhotoSizeExtraLarge for Katapult Pro default)
//
// Returns an error if the requested size is not available for the photo.
func (c *Client) GetPhotoURL(ctx context.Context, jobID, photoID string, size PhotoSize) (*PhotoURLResponse, error) {
	path := fmt.Sprintf("v2/jobs/%s/photoURL/%s", jobID, photoID)

	query := url.Values{}
	query.Set("api_key", c.apiKey)
	if size != "" && size.IsValid() {
		query.Set("file_size", size.String())
	}

	var resp PhotoURLResponse
	if err := c.get(ctx, path, query, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// get performs a GET request and decodes the JSON response into out.
func (c *Client) get(ctx context.Context, path string, query url.Values, out any) error {
	reqURL := c.baseURL.ResolveReference(&url.URL{Path: path, RawQuery: query.Encode()})

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	if out != nil && len(body) > 0 {
		if err := json.Unmarshal(body, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
