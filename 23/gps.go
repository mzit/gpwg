package main

import (
	"fmt"
	"math"
)

type gpsLocation struct {
	name      string
	lat, long float64
}

type gpsWorld struct {
	radius float64
}

type gps struct {
	currentLocation gpsLocation
	dstLocation     gpsLocation
	world           gpsWorld
}

func (gl gpsLocation) description() string {
	return fmt.Sprintf("%v (%.1fº, %.1fº)", gl.name, gl.lat, gl.long)
}

func (w gpsWorld) distance(p1, p2 gpsLocation) float64 {
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

func (g gps) distance() float64 {
	return g.world.distance(g.currentLocation, g.dstLocation)
}

func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.dstLocation.description())
}

type rover struct {
	gps
}

func main() {
	mars := gpsWorld{radius: 3389.5}
	bradbury := gpsLocation{"Bradbury Landing", -4.5895, 137.4417}
	elysium := gpsLocation{"Elysium Planitia", 4.5, 135.9}

	gps := gps{
		world:           mars,
		currentLocation: bradbury,
		dstLocation:     elysium,
	}

	curiosity := rover{
		gps: gps,
	}

	fmt.Println(curiosity.message())
}
