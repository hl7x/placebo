package message

import (
	"strings"
	"time"
)

// MSH field positions, counted the way a split on the field separator lands
// them. The MSH segment is offset by one against the usual HL7 numbering
// because MSH-1 is the field separator itself, so MSH-3 sits at index 2.
const (
	mshSendingApplication   = 2
	mshSendingFacility      = 3
	mshReceivingApplication = 4
	mshReceivingFacility    = 5
	mshMessageType          = 8
	mshMessageControlID     = 9
	mshProcessingID         = 10
	mshVersionID            = 11
)

// Ack builds an HL7 acknowledgement for a received message. Interface engines
// hold the connection open until they get one, so a listener that never
// replies looks like a hang to whatever is sending.
func Ack(received string) string {

	fields := mshFields(received)

	if fields == nil {
		// Nothing parseable to acknowledge. AR tells the sender the message
		// was rejected rather than leaving it waiting.
		return strings.Join([]string{
			"MSH|^~\\&|PLACEBO|PLACEBO|||" + hl7Now() + "||ACK|" + hl7Now() + "|P|2.3",
			"MSA|AR||Message could not be parsed",
		}, SegmentSeparator)
	}

	// The reply travels back the way the message came, so the sending and
	// receiving pairs swap places.
	msh := []string{
		"MSH",
		"^~\\&",
		field(fields, mshReceivingApplication),
		field(fields, mshReceivingFacility),
		field(fields, mshSendingApplication),
		field(fields, mshSendingFacility),
		hl7Now(),
		"",
		ackType(field(fields, mshMessageType)),
		field(fields, mshMessageControlID),
		processingID(field(fields, mshProcessingID)),
		versionID(field(fields, mshVersionID)),
	}

	msa := []string{
		"MSA",
		"AA",
		field(fields, mshMessageControlID),
	}

	return strings.Join([]string{
		strings.Join(msh, "|"),
		strings.Join(msa, "|"),
	}, SegmentSeparator)
}

// mshFields returns the MSH segment split into fields, or nil when the
// message has no MSH segment to reply to.
func mshFields(received string) []string {

	for _, segment := range SplitSegments(received) {
		if strings.HasPrefix(segment, "MSH|") {
			return strings.Split(segment, "|")
		}
	}

	return nil
}

func field(fields []string, i int) string {

	if i >= len(fields) {
		return ""
	}

	return fields[i]
}

// ackType mirrors the trigger event of the message being acknowledged, so
// ADT^A01 is answered with ACK^A01.
func ackType(messageType string) string {

	_, trigger, found := strings.Cut(messageType, "^")
	if !found || trigger == "" {
		return "ACK"
	}

	// Drop any structure component, keeping just the trigger event.
	trigger, _, _ = strings.Cut(trigger, "^")

	return "ACK^" + trigger
}

func processingID(id string) string {

	if id == "" {
		return "P"
	}

	return id
}

func versionID(version string) string {

	if version == "" {
		return "2.3"
	}

	return version
}

func hl7Now() string {
	return time.Now().Format("20060102150405")
}
