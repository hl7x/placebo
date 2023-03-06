package random

import (
	"fmt"
)

type Address struct {
	Street		string
	Number		string
	PostalCode	string
	State		string
	Country		string
	City		string
}

func (p *Patient) NewAddress() *Patient {
	
	address := Street().
		Number().
		RegionSpecific()
	
	return &Patient{}

}

func Street() *Address {

	streetAddress := []string{
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

	roadSign := []string{
			"street",
			"avenue",
			"way",
			"road"}

}

func (a *Address) Number() *Address {
	
	streetNumber := rand.Intn(1000)
	
	return &Address{Street: a.Street, Number: streetNumber, PostalCode: a.PostalCode, State: a.State}

}

func City(s string) string {

		pair := map[string][]string{
		"AL": {"Mobile", "Montgomery", "Birmingham"},
		"AK": {"Anchorage", "Juneau", "Fairbanks"},
		"AZ": {"Phoenix", "Tucson", "Mesa"},
		"AR": {"Little Rock", "Fayetteville", "Springdale"},
		"CA": {"Los Angeles", "San Diego", "San Jose"},
		"CO": {"Denver", "Colorado Springs", "Aurora"},
		"CT": {"Bridgeport", "Stamford", "New Haven"},
		"DE": {"Wilmington", "Dover", "Newark"},
		"DC": {"Washington DC"},
		"FL": {"Jacksonville", "Miami", "Tampa"},
		"GA": {"Atlanta", "Columbus", "Augusta"},
		"HI": {"Honolulu", "Pearl City", "Hilo"},
		"ID": {"Boise", "Meridian", "Nampa"},
		"IL": {"Chicago", "Aurora", "Naperville"},
		"IN": {"Indianapolis", "Fort Wayne", "Evansville"},
		"IA": {"Des Moines", "Cedar Rapids", "Davenport"},
		"KS": {"Wichita", "Overland Park", "Kansas City"},
		"KY": {"Louisville", "Lexington", "Bowling Green"},
		"LA": {"New Orleans", "Baton Rouge", "Sherveport"},
		"ME": {"Portland", "Bangor", "Lewiston"},
		"MD": {"Baltimore", "Columbia", "Germantown"},
		"MA": {"Boston", "Worcester", "Springfield"},
		"MI": {"Detroit", "Grand Rapids", "Warren"},
		"MN": {"Minneapolis", "Saint Paul", "Rochester"},
		"MS": {"Jackson", "Gulfport", "Southaven"},
		"MO": {"Kansas City", "Saint Louis", "Springfield"},
		"MT": {"Billings", "Missoula", "Great Falls"},
		"NE": {"Omaha", "Lincoln", "Bellevue"},
		"NV": {"Las Vegas", "Henderson", "Reno"},
		"NH": {"Manchester", "Nashua", "Concord"},
		"NJ": {"Newark", "Jersey City", "Paterson"},
		"NM": {"Albuquerque", "Las Cruces", "Santa Fe"},
		"NY": {"New York City", "Buffalo", "Yonkers"},
		"NC": {"Charlotte", "Raleigh", "Greensboro"},
		"ND": {"Fargo", "Bismark", "Grand Forks"},
		"OH": {"Columbus", "Cleveland", "Cincinnati"},
		"OK": {"Oklahoma City", "Tulsa", "Norman"},
		"OR": {"Portland", "Salem", "Eugene"},
		"PA": {"Philadelphia", "Pittsburgh", "Allentown"},
		"RI": {"Providence", "Cranston", "Warwick"},
		"SC": {"Charlseton", "Columbia", "Mount Pleasant"},
		"SD": {"Sioux Falls", "Rapid City", "Aberdeen"},
		"TN": {"Nashville", "Memphis", "Knoxville"},
		"TX": {"Dallas", "Houston", "Austin"},
		"UT": {"Salt Lake City", "Provo", "West Jordan"},
		"VT": {"Burlington", "Colchester", "Rutland"},
		"VA": {"Virginia Beach", "Norfolk", "Richmond"},
		"WA": {"Seattle", "Spokane", "Tacoma"},
		"WV": {"Charleston", "Huntington", "Morgantown"},
		"WI": {"Milwaukee", "Madison", "Green Bay"},
		"WY": {"Cheyenne", "Casper", "Laramie"}}
		
		rand.Seed(time.Now().UnixNano())

		for k, v := range pair {
			if strings.Contains(k, s) {
				return v[rand.Intn(len(v))]
			}
		}
}
