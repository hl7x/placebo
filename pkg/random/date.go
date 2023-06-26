package random

import (
	"math/rand"
	"fmt"
)

func Date() string {
	
	month := Month()
	day := Day()
	year := Year()

	dateOfBirth := fmt.Sprintf("%v-%v-%v", month, day, year)

	return dateOfBirth	

	
}

func Month() int {

	max := 5
	min := 1

	month := rand.Intn(max - min) + max

	return month
}

func Day() int {
	
	max := 15
	min := 1

	day := rand.Intn(max - min) + max
	
	return day

}

func Year() int {

	max := 1970
	min := 1920

	year := rand.Intn(max - min) + max

	return year
}
