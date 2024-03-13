package cmd

import (

	"testing"
)

func TestListener(t *testing.T) {

	port := ":9700"

	var tests = []struct {
		description 	string
		expected	error
		input		string
	}{
		{"Should Listen Handle A Net Port", nil, port},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Listener(tc.input)
			if got != tc.expected {
				t.Fatalf("Listener(%v)=%v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}
