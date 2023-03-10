package random

import (
	"fmt"
	"math/rand"
	"placebo/internal/tools"
	"sync"
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
	c := make(chan *Patient)
	var wg sync.WaitGroup

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			patientInstance := NewPatient()
			c <- patientInstance
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var tmp []*Patient
	for patientInstance := range c {
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

	return &Patient{FirstName: first[number], LastName: last[number2]}
}

func (p *Patient) MrnAndEncounterID() *Patient {

	rand.Seed(time.Now().UnixNano())
	randomMrn := rand.Intn(1000000000)

	rand.Seed(time.Now().UnixNano())
	randomEncounterID := rand.Intn(1000000000)

	return &Patient{FirstName: p.FirstName, LastName: p.LastName, MRN: randomMrn, EncounterId: randomEncounterID, Phone: p.Phone, DOB: p.DOB, PatientAddress: p.PatientAddress}

}

func (p *Patient) PhoneNumber() *Patient {

	number := 8065550109

	return &Patient{FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: number, DOB: p.DOB, PatientAddress: p.PatientAddress}

}

func (p *Patient) DateOfBirth() *Patient {

	currentTime := time.Now()

	month := currentTime.Month()
	day := currentTime.Day()
	year := currentTime.Year()

	date := fmt.Sprintf("%v-%v-%v", int(month), day, year)

	return &Patient{FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: p.Phone, DOB: date, PatientAddress: p.PatientAddress}

}

func (p *Patient) Arrival() *Patient {

	//stub
	return &Patient{FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: p.Phone, DOB: p.DOB, ArrivalDate: "10/10/2022", DischargeDate: p.DischargeDate, PatientAddress: p.PatientAddress}
}

func (p *Patient) Discharge() *Patient {

	//stub
	return &Patient{FirstName: p.FirstName, LastName: p.LastName, MRN: p.MRN, EncounterId: p.EncounterId, Phone: p.Phone, DOB: p.DOB, ArrivalDate: p.ArrivalDate, DischargeDate: "10/12/2022", PatientAddress: p.PatientAddress}
}
