package random

import (
	"fmt"
	"github.com/hl7x/placebo/internal/tools"
	"math/rand"
)

var FACILITY = []string{
	"WING",
	"LEVEL",
	"SECTION",
	"WARD",
	"MEDSURG",
	"INTENSIVE",
	"PREOP"}

var ORIENT = []string{
	"NORTH",
	"EAST",
	"SOUTH",
	"WEST"}

var BEDLETTER = []string{
	"A",
	"B",
	"C",
	"D",
	"E"}

type Location struct {
	Bed              string
	Room             string
	LocationFacility string
}

func NewLocation() *Location {

	loc := &Location{}

	return loc.LocationBed().LocationRoom().Facility()
}

func NewFacility() string {

	f := FACILITY[tools.RandomSelector(FACILITY)]

	orient := ORIENT[tools.RandomSelector(ORIENT)]

	return orient + " " + f
}

func newBed() string {

	number := fmt.Sprint(rand.Intn(500))

	bedLetter := BEDLETTER[tools.RandomSelector(BEDLETTER)]

	return number + bedLetter

}

func newRoom() string {

	number := fmt.Sprint(rand.Intn(1000))

	return number
}

func (l *Location) LocationBed() *Location {

	l.Bed = newBed()

	return l
}

func (l *Location) LocationRoom() *Location {

	l.Room = newRoom()

	return l
}

func (l *Location) Facility() *Location {

	l.LocationFacility = NewFacility()

	return l
}
