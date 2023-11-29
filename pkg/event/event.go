package event

import (

	"bytes"
	"text/template"

	"placebo/pkg/random"
	"placebo/pkg/templates"
)

/* Build HL7 message based on certain medical event for various scenarios */

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
	case "schedule":
		return templates.ScheduledPatientInfo()
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
