package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

var mars = world{radius: 3389.5}

func (w world) distance(p1, p2 location1) float64 {
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
	spirit := location1{-14.5684, 175.472636}
	oppo := location1{-1.9462, 354.4734}
	dist := mars.distance(spirit, oppo)
	fmt.Printf("%.2f km\n", dist)
}

type location1 struct {
	lat, long float64
}
