package random

import (
	"fmt"
	"math"
	"math/rand"
)

// LabTest is one orderable test and the range a lab reports it against, so a
// generated result lands near a value a real interface would carry.
type LabTest struct {
	Code  string
	Name  string
	Units string
	Low   float64
	High  float64
}

// LABTEST holds LOINC codes for tests from the chemistry and hematology
// panels that order and result messages most often carry.
var LABTEST = []LabTest{
	{"718-7", "HEMOGLOBIN", "g/dL", 13.5, 17.5},
	{"4544-3", "HEMATOCRIT", "%", 41, 53},
	{"6690-2", "LEUKOCYTES", "10*3/uL", 4.5, 11},
	{"777-3", "PLATELETS", "10*3/uL", 150, 400},
	{"2345-7", "GLUCOSE", "mg/dL", 70, 99},
	{"2160-0", "CREATININE", "mg/dL", 0.7, 1.3},
	{"3094-0", "UREA NITROGEN", "mg/dL", 7, 20},
	{"2951-2", "SODIUM", "mmol/L", 136, 145},
	{"2823-3", "POTASSIUM", "mmol/L", 3.5, 5.1},
	{"2075-0", "CHLORIDE", "mmol/L", 98, 107},
}

// ReferenceRange formats the range the way OBX-7 carries it.
func (l LabTest) ReferenceRange() string {
	return fmt.Sprintf("%v-%v", l.Low, l.High)
}

// DocumentType is a kind of transcribed clinical note, paired with the name a
// document notification would title it with.
type DocumentType struct {
	Code string
	Name string
}

var DOCUMENTTYPE = []DocumentType{
	{"DS", "DISCHARGE SUMMARY"},
	{"OP", "OPERATIVE NOTE"},
	{"HP", "HISTORY AND PHYSICAL"},
	{"CN", "CONSULTATION NOTE"},
	{"PN", "PROGRESS NOTE"},
	{"ED", "EMERGENCY DEPARTMENT NOTE"},
}

// Order is a lab order and the result that comes back for it. It fills the
// ORC, OBR and OBX segments that order and result messages are built around.
type Order struct {
	PlacerNumber string
	FillerNumber string
	Test         LabTest
	ResultValue  string
	AbnormalFlag string
	ResultStatus string
}

func NewOrder() *Order {

	o := &Order{}

	return o.OrderNumbers().OrderTest().OrderResult()
}

func (o *Order) OrderNumbers() *Order {

	o.PlacerNumber = fmt.Sprint(rand.Intn(1000000000))
	o.FillerNumber = fmt.Sprint(rand.Intn(1000000000))

	return o
}

func (o *Order) OrderTest() *Order {

	o.Test = LABTEST[rand.Intn(len(LABTEST))]

	return o
}

func (o *Order) OrderResult() *Order {

	spread := o.Test.High - o.Test.Low

	// Reach past both ends of the reference range so some results come back
	// flagged. A result set that is always normal is the less useful one to
	// test an interface against.
	value := o.Test.Low - spread*0.25 + rand.Float64()*spread*1.5

	// Round before flagging, not after. The flag has to describe the value
	// the message actually carries, or a result reads as abnormal next to a
	// number inside the range it is flagged against.
	value = math.Round(value*10) / 10

	o.ResultValue = fmt.Sprintf("%.1f", value)

	switch {
	case value < o.Test.Low:
		o.AbnormalFlag = "L"
	case value > o.Test.High:
		o.AbnormalFlag = "H"
	default:
		o.AbnormalFlag = "N"
	}

	// Results placebo sends are complete ones; there is no partial state to
	// report on a value it made up in full.
	o.ResultStatus = "F"

	return o
}

// Document is a transcribed clinical note. It fills the TXA segment that a
// document notification is built around.
type Document struct {
	Number string
	Type   DocumentType
	Status string
}

func NewDocument() *Document {

	d := &Document{}

	return d.DocumentNumber().Kind().CompletionStatus()
}

func (d *Document) DocumentNumber() *Document {

	d.Number = fmt.Sprint(rand.Intn(1000000000))

	return d
}

func (d *Document) Kind() *Document {

	d.Type = DOCUMENTTYPE[rand.Intn(len(DOCUMENTTYPE))]

	return d
}

func (d *Document) CompletionStatus() *Document {

	d.Status = "AU"

	return d
}
