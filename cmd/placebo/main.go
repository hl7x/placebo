package main

import (
	"flag"
	"fmt"

	"placebo/cmd/placebo/cmd"
)

var randomFile = flag.String("file", "", "Creates a csv file with a fake patient at /tmp/.\n\t'csv' command creates a random csv file.\n\t\tadding numbers to the 'csv' command will produce multiple fake patients i.e. 'placebo --file csv 4'")

var sendHl7 = flag.String("send", "", "Send a hl7 message with automatically generated fake patient data. \n\t'hl7' creates and sends an hl7 message. You will be prompted with a text editor to apply changes if you want.  i.e 'placebo --send hl7' [hl7s get sent to 127.0.0.1:9700 by default]\n\tSub command:\n\t\t'file' - Edit a hl7 file and send the file when you complete editing i.e. placebo --send hl7 /path/to/file/import_hl7636272.txt\n\t\t'last' - Interactive prompt to reopen the last sent hl7 message i.e. placebo --send hl7 last\n\n\t\tSub commands for preset scenarios:\n\t\t'post_admit' -  ADT^A01 event that admits a patient\n\t\t'post_discharge' - ADT^A01 & ADT^A03 events that admit and then discharge the patient\n\t\t'pre_admit' - ADT^A05 event that establishes preadmit info\n\t\t'referral' - REF^I12 event that informs a referral for a patient\n\t\t\t.e. 'placebo --send hl7 post_discharge'")

var listenHl7 = flag.String("listen", "", "Recieve and print an hl7 message sent to the designated port. (By default port 9700 is used).\n\n\ti.e. 'placebo --listen hl7'")

var readHl7 = flag.String("read", "", "Read HL7 message to assist with analyzing segments.\n\tSubcommand: sugarpill - use this to breakdown the message in a much more readible way")

var port = flag.String("port", "", "Override the port used for sending and listening (default 9700). Can also be set via PLACEBO_PORT in a .env file.")

func main() {
	flag.Parse()

	if *port != "" {
		cmd.Port = *port
	}

	err := cmd.File(*randomFile)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	err = cmd.SendHl7Message(*sendHl7)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	err = cmd.ListenHl7Message(*listenHl7)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	err = cmd.ReadHl7Message(*readHl7)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
