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
			if got != tc.want {
				t.Fatalf("File(%s)=%#v, want %#v", tc.input, got, tc.want)
			}
		})
	}
}
