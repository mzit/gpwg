package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	lightYear := distance / lightSpeed / 86400 / 365
	fmt.Println(lightYear)
	fmt.Printf("%T", lightYear)
}
