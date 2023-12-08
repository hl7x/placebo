package cmd

import (
	"testing"
	"net"
	"time"
	"os"
	"errors"

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
//	time.Sleep(10000 * time.Millisecond)

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
		expected	error
		input		string
		args		[]string
	}{
		{"Should Return Nil When No Command Is Given", nil, "", []string{""}},
//		{"Should Return Nil When 'hl7' Command Is Given", nil, "hl7", []string{"placebo", "--send", "hl7"}},
		{"Should Return Nil When Given Proper Sub Command", nil, "hl7", []string{"placebo","--send", "hl7", "post_discharge"}},
		{"Should Return Error When Given Bad Sub Command", errors.New("Command Not Found."), "hl7", []string{"placebo", "--send", "hl7", "taco"}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			os.Args = tc.args
			got := SendHl7Message(tc.input)

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
