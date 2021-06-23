package main

import (
	"fmt"
	"math"
)

type world1 struct {
	radius float64
}

type coordinate1 struct {
	//degrees度/minutes分/seconds秒
	d, m, s float64
	h       rune
}

var mars = world1{radius: 3389.5}
var earth = world1{radius: 6371}

func (w world1) distance(p1, p2 location1) float64 {
	//两个位置之间的距离
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

//rad 将角度转为弧度
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	spirit := newLocation1(coordinate1{14, 34, 6.2, 'S'}, coordinate1{175, 28, 21.5, 'E'})
	opportunity := newLocation1(coordinate1{1, 56, 46.3, 'S'}, coordinate1{354, 28, 24.2, 'E'})
	curiosity := newLocation1(coordinate1{4, 35, 22.2, 'S'}, coordinate1{137, 26, 30.12, 'E'})
	insight := newLocation1(coordinate1{4, 30, 0.0, 'N'}, coordinate1{135, 54, 0, 'E'})

	fmt.Printf("Spirit to Opportunity %.2f km\n", mars.distance(spirit, opportunity))
	fmt.Printf("Spirit to Curiosity %.2f km\n", mars.distance(spirit, curiosity))
	fmt.Printf("Spirit to InSight %.2f km\n", mars.distance(spirit, insight))

	fmt.Printf("Opportunity to Curiosity %.2f km\n", mars.distance(opportunity, curiosity))
	fmt.Printf("Opportunity to InSight %.2f km\n", mars.distance(opportunity, insight))

	fmt.Printf("Curiosity to InSight %.2f km\n", mars.distance(curiosity, insight))

	london := newLocation1(coordinate1{51, 30, 0, 'N'}, coordinate1{0, 8, 0, 'W'})
	paris := newLocation1(coordinate1{48, 51, 0, 'N'}, coordinate1{2, 21, 0, 'E'})
	fmt.Printf("London to Paris %.2f km\n", earth.distance(london, paris))

	edmonton := newLocation1(coordinate1{53, 32, 0, 'N'}, coordinate1{113, 30, 0, 'W'})
	ottawa := newLocation1(coordinate1{45, 25, 0, 'N'}, coordinate1{75, 41, 0, 'W'})
	fmt.Printf("Hometown to Capital %.2f km\n", earth.distance(edmonton, ottawa))

	mountSharp := newLocation1(coordinate1{5, 4, 48, 'S'}, coordinate1{137, 51, 0, 'E'})
	olympusMons := newLocation1(coordinate1{18, 39, 0, 'N'}, coordinate1{226, 12, 0, 'E'})
	fmt.Printf("Mount Sharp to Olympus Mons %.2f km\n", mars.distance(mountSharp, olympusMons))

}

type location1 struct {
	lat, long float64
}

func newLocation1(lat, long coordinate1) location1 {
	return location1{lat.decimal1(), long.decimal1()}
}
func (c coordinate1) decimal1() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
