package message

import "strings"

// SegmentSeparator is the HL7 segment terminator. The spec calls for a
// carriage return; a newline is not a valid separator and strict parsers
// reject a message that uses one.
const SegmentSeparator = "\r"

// SplitSegments breaks a message into its segments, accepting carriage
// returns, newlines, or both, since messages reach placebo from files people
// have edited as well as off the wire. Blank segments are dropped.
func SplitSegments(msg string) []string {

	normalized := strings.ReplaceAll(msg, "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")

	segments := []string{}

	for _, segment := range strings.Split(normalized, "\n") {
		if strings.TrimSpace(segment) != "" {
			segments = append(segments, segment)
		}
	}

	return segments
}

// Normalize puts a message into wire form: segments terminated by carriage
// returns. Anything placebo sends goes through here, so a message typed or
// edited with newlines still leaves as valid HL7.
func Normalize(msg string) string {
	return strings.Join(SplitSegments(msg), SegmentSeparator)
}

// ForDisplay puts a message into newline form for printing to a terminal or
// writing to a file someone is about to open in an editor. A bare carriage
// return overwrites the line it is on, which makes a valid message look
// truncated.
func ForDisplay(msg string) string {
	return strings.Join(SplitSegments(msg), "\n")
}
