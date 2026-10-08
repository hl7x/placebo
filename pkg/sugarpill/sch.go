package sugarpill

import (
	"fmt"

	"github.com/hl7x/placebo/pkg/random"
)

// SCH carries the appointment a scheduling message is about. It is the
// segment that makes an SIU message an SIU message.
type SCH struct {
	PlacerAppointmentID       *EntityIdentifier `json:"PlacerAppointmentID"`       // SCH-1
	FillerAppointmentID       *EntityIdentifier `json:"FillerAppointmentID"`       // SCH-2
	OccurrenceNumber          string            `json:"OccurrenceNumber"`          // SCH-3
	PlacerGroupNumber         *EntityIdentifier `json:"PlacerGroupNumber"`         // SCH-4
	ScheduleID                *ServiceCode      `json:"ScheduleID"`                // SCH-5
	EventReason               *ServiceCode      `json:"EventReason"`               // SCH-6
	AppointmentReason         *ServiceCode      `json:"AppointmentReason"`         // SCH-7
	AppointmentType           *ServiceCode      `json:"AppointmentType"`           // SCH-8
	AppointmentDuration       string            `json:"AppointmentDuration"`       // SCH-9
	AppointmentDurationUnits  *ServiceCode      `json:"AppointmentDurationUnits"`  // SCH-10
	AppointmentTimingQuantity *QuantityTiming   `json:"AppointmentTimingQuantity"` // SCH-11
	PlacerContactPerson       *XCN              `json:"PlacerContactPerson"`       // SCH-12
	PlacerContactPhoneNumber  *XTN              `json:"PlacerContactPhoneNumber"`  // SCH-13
	PlacerContactAddress      *XAD              `json:"PlacerContactAddress"`      // SCH-14
	PlacerContactLocation     *PatientLocation  `json:"PlacerContactLocation"`     // SCH-15
	FillerContactPerson       *XCN              `json:"FillerContactPerson"`       // SCH-16
	FillerContactPhoneNumber  *XTN              `json:"FillerContactPhoneNumber"`  // SCH-17
	FillerContactAddress      *XAD              `json:"FillerContactAddress"`      // SCH-18
	FillerContactLocation     *PatientLocation  `json:"FillerContactLocation"`     // SCH-19
	EnteredByPerson           *XCN              `json:"EnteredByPerson"`           // SCH-20
	EnteredByPhoneNumber      *XTN              `json:"EnteredByPhoneNumber"`      // SCH-21
	EnteredByLocation         *PatientLocation  `json:"EnteredByLocation"`         // SCH-22
	ParentPlacerAppointmentID *EntityIdentifier `json:"ParentPlacerAppointmentID"` // SCH-23
	ParentFillerAppointmentID *EntityIdentifier `json:"ParentFillerAppointmentID"` // SCH-24
	FillerStatusCode          *ServiceCode      `json:"FillerStatusCode"`          // SCH-25
	PlacerOrderNumber         *EntityIdentifier `json:"PlacerOrderNumber"`         // SCH-26
	FillerOrderNumber         *EntityIdentifier `json:"FillerOrderNumber"`         // SCH-27
}

func NewSCHSegment(p *random.Patient) *SCH {

	appointment := &EntityIdentifier{
		EntityIdentifier: fmt.Sprint(p.VisitId),
		NamespaceID:      "PLACEBO",
	}

	provider := &XCN{
		ID:         p.Provider.ID,
		FamilyName: p.Provider.LastName,
		GivenName:  p.Provider.FirstName,
	}

	sch := &SCH{
		PlacerAppointmentID:      appointment,
		FillerAppointmentID:      &EntityIdentifier{},
		PlacerGroupNumber:        &EntityIdentifier{},
		ScheduleID:               &ServiceCode{},
		EventReason:              &ServiceCode{},
		AppointmentReason:        &ServiceCode{Identifier: "ROUTINE", Text: "ROUTINE APPOINTMENT"},
		AppointmentType:          &ServiceCode{Identifier: "NORMAL"},
		AppointmentDuration:      "30",
		AppointmentDurationUnits: &ServiceCode{Identifier: "MIN", Text: "MINUTES"},
		// SCH-11 is a timing quantity, the same datatype OBR-27 carries. Its
		// start date is what a scheduling system reads the appointment off.
		AppointmentTimingQuantity: &QuantityTiming{StartDate: p.Appointment.HL7()},
		PlacerContactPerson:       &XCN{},
		PlacerContactPhoneNumber:  &XTN{},
		PlacerContactAddress:      &XAD{},
		PlacerContactLocation:     &PatientLocation{},
		FillerContactPerson:       provider,
		FillerContactPhoneNumber:  &XTN{},
		FillerContactAddress:      &XAD{},
		FillerContactLocation:     &PatientLocation{},
		EnteredByPerson:           &XCN{},
		EnteredByPhoneNumber:      &XTN{},
		EnteredByLocation:         &PatientLocation{},
		ParentPlacerAppointmentID: &EntityIdentifier{},
		ParentFillerAppointmentID: &EntityIdentifier{},
		FillerStatusCode:          &ServiceCode{Identifier: "BOOKED"},
		PlacerOrderNumber:         &EntityIdentifier{},
		FillerOrderNumber:         &EntityIdentifier{},
	}

	return sch
}
