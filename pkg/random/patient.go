package random

import (
	"fmt"
	"math/rand"
	"placebo/internal/tools"
	"strconv"
	"strings"
	"time"
)

type Collection struct {
	Patients []*Patient
}

type Patient struct {
	FirstName      string
	LastName       string
	MRN            string
	VisitId    	int
	Phone          string
	DOB            string
	Hl7DOB         string
	PatientAddress *Address
	ArrivalDate    string
	DischargeDate  string
	Appointment    string
	Hl7Info        *Hl7Dates
}

// This is a lazy way of handling HL7 structures of dates. This should be handled differently.
type Hl7Dates struct {
	HL7Arrival     string
	HL7Discharge   string
	HL7Event       string
	HL7DOB         string
	HL7Appointment string
}

func NewPatients(max int) Collection {
	var tmp []*Patient

	for i := 0; i < max; i++ {
		patientInstance := NewPatient()
		tmp = append(tmp, patientInstance)
	}

	return Collection{tmp}

}

func NewPatient() *Patient {

	fakePatient := Name().
		NewAddress().
		Mrn().
		VisitID().
		PhoneNumber().
		DateOfBirth().
		Hl7DateOfBirthFmt().
		Arrival().
		Discharge().
		AppointmentDate().
		HL7Info()

	return fakePatient

}

func Name() *Patient {

	p := &Patient{}

	first := FIRSTNAME

	last := LASTNAME

	number := tools.RandomSelector(first)

	number2 := tools.RandomSelector(last)

	p.FirstName = first[number]
	p.LastName = last[number2]

	return p
}

func (p *Patient) Mrn() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomMrn := rand.Intn(1000000000)

	p.MRN = fmt.Sprint(randomMrn)

	return p
}

func (p *Patient) VisitID() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomVisitID := rand.Intn(1000000000)

	p.VisitId = randomVisitID

	return p
}

func (p *Patient) DateOfBirth() *Patient {

	date := Date()

	p.DOB = date

	return p
}

// Obsolete
func (p *Patient) Hl7DateOfBirthFmt() *Patient {

	date := p.DOB

	split := strings.Split(date, "-")

	var digits []string

	for _, number := range split {
		parse, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		if parse < 10 {
			number = "0" + number
		}

		digits = append(digits, number)
	}

	format := fmt.Sprintf("%v%v%v", digits[2], digits[0], digits[1])

	p.Hl7DOB = format

	return p

}

func (p *Patient) Arrival() *Patient {

	currentTime := time.Now()

	Arrival := currentTime.AddDate(0, 0, -5)

	p.ArrivalDate = fmt.Sprintf("%v/%v/%v", int(Arrival.Month()), Arrival.Day(), Arrival.Year())

	return p
}

func (p *Patient) Discharge() *Patient {

	currentTime := time.Now()

	Discharge := currentTime.AddDate(0, 0, -3)

	p.DischargeDate = fmt.Sprintf("%v/%v/%v", int(Discharge.Month()), Discharge.Day(), Discharge.Year())

	return p
}

func (p *Patient) AppointmentDate() *Patient {

	date := p.ArrivalDate

	splitDate := strings.Split(date, "/")

	month, err := strconv.Atoi(splitDate[0])
	if err != nil {
		fmt.Println(err)
	}

	if month+1 > 12 {
		month = 1
	} else {
		month++
	}

	p.Appointment = fmt.Sprintf("%v/%v/%v", month, splitDate[1], splitDate[2])

	return p
}

func EventDate() string {

	currentTime := time.Now()

	time := fmt.Sprintf("%v/%v/%v", int(currentTime.Month()), currentTime.Day(), currentTime.Year())

	return time

}

func (p *Patient) HL7Info() *Patient {

	p.Hl7Info = HL7DateConstructor(p.ArrivalDate, p.DischargeDate, p.DOB, p.Appointment)

	return p

}

func HL7DateConstructor(arrival string, discharge string, dob string, appointment string) *Hl7Dates {

	dates := &Hl7Dates{}
	event := EventDate()

	dates.HL7Arrival = Hl7DateFormatter(arrival)
	dates.HL7Discharge = Hl7DateFormatter(discharge)
	dates.HL7DOB = Hl7DateFormatter(dob)
	dates.HL7Event = Hl7DateFormatter(event)
	dates.HL7Appointment = Hl7DateFormatter(appointment)

	return dates

}
