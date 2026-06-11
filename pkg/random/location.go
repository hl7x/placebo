package random

type Location struct {
	Bed              string
	Room             string
	LocationFacility string
}

func NewLocation() *Location {

	loc := &Location{}

	return loc.LocationBed().LocationRoom().Facility()
}

func (l *Location) LocationBed() *Location {

	l.Bed = "20A"

	return l
}

func (l *Location) LocationRoom() *Location {

	l.Room = "200"

	return l
}

func (l *Location) Facility() *Location {

	l.LocationFacility = "South Terminal"

	return l
}
