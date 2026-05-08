package photos

import (
	"encoding/json"
	"testing"
)

// Regression: a single photo whose poleHeight value arrives as an object
// (instead of the documented boolean marker) used to fail Photo decode and,
// because the photos list endpoint decodes into []Photo as a single batch,
// take the entire response down with it. PoleHeight is now typed as
// map[string]any so both shapes round-trip cleanly.
func TestPhotofirstData_UnmarshalJSON_PoleHeightObjectValue(t *testing.T) {
	body := []byte(`{
		"poleHeight": {
			"-OkLjkPOOYXkSSP8jTqJ": {
				"_measured_height": 343.14,
				"_routine_instance_id": "-OkLjkPXb4fw2yqckWs5"
			}
		}
	}`)
	var p PhotofirstData
	if err := json.Unmarshal(body, &p); err != nil {
		t.Fatalf("unmarshal poleHeight object value: %v", err)
	}
	if got := len(p.PoleHeight); got != 1 {
		t.Fatalf("PoleHeight length: got %d want 1", got)
	}
	v, ok := p.PoleHeight["-OkLjkPOOYXkSSP8jTqJ"]
	if !ok {
		t.Fatalf("PoleHeight key missing")
	}
	if _, ok := v.(map[string]any); !ok {
		t.Errorf("PoleHeight value: got %T want map[string]any", v)
	}
}

func TestPhotofirstData_UnmarshalJSON_PoleHeightBoolValue(t *testing.T) {
	body := []byte(`{
		"poleHeight": {
			"-OkLjkPOOYXkSSP8jTqJ": true
		}
	}`)
	var p PhotofirstData
	if err := json.Unmarshal(body, &p); err != nil {
		t.Fatalf("unmarshal poleHeight bool value: %v", err)
	}
	v, ok := p.PoleHeight["-OkLjkPOOYXkSSP8jTqJ"]
	if !ok {
		t.Fatalf("PoleHeight key missing")
	}
	if b, ok := v.(bool); !ok || !b {
		t.Errorf("PoleHeight value: got %v (%T) want true", v, v)
	}
}

// Photo decode must succeed end-to-end with either poleHeight shape so the
// list endpoint stays usable for jobs whose photos mix workflow states.
func TestPhoto_UnmarshalJSON_MixedPoleHeightShapesAcrossPhotos(t *testing.T) {
	body := []byte(`[
		{
			"id": "photo-1",
			"photofirst_data": {
				"poleHeight": {"anchor-a": true}
			}
		},
		{
			"id": "photo-2",
			"photofirst_data": {
				"poleHeight": {
					"anchor-b": {"_measured_height": 12.5}
				}
			}
		}
	]`)
	var photos []Photo
	if err := json.Unmarshal(body, &photos); err != nil {
		t.Fatalf("unmarshal photos list with mixed poleHeight shapes: %v", err)
	}
	if len(photos) != 2 {
		t.Fatalf("photos len: got %d want 2", len(photos))
	}
	if photos[0].ID != "photo-1" || photos[1].ID != "photo-2" {
		t.Errorf("photo ids not preserved: %+v, %+v", photos[0].ID, photos[1].ID)
	}
}
