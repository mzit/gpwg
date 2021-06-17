package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 41
	//marsAge := age
	marsAge := float64(age)

	marsDays := 687.0
	earthDay := 365.2425
	marsAge = marsAge * earthDay / marsDays
	fmt.Println("marsAge is ", marsAge)

	var bh float64 = 32768
	var h = int16(bh)
	fmt.Println(h)

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	countdown := 10
	str := "aaa" + strconv.Itoa(countdown) + "bbbbb"
	fmt.Println(str)

	countdown = 9
	str = fmt.Sprintf("aaaa%vbbbbb", countdown)
	fmt.Println(str)

	countdown, err := strconv.Atoi("10")
	if err != nil {

	}
	fmt.Println(countdown)
}
