package data

import "fmt"

type Distance float64
type DistanceKm float64

func (d Distance) ToKm() DistanceKm {
	return DistanceKm(d * 1.60934)
}

func (d DistanceKm) ToMiles() Distance {
	return Distance(d / 1.60934)
}

func Test() {
	d := Distance(34.5)
	km := d.ToKm()
	miles := km.ToMiles()

	fmt.Printf("%v miles is %v km\n", miles, km)
}
