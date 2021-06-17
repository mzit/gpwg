package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	//height行 width列
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}
func (u Universe) Show() {
	for _, height := range u {
		for _, width := range height {
			if width {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func (u Universe) Seed() {
	for i := 0; i < (height * width / 4); i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (x + height) % height
	return u[x][y]
}
func main() {
	u := NewUniverse()
	u.Seed()
	u.Show()
}
