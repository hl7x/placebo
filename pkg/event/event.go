package event

import (
	"bytes"
	"text/template"

	"placebo/pkg/random"
	"placebo/pkg/templates"
)

/* Build HL7 message based on certain medical event for various scenarios */


// Template approach is depricated and not supported anymore.

type Event struct {
	MessageCode 	string
	EventCode	string
	Patient		*random.Patient
}

var EventTypeAndMessageCode = map[string]map[string]string{
	"ADT": {
		"admit": "A01",
		"transfer": "A02",
		"discharge": "A03",
		"register": "A04",
		"pre-admit": "A05",
	},
	"SIU": {
		"referral": "I12",
	},
}

/* TODO: Starting place for event message handling
func Build(eventCode string, eventScenarios []string) []string {

	var messages []string

	if len(eventScanarios) > 1 {
		for _, message := range eventScenarios {
			event := NewEvent(patient, eventCode, message)

			messages
	} else {
		//there
	}

}
*/

/* TODO: Transform  Event to Message String
func EventMessage(e *Event) string {


}
*/

func NewEvent(p *random.Patient, messageCommand string, eventCommand string) *Event {

	event := Event{}

	eventPatient := event.EventPatient(p)

	message := eventPatient.EventMessageCode(messageCommand)

	messageEventType := message.TypeEventCode(eventCommand)

	return messageEventType
}

// MessageCode Setter
func (e *Event) EventMessageCode(code string) *Event {

	e.MessageCode = code

	return e
}

// EventCode Setter
func (e *Event) TypeEventCode(code string) *Event {

	eventCode := EventTypeAndMessageCode[e.MessageCode][code]

	e.EventCode = eventCode

	return e

}

// Patient Setter
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

	return nil

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
