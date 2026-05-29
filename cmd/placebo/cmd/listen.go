package cmd

import (
	"flag"
	"fmt"
	"os"

	"placebo/internal/network"
)

func ListenHl7Message(f string) error {

	switch f {
	case "hl7":

		command := flag.Args()

		if len(command) == 0 {

			err := Listener(":" + Port)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		} else {
			return nil
		}

		return nil

	default:
		return nil
	}

	return nil

}

func Listener(p string) error {

	err := network.ListenClient(p)
	if err != nil {
		return err
	}

	return nil
}
