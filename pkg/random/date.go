package random

import (
	"math/rand"
	"fmt"
	"strconv"
)

func Date() string {
	
	month := Month()
	day := Day()
	year := Year()

	dateOfBirth := fmt.Sprintf("%v-%v-%v", month, day, year)

	return dateOfBirth	

	
}

func Month() int {

	max := 6
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

func Hl7Formatter(date string) string {
	
	split := strings.Split(date, "-")

	var digits []string

	for _, number := range split {
		parse, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		if parse < 10 {
			number = "0" + number
		}

		digits = append(digits, number)
	}

	format := fmt.Sprintf("%v%v%v", digits[2], digits[0], digits[1])

	return format

}

