package random

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hl7x/placebo/internal/tools"
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
	DOB            PatientDate
	Sex            string
	PatientAddress *Address
	ArrivalDate    PatientDate
	DischargeDate  PatientDate
	Appointment    PatientDate
	EventDate      PatientDate
	Provider       *Provider
	Location       *Location
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
		SetEventDate().
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

	p.MRN = fmt.Sprint(rand.Intn(1000000000))

	return p
}

func (p *Patient) Id() *Patient {

	p.PatientId = fmt.Sprint(rand.Intn(1000000000))

	return p
}

func (p *Patient) VisitID() *Patient {

	p.VisitId = rand.Intn(1000000000)

	return p
}

func (p *Patient) DateOfBirth() *Patient {

	p.DOB = Date()

	return p
}

func (p *Patient) PatientSex() *Patient {
	sexOptions := []string{"M", "F"}

	p.Sex = sexOptions[tools.RandomSelector(sexOptions)]

	return p
}

func (p *Patient) Arrival() *Patient {

	p.ArrivalDate = PatientDate(time.Now().AddDate(0, 0, -5))

	return p
}

func (p *Patient) Discharge() *Patient {

	p.DischargeDate = PatientDate(time.Now().AddDate(0, 0, -3))

	return p
}

func (p *Patient) AppointmentDate() *Patient {

	p.Appointment = PatientDate(time.Time(p.ArrivalDate).AddDate(0, 1, 0))

	return p
}

func (p *Patient) SetEventDate() *Patient {

	p.EventDate = PatientDate(time.Now())

	return p
}

func (p *Patient) PatientProvider() *Patient {

	p.Provider = NewProvider()

	return p
}

func (p *Patient) PatientLocation() *Patient {

	p.Location = NewLocation()

	return p
}
