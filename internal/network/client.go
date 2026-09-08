package network

import (
	"errors"
	"fmt"
	"io"
	"net"
	"syscall"
	"time"

	"github.com/hl7x/placebo/internal/mllp"
	"github.com/hl7x/placebo/pkg/message"
)

// AckTimeout bounds the wait for an acknowledgement. Not every receiver
// sends one, so the wait has to end on its own rather than hang the CLI.
var AckTimeout = 10 * time.Second

// SendClient sends data as an MLLP block and returns the acknowledgement the
// receiver sent back. An empty ack means the receiver stayed silent or closed
// the connection, which is not treated as a failure.
func SendClient(ip string, port string, data string) (string, error) {

	serverAddr := ip + ":" + port

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	_, err = conn.Write(mllp.Encode(message.Normalize(data)))
	if err != nil {
		return "", err
	}

	err = conn.SetReadDeadline(time.Now().Add(AckTimeout))
	if err != nil {
		return "", err
	}

	ack, _, err := mllp.NewReader(conn).ReadMessage()
	if err != nil {
		if isSilence(err) {
			return "", nil
		}

		return "", err
	}

	return ack, nil
}

// isSilence reports whether an error just means no acknowledgement arrived.
// The message is already on the wire at that point, so a receiver that stays
// quiet, hangs up, or resets the connection is not a send failure.
func isSilence(err error) bool {

	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return true
	}

	if errors.Is(err, syscall.ECONNRESET) {
		return true
	}

	var netErr net.Error

	return errors.As(err, &netErr) && netErr.Timeout()
}

func RequestHandler(conn net.Conn) {

	defer conn.Close()

	reader := mllp.NewReader(conn)

	// A sender may put many messages on one connection, so keep reading
	// until it hangs up.
	for {
		received, framed, err := reader.ReadMessage()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Println("ERROR: ", err)
			}

			return
		}

		if !framed {
			fmt.Println("Warning: message arrived without MLLP framing. HL7 senders wrap messages as <VT>message<FS><CR>.")
		}

		fmt.Printf("Recieved: %q\n", received)

		_, err = conn.Write(mllp.Encode(message.Ack(received)))
		if err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
	}
}

func ListenClient(port string) error {

	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Printf("Listening on TCP port %v\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go RequestHandler(conn)
	}
}
