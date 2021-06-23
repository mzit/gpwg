package main

import "fmt"

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v,%v", l.lat, l.long)
}

func main() {
	cur := location{
		lat:  -4,
		long: 137,
	}
	fmt.Println(cur)
}
