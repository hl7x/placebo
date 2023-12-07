package cmd

import (	
	"fmt"
	"os"
	"bufio"
	"strings"
	
	"placebo/internal/network"
	"placebo/pkg/event"
	"placebo/file"
	"placebo/internal/sysCmd"
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
			
			eventHolder := []string{"admit"}
			hl7 := event.Builder(eventHolder)

			openPath := file.CreateInteractiveHl7(hl7[0])
			sysCmd.TextEditorOpen(openPath)

			fmt.Println("\nSend The Message Out? [Y/n]")
			reader := bufio.NewReader(os.Stdin)

			input, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			input = strings.TrimSpace(input)

			if input == "Y" {
				messageText := file.ReadInteractiveHl7(openPath)
				err = DefaultSend(messageText)
				if err != nil {
					return err
				}
			} else {
				os.Exit(1)
				return nil
			}


			return nil
		} else if command[0] == "file" {
			path := command[1]

			sysCmd.TextEditorOpen(path)

			fmt.Println("\nSend The Message Out? [Y/n]")
			reader := bufio.NewReader(os.Stdin)

			input, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			input = strings.TrimSpace(input)

			if input == "Y" {
				messageText := file.ReadInteractiveHl7(path)
				err = DefaultSend(messageText)
				if err != nil {
					return nil
				}
			} else {
				os.Exit(1)
				return nil
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
	preadmitEvent	:= []string{"preadmit"}

	switch s {
	case "post_discharge":
		messages := event.Builder(dischargeEvent)
		sent := MultiSender(messages)
		
		return sent
	case "post_admit":
		message := event.Builder(admitEvent)
		sent := MultiSender(message)

		return sent
	case "pre_admit":
		message := event.Builder(preadmitEvent)
		sent := MultiSender(message)

		return sent
	default:
	/*:
		message := event.Builder(admitEvent)
		sent := MultiSender(message)

		return sent
	*/
		fmt.Println("Command Not Found.")
		return nil
	}

	return nil
}

