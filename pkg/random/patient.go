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
	PatientAddress 	*Address
	ArrivalDate    	string
	DischargeDate  	string
	Appointment	string
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
		Arrival().
		Discharge().
		AppointmentDate()

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

func ( p *Patient) AppointmentDate() *Patient {

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

