package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)
	fmt.Printf("%T\n", &answer)

	address := &answer
	fmt.Println(*address)

	canada := "canada"
	var home *string

	fmt.Printf("%T\n", home)
	home = &canada
	fmt.Println(*home)
}
