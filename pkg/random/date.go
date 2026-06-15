package random

import (
	"math/rand"
	"time"
)

// PatientDate wraps time.Time and exposes format methods so each consumer
// (HL7 segments, CSV templates) can request the format it needs.
type PatientDate time.Time

// HL7 returns the date formatted as YYYYMMDD, the standard HL7 date format.
func (d PatientDate) HL7() string {
	return time.Time(d).Format("20060102")
}

// CSV returns the date formatted as M/D/YYYY for use in CSV output.
func (d PatientDate) CSV() string {
	return time.Time(d).Format("1/2/2006")
}

// String returns the CSV format so template rendering uses the human-readable form by default.
func (d PatientDate) String() string {
	return d.CSV()
}

func Date() PatientDate {
	return PatientDate(time.Date(Year(), time.Month(Month()), Day(), 0, 0, 0, 0, time.UTC))
}

func Month() int {
	max := 6
	min := 1
	return rand.Intn(max-min) + max
}

func Day() int {
	max := 15
	min := 1
	return rand.Intn(max-min) + max
}

func Year() int {
	max := 1970
	min := 1920
	return rand.Intn(max-min) + max
}
