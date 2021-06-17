package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "ajdklak", "akjlsd", "kjk", "jkkjkl", "jklkljkl")
	fmt.Println(dwarfs, len(dwarfs), cap(dwarfs))
}
