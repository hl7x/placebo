package network

import (
	"testing"
	"net"
	"time"

)

func TestRequestHandler(t *testing.T) {
	
	duration := 2 * time.Second
	timer := time.After(duration)

	var tests = []struct {
		description	string
	}{
		{"Should Accept Connection"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			go func() {

				ln, err := net.Listen("tcp", ":9702")
				if err != nil {
					t.Fatalf("Connection Not starting...")
				}

				conn, _ := ln.Accept()
				RequestHandler(conn)
			}()

			select {
			case <- timer:
				return
			default:
				return
			}
		})
	}
}
