package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"placebo/file"
	"placebo/internal/network"
	"placebo/internal/sysCmd"
	"placebo/pkg/event"
	"placebo/pkg/random"
	"placebo/pkg/sugarpill"
)

var Address = "127.0.0.1"
var Port = "9700"

func init() {
	if p := loadEnvPort(); p != "" {
		Port = p
	}
}

func loadEnvPort() string {
	data, err := os.ReadFile(".env")
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "PLACEBO_PORT=") {
			return strings.TrimPrefix(line, "PLACEBO_PORT=")
		}
	}
	return ""
}

func SendHl7Message(f string) error {

	switch f {
	case "":
		return nil
	case "hl7":

		command := os.Args[3:]

		if len(command) == 0 {

			triggerType := "admit"
			messageType := "ADT"
			patient := random.NewPatient()

			hl7 := event.Build(patient, messageType, triggerType)

			openPath := file.CreateInteractiveHl7(hl7)
			sent, err := InteractivePrompt(openPath)
			if err != nil {
				return err
			}

			if sent != "" {
				fmt.Println("HL7 Sent: \n", sent)
			} else {
				return nil
			}

			return nil

		} else if command[0] == "file" {
			path := command[1]

			sent, err := InteractivePrompt(path)
			if err != nil {
				return err
			}

			if sent != "" {
				fmt.Println("HL7 Sent: \n", sent)
			} else {
				return nil
			}

			return nil

		} else if command[0] == "last" {
			lastFile := file.Tempdir + file.IntFile

			sent, err := InteractivePrompt(lastFile)
			if err != nil {
				return err
			}

			if sent != "" {
				fmt.Println("HL7 Sent: \n", sent)
			} else {
				return nil
			}

		} else if command[0] == "sugarpill" {
			//testing
			patient := random.NewPatient()

			er := SugarpillProcess(patient)
			if er != nil {
				return er
			}

			return nil

		} else {
			triggerType := command[0]
			messageType := "ADT"

			err := EventAndMessage(messageType, triggerType)
			if err != nil {
				return err
			}
		}

	}

	return nil

}

// Note: 9700 is the default sending port
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

func EventAndMessage(e string, s string) error {

	patient := random.NewPatient()

	if e != "" {

		evt := event.MessageAndTriggerEvent[e][s]

		if evt != "" {

			evn := event.Build(patient, e, s)

			err := DefaultSend(evn)
			if err != nil {
				return err
			}

			fmt.Printf("HL7 Sent: \n%v\n", evn)
		} else {
			return errors.New("Command Not Found.")
		}
	} else {
		return errors.New("Command Not Found.")
	}

	return nil
}

func InteractivePrompt(filePath string) (string, error) {

	sysCmd.TextEditorOpen(filePath)

	fmt.Println("\nSend The Message Out? [Y/n]")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	if input == "Y" || input == "\r" {

		fileText, err := file.ReadFile(filePath)
		if err != nil {
			return "", err
		}

		err = DefaultSend(fileText)
		if err != nil {
			return "", err
		}

		return fileText, nil

	} else {
		os.Exit(1)
		return "", nil
	}

	return "", nil

}

func PostPrompt(filePath string) (string, error) {

	fmt.Println("\nSend The Message Out? [Y/n]")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	if input == "Y" || input == "\r" {

		fileText, err := file.ReadFile(filePath)
		if err != nil {
			return "", err
		}

		return fileText, nil

	} else {
		os.Exit(1)
		return "", nil
	}

	return "", nil

}

func SugarpillProcess(patient *random.Patient) error {

	//bind created patient to HL7 message in JSON format
	message := sugarpill.NewHL7Message(patient).MessageToJson()

	sp := file.SugarPillInteractive(message)
	sysCmd.TextEditorOpen(sp)

	convert, err := PostPrompt(sp)
	if err != nil {
		return err
	}

	//convert back to HL7 format from the JSON
	jsonMessage := sugarpill.JsonToMessage(convert)

	send := sugarpill.MessageBuilder(jsonMessage)

	DefaultSend(send)
	fmt.Println(send)

	return nil

}
