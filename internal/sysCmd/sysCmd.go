package sysCmd

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func TextEditorOpen(path string) {

	var stderr bytes.Buffer

	cmd := exec.Command("vi", path)
	cmd.Stderr = &stderr

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("vi exited with error: %v, stderr: %s", err, stderr.String())
	}

}
