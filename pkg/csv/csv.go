package csv

import (
	"reflect"
	
	"placebo/pkg/random"
)

// Build a CSV file from patient data
func Builder(p *random.Collection, d string) (string, error) {

	err :=  delimiterValidation(d)
	if err != nil {
		return "", err
	}

	data, err := DataProcess(p)
	if err != nil {
		return "", err
	}

	attached := strings.Join(data, d)

	return attached, nil

}

// Abstract the patient data into a slice of collected strings through iterrating
func DataProcess(p *random.Collection) ([]string, error) {

	leadingPatient := p[0]

	header := FieldExtraction(leadingPatient)

	var collate []string
	for _, v := range p {
		body := ValueExtraction(v)
		collate = append(collate, body)
	}

	header = append(header, collate)

	return header, nil
}

// TODO: Function to return values from a provided struct
func ValueExtraction(v interface{}) []string {
	//here
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

	fields = append(fields, "\n")

	return fields
}

// Check the provided delimiter. Private.
func delimiterValidation(d string) error {
	
	switch d {
	case "|":
		nil
	case ",":
		nil
	default:
		errors.New("Delimiter Format Not Supported.")
	}

	return nil

}
