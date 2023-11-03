package main

import (
	"flag"
	"fmt"

	"placebo/cmd/placebo/cmd"
)

var randomFile = flag.String("file", "", "Creates a csv file with a fake patient at /tmp/.\n\t'csv' command creates a random csv file.\n\t\tadding numbers to the 'csv' command will produce multiple fake patients i.e. 'placebo --file csv 4'")

var sendHl7 = flag.String("send", "", "Send a hl7 message with fake patient data.\n\t'hl7' creates an hl7 message i.e placebo --send hl7\n\t\thl7s get sent to 127.0.0.1:9700 by default")

func main() {
	flag.Parse()

	err := cmd.File(*randomFile)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	err = cmd.SendHl7Message(*sendHl7)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
