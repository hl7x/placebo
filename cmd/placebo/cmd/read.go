package cmd

import (
	"os"
	"fmt"
//	"io"

	"placebo/pkg/sugarpill"
	"placebo/file"
//	"placebo/pkg/message"

)

func ReadHl7Message(f string) error {

	switch f {
	case "":
		return nil
	case "sugarpill":

		filePath := os.Args[3:]

		if len(filePath) == 0 {
			return nil
		}

		hl7Content, err := file.ReadFile(filePath[0])
		if err != nil {
			return err
		}

		message := sugarpill.ReadHL7(hl7Content)

		fmt.Println(message)

		return nil
	default:
		return nil
	}

	return nil
}


