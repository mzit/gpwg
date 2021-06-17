package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg."
	for _, i := range message {
		if i >= 'a' && i <= 'z' {
			i -= 3
			if i < 'a' {
				i += 26
			}
		} else if i >= 'A' && i <= 'Z' {
			i -= 3
			if i < 'A' {
				i += 26
			}
		}
		fmt.Printf("%c", i)
	}
}
