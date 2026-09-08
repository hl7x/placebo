package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hl7x/placebo/file"
	"github.com/hl7x/placebo/internal/network"
	"github.com/hl7x/placebo/internal/sysCmd"
	"github.com/hl7x/placebo/pkg/event"
	"github.com/hl7x/placebo/pkg/message"
	"github.com/hl7x/placebo/pkg/random"
	"github.com/hl7x/placebo/pkg/sugarpill"
)

var Address = "127.0.0.1"
var Port = "9700"

func init() {
	applyEnvPort()
}

func applyEnvPort() {
	if p := os.Getenv("PLACEBO_PORT"); p != "" {
		Port = p
	}
}

func SendHl7Message(args []string) error {

	if wantsHelp(args) {
		fmt.Println(commandHelp["send"])
		return nil
	}

	if len(args) == 0 {
		return missingSubcommand("send")
	}

	if args[0] != "hl7" {
		return unknownSubcommand("send", args[0])
	}

	command := args[1:]

	if len(command) == 0 {

		triggerType := "admit"
		messageType := "ADT"
		patient := random.NewPatient()

		hl7 := event.Build(patient, messageType, triggerType)

		openPath := file.CreateInteractiveHl7(hl7)

		return sendFromPrompt(openPath)
	}

	switch command[0] {
	case "file":

		if len(command) < 2 {
			return errors.New("'placebo send hl7 file' needs a file\n\n\tplacebo send hl7 file <path/to/hl7_file.txt>")
		}

		path := command[1]

		_, err := os.Stat(path)
		if err != nil {
			return err
		}

		return sendFromPrompt(path)

	case "last":
		return sendFromPrompt(file.Tempdir + file.IntFile)

	case "sugarpill":
		return SugarpillProcess(random.NewPatient())

	default:
		return EventAndMessage("ADT", command[0])
	}
}

// The message and any ACK are printed by DefaultSend, once the prompt has
// actually sent it.
func sendFromPrompt(filePath string) error {

	_, err := InteractivePrompt(filePath)

	return err
}

// Note: 9700 is the default sending port
func DefaultSend(templatePatient string) error {

	ack, err := network.SendClient(Address, Port, templatePatient)
	if err != nil {
		return err
	}

	fmt.Printf("HL7 Sent: \n%v\n", message.ForDisplay(templatePatient))

	if ack == "" {
		fmt.Println("No ACK Received. The receiver took the message without acknowledging it.")
		return nil
	}

	fmt.Printf("ACK Received: \n%v\n", message.ForDisplay(ack))

	return nil
}

func MultiSender(mes []string) error {

	for _, msg := range mes {
		err := DefaultSend(msg)
		if err != nil {
			return err
		}
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
		} else {
			return errors.New(fmt.Sprintf("Command %v Not Found", s))
		}
	} else {
		return errors.New(fmt.Sprintf("Command %v Not Found", s))
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
		os.Exit(0)
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
		os.Exit(0)
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
