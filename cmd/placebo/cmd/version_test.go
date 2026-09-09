package cmd

import (
	"strings"
	"testing"
)

func TestSetBuildInfo(t *testing.T) {

	originals := [3]string{version, commit, date}
	defer func() { version, commit, date = originals[0], originals[1], originals[2] }()

	SetBuildInfo("1.4.0", "3cf305b", "2026-09-08")

	got := versionString()

	for _, want := range []string{"1.4.0", "3cf305b", "2026-09-08"} {
		if !strings.Contains(got, want) {
			t.Fatalf("versionString()=%q, want it to contain %q", got, want)
		}
	}
}

func TestVersion(t *testing.T) {

	var tests = []struct {
		description string
		args        []string
		wantErr     string
	}{
		{"No Args", []string{}, ""},
		{"Help", []string{"help"}, ""},
		{"Unknown Subcommand", []string{"taco"}, `unknown subcommand "taco"`},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Version(tc.args)

			if tc.wantErr == "" {
				if got != nil {
					t.Fatalf("Version(%v)=%v, want nil", tc.args, got)
				}
				return
			}

			if got == nil || !strings.Contains(got.Error(), tc.wantErr) {
				t.Fatalf("Version(%v)=%v, want error containing %q", tc.args, got, tc.wantErr)
			}
		})
	}
}
