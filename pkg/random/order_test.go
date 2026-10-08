package random

import (
	"strconv"
	"testing"
)

func TestNewOrder(t *testing.T) {

	// The result is generated against the test's own reference range, so it
	// takes a few runs to see each of the flags it can produce.
	for i := 0; i < 200; i++ {
		order := NewOrder()

		if order.PlacerNumber == "" || order.FillerNumber == "" {
			t.Fatal("NewOrder() produced an order without identifying numbers")
		}

		if order.Test.Code == "" || order.Test.Name == "" {
			t.Fatalf("NewOrder() produced an order without a test: %v", order.Test)
		}

		value, err := strconv.ParseFloat(order.ResultValue, 64)
		if err != nil {
			t.Fatalf("NewOrder() result %q is not a number", order.ResultValue)
		}

		var expected string

		switch {
		case value < order.Test.Low:
			expected = "L"
		case value > order.Test.High:
			expected = "H"
		default:
			expected = "N"
		}

		if order.AbnormalFlag != expected {
			t.Fatalf("NewOrder() flagged %v against %v-%v as %v, expected %v", value, order.Test.Low, order.Test.High, order.AbnormalFlag, expected)
		}
	}
}

func TestLabTestReferenceRange(t *testing.T) {

	var tests = []struct {
		description string
		input       LabTest
		expected    string
	}{
		{"Whole Numbers", LabTest{Low: 41, High: 53}, "41-53"},
		{"Fractional Numbers", LabTest{Low: 3.5, High: 5.1}, "3.5-5.1"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := tc.input.ReferenceRange()

			if got != tc.expected {
				t.Fatalf("ReferenceRange()=%v, expected %v", got, tc.expected)
			}
		})
	}
}

func TestNewDocument(t *testing.T) {

	document := NewDocument()

	if document.Number == "" {
		t.Fatal("NewDocument() produced a document without a number")
	}

	if document.Type.Code == "" || document.Type.Name == "" {
		t.Fatalf("NewDocument() produced a document without a type: %v", document.Type)
	}

	if document.Status == "" {
		t.Fatal("NewDocument() produced a document without a completion status")
	}
}

// Every table a segment picks from has to have something in it, or the
// segment comes out blank.
func TestOrderTablesArePopulated(t *testing.T) {

	if len(LABTEST) == 0 {
		t.Fatal("LABTEST is empty")
	}

	if len(DOCUMENTTYPE) == 0 {
		t.Fatal("DOCUMENTTYPE is empty")
	}

	for _, test := range LABTEST {
		if test.Low >= test.High {
			t.Fatalf("%v has a reference range of %v-%v", test.Name, test.Low, test.High)
		}

		if test.Units == "" {
			t.Fatalf("%v has no units", test.Name)
		}
	}
}
