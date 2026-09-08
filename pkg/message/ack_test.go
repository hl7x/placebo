package message

import (
	"strings"
	"testing"
)

const admit = "MSH|^~\\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT^A01|12345|P|2.3\rEVN|A01|202405290800\rPID|1|56789"

func TestAck(t *testing.T) {

	ack := Ack(admit)

	segments := SplitSegments(ack)
	if len(segments) != 2 {
		t.Fatalf("Ack() produced %d segments, want 2:\n%v", len(segments), ack)
	}

	if !strings.Contains(ack, SegmentSeparator) || strings.Contains(ack, "\n") {
		t.Fatalf("Ack() must separate segments with carriage returns only, got %q", ack)
	}

	msh := strings.Split(segments[0], "|")

	var tests = []struct {
		description string
		index       int
		want        string
	}{
		{"Sending Application Is The Original Receiver", mshSendingApplication, "RECVAPP"},
		{"Sending Facility Is The Original Receiver", mshSendingFacility, "LAB"},
		{"Receiving Application Is The Original Sender", mshReceivingApplication, "SENDAPP"},
		{"Receiving Facility Is The Original Sender", mshReceivingFacility, "PLACEBO"},
		{"Message Type Mirrors The Trigger", mshMessageType, "ACK^A01"},
		{"Control ID Is Carried Over", mshMessageControlID, "12345"},
		{"Processing ID Is Carried Over", mshProcessingID, "P"},
		{"Version Is Carried Over", mshVersionID, "2.3"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if field(msh, tc.index) != tc.want {
				t.Errorf("ACK MSH field %d=%q, want %q", tc.index, field(msh, tc.index), tc.want)
			}
		})
	}

	msa := strings.Split(segments[1], "|")

	if msa[0] != "MSA" || msa[1] != "AA" || msa[2] != "12345" {
		t.Errorf("MSA=%q, want MSA|AA|12345", segments[1])
	}
}

func TestAckAcceptsNewlineSeparatedInput(t *testing.T) {

	ack := Ack(strings.ReplaceAll(admit, "\r", "\n"))

	if !strings.Contains(ack, "|12345|") {
		t.Fatalf("Ack() did not pick up the control ID from a newline-separated message:\n%v", ack)
	}
}

func TestAckRejectsUnparseableMessage(t *testing.T) {

	ack := Ack("this is not an HL7 message")

	segments := SplitSegments(ack)
	if len(segments) != 2 {
		t.Fatalf("Ack() produced %d segments, want 2", len(segments))
	}

	if !strings.HasPrefix(segments[1], "MSA|AR|") {
		t.Errorf("MSA=%q, want an AR reject", segments[1])
	}
}

func TestAckWithoutATriggerEvent(t *testing.T) {

	ack := Ack("MSH|^~\\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT|12345|P|2.3")

	msh := strings.Split(SplitSegments(ack)[0], "|")

	if field(msh, mshMessageType) != "ACK" {
		t.Errorf("message type=%q, want ACK", field(msh, mshMessageType))
	}
}

func TestAckFillsMissingTrailingFields(t *testing.T) {

	ack := Ack("MSH|^~\\&|SENDAPP|PLACEBO|RECVAPP|LAB|202405290800||ADT^A01|12345")

	msh := strings.Split(SplitSegments(ack)[0], "|")

	if field(msh, mshProcessingID) != "P" {
		t.Errorf("processing ID=%q, want the P default", field(msh, mshProcessingID))
	}

	if field(msh, mshVersionID) != "2.3" {
		t.Errorf("version=%q, want the 2.3 default", field(msh, mshVersionID))
	}
}
