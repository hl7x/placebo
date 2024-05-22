package message

import (
	"fmt"
	"reflect"
	"strings"

//	"placebo/file"
//	"placebo/pkg/sugarpill"
)

func CreateHL7(v interface{}, segment string) string {
    val := reflect.ValueOf(v)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    fields := []string{segment}
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        if field.Kind() == reflect.Ptr {
            field = field.Elem()
        }

        // Handle nested structs by concatenating their fields with '^' and ensuring no trailing caret.
        if field.Kind() == reflect.Struct {
            nestedFields := make([]string, 0, field.NumField())
            for j := 0; j < field.NumField(); j++ {
                nestedField := field.Field(j)
                if nestedField.Kind() == reflect.Ptr {
                    nestedField = nestedField.Elem()
                }
                nestedValue := fmt.Sprintf("%v", nestedField.Interface())
                nestedFields = append(nestedFields, nestedValue) // Always append, even if empty
            }
            fields = append(fields, strings.Join(nestedFields, "^"))
        } else {
            fieldValue := fmt.Sprintf("%v", field.Interface())
            fields = append(fields, fieldValue) // Directly append the field value, empty or not.
        }
    }

    //the trailing pipe is necessary for how hl7 segments are constructed here
    allPipes := strings.Join(fields, "|") + "|"

    return allPipes
}



// TODO: For Future PR to Setup Reading an HL7 File Through Sugarpill to Easily Read HL7
/*
func ReadHL7(lines string) string {

	splitLine := strings.Split(lines, "\n")
	
	var segType interface{}
	for line := range splitLine {
		segType, value = HL7Line(line)
		newMessage.SegmentAssigner(segType, value)

	}

	return newMessage
}
*/

/*
func HL7Line(s string) string {
	
	segments := strings.Split(s, "|")

	segmentType := sugarpill.SegmentFinder(segments[0])

	mapped := MarshallHL7(segments, segmentType)

	messageFmt := json.MashalIndent(mapped, "", " ")

	return string(messageFmt)

}
*/

func MarshallHL7(segs []string, segType interface{}) interface{} {

	val := reflect.ValueOf(segType)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	

	segments := segs[1:]
	for i := 0; i < val.NumField() && i < len(segments); i++ {
		structFieldValue := val.Field(i)
		if structFieldValue.CanSet() {
			switch structFieldValue.Kind() {
			case reflect.String:
				structFieldValue.SetString(segments[i])
			}
		}
//		structFieldValue.Set(segments[i])
	}

	return segType


}

