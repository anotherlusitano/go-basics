package data

import "fmt"

// ----------------------------------
// ALIAS TO TYPE

type (
	integer = int
	json    = map[string]string
)

var x integer

// ----------------------------------
// NEW TYPE

type (
	distance   float64 // miles
	distanceKm float64
)

// Method of distance type
func (miles distance) toKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (miles distanceKm) toMiles() distance {
	return distance(miles / 1.60934)
}

func Test() {
	d := distance(34.5)
	e := 34.5

	f := d.toKm()
	g := f.toMiles()

	fmt.Printf("%v is from type: %T\n", d, d)
	fmt.Printf("%v is from type: %T\n", e, e)
	fmt.Printf("%v is from type: %T\n", f, f)
	fmt.Printf("%v is from type: %T\n", g, g)
}
