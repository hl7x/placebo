package cmd

import (
	"fmt"
	"os"
	
	"placebo/internal/network"
)

//var Address = "127.0.0.1"
var listenPort = ":9700"

func ListenHl7Message(f string) error {

	switch f {
	case "hl7":
		
		command := os.Args[3:]

		if len(command) == 0 {
			
			//default listening on port 9700
			err := Listener(listenPort)
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
