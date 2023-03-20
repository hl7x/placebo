package cmd

import (
	"errors"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	var tests = []struct {
		description string
		want        error
		input       string
		args        []string
	}{
		{"Empty String", nil, "", []string{""}},
		{"Invalid Input", errors.New("not a Valid Command"), "taco", []string{""}},
		{"Default 'csv' command", nil, "csv", []string{"placebo", "csv", ""}},
		{"Pass in Numbers", nil, "csv", []string{"placebo", "csv", "3"}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			os.Args = tc.args
			got := File(tc.input)
			if (got == nil) != (tc.want == nil) || (got != nil && got.Error() != tc.want.Error()) {
				t.Errorf("File(%s)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
