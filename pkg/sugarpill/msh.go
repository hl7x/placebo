package sugarpill

import (
	"fmt"
	"math/rand"
	"placebo/pkg/random"
)

type MSH struct {
	Encode               string       `json:Encode`                 // MSH-2
	SendingApplication   string       `json:"SendingApplication"`   // MSH-3
	SendingFacility      string       `json:"SendingFacility"`      // MSH-4
	ReceivingApplication string       `json:"ReceivingApplication"` // MSH-5
	ReceivingFacility    string       `json:"ReceivingFacility"`    // MSH-6
	DateTimeOfMessage    string       `json:"DateTimeOfMessage"`    // MSH-7
	Security             string       `json:"Security"`             // MSH-8
	MessageType          *MessageType `json:"MessageType"`          // MSH-9
	MessageControlID     string       `json:"MessageControlID"`     // MSH-10
	ProcessingID         string       `json:"ProcessingID"`         // MSH-11
	VersionID            string       `json:"VersionID"`            // MSH-12
}

type MessageType struct {
	MessageCode  string `json:"MessageCode"`  // MSH-9.1
	TriggerEvent string `json:"TriggerEvent"` // MSH-9.2
}

func NewMSHSegment(p *random.Patient) *MSH {

	messageType := &MessageType{
		MessageCode:  "ADT",
		TriggerEvent: "A01",
	}

	msh := &MSH{
		Encode:               "^~\\&",
		SendingApplication:   "PLACEBO",
		SendingFacility:      "PLACEBO",
		ReceivingApplication: "DEMO",
		ReceivingFacility:    "DEMO",
		DateTimeOfMessage:    p.EventDate.HL7(),
		Security:             "",
		MessageType:          messageType,
		MessageControlID:     MessageControlID(),
		ProcessingID:         "P",
		VersionID:            "2.3",
	}

	return msh
}

func MessageControlID() string {
	id := fmt.Sprintf("%d", rand.Int31())

	return id
}
