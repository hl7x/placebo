package main

import (
	"flag"
	"fmt"

	"placebo/cmd/placebo/cmd"
)

var randomFile = flag.String("file", "", "Creates a csv file with a fake patient at /tmp/.\n\t'csv' command creates a random csv file.\n\t\tadding numbers to the 'csv' command will produce multiple fake patients i.e. 'placebo --file csv 4'")

var sendHl7 = flag.String("send", "", "Send a hl7 message with automatically generated fake patient data.\n\t'hl7' creates and sends an hl7 message i.e 'placebo --send hl7' [hl7s get sent to 127.0.0.1:9700 by default]\n\tSub commands:\n\t\t'post_admit' -  ADT^A01 event that admits a patient\n\t\t'post_discharge' - ADT^A01 & ADT^A03 events that admit and then discharge the patient\n\t\t'pre_admit' - ADT^A05 event that establishes preadmit info\n\ti.e. 'placebo --send hl7 post_discharge'")

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
