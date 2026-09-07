package cmd

import (
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {

	var tests = []struct {
		description string
		args        []string
		wantErr     string
	}{
		{"No Args Prints Usage", []string{}, ""},
		{"Help Command", []string{"help"}, ""},
		{"Help For A Command", []string{"help", "send"}, ""},
		{"Help Flag", []string{"--help"}, ""},
		{"Help For An Unknown Command", []string{"help", "taco"}, `unknown command "taco"`},
		{"Unknown Command", []string{"taco"}, `unknown command "taco"`},
		{"Unknown Flag", []string{"--taco"}, "unknown flag --taco"},
		{"Retired Flag Points At The Command", []string{"--send", "hl7"}, `"send" is a command now, not a flag`},
		{"Retired Flag Without Dashes Is A Command", []string{"file", "hl7"}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Execute(tc.args)

			if tc.wantErr == "" {
				if got != nil {
					t.Fatalf("Execute(%v)=%v, want nil", tc.args, got)
				}
				return
			}

			if got == nil || !strings.Contains(got.Error(), tc.wantErr) {
				t.Fatalf("Execute(%v)=%v, want error containing %q", tc.args, got, tc.wantErr)
			}
		})
	}
}

func TestExtractOptions(t *testing.T) {

	var tests = []struct {
		description string
		args        []string
		wantArgs    []string
		wantPort    string
		wantErr     string
	}{
		{"No Options", []string{"send", "hl7"}, []string{"send", "hl7"}, "9700", ""},
		{"Port Before The Command", []string{"--port", "8500", "listen", "hl7"}, []string{"listen", "hl7"}, "8500", ""},
		{"Port After The Subcommand", []string{"listen", "hl7", "--port", "8500"}, []string{"listen", "hl7"}, "8500", ""},
		{"Port With Equals", []string{"listen", "hl7", "--port=8500"}, []string{"listen", "hl7"}, "8500", ""},
		{"Single Dash Port", []string{"-port", "8500", "listen"}, []string{"listen"}, "8500", ""},
		{"Port Without A Value", []string{"listen", "hl7", "--port"}, nil, "9700", "needs a port number"},
		{"Port That Is Not A Number", []string{"--port", "taco"}, nil, "9700", `invalid port "taco"`},
		{"Port Out Of Range", []string{"--port", "70000"}, nil, "9700", `invalid port "70000"`},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			original := Port
			Port = "9700"
			defer func() { Port = original }()

			got, err := extractOptions(tc.args)

			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("extractOptions(%v) error=%v, want error containing %q", tc.args, err, tc.wantErr)
				}
			} else if err != nil {
				t.Fatalf("extractOptions(%v) error=%v, want nil", tc.args, err)
			}

			if Port != tc.wantPort {
				t.Errorf("extractOptions(%v) set Port=%s, want %s", tc.args, Port, tc.wantPort)
			}

			if tc.wantErr == "" && strings.Join(got, " ") != strings.Join(tc.wantArgs, " ") {
				t.Errorf("extractOptions(%v)=%v, want %v", tc.args, got, tc.wantArgs)
			}
		})
	}
}
