package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}
	//sort.StringSlice(planets).Sort()
	sort.Strings(planets)
	fmt.Println(planets)
}
