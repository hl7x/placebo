package message

import "testing"

// A message built for a type leaves out the segments that type does not
// carry, so the builder is handed nil segments as a matter of course.
func TestCreateHL7NilSegment(t *testing.T) {

	type segment struct {
		Field string
	}

	var absent *segment

	var tests = []struct {
		description string
		input       interface{}
		expected    string
	}{
		{"A Nil Segment Renders As Nothing", absent, ""},
		{"An Untyped Nil Renders As Nothing", nil, ""},
		{"A Present Segment Still Renders", &segment{Field: "value"}, "ZZZ|value|"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := CreateHL7(tc.input, "ZZZ")

			if got != tc.expected {
				t.Fatalf("CreateHL7(%v)=%q, expected %q", tc.input, got, tc.expected)
			}
		})
	}
}
