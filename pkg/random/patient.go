package random

import (
	"fmt"
	"math/rand"
	"time"
)

type Collection struct { 
	Patients []*Patient
}

type Patient struct {
	FirstName	string
	LastName	string
	MRN		int
	EncounterId	int
	Phone		int
	DOB		string
	Address		*Address
	ArrivalDate	string
	DischargeDate	string
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
		MrnAndEncounterID().
		PhoneNumber().
		DateOfBirth().
		Arrival().
		Discharge()

	return fakePatient

}

func Name() *Patient {

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

	min := 0
	max  := len(first)
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(max - min)
	
	rand.Seed(time.Now().UnixNano())

	number2 := rand.Intn(max - min)

	return &Patient{ FirstName: first[number], LastName: last[number2] }
}

func (p *Patient) MrnAndEncounterID() *Patient {
	
	rand.Seed(time.Now().UnixNano())
	randomMrn := rand.Intn(1000000000)

	rand.Seed(time.Now().UnixNano())
	randomEncounterID := rand.Intn(1000000000)

	return &Patient{FirstName: p.FirstName, LastName: p.LastName, MRN: randomMrn, EncounterId: randomEncounterID, Phone: p.Phone, DOB: p.DOB }

}

func (p *Patient) PhoneNumber() *Patient {

	number := 8065550109

	return &Patient{ FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: number, DOB: p.DOB }

}

func (p *Patient) DateOfBirth() *Patient {
	
	currentTime := time.Now()
	
	month := currentTime.Month()
	day := currentTime.Day()
	year := currentTime.Year()

	date := fmt.Sprintf("%v-%v-%v", int(month), day, year)

	return &Patient{ FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: p.Phone, DOB: date }

}

func (p *Patient) Arrival() *Patient {

	//stub
	return &Patient{ FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: p.Phone, DOB: p.DOB, ArrivalDate: "10/10/2022", DischargeDate: p.DischargeDate }
}

func (p *Patient) Discharge() *Patient {

	//stub
	return &Patient{ FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: p.Phone, DOB: p.DOB, ArrivalDate: p.ArrivalDate, DischargeDate: "10/12/2022" }
}
