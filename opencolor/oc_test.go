package opencolor

import (
	"image/color"
	"testing"
)

// TestNumberOfColors tests that there is an expected number of colors defined.
// This test is more for fun than anything else.
func TestNumberOfColors(t *testing.T) {
	// Simply test that there are 130 colors in total: 13 sets, 10 colors per set.
	allCount := 130
	want := allCount
	got := len(All)
	if got != want {
		t.Errorf("unexpected number of defined colors (All): got=%d, want=%d", got, want)
	}

	// Check that every set has 10 colors defined.
	setCount := 0
	sets := []struct {
		name   string
		colors []color.RGBA
	}{
		{"Grays", Grays},
		{"Reds", Reds},
		{"Pinks", Pinks},
		{"Grapes", Grapes},
		{"Violets", Violets},
		{"Indigos", Indigos},
		{"Blues", Blues},
		{"Cyans", Cyans},
		{"Teals", Teals},
		{"Greens", Greens},
		{"Limes", Limes},
		{"Yellows", Yellows},
		{"Oranges", Oranges},
	}
	want = 10
	for _, set := range sets {
		got := len(set.colors)
		if got != want {
			t.Errorf("unexpected number of defined colors (Set: %s): got=%d, want=%d",
				set.name, got, want)
		}
		setCount += got
	}

	// Check that all colors in sets add up to the expected number.
	// This is quite silly, but still...
	if allCount != setCount {
		t.Errorf("defined colors in sets don't add up: sets=%d, all=%d", setCount, allCount)
	}
}
