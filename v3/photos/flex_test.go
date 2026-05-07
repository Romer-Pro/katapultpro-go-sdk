package photos

import (
	"encoding/json"
	"math"
	"testing"
)

func TestFlexFloat64_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{"plain number", `5`, 5, false},
		{"decimal number", `5.5`, 5.5, false},
		{"zero", `0`, 0, false},
		{"negative", `-3.25`, -3.25, false},
		{"quoted integer", `"5"`, 5, false},
		{"quoted decimal", `"5.5"`, 5.5, false},
		{"quoted negative", `"-3.25"`, -3.25, false},
		{"empty string", `""`, 0, false},
		{"null", `null`, 0, false},
		{"invalid quoted", `"abc"`, 0, true},
		{"invalid raw", `{}`, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var f flexFloat64
			err := f.UnmarshalJSON([]byte(tc.input))
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tc.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tc.input, err)
			}
			if math.Abs(float64(f)-tc.want) > 1e-9 {
				t.Fatalf("input %q: got %v, want %v", tc.input, float64(f), tc.want)
			}
		})
	}
}

func TestFlexFloat64_MarshalJSON_RoundTrip(t *testing.T) {
	original := flexFloat64(5.5)
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(b) != "5.5" {
		t.Fatalf("expected 5.5, got %q", string(b))
	}
	var got flexFloat64
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got != original {
		t.Fatalf("round-trip mismatch: got %v, want %v", got, original)
	}
}
