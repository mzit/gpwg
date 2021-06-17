package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var bucket []int
	for true {
		beforeCap := cap(bucket)
		bucket = append(bucket, rand.Intn(10))
		afterCap := cap(bucket)
		if beforeCap != afterCap {
			fmt.Println(afterCap)
		}
	}
}
