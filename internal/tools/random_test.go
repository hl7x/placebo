package tools

import (
	"testing"
)

func TestRandomSelector(t *testing.T) {

	var tests = []struct {
		description string
		want        int
		input       []string
	}{
		{"List of Names", 1, []string{"John", "Bill"}},
		{"List of Places", 2, []string{"NEW YORK", "DENVER", "DALLAS"}},
		{"Numbers as Strings", 0, []string{"1"}},
		{"Empty String", 0, []string{""}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := RandomSelector(tc.input)
			if got != tc.want && got >= tc.want {
				t.Fatalf("RandomSelector(%s)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
