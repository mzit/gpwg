package main

import "fmt"

func main() {
	var administrator *string
	scolese := "aajklaksdj"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "bbbbbb"
	administrator = &bolden
	fmt.Println(*administrator)

	*administrator = "ccccc"

	fmt.Println(bolden)

	timmy := person{
		name: "Timo",
		age:  10,
	}
	timmy.superpower = "flying"
	//birthday(&timmy)
	timmy.birthday()
	fmt.Printf("%+v\n", timmy)

	superpowers := &[3]string{"flight", "invisibility", "super strength"}

	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

}

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}
