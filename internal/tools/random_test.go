package tools

import (
	"testing"
)

func TestRandomSelector(t *testing.T) {

	var tests = []struct {
		description string
		expectedMax int
		input       []string
	}{
		{"List of Names", 4, []string{"John", "Bill", "Jacob", "Rita", "Beth"}},
		{"List of Places", 2, []string{"NEW YORK", "DENVER", "DALLAS"}},
		{"Numbers as Strings", 0, []string{"1"}},
		{"Empty String", 0, []string{""}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := RandomSelector(tc.input)
			if got < 0 && got > tc.expectedMax {
				t.Fatalf("RandomSelector(%s)=%v, expected a value between 0 and %v", tc.input, got, tc.expectedMax)
			}
		})
	}
}
