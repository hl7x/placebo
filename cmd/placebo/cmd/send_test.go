package cmd

import (
	"testing"
	"net"
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

	// Delay for GitHub Actions test assessment
	time.Sleep(10000 * time.Millisecond)

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
			// You can add logic here to respond to certain messages if necessary
		}(conn)
	}
}

func TestSendHl7Message(t *testing.T) {

	var tests = []struct{
		description	string
		input		string
		expected	error
	}{
		{"Should Return Nil When No Command Is Given", "", nil},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := SendHl7Message(tc.input)

			if got != tc.expected {
				t.Fatalf("SendHl7Message(%v)=%v expected %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestMultiSender(t *testing.T) {

	message := []string{"admit"}

	var tests = []struct{
		description	string
		input		[]string
		expected	error
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

func TestEventSelector(t *testing.T) {
	
	command := "post_discharge"

	var tests = []struct{
		description	string
		input		string
		expected	error
	}{
		{"Should Return Nil When command selector is correct", command, nil},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := EventSelector(tc.input)

			if got != tc.expected {
				t.Fatalf("EventSelector(%v)=%v expected %v", tc.input, got, tc.expected)
			}
		})
	}

}
