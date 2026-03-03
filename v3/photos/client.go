package photos

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/romer-pro/katapultpro-go-sdk/v3/internal/request"
)

// Client performs photo API operations for a single job. Create with NewClient(do, jobID).
type Client struct {
	do    request.Doer
	jobID string
}

// NewClient returns a photos client for the given job.
func NewClient(do request.Doer, jobID string) *Client {
	return &Client{do: do, jobID: jobID}
}

// Photo returns a PhotoScope for the given photo, enabling drill-down to elements and anchors.
// Example: client.Job(jobID).Photos().Photo(photoID).Elements().List(ctx)
func (c *Client) Photo(photoID string) *PhotoScope {
	return &PhotoScope{do: c.do, jobID: c.jobID, photoID: photoID}
}

// PhotoScope scopes operations to a single photo. Use Photos().Photo(photoID) to create one.
type PhotoScope struct {
	do      request.Doer
	jobID   string
	photoID string
}

// PhotoID returns the scoped photo ID.
func (s *PhotoScope) PhotoID() string { return s.photoID }

// Get returns the photo record.
func (s *PhotoScope) Get(ctx context.Context) (*Photo, error) {
	path := "v3/jobs/" + s.jobID + "/photos/" + s.photoID
	var photo Photo
	if err := s.do.Do(ctx, http.MethodGet, path, nil, nil, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Associate associates (or unassociates) the photo to a node or section.
func (s *PhotoScope) Associate(ctx context.Context, req *AssociatePhotoRequest) error {
	path := "v3/jobs/" + s.jobID + "/photos/" + s.photoID + "/associate"
	return s.do.Do(ctx, http.MethodPost, path, nil, req, nil)
}

// Elements returns an ElementsClient for this photo.
func (s *PhotoScope) Elements() *ElementsClient {
	return &ElementsClient{do: s.do, jobID: s.jobID, photoID: s.photoID}
}

// Anchors returns an AnchorsClient for this photo's calibration anchors.
func (s *PhotoScope) Anchors() *AnchorsClient {
	return &AnchorsClient{do: s.do, jobID: s.jobID, photoID: s.photoID}
}

// ElementsClient performs photo element operations for a single photo.
type ElementsClient struct {
	do      request.Doer
	jobID   string
	photoID string
}

// List returns all photo elements on the photo.
func (c *ElementsClient) List(ctx context.Context) ([]PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/photo_elements"
	var out []PhotoElement
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified photo element.
func (c *ElementsClient) Get(ctx context.Context, elementID string) (*PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/photo_elements/" + elementID
	var el PhotoElement
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// Create creates a new photo element.
func (c *ElementsClient) Create(ctx context.Context, req *CreatePhotoElementRequest) (*PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/photo_elements"
	var el PhotoElement
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// Update updates the specified photo element. Use opts.OnlyIfExists to avoid creating with the given id.
func (c *ElementsClient) Update(ctx context.Context, elementID string, req *UpdatePhotoElementRequest, opts *UpdatePhotoElementOptions) (*PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/photo_elements/" + elementID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var el PhotoElement
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// Delete deletes the specified photo element.
func (c *ElementsClient) Delete(ctx context.Context, elementID string) error {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/photo_elements/" + elementID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// AnchorsClient performs calibration anchor operations for a single photo.
type AnchorsClient struct {
	do      request.Doer
	jobID   string
	photoID string
}

// List returns all calibration anchors on the photo.
func (c *AnchorsClient) List(ctx context.Context) ([]PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/calibration_anchors"
	var out []PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified calibration anchor.
func (c *AnchorsClient) Get(ctx context.Context, anchorID string) (*PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/calibration_anchors/" + anchorID
	var anchor PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &anchor); err != nil {
		return nil, err
	}
	return &anchor, nil
}

// Create creates a new calibration anchor on the photo.
func (c *AnchorsClient) Create(ctx context.Context, req *CreatePhotoCalibrationAnchorRequest) (*PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/calibration_anchors"
	var anchor PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &anchor); err != nil {
		return nil, err
	}
	return &anchor, nil
}

// Update updates the specified calibration anchor. Use opts.OnlyIfExists to avoid creating with the given id.
func (c *AnchorsClient) Update(ctx context.Context, anchorID string, req *UpdatePhotoCalibrationAnchorRequest, opts *UpdatePhotoCalibrationAnchorOptions) (*PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/calibration_anchors/" + anchorID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var anchor PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &anchor); err != nil {
		return nil, err
	}
	return &anchor, nil
}

// Delete deletes the specified calibration anchor.
func (c *AnchorsClient) Delete(ctx context.Context, anchorID string) error {
	path := "v3/jobs/" + c.jobID + "/photos/" + c.photoID + "/calibration_anchors/" + anchorID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// List returns all photo records in the job (v3).
func (c *Client) List(ctx context.Context) ([]Photo, error) {
	path := "v3/jobs/" + c.jobID + "/photos"
	var out []Photo
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns the specified photo record (v3).
func (c *Client) Get(ctx context.Context, photoID string) (*Photo, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID
	var photo Photo
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Upload uploads a photo (image/jpeg) to the job (v3).
func (c *Client) Upload(ctx context.Context, imageData io.Reader) (*Photo, error) {
	path := "v3/jobs/" + c.jobID + "/photos"
	var photo Photo
	if err := c.do.DoWithBody(ctx, http.MethodPost, path, nil, "image/jpeg", imageData, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// Associate associates (or unassociates) the photo to a node or section (v3).
func (c *Client) Associate(ctx context.Context, photoID string, req *AssociatePhotoRequest) error {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/associate"
	return c.do.Do(ctx, http.MethodPost, path, nil, req, nil)
}

// ListElements returns all photo elements on the photo (v3). API path: photo_elements.
func (c *Client) ListElements(ctx context.Context, photoID string) ([]PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/photo_elements"
	var out []PhotoElement
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetElement returns the specified photo element (v3).
func (c *Client) GetElement(ctx context.Context, photoID, elementID string) (*PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/photo_elements/" + elementID
	var el PhotoElement
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// CreateElement creates a new photo element (v3).
func (c *Client) CreateElement(ctx context.Context, photoID string, req *CreatePhotoElementRequest) (*PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/photo_elements"
	var el PhotoElement
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// UpdateElement updates the specified photo element (v3). Use opts.OnlyIfExists to avoid creating with the given id.
func (c *Client) UpdateElement(ctx context.Context, photoID, elementID string, req *UpdatePhotoElementRequest, opts *UpdatePhotoElementOptions) (*PhotoElement, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/photo_elements/" + elementID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var el PhotoElement
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &el); err != nil {
		return nil, err
	}
	return &el, nil
}

// DeleteElement deletes the specified photo element (v3).
func (c *Client) DeleteElement(ctx context.Context, photoID, elementID string) error {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/photo_elements/" + elementID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// ListCalibrationAnchors returns all calibration anchors on the photo (v3).
func (c *Client) ListCalibrationAnchors(ctx context.Context, photoID string) ([]PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/calibration_anchors"
	var out []PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetCalibrationAnchor returns the specified calibration anchor (v3).
func (c *Client) GetCalibrationAnchor(ctx context.Context, photoID, anchorID string) (*PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/calibration_anchors/" + anchorID
	var anchor PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodGet, path, nil, nil, &anchor); err != nil {
		return nil, err
	}
	return &anchor, nil
}

// CreateCalibrationAnchor creates a new calibration anchor on the photo (v3).
func (c *Client) CreateCalibrationAnchor(ctx context.Context, photoID string, req *CreatePhotoCalibrationAnchorRequest) (*PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/calibration_anchors"
	var anchor PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodPost, path, nil, req, &anchor); err != nil {
		return nil, err
	}
	return &anchor, nil
}

// UpdateCalibrationAnchor updates the specified calibration anchor (v3). Use opts.OnlyIfExists to avoid creating with the given id.
func (c *Client) UpdateCalibrationAnchor(ctx context.Context, photoID, anchorID string, req *UpdatePhotoCalibrationAnchorRequest, opts *UpdatePhotoCalibrationAnchorOptions) (*PhotoCalibrationAnchor, error) {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/calibration_anchors/" + anchorID
	var q url.Values
	if opts != nil && opts.OnlyIfExists {
		q = url.Values{}
		q.Set("onlyIfExists", "true")
	}
	var anchor PhotoCalibrationAnchor
	if err := c.do.Do(ctx, http.MethodPost, path, q, req, &anchor); err != nil {
		return nil, err
	}
	return &anchor, nil
}

// DeleteCalibrationAnchor deletes the specified calibration anchor (v3).
func (c *Client) DeleteCalibrationAnchor(ctx context.Context, photoID, anchorID string) error {
	path := "v3/jobs/" + c.jobID + "/photos/" + photoID + "/calibration_anchors/" + anchorID
	return c.do.Do(ctx, http.MethodDelete, path, nil, nil, nil)
}
