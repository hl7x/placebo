package sugarpill

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hl7x/placebo/pkg/random"
)

func TestNewLayoutMessage(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	var tests = []struct {
		description string
		layout      []string
		present     []string
		absent      []string
	}{
		{
			"Builds Only The Segments Named",
			[]string{"MSH", "PID", "SCH"},
			[]string{"MSH", "PID", "SCH"},
			[]string{"EVN", "PV1", "ORC", "RF1", "TXA"},
		},
		{
			"Ignores A Segment It Cannot Build",
			[]string{"MSH", "ZZZ", "PID"},
			[]string{"MSH", "PID"},
			[]string{"ZZZ"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewLayoutMessage(patient, tc.layout)

			for _, name := range tc.present {
				if got.Segment(name) == nil {
					t.Fatalf("NewLayoutMessage(%v) is missing %v", tc.layout, name)
				}
			}

			for _, name := range tc.absent {
				if got.Segment(name) != nil {
					t.Fatalf("NewLayoutMessage(%v) should not carry %v", tc.layout, name)
				}
			}
		})
	}
}

// A segment the message does not carry is left out of the rendered message
// rather than sent as an empty one.
func TestMessageBuilderForSkipsAbsentSegments(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	message := NewLayoutMessage(patient, []string{"MSH", "PID"})

	got := MessageBuilderFor(message, []string{"MSH", "EVN", "PID", "ORC"})

	if strings.Contains(got, "EVN|") || strings.Contains(got, "ORC|") {
		t.Fatalf("MessageBuilderFor()=%v, expected absent segments to be left out", got)
	}

	if !strings.Contains(got, "MSH|") || !strings.Contains(got, "PID|") {
		t.Fatalf("MessageBuilderFor()=%v, expected the segments the message carries", got)
	}

	// Two segments means one separator, with none left dangling where a
	// skipped segment used to be.
	if count := strings.Count(got, "\r"); count != 1 {
		t.Fatalf("MessageBuilderFor()=%q, expected 1 segment separator, got %v", got, count)
	}
}

func TestMessageBuilderForOrder(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	layout := []string{"MSH", "SCH", "PID", "PV1"}

	got := MessageBuilderFor(NewLayoutMessage(patient, layout), layout)

	segments := strings.Split(got, "\r")

	if len(segments) != len(layout) {
		t.Fatalf("MessageBuilderFor()=%q, expected %v segments, got %v", got, len(layout), len(segments))
	}

	for i, name := range layout {
		if !strings.HasPrefix(segments[i], name+"|") {
			t.Fatalf("MessageBuilderFor() segment %v is %q, expected it to start with %v", i, segments[i], name)
		}
	}
}

// A message that came back from JSON with segments removed still renders,
// since the sugarpill flow hands the whole message over to be edited.
func TestMessageBuilderAfterSegmentsRemoved(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	message := NewHL7Message(patient)
	message.EVN = nil
	message.PV1 = nil

	got := MessageBuilder(message)

	if strings.Contains(got, "EVN|") || strings.Contains(got, "PV1|") {
		t.Fatalf("MessageBuilder()=%v, expected removed segments to be left out", got)
	}

	if !strings.Contains(got, "MSH|") {
		t.Fatalf("MessageBuilder()=%v, expected the segments that remain", got)
	}
}

// A segment filled in by hand is rendered even though the default build does
// not include it, so sugarpill can be used to craft any message type.
func TestMessageBuilderRendersSegmentsAddedByHand(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	message := NewHL7Message(patient)

	if message.SCH != nil {
		t.Fatal("NewHL7Message() should not build a scheduling segment by default")
	}

	message.SCH = NewSCHSegment(patient)

	if !strings.Contains(MessageBuilder(message), "SCH|") {
		t.Fatal("MessageBuilder() should render a segment that was added by hand")
	}
}

func TestApplyEvent(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	var tests = []struct {
		description string
		messageType string
		trigger     string
		check       func(*HL7Message) error
	}{
		{"Stamps The Message Type And Trigger Onto MSH", "ADT", "A04", func(m *HL7Message) error {
			if m.MSH.MessageType.MessageCode != "ADT" || m.MSH.MessageType.TriggerEvent != "A04" {
				return fmt.Errorf("MSH-9 is %v^%v", m.MSH.MessageType.MessageCode, m.MSH.MessageType.TriggerEvent)
			}
			return nil
		}},
		{"Records The Trigger Event In EVN", "ADT", "A04", func(m *HL7Message) error {
			if m.EVN.EventTypeCode != "A04" {
				return fmt.Errorf("EVN-1 is %v, expected the trigger event", m.EVN.EventTypeCode)
			}
			return nil
		}},
		{"Reports Results Against An Existing Order", "ORU", "R01", func(m *HL7Message) error {
			if m.ORC.OrderControl != "RE" {
				return fmt.Errorf("ORC-1 is %v, expected RE", m.ORC.OrderControl)
			}
			return nil
		}},
		{"Leaves An Order Without A Result Status", "ORM", "O01", func(m *HL7Message) error {
			if m.OBR.ResultStatus != "" {
				return fmt.Errorf("OBR-25 is %v, expected no result status on an order", m.OBR.ResultStatus)
			}
			return nil
		}},
		{"Carries Document Text Rather Than A Lab Value", "MDM", "T02", func(m *HL7Message) error {
			if m.OBX.ValueType != "TX" {
				return fmt.Errorf("OBX-2 is %v, expected TX", m.OBX.ValueType)
			}
			if m.OBX.AbnormalFlags != "" {
				return fmt.Errorf("OBX-8 is %v, expected no abnormal flag on a document", m.OBX.AbnormalFlags)
			}
			return nil
		}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			message := NewLayoutMessage(patient, SegmentOrder)

			message.ApplyEvent(patient, tc.messageType, tc.trigger)

			if err := tc.check(message); err != nil {
				t.Fatalf("ApplyEvent(%v, %v): %v", tc.messageType, tc.trigger, err)
			}
		})
	}
}

// ApplyEvent reaches for segments the message may not carry, so it has to
// cope with a message laid out without them.
func TestApplyEventOnPartialMessage(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	message := NewLayoutMessage(patient, []string{"MSH"})

	message.ApplyEvent(patient, "MDM", "T02")

	if message.MSH.MessageType.TriggerEvent != "T02" {
		t.Fatalf("ApplyEvent() left MSH-9.2 as %v", message.MSH.MessageType.TriggerEvent)
	}
}

func TestNewSegment(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	for _, name := range SegmentOrder {
		if NewSegment(name, patient) == nil {
			t.Fatalf("NewSegment(%v) returned nil, expected every listed segment to be buildable", name)
		}
	}

	if NewSegment("ZZZ", patient) != nil {
		t.Fatal("NewSegment(ZZZ) should return nil for a segment placebo does not know")
	}
}

func TestSegmentFinderKnowsEverySegment(t *testing.T) {

	for _, name := range SegmentOrder {
		if SegmentFinder(name) == nil {
			t.Fatalf("SegmentFinder(%v) returned nil, so a message carrying it cannot be read back", name)
		}
	}
}

func TestNewPatientHasEverySegmentsData(t *testing.T) {

	patient := random.NewPatient()

	// A segment constructor reads straight off these, so a patient missing
	// one panics rather than producing a message.
	if patient.Provider == nil || patient.Location == nil || patient.Order == nil || patient.Document == nil {
		t.Fatal("random.NewPatient() is missing data a segment is built from")
	}
}
