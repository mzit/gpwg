package main

import "fmt"

type coordinate struct {
	//degrees度/minutes分/seconds秒
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func main() {
	//lat := coordinate{
	//	d: 4,
	//	m: 35,
	//	s: 22.2,
	//	h: 'S',
	//}
	//long := coordinate{
	//	137, 26, 30.12, 'e',
	//}
	//fmt.Println(lat.decimal(), long.decimal())
	//curiosity := location{
	//	lat.decimal(), long.decimal(),
	//}
	//fmt.Println(curiosity)

	curiosity := newLocation(coordinate{4, 35, 22.2, 's'}, coordinate{137, 26, 30.12, 'e'})
	fmt.Println(curiosity)
}

type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
