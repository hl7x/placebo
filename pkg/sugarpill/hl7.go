package sugarpill

import (
	"encoding/json"
	"strings"

	"placebo/pkg/message"
	"placebo/pkg/random"
)

type HL7Message struct {
	MSH *MSH `json:"MSH"`
	EVN *EVN `json:"EVN"`
	PID *PID `json:"PID"`
	ROL *ROL `json:"ROL"`
	DB1 *DB1 `json:"DB1"`
	ARV *ARV `json:"ARV"`
	PV1 *PV1 `json:"PV1"`
	PV2 *PV2 `json:"PV2"`
	DG1 *DG1 `json:"DG1"`
	OBX *OBX `json:"OBX"`
	NK1 *NK1 `json:"NK1"`
	AL1 *AL1 `json:"AL1"`
	OBR *OBR `json:"OBR"`
}

func NewHL7Message(p *random.Patient) *HL7Message {

	message := &HL7Message{
		MSH: NewMSHSegment(p),
		EVN: NewEVNSegment(p),
		PID: NewPIDSegment(p),
		ROL: NewROLSegment(p),
		DB1: NewDB1Segment(p),
		ARV: NewARVSegment(p),
		PV1: NewPV1Segment(p),
		PV2: NewPV2Segment(p),
		DG1: NewDG1Segment(p),
		OBX: NewOBXSegment(p),
		NK1: NewNK1Segment(p),
		AL1: NewAL1Segment(p),
		OBR: NewOBRSegment(p),
	}

	return message
}

func NewHL7EventMessage(p *random.Patient, t string, e string) string {

	message := NewHL7Message(p)
	message.MSH.MessageType.MessageCode = t
	message.MSH.MessageType.TriggerEvent = e

	message.EVN.EventTypeCode = t

	product := MessageBuilder(message)

	return product
}

func (m *HL7Message) MessageToJson() string {

	message, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}

	return string(message)
}

func JsonToMessage(s string) *HL7Message {
	jsonMessage := []byte(s)

	var msg *HL7Message
	err := json.Unmarshal(jsonMessage, &msg)
	if err != nil {
		panic(err)
	}

	return msg
}

func (msg *HL7Message) ReadFromFile(s string) *HL7Message {
	lines := LinesFromFile(s)

	for _, line := range lines {
		segs := strings.Split(line, "|")
		segType := SegmentFinder(segs[0])
		onlyMessageSegs := segs[1:]
		bind := message.MarshallHL7(onlyMessageSegs, segType)
		msg.SegmentAssigner(bind)
	}

	return msg
}

func ReadHL7(content string) string {

	messageStruct := &HL7Message{}
	msg := messageStruct.ReadFromFile(content)

	messageFmt, _ := json.MarshalIndent(msg, "", " ")

	jsonMessage := SanitizeLines(string(messageFmt))

	return jsonMessage
}

func SegmentFinder(s string) interface{} {

	switch s {
	case "MSH":
		return &MSH{}
	case "EVN":
		return &EVN{}
	case "PID":
		return &PID{}
	case "ROL":
		return &ROL{}
	case "DB1":
		return &DB1{}
	case "ARV":
		return &ARV{}
	case "PV1":
		return &PV1{}
	case "PV2":
		return &PV2{}
	case "DG1":
		return &DG1{}
	case "OBX":
		return &OBX{}
	case "NK1":
		return &NK1{}
	case "AL1":
		return &AL1{}
	case "OBR":
		return &OBR{}
	default:
		return nil
	}
}

func (msg *HL7Message) SegmentAssigner(s interface{}) *HL7Message {

	switch v := s.(type) {
	case *MSH:
		msg.MSH = v
	case *EVN:
		msg.EVN = v
	case *PID:
		msg.PID = v
	case *ROL:
		msg.ROL = v
	case *DB1:
		msg.DB1 = v
	case *ARV:
		msg.ARV = v
	case *PV1:
		msg.PV1 = v
	case *PV2:
		msg.PV2 = v
	case *DG1:
		msg.DG1 = v
	case *OBX:
		msg.OBX = v
	case *NK1:
		msg.NK1 = v
	case *AL1:
		msg.AL1 = v
	case *OBR:
		msg.OBR = v
	default:
		return nil
	}

	return msg
}

func MessageBuilder(msg *HL7Message) string {
	segments := []string{
		message.CreateHL7(msg.MSH, "MSH"),
		message.CreateHL7(msg.EVN, "EVN"),
		message.CreateHL7(msg.PID, "PID"),
		message.CreateHL7(msg.ROL, "ROL"),
		message.CreateHL7(msg.DB1, "DB1"),
		message.CreateHL7(msg.ARV, "ARV"),
		message.CreateHL7(msg.PV1, "PV1"),
		message.CreateHL7(msg.PV2, "PV2"),
		message.CreateHL7(msg.DG1, "DG1"),
		message.CreateHL7(msg.OBX, "OBX"),
		message.CreateHL7(msg.NK1, "NK1"),
		message.CreateHL7(msg.AL1, "AL1"),
		message.CreateHL7(msg.OBR, "OBR"),
	}

	return strings.Join(segments, "\n")
}

func LinesFromFile(s string) []string {

	splitLines := strings.Split(s, "\n")

	return splitLines
}

// Remove 'null' segments for easier JSON display
func SanitizeLines(lines string) string {
	split := strings.Split(lines, "\n")

	var newLines []string

	for _, line := range split {
		if !strings.Contains(line, ": null") {
			newLines = append(newLines, line)
		}
	}

	join := strings.Join(newLines, "\n")

	return join
}
