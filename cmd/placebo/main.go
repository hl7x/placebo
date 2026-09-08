package main

import (
	"fmt"
	"os"

	"github.com/hl7x/placebo/cmd/placebo/cmd"
)

// Stamped in at release time by goreleaser through -ldflags -X.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	cmd.SetBuildInfo(version, commit, date)

	err := cmd.Execute(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
