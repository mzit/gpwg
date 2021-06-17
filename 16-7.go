package main

import "fmt"

func terraform(planets [3]string) [3]string {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
	return planets
}

func main() {
	planets := [...]string{
		"AAA",
		"BBB",
		"CCC",
	}

	newPlanet := terraform(planets)
	fmt.Println(planets)
	fmt.Println(newPlanet)
}
