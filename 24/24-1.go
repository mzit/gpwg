package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct {
}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", 12)
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type crater struct {
}

type starship struct {
	laser
}

func main() {
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())
	shout(martian{})
	shout(laser(1))
	s := starship{
		laser(3),
	}
	fmt.Println(s.talk())
	shout(s)
}
