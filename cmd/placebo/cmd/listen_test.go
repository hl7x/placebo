package cmd

import (
	"net"
	"testing"
	"time"
)

func TestListenOnCustomPort(t *testing.T) {
	customPort := "9803"

	original := Port
	Port = customPort
	defer func() { Port = original }()

	ready := make(chan struct{})
	go func() {
		// Brief pause to let the goroutine below start listening
		time.Sleep(100 * time.Millisecond)
		close(ready)
	}()

	done := make(chan error, 1)
	go func() {
		done <- ListenHl7Message("hl7", []string{})
	}()

	<-ready

	conn, err := net.DialTimeout("tcp", "127.0.0.1:"+customPort, 2*time.Second)
	if err != nil {
		t.Fatalf("expected listener on custom port %s, got error: %v", customPort, err)
	}
	conn.Close()
}

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
