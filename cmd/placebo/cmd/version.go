package cmd

import "fmt"

// Build information stamped in by goreleaser through -ldflags -X on the main
// package. The defaults are what a plain 'go build' leaves behind.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// SetBuildInfo hands the values main was stamped with to the command package.
func SetBuildInfo(v string, c string, d string) {
	version, commit, date = v, c, d
}

// Version reports the build placebo is running.
func Version(args []string) error {

	if wantsHelp(args) {
		fmt.Println(commandHelp["version"])
		return nil
	}

	if len(args) > 0 {
		return unknownSubcommand("version", args[0])
	}

	fmt.Println(versionString())

	return nil
}

func versionString() string {
	return fmt.Sprintf("placebo %s (commit %s, built %s)", version, commit, date)
}
