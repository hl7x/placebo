package sysCmd

import (
	"os/exec"
//	"fmt"
	"bytes"
	"log"
	"os"

)

func TextEditorOpen(path string) {

	var stderr bytes.Buffer

	cmd := exec.Command("vi", path)
	cmd.Stderr = &stderr

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("vi exited with error: %v, stderr: %s", err, stderr.String())
	}
/*
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("vi exited with error: %v, stderr: %s", err, stderr.String())	
	}
*/
}
