package event

import (
	"strings"
	"testing"

	"github.com/hl7x/placebo/pkg/random"
)

func TestLookup(t *testing.T) {

	var tests = []struct {
		description string
		input       string
		found       bool
		expected    string
	}{
		{"Finds A Message Type By Code", "ADT", true, "ADT"},
		{"Matches A Code Regardless Of Case", "oru", true, "ORU"},
		{"Does Not Find An Unknown Code", "XYZ", false, ""},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got, ok := Lookup(tc.input)

			if ok != tc.found {
				t.Fatalf("Lookup(%v) found=%v, expected %v", tc.input, ok, tc.found)
			}

			if got.Code != tc.expected {
				t.Fatalf("Lookup(%v)=%v, expected %v", tc.input, got.Code, tc.expected)
			}
		})
	}
}

func TestResolve(t *testing.T) {

	var tests = []struct {
		description     string
		input           []string
		expectedType    string
		expectedTrigger string
		wantErr         string
	}{
		{"Nothing Given Builds An Admit", nil, "ADT", "admit", ""},
		{"A Scenario Alone Stays An ADT Event", []string{"discharge"}, "ADT", "discharge", ""},
		{"A Message Type Alone Uses Its First Scenario", []string{"ORU"}, "ORU", "result", ""},
		{"A Message Type Is Matched Regardless Of Case", []string{"siu"}, "SIU", "schedule", ""},
		{"A Message Type And Scenario Resolve Together", []string{"siu", "reschedule"}, "SIU", "reschedule", ""},
		{"A Scenario Shared By Two Types Resolves Per Type", []string{"ref", "cancel"}, "REF", "cancel", ""},
		{"An Unknown Scenario Reports The Ones That Exist", []string{"oru", "taco"}, "", "", "ORU scenarios: result"},
		{"An Unknown Token Reports Types And ADT Scenarios", []string{"taco"}, "", "", "Message types:"},
		{"An Unknown Type With A Scenario Is Reported", []string{"taco", "admit"}, "", "", "unknown message type"},
		{"More Than One Scenario Is Refused", []string{"adt", "admit", "extra"}, "", "", "too many arguments"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			gotType, gotTrigger, err := Resolve(tc.input)

			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("Resolve(%v)=%v, want error containing %q", tc.input, err, tc.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Resolve(%v) returned %v, want no error", tc.input, err)
			}

			if gotType != tc.expectedType || gotTrigger != tc.expectedTrigger {
				t.Fatalf("Resolve(%v)=%v %v, expected %v %v", tc.input, gotType, gotTrigger, tc.expectedType, tc.expectedTrigger)
			}
		})
	}
}

// A message type is only worth having if the message it builds announces that
// type and carries the segment the type is defined around.
func TestBuildMessageTypes(t *testing.T) {

	patient := random.NewPatient()

	var tests = []struct {
		description string
		messageType string
		trigger     string
		expected    string
		defining    string
	}{
		{"Admit", "ADT", "admit", "ADT^A01", "EVN|A01"},
		{"Discharge", "ADT", "discharge", "ADT^A03", "EVN|A03"},
		{"Order", "ORM", "order", "ORM^O01", "ORC|NW"},
		{"Result", "ORU", "result", "ORU^R01", "OBX|1|NM"},
		{"Appointment", "SIU", "schedule", "SIU^S12", "SCH|"},
		{"Referral", "REF", "referral", "REF^I12", "RF1|"},
		{"Document", "MDM", "document", "MDM^T02", "TXA|"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Build(patient, tc.messageType, tc.trigger)
			if err != nil {
				t.Fatalf("Build(%v, %v) returned %v", tc.messageType, tc.trigger, err)
			}

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("Build(%v, %v)=%v, expected to contain %v", tc.messageType, tc.trigger, got, tc.expected)
			}

			if !strings.Contains(got, tc.defining) {
				t.Fatalf("Build(%v, %v)=%v, expected to contain %v", tc.messageType, tc.trigger, got, tc.defining)
			}
		})
	}
}

// A message carries the segments its type is defined around and no others.
// Sending an order with an insurance segment attached is what this guards.
func TestBuildLeavesOutUnrelatedSegments(t *testing.T) {

	patient := random.NewPatient()

	var tests = []struct {
		description string
		messageType string
		trigger     string
		absent      []string
	}{
		{"An Admit Carries No Order Segments", "ADT", "admit", []string{"ORC|", "OBR|", "OBX|", "SCH|", "RF1|", "TXA|"}},
		{"A Result Carries No Administrative Segments", "ORU", "result", []string{"EVN|", "IN1|", "GT1|", "NK1|", "SCH|"}},
		{"An Appointment Carries No Order Segments", "SIU", "schedule", []string{"ORC|", "OBR|", "RF1|", "TXA|"}},
		{"A Referral Carries No Scheduling Segments", "REF", "referral", []string{"SCH|", "TXA|", "ORC|"}},
		{"A Document Carries No Referral Segments", "MDM", "document", []string{"RF1|", "SCH|", "ORC|", "IN1|"}},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Build(patient, tc.messageType, tc.trigger)
			if err != nil {
				t.Fatalf("Build(%v, %v) returned %v", tc.messageType, tc.trigger, err)
			}

			for _, segment := range tc.absent {
				if strings.Contains(got, segment) {
					t.Fatalf("Build(%v, %v)=%v, expected no %v segment", tc.messageType, tc.trigger, got, segment)
				}
			}
		})
	}
}

// EVN-1 is the trigger event. Reporting the message code there leaves the
// segments of one message disagreeing about which event it describes.
func TestBuildAgreesOnTheTriggerEvent(t *testing.T) {

	patient := random.NewPatient()

	got, err := Build(patient, "ADT", "discharge")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(got, "EVN|A03") {
		t.Fatalf("Build(ADT, discharge)=%v, expected EVN to record the A03 trigger", got)
	}

	if strings.Contains(got, "EVN|ADT") {
		t.Fatalf("Build(ADT, discharge)=%v, expected EVN to record a trigger event, not the message code", got)
	}
}

func TestBuildUnknown(t *testing.T) {

	patient := random.NewPatient()

	var tests = []struct {
		description string
		messageType string
		trigger     string
		wantErr     string
	}{
		{"Unknown Message Type", "XYZ", "admit", "unknown message type"},
		{"Unknown Trigger", "ADT", "taco", "Command taco Not Found"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			_, err := Build(patient, tc.messageType, tc.trigger)

			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("Build(%v, %v)=%v, want error containing %q", tc.messageType, tc.trigger, err, tc.wantErr)
			}
		})
	}
}

// Every registered type has to be buildable, so a type added to the registry
// without a layout or a trigger fails here rather than at the command line.
func TestEveryDefinitionBuilds(t *testing.T) {

	patient := random.NewPatient()

	for _, definition := range Definitions {
		t.Run(definition.Code, func(t *testing.T) {

			if len(definition.Triggers) == 0 {
				t.Fatalf("%v has no scenarios", definition.Code)
			}

			if len(definition.Layout) == 0 {
				t.Fatalf("%v has no segment layout", definition.Code)
			}

			for _, trigger := range definition.Triggers {
				got, err := Build(patient, definition.Code, trigger.Name)
				if err != nil {
					t.Fatalf("Build(%v, %v) returned %v", definition.Code, trigger.Name, err)
				}

				expected := definition.Code + "^" + trigger.Code

				if !strings.Contains(got, expected) {
					t.Fatalf("Build(%v, %v)=%v, expected to contain %v", definition.Code, trigger.Name, got, expected)
				}
			}
		})
	}
}

func TestCatalog(t *testing.T) {

	catalog := Catalog()

	for _, definition := range Definitions {
		if !strings.Contains(catalog, definition.Code) {
			t.Fatalf("Catalog()=%v, expected to list %v", catalog, definition.Code)
		}

		for _, trigger := range definition.Triggers {
			if !strings.Contains(catalog, trigger.Name) {
				t.Fatalf("Catalog()=%v, expected to list the %v scenario", catalog, trigger.Name)
			}
		}
	}
}
