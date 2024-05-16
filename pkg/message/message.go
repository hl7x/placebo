package message

import (
	"fmt"
	"reflect"
	"strings"
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

/*
func ReadHL7(message string) {

}
*/
