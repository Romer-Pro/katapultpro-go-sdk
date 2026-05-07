package photos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// flexFloat64 decodes JSON values that the Katapult Pro API quotes as strings
// (e.g. `"5"`, `"5.0"`) but which are semantically numeric. Several response
// fields — notably calibration-anchor heights — round-trip through Firebase
// Realtime Database, which sometimes preserves the original entry as a string
// even when the field is documented as a decimal number.
//
// The type accepts: JSON numbers, JSON strings parseable as float64, JSON
// null, and the empty string (treated as zero / no-op). Marshalling emits a
// plain number so writes back to the API stay numeric.
type flexFloat64 float64

// UnmarshalJSON parses either a JSON number or a JSON-quoted numeric string.
func (f *flexFloat64) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return fmt.Errorf("flexFloat64: unquote: %w", err)
		}
		if s == "" {
			return nil
		}
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fmt.Errorf("flexFloat64: parse %q: %w", s, err)
		}
		*f = flexFloat64(v)
		return nil
	}
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return fmt.Errorf("flexFloat64: parse number: %w", err)
	}
	*f = flexFloat64(v)
	return nil
}

// MarshalJSON emits a plain JSON number so encode/decode round-trips do not
// silently turn a numeric field into a quoted string.
func (f flexFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}
