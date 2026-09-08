// Package mllp implements HL7's Minimal Lower Layer Protocol, the framing
// every real HL7 interface engine expects on a TCP connection. A message is
// wrapped as <VT>message<FS><CR>. Without the wrapper a receiver either
// discards the bytes or waits forever for an end block that never arrives.
package mllp

import (
	"bufio"
	"fmt"
	"io"
)

const (
	// StartBlock marks the beginning of a message.
	StartBlock = 0x0B
	// EndBlock marks the end of a message.
	EndBlock = 0x1C
	// CarriageReturn always follows the end block.
	CarriageReturn = 0x0D
)

// Encode wraps a message in an MLLP block ready to be written to a connection.
func Encode(message string) []byte {

	block := make([]byte, 0, len(message)+3)

	block = append(block, StartBlock)
	block = append(block, message...)
	block = append(block, EndBlock, CarriageReturn)

	return block
}

// Reader pulls MLLP blocks off a stream. A single connection may carry any
// number of them, so callers should keep reading until EOF.
type Reader struct {
	buf *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{buf: bufio.NewReader(r)}
}

// ReadMessage returns the next message on the stream. Anything that does not
// start with a start block is read whole and reported with framed set to
// false, so an unframed sender gets a warning rather than silence.
func (r *Reader) ReadMessage() (message string, framed bool, err error) {

	first, err := r.buf.Peek(1)
	if err != nil {
		return "", false, err
	}

	if first[0] != StartBlock {
		raw, err := io.ReadAll(r.buf)
		if err != nil {
			return "", false, err
		}

		return string(raw), false, nil
	}

	_, err = r.buf.Discard(1)
	if err != nil {
		return "", true, err
	}

	message, err = r.buf.ReadString(EndBlock)
	if err != nil {
		return "", true, fmt.Errorf("unterminated MLLP block: %w", err)
	}

	message = message[:len(message)-1]

	trailer, err := r.buf.ReadByte()
	if err != nil {
		return "", true, fmt.Errorf("MLLP block missing its carriage return: %w", err)
	}

	if trailer != CarriageReturn {
		return "", true, fmt.Errorf("malformed MLLP block: expected 0x%02X after the end block, got 0x%02X", CarriageReturn, trailer)
	}

	return message, true, nil
}
