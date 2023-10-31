package cmd

import (	
	"text/template"
//	"os"
	"bytes"
	"fmt"
	
//	"placebo/internal/network"
	"placebo/pkg/random"
	"placebo/pkg/templates"
)

func SendHl7Message(f string) error {
	
	switch f {
	case "":
		return nil
	case "hl7":
		patient := random.NewPatient()
		hl7 := Hl7Format(patient)

		fmt.Println(hl7)

		return nil

	}

	return nil

}

func Hl7Format(p *random.Patient) string {

	temp := templates.SimpleHl7Info()

	t, err := template.New("hl7").Parse(string(temp))
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	t.Execute(&tpl, p)

	result := tpl.String()

	return result

}
