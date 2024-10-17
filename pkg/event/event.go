package event

import (
	"bytes"
	"text/template"

	"placebo/pkg/random"
	"placebo/pkg/templates"
)

/* Build HL7 message based on certain medical event for various scenarios */

var MessageTypeCode = map[string]string{
	"AdmitDischargeTransfer": "ADT",
	"SchedulingInformationUnsolicited": "SIU",
}

var EventTypeCode = map[string]string{
	"admit": "A01",
	"transfer": "A02",
	"discharge": "A03",
	"register": "A04",
	"preadmit": "A05",
	"referral": "I12",
}

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
