package cmd

import (
	"errors"
	"net"
	"os"
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

func TestDefaultPort(t *testing.T) {
	if Port != "9700" {
		t.Fatalf("expected default port 9700, got %s", Port)
	}
}

func TestEnvPort(t *testing.T) {
	original := Port
	defer func() { Port = original }()

	os.Setenv("PLACEBO_PORT", "9801")
	defer os.Unsetenv("PLACEBO_PORT")

	applyEnvPort()

	if Port != "9801" {
		t.Fatalf("expected port 9801 from PLACEBO_PORT env var, got %s", Port)
	}
}

func TestCustomPort(t *testing.T) {
	customPort := "9802"

	l, err := net.Listen("tcp", "127.0.0.1:"+customPort)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	received := make(chan bool, 1)
	go func() {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		buf := make([]byte, 1024)
		conn.Read(buf)
		received <- true
	}()

	original := Port
	Port = customPort
	defer func() { Port = original }()

	if err := DefaultSend("test"); err != nil {
		t.Fatalf("DefaultSend failed on custom port: %v", err)
	}

	select {
	case <-received:
	case <-time.After(2 * time.Second):
		t.Fatal("message not received on custom port")
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
