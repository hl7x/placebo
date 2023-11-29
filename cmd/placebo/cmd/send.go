package cmd

import (	
//	"text/template"
//	"bytes"
	"fmt"
	"os"
	
	"placebo/internal/network"
//	"placebo/pkg/random"
//	"placebo/pkg/templates"
	"placebo/pkg/event"
)

var Address = "127.0.0.1"
var Port = "9700"

func SendHl7Message(f string) error {
	
	switch f {
	case "":
		return nil
	case "hl7":
		
		command := os.Args[3:]

		if len(command) == 0 {
			err := EventSelector("")
			if err != nil {
				return err
			}

		} else {
			err := EventSelector(command[0])
			if err != nil {
				return err
			}
		}

	}

	return nil

}

/* Local dev environment hl7_server to accept messages over 127.0.0.1:9700
this can be used as default
*/
func DefaultSend(templatePatient string) error {


	err := network.SendClient(Address, Port, templatePatient)
	if err != nil {
		return err
	}

	return nil
}

func MultiSender(mes []string) error {

	for _, message := range mes {
		err := DefaultSend(message)
		if err != nil {
			return err
		}

		fmt.Printf("HL7 Sent: \n%v\n", message)
	}

	return nil
}

func EventSelector(s string) error {
	
	dischargeEvent  := []string{"admit", "discharge"}
	admitEvent	:= []string{"admit"}
//	scheduleEvent	:= []string{"schedule"}

	switch s {
	case "post_discharge":
		messages := event.Builder(dischargeEvent)
		sent := MultiSender(messages)
		
		return sent
	default:
		message := event.Builder(admitEvent)
		sent := MultiSender(message)

		return sent
	}

	return nil
}

