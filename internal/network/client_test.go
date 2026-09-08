package network

import (
	"net"
	"strings"
	"testing"
	"time"

	"github.com/hl7x/placebo/internal/mllp"
)

//TODO: TestSendClient

func TestRequestHandler(t *testing.T) {

	duration := 2 * time.Second
	timer := time.After(duration)

	var tests = []struct {
		description string
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
			case <-timer:
				return
			default:
				return
			}
		})
	}
}

//TODO: TestListenClient

// readFrame reads one MLLP block off a connection the way a real interface
// engine would, without going through placebo's own reader.
func readFrame(t *testing.T, conn net.Conn) string {
	t.Helper()

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("reading frame: %v", err)
	}

	frame := string(buf[:n])

	if frame[0] != 0x0B {
		t.Fatalf("frame does not start with 0x0B: %q", frame)
	}

	if !strings.HasSuffix(frame, "\x1c\x0d") {
		t.Fatalf("frame does not end with 0x1C 0x0D: %q", frame)
	}

	return frame[1 : len(frame)-2]
}

// TestSendClientFramesAndReadsAck stands a bare TCP listener in for a real
// receiver: it checks the framing byte for byte and replies with an ACK.
func TestSendClientFramesAndReadsAck(t *testing.T) {

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	received := make(chan string, 1)

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		received <- readFrame(t, conn)

		conn.Write(mllp.Encode("MSH|^~\\&|RECVAPP|LAB|SENDAPP|PLACEBO|202405290800||ACK^A01|12345|P|2.3\rMSA|AA|12345"))
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	ack, err := SendClient("127.0.0.1", port, "MSH|^~\\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT^A01|12345|P|2.3\nEVN|A01")
	if err != nil {
		t.Fatalf("SendClient() error=%v", err)
	}

	select {
	case got := <-received:
		// Newlines in the source message must leave as carriage returns.
		if strings.Contains(got, "\n") {
			t.Errorf("message went out with newline separators: %q", got)
		}

		if !strings.Contains(got, "MSH|^~\\&|SENDAPP") || !strings.Contains(got, "\rEVN|A01") {
			t.Errorf("unexpected message on the wire: %q", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("receiver never got a message")
	}

	if !strings.Contains(ack, "MSA|AA|12345") {
		t.Errorf("SendClient() ack=%q, want the MSA carried back", ack)
	}
}

// TestSendClientToleratesSilentReceiver covers receivers that never
// acknowledge; the message is already delivered, so this is not an error.
func TestSendClientToleratesSilentReceiver(t *testing.T) {

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		buf := make([]byte, 4096)
		conn.Read(buf)
		conn.Close()
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())

	original := AckTimeout
	AckTimeout = 500 * time.Millisecond
	defer func() { AckTimeout = original }()

	ack, err := SendClient("127.0.0.1", port, "MSH|^~\\&|SENDAPP")
	if err != nil {
		t.Fatalf("SendClient() error=%v, want nil for a silent receiver", err)
	}

	if ack != "" {
		t.Errorf("SendClient() ack=%q, want empty", ack)
	}
}

// TestRequestHandlerAcksSender covers the other direction: a real sender
// pointed at 'placebo listen' must get an ACK back or it hangs.
func TestRequestHandlerAcksSender(t *testing.T) {

	server, client := net.Pipe()

	go RequestHandler(server)

	go func() {
		client.Write(mllp.Encode("MSH|^~\\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT^A01|999|P|2.3\rEVN|A01"))
	}()

	client.SetReadDeadline(time.Now().Add(2 * time.Second))

	ack, framed, err := mllp.NewReader(client).ReadMessage()
	if err != nil {
		t.Fatalf("reading ack: %v", err)
	}

	if !framed {
		t.Error("ack came back without MLLP framing")
	}

	if !strings.Contains(ack, "MSA|AA|999") {
		t.Errorf("ack=%q, want MSA|AA|999", ack)
	}

	if !strings.HasPrefix(ack, "MSH|") {
		t.Errorf("ack=%q, want it to start with an MSH segment", ack)
	}
}
