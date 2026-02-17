// Package envelope parses the Katapult Pro API v3 response envelope.
// It is used by the v3 client and is not part of the public API.
package envelope

import "encoding/json"

// Envelope matches the Katapult Pro v3 response shape: { status, data?, message?, type?, meta? }.
// Meta is left as raw JSON so the caller can unmarshal into their own type.
type Envelope struct {
	Status  string          `json:"status"`
	Data    json.RawMessage `json:"data,omitempty"`
	Message string          `json:"message,omitempty"`
	Type    string          `json:"type,omitempty"`
	Meta    json.RawMessage `json:"meta,omitempty"`
}

// Parse unmarshals the response body into an Envelope.
// It does not validate status or status code; the caller does that.
func Parse(body []byte) (*Envelope, error) {
	var e Envelope
	if err := json.Unmarshal(body, &e); err != nil {
		return nil, err
	}
	return &e, nil
}
