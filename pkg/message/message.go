package message

import (
	"fmt"
	"reflect"
	"strings"
)

func HL7(v interface{}, segment string) string {
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
            nestedFields := []string{}
            for j := 0; j < field.NumField(); j++ {
                nestedField := field.Field(j)
                if nestedField.Kind() == reflect.Ptr {
                    nestedField = nestedField.Elem()
                }
                nestedValue := fmt.Sprintf("%v", nestedField.Interface())
                if nestedValue == "" {
                    nestedValue = "\"\""  // Represents an empty value in HL7
                }
                nestedFields = append(nestedFields, nestedValue)
            }
            fields = append(fields, strings.Join(nestedFields, "^"))
        } else {
            fieldValue := fmt.Sprintf("%v", field.Interface())
            if fieldValue == "" {
                fieldValue = "\"\""
            }
            fields = append(fields, fieldValue)
        }
    }

    return strings.Join(fields, "|")
}
