package main

import (
	"fmt"
	"os"

	"github.com/hl7x/placebo/cmd/placebo/cmd"
)

func main() {

	err := cmd.Execute(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
