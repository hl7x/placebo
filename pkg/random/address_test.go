package random

import "testing"

func TestPatient_NewAddress(t *testing.T) {

	t.Run("Return Address Properties", func(t *testing.T) {
		testPatient := &Patient{}
		address := testPatient.NewAddress()
		got := address
		if got.PatientAddress.Street == "" {
			t.Fatalf("Got %v, Street should not be empty", got.PatientAddress.Street)
		} else if got.PatientAddress.StructureNumber > 10000 || got.PatientAddress.StructureNumber < 0 {
			t.Fatalf("Got %v, number should be greater than 0 and less than 10000", got.PatientAddress.StructureNumber)
		} else if got.PatientAddress.RegionInfo.City == "" {
			t.Fatalf("Got %v, City Should not be empty", got.PatientAddress.RegionInfo.City)
		}
	})
}
