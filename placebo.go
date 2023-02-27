package main

import (
	"flag"
//	"os"
	
	"placebo/cmd"
)

var randomPatientFile = flag.String("file", "", "Creates a csv file with a fake patient at /tmp/.\n\t'init' command creates a random file quickly.\n\t\tadding numbers to the 'init' command will produce multiple fake patients i.e. 'placebo --file init 4'")

func main() {
//	flag := os.Args[2]
//	cmd := os.Args[3]
	flag.Parse()
	
	cmd.PatientFile(*randomPatientFile)
//	cmd.PatientFile(flag, cmd)

}
