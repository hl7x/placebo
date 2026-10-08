package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hl7x/placebo/pkg/event"
)

const Usage = `placebo generates fake patient data for testing healthcare applications.

Usage:
  placebo <command> [subcommand] [arguments] [options]

Commands:
  file     Create a file of fake patient data
  send     Send an HL7 message built from fake patient data
  listen   Receive and print HL7 messages
  read     Break an HL7 message down into a readable structure
  version  Show the version of placebo you are running
  help     Show help for a command

Options:
  --port <port>   Port used for sending and listening (default 9700).
                  Can also be set with the PLACEBO_PORT environment variable.
  --version       Show the version of placebo you are running.

Run 'placebo help <command>' for details on a command.`

var commandHelp = map[string]string{
	"file": `Create a file of fake patient data at /tmp/.

Usage:
  placebo file csv [number of patients]
  placebo file hl7 [message type] [scenario]

Subcommands:
  csv   Create a CSV file. Pass a number to generate more than one patient.
  hl7   Create an HL7 message file for a single patient. Defaults to an
        ADT^A01 admit. Name a message type to build another kind of message,
        and a scenario to pick the trigger event within it.
  hl7 types
        List every message type and scenario placebo can build.

Examples:
  placebo file csv
  placebo file csv 4
  placebo file hl7
  placebo file hl7 discharge
  placebo file hl7 oru
  placebo file hl7 siu reschedule`,

	"send": `Send an HL7 message with automatically generated fake patient data.
Messages go to 127.0.0.1:9700 unless --port says otherwise.

Usage:
  placebo send hl7 [message type] [scenario]
  placebo send hl7 [subcommand]

With nothing further, an ADT^A01 admit is built and opened in your text
editor so you can edit it before it is sent.

A message type builds a different kind of message, and a scenario picks the
trigger event within it. A scenario on its own is read as an ADT event, so
'placebo send hl7 discharge' still means ADT^A03. Each message type carries
only the segments that type is defined around, so an order message is not
sent with an insurance segment stapled to it.

Subcommands:
  file <path>      Edit an existing HL7 file and send it when you are done.
  last             Reopen the last sent HL7 message in an interactive prompt.
  sugarpill        Build a message through an easy-to-read interactive prompt.
  types            List every message type and scenario placebo can build.

Message types and scenarios:
{{types}}

Examples:
  placebo send hl7
  placebo send hl7 discharge
  placebo send hl7 oru
  placebo send hl7 siu reschedule
  placebo send hl7 file /tmp/import_hl7636272.txt
  placebo send hl7 --port 8500`,

	"listen": `Receive and print an HL7 message sent to the designated port.
Port 9700 is used unless --port says otherwise.

Usage:
  placebo listen hl7

Examples:
  placebo listen hl7
  placebo listen hl7 --port 8500`,

	"read": `Read an HL7 message to assist with analyzing segments.

Usage:
  placebo read sugarpill <path/to/hl7_file.txt>

Subcommands:
  sugarpill   Break the message down into a much more readable structure.

Examples:
  placebo read sugarpill hl7_message.txt`,

	"version": `Show the version of placebo you are running, along with the commit it
was built from and the date it was built.

Usage:
  placebo version

Examples:
  placebo version
  placebo --version`,
}

// The message type catalog is generated from the registry rather than written
// out here, so a type added to the registry shows up in help on its own.
func init() {
	for command, help := range commandHelp {
		commandHelp[command] = strings.ReplaceAll(help, "{{types}}", event.Catalog())
	}
}

// The commands below used to be spelled as flags. Point anyone still using
// the old spelling at the command that replaced it.
var retiredFlags = map[string]string{
	"file":   "file",
	"send":   "send",
	"listen": "listen",
	"read":   "read",
}

// Execute routes args (everything after the program name) to a command.
func Execute(args []string) error {

	args, err := extractOptions(args)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Println(Usage)
		return nil
	}

	command, rest := args[0], args[1:]

	if strings.HasPrefix(command, "-") {
		if isHelp(command) {
			fmt.Println(Usage)
			return nil
		}

		if isVersion(command) {
			fmt.Println(versionString())
			return nil
		}

		return unknownFlag(command)
	}

	switch command {
	case "file":
		return File(rest)
	case "send":
		return SendHl7Message(rest)
	case "listen":
		return ListenHl7Message(rest)
	case "read":
		return ReadHl7Message(rest)
	case "version":
		return Version(rest)
	case "help":
		return Help(rest)
	default:
		return fmt.Errorf("unknown command %q\n\n%s", command, Usage)
	}
}

func Help(args []string) error {

	if len(args) == 0 {
		fmt.Println(Usage)
		return nil
	}

	help, ok := commandHelp[args[0]]
	if !ok {
		return fmt.Errorf("unknown command %q\n\n%s", args[0], Usage)
	}

	fmt.Println(help)

	return nil
}

// extractOptions pulls options out of args from any position, so they can be
// given before the command or after the subcommand, and returns what is left.
func extractOptions(args []string) ([]string, error) {

	rest := make([]string, 0, len(args))

	for i := 0; i < len(args); i++ {
		name, value, hasValue := strings.Cut(args[i], "=")

		if name != "--port" && name != "-port" {
			rest = append(rest, args[i])
			continue
		}

		if !hasValue {
			if i+1 == len(args) {
				return nil, fmt.Errorf("flag %s needs a port number", name)
			}

			i++
			value = args[i]
		}

		port, err := strconv.Atoi(value)
		if err != nil || port < 1 || port > 65535 {
			return nil, fmt.Errorf("invalid port %q: expected a number between 1 and 65535", value)
		}

		Port = value
	}

	return rest, nil
}

func unknownFlag(arg string) error {

	name, _, _ := strings.Cut(arg, "=")

	if command, ok := retiredFlags[strings.TrimLeft(name, "-")]; ok {
		return fmt.Errorf("unknown flag %s: %q is a command now, not a flag\n\n\tplacebo %s ...\n\n%s", name, command, command, Usage)
	}

	return fmt.Errorf("unknown flag %s\n\n%s", name, Usage)
}

func isHelp(arg string) bool {
	return arg == "help" || arg == "-h" || arg == "--help"
}

func isVersion(arg string) bool {
	return arg == "-v" || arg == "--version"
}

// wantsHelp reports whether a command was asked for its own help instead of
// being asked to do work, as in 'placebo send help'.
func wantsHelp(args []string) bool {
	return len(args) > 0 && isHelp(args[0])
}

func missingSubcommand(command string) error {
	return fmt.Errorf("'placebo %s' needs a subcommand\n\n%s", command, commandHelp[command])
}

func unknownSubcommand(command string, sub string) error {
	return fmt.Errorf("unknown subcommand %q for 'placebo %s'\n\n%s", sub, command, commandHelp[command])
}
