package random

import (
	"math/rand"
	"placebo/internal/tools"
	"time"
	"fmt"
	"strings"
	"strconv"
)

type Collection struct {
	Patients []*Patient
}

type Patient struct {
	FirstName      	string
	LastName       	string
	MRN            	int
	EncounterId    	int
	Phone          	string
	DOB            	string
	Hl7DOB		string
	PatientAddress 	*Address
	ArrivalDate    	string
	DischargeDate  	string
	Appointment	string
	Hl7Info		*Hl7Dates
}

type Hl7Dates struct {
	HL7Arrival	string
	HL7Discharge	string
	HL7Event	string
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
		EncounterID().
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

	p.MRN = randomMrn

	return p
}

func (p *Patient) EncounterID() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomEncounterID := rand.Intn(1000000000)
	
	p.EncounterId = randomEncounterID

	return p
}

func (p *Patient) DateOfBirth() *Patient {

	date := Date()

	p.DOB = date

	return p
}

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

	if month +1 > 12 {
		month = 1
	} else {
		month++
	}

	p.Appointment = fmt.Sprintf("%v/%v/%v", month, splitDate[1], splitDate[2])

	return p
}

func (p *Patient) HL7Info() *Patient {
	
	p.Hl7Info = HL7DateConstructor(p.ArrivalDate, p.DischargeDate)

	return p

}

func HL7DateConstructor(arrival string, discharge string) *Hl7Dates {

	dates := &Hl7Dates{}

	a := Hl7DateFormatter(arrival)
	d := Hl7DateFormatter(discharge)

	dates.HL7Arrival = a
	dates.HL7Discharge = d

	return dates

}

