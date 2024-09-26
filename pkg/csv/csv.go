package csv

import (
	"reflect"
	"strings"
	"errors"
	"fmt"
	
	"placebo/pkg/random"
)

// Build a CSV file from patient data
func Builder(p *random.Collection, d string) (string, error) {

	err :=  delimiterValidation(d)
	if err != nil {
		return "", err
	}

	data, err := DataProcess(p, d)
	if err != nil {
		return "", err
	}


	return data, nil

}

// Abstract the patient data into a slice of collected strings through iterrating
func DataProcess(p *random.Collection, d string) (string, error) {

	leadingPatient := p.Patients[0]
	allPatients := p.Patients

	header := FieldExtraction(leadingPatient)

	headerFmt := CsvFormatter(header, d)

	var collate []string
	for _, v := range allPatients {
		body := ValueExtraction(v)
		bodyFmt := CsvFormatter(body, d)
		collate = append(collate, bodyFmt)
	}


	result := headerFmt + "\n" + strings.Join(collate, "\n")

	return result, nil
}


// TODO: Function to return values from a provided struct
func ValueExtraction(v interface{}) []string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var fields []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Check if the field is valid before trying to use it.
		if !field.IsValid() {
			continue
		}

		// Check for pointer types and dereference if needed
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				// Handle nil pointer case
				fields = append(fields, "")
				continue
			}
			field = field.Elem()
		}

		// Handle nested structs
		if field.Kind() == reflect.Struct {
			nestedFields := make([]string, 0, field.NumField())
			for j := 0; j < field.NumField(); j++ {
				nestedField := field.Field(j)

				// Again, check for valid field and nil pointers
				if !nestedField.IsValid() {
					nestedFields = append(nestedFields, "")
					continue
				}
				if nestedField.Kind() == reflect.Ptr {
					if nestedField.IsNil() {
						nestedFields = append(nestedFields, "")
						continue
					}
					nestedField = nestedField.Elem()
				}

				// Convert field to string
				nestedValue := fmt.Sprintf("%v", nestedField.Interface())
				nestedFields = append(nestedFields, nestedValue)
			}
			fields = append(fields, strings.Join(nestedFields, ","))
		} else {
			// Handle non-struct fields
			fieldValue := fmt.Sprintf("%v", field.Interface())
			fields = append(fields, fieldValue)
		}
	}

	return fields
}

// Return the Fields Associated with the Provided Struct
func FieldExtraction(v interface{}) []string {
	var fields []string
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type

		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		if fieldType.Kind() == reflect.Struct {
			nestedFields := FieldExtraction(reflect.New(fieldType).Elem().Interface())
			for _, nestedField := range nestedFields {
				fields = append(fields, nestedField)
			}
		} else {
			fields = append(fields, field.Name)
		}
	}


	return fields
}

// Check the provided delimiter. Private.
func delimiterValidation(d string) error {

	// switch to make any future editions easier to add.
	switch d {
	case "|":
		return nil
	case ",":
		return nil
	default:
		err := errors.New("Delimiter Format Not Supported.")
		return err
	}

	return nil

}

func CsvFormatter(slice []string, d string) string {
	
	result := strings.Join(slice, d)

	return result
}
