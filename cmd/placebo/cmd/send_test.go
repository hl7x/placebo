package cmd

import (
	"errors"
	"net"
	"testing"
	"time"
)

func init() {
	go mockServer()
}

func mockServer() {
	l, err := net.Listen("tcp", "127.0.0.1:9700")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			_, err := c.Read(buf)
			if err != nil {
				return
			}

		}(conn)
	}
}

func TestSendHl7Message(t *testing.T) {

	time.Sleep(1000 * time.Millisecond)

	var tests = []struct {
		description string
		expected    error
		input       string
		args        []string
	}{
		{"Should Return Nil When No Command Is Given", nil, "", []string{}},
		{"Should Return Nil When Given Proper Sub Command", nil, "hl7", []string{"discharge"}},
		{"Should Return Error When Given Bad Sub Command", errors.New("Command Not Found."), "hl7", []string{"taco"}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := SendHl7Message(tc.input, tc.args)

			if got != nil && tc.expected != nil {
				if got.Error() != tc.expected.Error() {
					t.Fatalf("SendHl7Message(%v)=%v expected %v", tc.input, got, tc.expected)
				}
			} else if got != tc.expected {
				t.Fatalf("SendHl7Message(%v)=%v, expected %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestMultiSender(t *testing.T) {

	time.Sleep(1000 * time.Millisecond)

	message := []string{"admit"}

	var tests = []struct {
		description string
		input       []string
		expected    error
	}{
		{"Send Multiple Messages", message, nil},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := MultiSender(tc.input)

			if got != tc.expected {
				t.Fatalf("MultiSender(%v)=%v expected %v", tc.input, got, tc.expected)
			}
		})
	}

}
