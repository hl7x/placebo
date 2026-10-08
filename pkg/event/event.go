package event

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/hl7x/placebo/pkg/random"
	"github.com/hl7x/placebo/pkg/sugarpill"
	"github.com/hl7x/placebo/pkg/templates"
)

/* Build HL7 message based on certain medical event for various scenarios */

// Template approach is depricated and not supported anymore.

type Event struct {
	MessageEvent string
	TriggerEvent string
	Patient      *random.Patient
}

// Trigger is one scenario inside a message type: the name placebo takes on
// the command line, and the trigger event code it writes into MSH-9.2.
type Trigger struct {
	Name        string
	Code        string
	Description string
}

// Definition is one HL7 message type. Layout is the ordered set of segments a
// message of this type carries, which is what distinguishes the types from
// each other once the patient data is filled in - an admit and a lab result
// describe the same patient through entirely different segments.
type Definition struct {
	Code        string
	Description string
	Triggers    []Trigger
	Layout      []string
}

// Definitions holds every message type placebo can generate. The first
// trigger of a type is the one used when none is named.
var Definitions = []Definition{
	{
		Code:        "ADT",
		Description: "Patient administration",
		Layout: []string{
			"MSH", "EVN", "PID", "PD1", "ROL", "DB1", "ARV", "NK1",
			"PV1", "PV2", "GT1", "IN1", "AL1", "DG1",
		},
		Triggers: []Trigger{
			{"admit", "A01", "Admit a patient"},
			{"transfer", "A02", "Transfer a patient"},
			{"discharge", "A03", "Discharge a patient"},
			{"register", "A04", "Register a patient"},
			{"pre-admit", "A05", "Establish preadmit information"},
			{"update", "A08", "Update patient information"},
			{"cancel-admit", "A11", "Cancel an admit"},
			{"cancel-discharge", "A13", "Cancel a discharge"},
		},
	},
	{
		Code:        "ORM",
		Description: "Order",
		Layout: []string{
			"MSH", "PID", "PV1", "ORC", "OBR", "DG1", "NTE",
		},
		Triggers: []Trigger{
			{"order", "O01", "Place a new order"},
		},
	},
	{
		Code:        "ORU",
		Description: "Observation result",
		Layout: []string{
			"MSH", "PID", "PV1", "ORC", "OBR", "OBX", "NTE",
		},
		Triggers: []Trigger{
			{"result", "R01", "Report an observation result"},
		},
	},
	{
		Code:        "SIU",
		Description: "Scheduling",
		Layout: []string{
			"MSH", "SCH", "PID", "PV1", "PV2", "NTE",
		},
		Triggers: []Trigger{
			{"schedule", "S12", "Book a new appointment"},
			{"reschedule", "S13", "Reschedule an appointment"},
			{"cancel", "S15", "Cancel an appointment"},
			{"no-show", "S26", "Record a patient who did not arrive"},
		},
	},
	{
		Code:        "REF",
		Description: "Patient referral",
		Layout: []string{
			"MSH", "RF1", "PID", "PV1", "DG1", "NTE",
		},
		Triggers: []Trigger{
			{"referral", "I12", "Refer a patient"},
			{"modify", "I13", "Modify a referral"},
			{"cancel", "I14", "Cancel a referral"},
		},
	},
	{
		Code:        "MDM",
		Description: "Document notification",
		Layout: []string{
			"MSH", "EVN", "PID", "PV1", "TXA", "OBX",
		},
		Triggers: []Trigger{
			{"document", "T02", "Notify of a new document"},
			{"status-change", "T04", "Notify of a document status change"},
			{"addendum", "T06", "Notify of a document addendum"},
		},
	},
}

// DefaultMessageType is the type placebo builds when a command names only a
// scenario. Patient administration was the only thing placebo could send
// before message types became selectable, so a bare scenario still means ADT.
const DefaultMessageType = "ADT"

// MessageAndTriggerEvent exposes the registry as the nested map that earlier
// versions of placebo carried, for callers that only need the trigger code.
var MessageAndTriggerEvent = messageAndTriggerEvent()

func messageAndTriggerEvent() map[string]map[string]string {

	all := map[string]map[string]string{}

	for _, definition := range Definitions {
		triggers := map[string]string{}

		for _, trigger := range definition.Triggers {
			triggers[trigger.Name] = trigger.Code
		}

		all[definition.Code] = triggers
	}

	return all
}

// Lookup finds a message type by its code, matched regardless of case so
// 'oru' works as well as 'ORU' on the command line.
func Lookup(code string) (Definition, bool) {

	for _, definition := range Definitions {
		if strings.EqualFold(definition.Code, code) {
			return definition, true
		}
	}

	return Definition{}, false
}

// Trigger finds a scenario within a message type by name.
func (d Definition) Trigger(name string) (Trigger, bool) {

	for _, trigger := range d.Triggers {
		if strings.EqualFold(trigger.Name, name) {
			return trigger, true
		}
	}

	return Trigger{}, false
}

// DefaultTrigger is the scenario a message type generates when none is named.
func (d Definition) DefaultTrigger() Trigger {

	if len(d.Triggers) == 0 {
		return Trigger{}
	}

	return d.Triggers[0]
}

// TriggerNames lists a type's scenarios in the order they are defined, for
// reporting back which ones a command could have asked for.
func (d Definition) TriggerNames() []string {

	names := make([]string, 0, len(d.Triggers))

	for _, trigger := range d.Triggers {
		names = append(names, trigger.Name)
	}

	return names
}

// MessageTypes lists the message type codes placebo can generate.
func MessageTypes() []string {

	codes := make([]string, 0, len(Definitions))

	for _, definition := range Definitions {
		codes = append(codes, definition.Code)
	}

	sort.Strings(codes)

	return codes
}

// Resolve reads a message type and scenario out of command arguments. A
// scenario on its own keeps working without naming a type, since that is how
// every placebo release before message types spelled an ADT event.
func Resolve(args []string) (string, string, error) {

	if len(args) == 0 {
		definition, _ := Lookup(DefaultMessageType)

		return DefaultMessageType, definition.DefaultTrigger().Name, nil
	}

	// A first argument naming a message type means the second, if there is
	// one, names a scenario within it.
	if definition, ok := Lookup(args[0]); ok {

		if len(args) > 2 {
			return "", "", fmt.Errorf("too many arguments for %v: expected a single scenario, got %v", definition.Code, strings.Join(args[1:], " "))
		}

		if len(args) == 1 {
			return definition.Code, definition.DefaultTrigger().Name, nil
		}

		if _, ok := definition.Trigger(args[1]); !ok {
			return "", "", unknownTrigger(definition, args[1])
		}

		return definition.Code, args[1], nil
	}

	if len(args) > 1 {
		return "", "", fmt.Errorf("unknown message type %q\n\nMessage types: %v", args[0], strings.Join(MessageTypes(), ", "))
	}

	definition, _ := Lookup(DefaultMessageType)

	if _, ok := definition.Trigger(args[0]); !ok {
		// A lone argument could have been meant as either a message type or
		// an ADT scenario, so report both sets rather than guessing which.
		return "", "", fmt.Errorf("%v\n\nMessage types: %v", unknownTrigger(definition, args[0]), strings.Join(MessageTypes(), ", "))
	}

	return definition.Code, args[0], nil
}

// Build renders an HL7 message of one type and scenario for a patient.
func Build(patient *random.Patient, e string, t string) (string, error) {

	definition, ok := Lookup(e)
	if !ok {
		return "", fmt.Errorf("unknown message type %q\n\nMessage types: %v", e, strings.Join(MessageTypes(), ", "))
	}

	trigger, ok := definition.Trigger(t)
	if !ok {
		return "", unknownTrigger(definition, t)
	}

	return sugarpill.NewHL7EventMessage(patient, definition.Code, trigger.Code, definition.Layout), nil
}

// unknownTrigger keeps the wording earlier placebo versions used, so anyone
// matching on it still sees what they expect, and adds the scenarios the
// message type actually offers.
func unknownTrigger(d Definition, name string) error {
	return fmt.Errorf("Command %v Not Found\n\n%v scenarios: %v", name, d.Code, strings.Join(d.TriggerNames(), ", "))
}

// Catalog lists every message type with its scenarios, for help output.
func Catalog() string {

	var out strings.Builder

	for i, definition := range Definitions {
		if i > 0 {
			out.WriteString("\n")
		}

		fmt.Fprintf(&out, "  %-4v %v\n", definition.Code, definition.Description)

		for _, trigger := range definition.Triggers {
			fmt.Fprintf(&out, "       %-17v %v^%v, %v\n", trigger.Name, definition.Code, trigger.Code, trigger.Description)
		}
	}

	return strings.TrimRight(out.String(), "\n")
}

// TODO: Intialize Event object from Patient and Event Commands
func NewEvent(p *random.Patient, messageCommand string, eventCommand string) *Event {

	event := Event{}

	eventPatient := event.EventPatient(p)

	message := eventPatient.MessageEventCode(messageCommand)

	messageEventType := message.TriggerEventCode(eventCommand)

	return messageEventType
}

// TODO: MessageEvent Setter (ie ADT or SIU)
func (e *Event) MessageEventCode(code string) *Event {

	e.MessageEvent = code

	return e
}

// TODO: TriggerEventCode Setter (ie A01 or A05)
func (e *Event) TriggerEventCode(code string) *Event {

	definition, ok := Lookup(e.MessageEvent)
	if !ok {
		return e
	}

	trigger, ok := definition.Trigger(code)
	if !ok {
		return e
	}

	e.TriggerEvent = trigger.Code

	return e

}

// TODO: Patient Setter
func (e *Event) EventPatient(p *random.Patient) *Event {

	e.Patient = p

	return e

}

// Obsolete: for template construction
func Builder(scenario []string) []string {

	var messages []string

	patient := random.NewPatient()

	for _, message := range scenario {

		template := TemplateFinder(message)
		mapped := TemplateMapper(patient, template)

		messages = append(messages, mapped)

	}

	return messages

}

// Obsolete: for template construction
func TemplateFinder(s string) []byte {

	switch s {
	case "admit":
		return templates.SimpleHl7Info()
	case "discharge":
		return templates.DischargeHl7Info()
	case "preadmit":
		return templates.PreadmitHl7Info()
	case "referral":
		return templates.ReferralHl7()
	default:
		return templates.SimpleHl7Info()
	}

}

// Obsolete: for template construction
func TemplateMapper(p *random.Patient, temp []byte) string {

	t, err := template.New("hl7").Parse(string(temp))
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	t.Execute(&tpl, p)

	result := tpl.String()

	return result

}
