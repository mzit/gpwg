package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}
	locations := []location{
		{name: "name", lat: -4.58, long: 137.4},
		{name: "bbb", lat: -14.58, long: 175.4},
		{name: "ccc", lat: -122.58, long: 128.4},
	}
	fmt.Println(locations)
}
