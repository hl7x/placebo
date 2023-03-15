package random

import (
	"fmt"
	"math/rand"
	"placebo/internal/tools"
	"time"
)

type Collection struct {
	Patients []*Patient
}

type Patient struct {
	FirstName      string
	LastName       string
	MRN            int
	EncounterId    int
	Phone          int
	DOB            string
	PatientAddress *Address
	ArrivalDate    string
	DischargeDate  string
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
		MrnAndEncounterID().
		PhoneNumber().
		DateOfBirth().
		Arrival().
		Discharge()

	return fakePatient

}

func Name() *Patient {

	p := &Patient{}

	first := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
		"iota",
		"kappa",
		"lambda",
		"mu",
		"nu",
		"xi",
		"omicron",
		"pi",
		"rho",
		"sigma",
		"tau",
		"upsilon",
		"phi",
		"chi",
		"psi",
		"omega"}

	last := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
		"iota",
		"kappa",
		"lambda",
		"mu",
		"nu",
		"xi",
		"omicron",
		"pi",
		"rho",
		"sigma",
		"tau",
		"upsilon",
		"phi",
		"chi",
		"psi",
		"omega"}

	number := tools.RandomSelector(first)

	number2 := tools.RandomSelector(last)

	p.FirstName = first[number]
	p.LastName = last[number2]

	return p
}

func (p *Patient) MrnAndEncounterID() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomMrn := rand.Intn(1000000000)

	rand.Seed(time.Now().UnixNano())
	randomEncounterID := rand.Intn(1000000000)

	p.MRN = randomMrn
	p.EncounterId = randomEncounterID

	return p
}

func (p *Patient) PhoneNumber() *Patient {

	number := 8065550109

	p.Phone = number

	return p
}

func (p *Patient) DateOfBirth() *Patient {

	currentTime := time.Now()

	month := currentTime.Month()
	day := currentTime.Day()
	year := currentTime.Year()

	date := fmt.Sprintf("%v-%v-%v", int(month), day, year)

	p.DOB = date

	return p
}

func (p *Patient) Arrival() *Patient {

	//stub
	p.ArrivalDate = "10/10/2022"

	return p
}

func (p *Patient) Discharge() *Patient {

	//stub
	p.DischargeDate = "10/12/2022"

	return p
}
