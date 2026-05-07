package photos

import "encoding/json"

// UnmarshalJSON tolerates `"height"` arriving as either a JSON number (`5.5`)
// or a JSON string (`"5.5"`). Katapult Pro's API has been observed to return
// the latter for calibration-anchor responses sourced from Firebase Realtime
// Database, where numeric leaf values can persist as strings. The public
// Height field stays a plain float64 so callers don't need to change.
func (a *PhotoCalibrationAnchor) UnmarshalJSON(b []byte) error {
	type alias PhotoCalibrationAnchor
	aux := &struct {
		Height flexFloat64 `json:"height,omitempty"`
		*alias
	}{
		alias: (*alias)(a),
	}
	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}
	a.Height = float64(aux.Height)
	return nil
}

// UnmarshalJSON applies the same string-or-number tolerance for height as
// PhotoCalibrationAnchor — AnchorCalibration values are sourced from the
// same upstream and exhibit the same string-quoted-number quirk.
func (a *AnchorCalibration) UnmarshalJSON(b []byte) error {
	type alias AnchorCalibration
	aux := &struct {
		Height flexFloat64 `json:"height,omitempty"`
		*alias
	}{
		alias: (*alias)(a),
	}
	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}
	a.Height = float64(aux.Height)
	return nil
}
