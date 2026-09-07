package cmd

import (
	"strings"
	"testing"
)

func TestFile(t *testing.T) {

	var tests = []struct {
		description string
		args        []string
		wantErr     string
	}{
		{"No Subcommand", []string{}, "'placebo file' needs a subcommand"},
		{"Invalid Subcommand", []string{"taco"}, `unknown subcommand "taco" for 'placebo file'`},
		{"Default 'csv' subcommand", []string{"csv"}, ""},
		{"Default HL7 file subcommand", []string{"hl7"}, ""},
		{"Pass in Numbers", []string{"csv", "3"}, ""},
		{"Passing Wrong Arg Type", []string{"csv", "test"}, `invalid patient count "test"`},
		{"Passing Zero", []string{"csv", "0"}, `invalid patient count "0"`},
		{"Help", []string{"help"}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := File(tc.args)

			if tc.wantErr == "" {
				if got != nil {
					t.Errorf("File(%v)=%v, want nil", tc.args, got)
				}
				return
			}

			if got == nil || !strings.Contains(got.Error(), tc.wantErr) {
				t.Errorf("File(%v)=%v, want error containing %q", tc.args, got, tc.wantErr)
			}
		})
	}
}
