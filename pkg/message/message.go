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

func MarshallHL7(segs []string, segType interface{}) interface{} {

	if segType == nil {
		fmt.Println("Warning: segType is nil")
		return nil
	}

	val := reflect.ValueOf(segType)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fmt.Printf("%+v\n", segType)
	
	//Testing validations on structs
	if !val.IsValid() {
		panic("Invalid segType: must be a non-nil pointer to a struct")
	}

	if val.Kind() != reflect.Struct {
		panic("segType must be a pointer to a struct")
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
	}

	return segType


}

