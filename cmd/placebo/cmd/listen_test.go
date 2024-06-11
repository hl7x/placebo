package cmd

import (
	"testing"
	"time"
)

// Very basic
func TestListener(t *testing.T) {

	time.Sleep(1000 * time.Millisecond)

	duration := 2 * time.Second
	timer := time.After(duration)

	port := ":9701"

	var tests = []struct {
		description string
		expected    error
		input       string
	}{
		{"Should Handle Listening  On A Net Port", nil, port},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {

			done := make(chan bool)

			go func() {
				got := Listener(tc.input)
				if got != tc.expected {
					t.Fatalf("Listener(%v)=%v, want %v", tc.input, got, tc.expected)
				}
			}()

			select {
			case <-timer:
				return
			case <-done:
				return
			}
		})
	}
}
