package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
	}
	opportunity := location{
		lat:  -1.9462,
		long: 354.4734,
	}
	fmt.Println(opportunity)
	spirit := location{-14, 175}
	fmt.Println(spirit)
	fmt.Printf("%v\n", spirit)
	fmt.Printf("%+v\n", spirit)

}
