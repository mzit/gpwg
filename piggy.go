package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bank := 0
	for bank < 2000 {
		switch rand.Intn(3) {
		case 0:
			bank += 5
		case 1:
			bank += 10
		case 2:
			bank += 25
		}
		dollars := bank / 100
		cents := bank % 100
		fmt.Printf("$%d.%02d\n", dollars, cents)

	}
}
