package main

import "fmt"

type Planets []string

func (p Planets) terraform() {
	for i, s := range p {
		p[i] = "New " + s
	}
}

func main() {
	planets := Planets{
		"aaa", "bbb", "ccc",
	}
	Planets(planets[1:]).terraform()
	fmt.Println(planets)
}
