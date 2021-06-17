package main

import "fmt"

func main() {
	planets := []string{
		"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh",
	}
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)

	//terrestrial = planets[0:4]
	//worlds = append(terrestrial,"Ceres")
	//fmt.Println(planets)
}
