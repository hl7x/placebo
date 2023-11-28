package cmd

import (	
	"text/template"
	"bytes"
	"fmt"
	"os"
	
	"placebo/internal/network"
	"placebo/pkg/random"
	"placebo/pkg/templates"
)

func SendHl7Message(f string) error {
	
	switch f {
	case "":
		return nil
	case "hl7":
		
		command := os.Args[3:]

		if len(command) == 0 {
			
			patient := random.NewPatient()
			template := templates.SimpleHl7Info()

			hl7 := Hl7Format(patient, template)

			err := DefaultSend(hl7)
			if err != nil {
				return err
			}

			fmt.Printf("HL7 Sent:\n %v", hl7)

			return nil

		} else if command[0] == "post_admit" {
			//post_admit hl7 event
		} else if command[0] == "appointment" {
			//appointment siu event
		} else if command[0] == "pre_admit" {
			// scheduled pre-admit event
		} else if command[0] == "post_discharge" {
			//post_discharge hl7 event
			err := EventSelector(command[0])
			if err != nil {
				return err
			}
		}

	}

	return nil

}

func Hl7Format(p *random.Patient, temp []byte) string {

//	temp := templates.SimpleHl7Info()

	t, err := template.New("hl7").Parse(string(temp))
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	t.Execute(&tpl, p)

	result := tpl.String()

	return result

}

/* Local dev environment hl7_server to accept messages over 127.0.0.1:9700
this can be used as default
*/
func DefaultSend(templatePatient string) error {

	address := "127.0.0.1"
	port := "9700"

	err := network.SendClient(address, port, templatePatient)
	if err != nil {
		return err
	}

	return nil
}

func EventSelector(s string) error {

	switch s {
	case "":
		return nil
	case "post_discharge":
		patient := random.NewPatient()
		admitTemp := templates.SimpleHl7Info()

		admitHl7 := Hl7Format(patient, admitTemp)
		err := DefaultSend(admitHl7)
		if err != nil {
			return err
		}

		fmt.Println("hl7 sent: \n", admitHl7)
		fmt.Println("\n")

		dischargeTemp := templates.DischargeHl7Info()

		dischargeHl7 := Hl7Format(patient, dischargeTemp)
		
		err = DefaultSend(dischargeHl7)
		if err != nil {
			return nil
		}

		fmt.Println("hl7 sent: \n", dischargeHl7)
	}

	return nil
}
