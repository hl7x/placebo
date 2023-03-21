package file

import (
	"fmt"
	"testing"
	"time"
)

func TestFileName(t *testing.T) {
	currentTime := time.Now()
	date := fmt.Sprintf("%v%v%v", currentTime.Year(), int(currentTime.Month()), currentTime.Day())

	var tests = []struct {
		description string
		expected    string
	}{
		{"Default Case", date + "_patient_import.csv"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := FileName()
			if got != tc.expected {
				t.Fatalf("FileName()=%v expected %v", got, tc.expected)
			}
		})
	}
}