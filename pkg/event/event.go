package event

import (
	"bytes"
	"text/template"

	"placebo/pkg/random"
	"placebo/pkg/templates"
	"placebo/pkg/sugarpill"
)

/* Build HL7 message based on certain medical event for various scenarios */


// Template approach is depricated and not supported anymore.

type Event struct {
	MessageEvent 	string
	TriggerEvent	string
	Patient		*random.Patient
}

var MessageAndTriggerEvent = map[string]map[string]string{
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

// Starting place for event message handling
func Build(patient *random.Patient, e string, t string) string {

	mess := MessageAndTriggerEvent[e][t]
	message := sugarpill.NewHL7EventMessage(patient, e, mess)

	return message
}


/* TODO: Transform  Event to Message String
func EventMessage(e *Event) string {


}
*/

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

// TODO: TriggerEvent Setter (ie A01 or A05)
func (e *Event) TriggerEventCode(code string) *Event {

	eventCode := MessageAndTriggerEvent[e.MessageEvent][code]

	e.TriggerEvent = eventCode

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
