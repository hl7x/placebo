package message

import (
	"strings"
	"testing"
)

func TestSplitSegments(t *testing.T) {

	var tests = []struct {
		description string
		input       string
		want        []string
	}{
		{"Carriage Returns", "MSH|1\rEVN|2\rPID|3", []string{"MSH|1", "EVN|2", "PID|3"}},
		{"Newlines", "MSH|1\nEVN|2\nPID|3", []string{"MSH|1", "EVN|2", "PID|3"}},
		{"Windows Line Endings", "MSH|1\r\nEVN|2\r\nPID|3", []string{"MSH|1", "EVN|2", "PID|3"}},
		{"Trailing Separator Is Dropped", "MSH|1\rEVN|2\r", []string{"MSH|1", "EVN|2"}},
		{"Blank Segments Are Dropped", "MSH|1\r\rEVN|2\n\n", []string{"MSH|1", "EVN|2"}},
		{"Empty Message", "", []string{}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := SplitSegments(tc.input)

			if strings.Join(got, "|SEP|") != strings.Join(tc.want, "|SEP|") {
				t.Errorf("SplitSegments(%q)=%v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {

	got := Normalize("MSH|1\nEVN|2\r\nPID|3\n")
	want := "MSH|1\rEVN|2\rPID|3"

	if got != want {
		t.Fatalf("Normalize()=%q, want %q", got, want)
	}
}

func TestForDisplay(t *testing.T) {

	got := ForDisplay("MSH|1\rEVN|2\rPID|3")
	want := "MSH|1\nEVN|2\nPID|3"

	if got != want {
		t.Fatalf("ForDisplay()=%q, want %q", got, want)
	}
}

func TestNormalizeIsStable(t *testing.T) {

	once := Normalize("MSH|1\nEVN|2")

	if twice := Normalize(once); twice != once {
		t.Fatalf("Normalize() is not idempotent: %q then %q", once, twice)
	}
}
