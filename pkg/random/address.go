package random

import (
	"math/rand"
	"placebo/internal/tools"
	"strings"
)

type Address struct {
	Street          string
	StructureNumber int
	RegionInfo      *Region
}

// PostalCode will be added later
type Region struct {
	State      string
	City       string
	PostalCode int
}

func (p *Patient) NewAddress() *Patient {

	address := Street().
		Number().
		RegionSpecific()

	p.PatientAddress = address

	return p
}

func (a *Address) RegionSpecific() *Address {

	r := &Region{}
	state := State()
	city := City(state)

	r.State = state
	r.City = city

	a.RegionInfo = r
	return a

}

func Street() *Address {

	streetAddress := []string{
		"MAIN",
		"FIRST",
		"SECOND",
		"THIRD",
		"FOURTH",
		"FIFTH",
		"SIXTH",
		"SEVENTH",
		"EIGTH",
		"NINTH",
		"TENTH",
		"ELEVENTH",
		"TWELFTH",
		"MAPLE",
		"ELM",
		"BIRCH",
		"RED",
		"ORANGE",
		"YELLOW",
		"BLUE",
		"PURPLE"}

	roadSign := []string{
		"STREET",
		"AVENUE",
		"WAY",
		"ROAD"}

	streetInt := tools.RandomSelector(streetAddress)
	roadInt := tools.RandomSelector(roadSign)

	return &Address{Street: streetAddress[streetInt] + " " + roadSign[roadInt]}
}

func (a *Address) Number() *Address {

	streetNumber := rand.Intn(10000)

	a.StructureNumber = int(streetNumber)

	return a

}

func City(s string) string {

	pair := map[string][]string{
		"AL": {"MOBILE", "MONTGOMERY", "BIRMINGHAM"},
		"AK": {"ANCHORAGE", "JUNEAU", "FAIRBANKS"},
		"AZ": {"PHOENIX", "TUCSON", "MESA"},
		"AR": {"LITTLE ROCK", "FAYETTEVILLE", "SPRINGDALE"},
		"CA": {"LOS ANGELES", "SAN DIEGO", "SAN JOSE"},
		"CO": {"DENVER", "COLORADO SPRINGS"},
		"CT": {"BRIDGEPORT", "STAMFORD", "NEW HAVEN"},
		"DE": {"WILMINGTON", "DOVER"},
		"DC": {"WASHINGTON DC"},
		"FL": {"JACKSONVILLE", "MIAMI", "TAMPA"},
		"GA": {"ATLANTA", "AUGUSTA"},
		"HI": {"HONOLULU", "PEARL CITY", "HILO"},
		"ID": {"BOISE", "MERIDIAN", "NAMPA"},
		"IL": {"CHICAGO", "AURORA", "NAPERVILLE"},
		"IN": {"INDIANAPOLIS", "FORT WAYNE", "EVANSVILLE"},
		"IA": {"DES MOINES", "CEDAR RAPIDS", "DAVENPORT"},
		"KS": {"WICHITA", "OVERLAND PARK"},
		"KY": {"LOUISVILLE", "LEXINGTON", "BOWLING GREEN"},
		"LA": {"NEW ORLEANS", "BATON ROUGE", "SHERVEPORT"},
		"ME": {"BANGOR", "LEWISTON"},
		"MD": {"BALTIMORE", "GERMANTOWN"},
		"MA": {"BOSTON", "WORCESTER"},
		"MI": {"DETROIT", "GRAND RAPIDS", "WARREN"},
		"MN": {"MINNEAPOLIS", "SAINT PAUL", "ROCHESTER"},
		"MS": {"JACKSON", "GULFPORT", "SOUTHAVEN"},
		"MO": {"KANSAS CITY", "SAINT LOUIS", "SPRINGFIELD"},
		"MT": {"BILLINGS", "MISSOULA", "GREAT FALLS"},
		"NE": {"OMAHA", "LINCOLN", "BELLEVUE"},
		"NV": {"LAS VEGAS", "HENDERSON", "RENO"},
		"NH": {"MANCHESTER", "NASHUA", "CONCORD"},
		"NJ": {"NEWARK", "JERSEY CITY", "PATERSON"},
		"NM": {"ALBUQUERQUE", "LAS CRUCES", "SANTA FE"},
		"NY": {"NEW YORK CITY", "BUFFALO", "YONKERS"},
		"NC": {"CHARLOTTE", "RALEIGH", "GREENSBORO"},
		"ND": {"FARGO", "BISMARK", "GRAND FORKS"},
		"OH": {"COLUMBUS", "CLEVELAND", "CINCINNATI"},
		"OK": {"OKLAHOMA CITY", "TULSA", "NORMAN"},
		"OR": {"PORTLAND", "SALEM", "EUGENE"},
		"PA": {"PHILADELPHIA", "PITTSBURGH", "ALLENTOWN"},
		"RI": {"PROVIDENCE", "CRANSTON", "WARWICK"},
		"SC": {"COLUMBIA", "MOUNT PLEASANT"},
		"SD": {"SIOUX FALLS", "RAPID CITY", "ABERDEEN"},
		"TN": {"NASHVILLE", "MEMPHIS", "KNOXVILLE"},
		"TX": {"DALLAS", "HOUSTON", "AUSTIN"},
		"UT": {"SALT LAKE CITY", "PROVO", "WEST JORDAN"},
		"VT": {"BURLINGTON", "COLCHESTER", "RUTLAND"},
		"VA": {"VIRGINIA BEACH", "NORFOLK", "RICHMOND"},
		"WA": {"SEATTLE", "SPOKANE", "TACOMA"},
		"WV": {"CHARLESTON", "HUNTINGTON", "MORGANTOWN"},
		"WI": {"MILWAUKEE", "MADISON", "GREEN BAY"},
		"WY": {"CHEYENNE", "CASPER", "LARAMIE"}}

	for k, v := range pair {
		valueInt := tools.RandomSelector(v)
		if strings.Contains(k, s) {
			return v[valueInt]
		}
	}
	return ""
}

func State() string {

	stateAbbr := []string{
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

	randomNumber := tools.RandomSelector(stateAbbr)

	return string(stateAbbr[randomNumber])

}
