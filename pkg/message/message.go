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
                if nestedValue != "" {  // Only append non-empty values
                    nestedFields = append(nestedFields, nestedValue)
                }
            }
            // Only append nested fields if there's something to append to avoid stray carets.
            if len(nestedFields) > 0 {
                fields = append(fields, strings.Join(nestedFields, "^"))
            } else {
                fields = append(fields, "") // Ensure the segment structure is maintained even if empty.
            }
        } else {
            fieldValue := fmt.Sprintf("%v", field.Interface())
            fields = append(fields, fieldValue)  // Directly append the field value, empty or not.
        }
    }

    return strings.Join(fields, "|")
}

