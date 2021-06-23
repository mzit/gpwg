package main

import "fmt"

type report struct {
	sol
	temperature
	location
}
type sol int
type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}
type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{
		high: -1,
		low:  -78,
	}
	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}
	fmt.Printf("%+v\n", report)
	fmt.Println(report.temperature.high)
	fmt.Println(report.high)
	report.high = 3333
	fmt.Println(report.temperature.high)
	fmt.Println(t.average())
	fmt.Println(report.temperature.average())
	fmt.Println(report.days(1234))
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}
func (l location) days(l2 location) int {
	return 5
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}
