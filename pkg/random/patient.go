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
	MiddleName     string
	LastName       string
	MiddleInitial  string
	MRN            string
	PatientId      string
	VisitId        int
	Phone          string
	DOB            string
	Sex            string
	PatientAddress *Address
	ArrivalDate    string
	DischargeDate  string
	Appointment    string
	Hl7Info        *Hl7Dates
	Provider       *Provider
	Location       *Location
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

	fakePatient := PatientName().
		NewAddress().
		Mrn().
		Id().
		VisitID().
		PhoneNumber().
		DateOfBirth().
		PatientSex().
		Arrival().
		Discharge().
		AppointmentDate().
		HL7Info().
		PatientProvider().
		PatientLocation()

	return fakePatient

}

func PatientName() *Patient {

	p := &Patient{}

	name := NewName()

	p.FirstName = name.FirstName
	p.MiddleName = name.MiddleName
	p.LastName = name.LastName
	p.MiddleInitial = name.MiddleInitial

	return p
}

func (p *Patient) Mrn() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomMrn := rand.Intn(1000000000)

	p.MRN = fmt.Sprint(randomMrn)

	return p
}

func (p *Patient) Id() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomId := rand.Intn(1000000000)

	p.PatientId = fmt.Sprint(randomId)

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

func (p *Patient) PatientSex() *Patient {
	sexOptions := []string{"M", "F"}

	sex := sexOptions[tools.RandomSelector(sexOptions)]

	p.Sex = sex

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

func (p *Patient) PatientProvider() *Patient {

	provider := NewProvider()

	p.Provider = provider

	return p

}

func (p *Patient) PatientLocation() *Patient {

	location := NewLocation()

	p.Location = location

	return p
}
