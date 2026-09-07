package cmd

import (
	"fmt"

	"github.com/hl7x/placebo/internal/network"
)

func ListenHl7Message(args []string) error {

	if wantsHelp(args) {
		fmt.Println(commandHelp["listen"])
		return nil
	}

	if len(args) == 0 {
		return missingSubcommand("listen")
	}

	if args[0] != "hl7" {
		return unknownSubcommand("listen", args[0])
	}

	if len(args) > 1 {
		return fmt.Errorf("unexpected argument %q\n\n%s", args[1], commandHelp["listen"])
	}

	return Listener(":" + Port)
}

func Listener(p string) error {

	err := network.ListenClient(p)
	if err != nil {
		return err
	}

	return nil
}
