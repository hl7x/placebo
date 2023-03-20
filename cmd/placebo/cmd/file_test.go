package cmd

import (
	"errors"
	"testing"
)

func TestFile(t *testing.T) {
	var tests = []struct {
		description string
		want        error
		input       string
	}{
		{"Empty String", nil, ""},
		{"Invalid Input", errors.New("not a Valid Command"), "taco"},
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
