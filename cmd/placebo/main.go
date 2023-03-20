package main

import (
	"flag"
	"fmt"

	"placebo/cmd/placebo/cmd"
)

var randomFile = flag.String("file", "", "Creates a csv file with a fake patient at /tmp/.\n\t'csv' command creates a random csv file.\n\t\tadding numbers to the 'csv' command will produce multiple fake patients i.e. 'placebo --file csv 4'")

func main() {
	flag.Parse()

	err := cmd.File(*randomFile)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}
