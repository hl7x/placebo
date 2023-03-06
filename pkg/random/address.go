package random

import (
	"fmt"
)

type Address struct {
	Street		string
	Number		string
	PostalCode	string
	State		string
}

func (p *Patient) NewAddress() *Patient {
	
	address := Street().
		Number().
		PostalCodeAndState()
	
	return &Patient{}

}

func Street() *Address {

	address1 := []string{
			"main",
			"first",	
			"second",
			"third",
			"fourth",
			"fifth",
			"sixth",
			"seventh",
			"eigth",
			"ninth",
			"tenth",
			"eleventh",
			"twelfth",
			"maple"
			"elm",
			"birch",
			"red",
			"orange",
			"yellow",
			"blue",
			"purple"}

	address2 := []string{
			"street",
			"avenue",
			"way",
			"road"}

}

func (a *Address) Number() *Address {


}

func (a *Address) PostalCodeAndState() *Address {

	state := []string{
		"AL",
		"AK",
		"AZ",
		"AR",
		"CA",
		"CO",
		"CT",
		"DE",
		"DC",
		"FL",
		"GA",
		"HI",
		"ID",
		"IL",
		"IN",
		"IA",
		"KS",
		"KY",
		"LA",
		"ME",
		"MD",
		"MA",
		"MI",
		"MN",
		"MS",
		"MO",
		"MT",
		"NE",
		"NV",
		"NH",
		"NJ",
		"NM",
		"NY",
		"NC",
		"ND",
		"OH",
		"OK",
		"OR",
		"PA",
		"RI",
		"SC",
		"SD",
		"TN",
		"TX",
		"UT",
		"VT",
		"VA",
		"WA",
		"WV",
		"WI",
		"WY"}

}
