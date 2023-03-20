package cmd

import (
	"errors"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	os.Args = []string{"placebo", "csv", ""}

	var tests = []struct {
		description string
		want        error
		input       string
	}{
		{"Empty String", nil, ""},
		{"Invalid Input", errors.New("not a Valid Command"), "taco"},
		{"Default 'csv' command", nil, "csv"},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := File(tc.input)
			if (got == nil) != (tc.want == nil) || (got != nil && got.Error() != tc.want.Error()) {
				t.Errorf("File(%s)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
