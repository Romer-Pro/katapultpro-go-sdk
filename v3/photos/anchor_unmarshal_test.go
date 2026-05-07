package photos

import (
	"encoding/json"
	"math"
	"testing"
)

func TestPhotoCalibrationAnchor_UnmarshalJSON_NumericHeight(t *testing.T) {
	body := []byte(`{
		"id": "anchor-1",
		"height": 5.5,
		"_score": 0.95,
		"pixel_selection": [{"percentX": 50, "percentY": 50}]
	}`)
	var a PhotoCalibrationAnchor
	if err := json.Unmarshal(body, &a); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if a.ID != "anchor-1" {
		t.Errorf("id: got %q want %q", a.ID, "anchor-1")
	}
	if math.Abs(a.Height-5.5) > 1e-9 {
		t.Errorf("height: got %v want %v", a.Height, 5.5)
	}
	if math.Abs(a.Score-0.95) > 1e-9 {
		t.Errorf("score: got %v want %v", a.Score, 0.95)
	}
	if len(a.PixelSelection) != 1 || a.PixelSelection[0].PercentX != 50 {
		t.Errorf("pixel selection not preserved: %+v", a.PixelSelection)
	}
}

func TestPhotoCalibrationAnchor_UnmarshalJSON_StringHeight(t *testing.T) {
	body := []byte(`{
		"id": "anchor-2",
		"height": "5.5",
		"pixel_selection": [{"percent_x": 25, "percent_y": 75}]
	}`)
	var a PhotoCalibrationAnchor
	if err := json.Unmarshal(body, &a); err != nil {
		t.Fatalf("unmarshal with string height: %v", err)
	}
	if math.Abs(a.Height-5.5) > 1e-9 {
		t.Errorf("height: got %v want %v", a.Height, 5.5)
	}
}

func TestPhotoCalibrationAnchor_UnmarshalJSON_MissingHeight(t *testing.T) {
	body := []byte(`{"id": "anchor-3"}`)
	var a PhotoCalibrationAnchor
	if err := json.Unmarshal(body, &a); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if a.Height != 0 {
		t.Errorf("expected height 0, got %v", a.Height)
	}
}

func TestPhotoCalibrationAnchor_MarshalJSON_NumericRoundTrip(t *testing.T) {
	original := PhotoCalibrationAnchor{
		ID:     "anchor-4",
		Height: 12.5,
	}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	// Default marshal goes through the struct tag path — height should still be a number on the wire.
	if string(b) == "" {
		t.Fatalf("empty marshal output")
	}
	var got PhotoCalibrationAnchor
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("round-trip unmarshal: %v", err)
	}
	if got.Height != original.Height || got.ID != original.ID {
		t.Errorf("round-trip mismatch: got %+v, want %+v", got, original)
	}
}

func TestAnchorCalibration_UnmarshalJSON_StringHeight(t *testing.T) {
	body := []byte(`{"height": "10", "_score": 0.5}`)
	var a AnchorCalibration
	if err := json.Unmarshal(body, &a); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if math.Abs(a.Height-10) > 1e-9 {
		t.Errorf("height: got %v want %v", a.Height, 10.0)
	}
	if math.Abs(a.Score-0.5) > 1e-9 {
		t.Errorf("score: got %v want %v", a.Score, 0.5)
	}
}
