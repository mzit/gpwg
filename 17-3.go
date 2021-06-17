package main

import "fmt"

func main() {
	dwarfs := []string{"aaa", "bbb", "ccc"}
	dwarfArray := [...]string{"aaa", "bbb", "ccc"}

	fmt.Printf("slice %T \n", dwarfs)
	fmt.Printf("array %T \n", dwarfArray)
}
